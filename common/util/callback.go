package util

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/parnurzeal/gorequest"
	"github.com/rohanchauhan02/clean/common/schemas"
	"github.com/tidwall/gjson"
)

func restAuth(joltInput []byte, generateKey schemas.GeneratedKey, mappedGeneratedKey map[string]string) (err error) {
	var isSuccess bool = true
	var payload []byte

	config := *generateKey.AuthConfig
	restConfig := config.RestConfig
	constructedHeaders := map[string]string{}
	callbackHeaders := restConfig.Headers
	callbackURL := fmt.Sprintf("%s/%s", restConfig.BaseUrl, restConfig.EndPoint)

	// Response mapping should be configured in the product config
	if restConfig.ResponseMapping == nil {
		return errors.New("response mapping auth is not correctly configured in product config")
	}

	if restConfig.BodyMapping != nil && restConfig.StaticBodyMapping != nil {
		return fmt.Errorf("rest config should not be contain both for body mapping and static body mapping")
	}

	// Construct headers
	for key, val := range callbackHeaders {
		constructedHeaders[key] = val
	}

	if restConfig.StaticBodyMapping != nil {
		StaticBodyMapping := restConfig.StaticBodyMapping
		// Check if the static body mapping having Payloads
		// If so then get from that instead

		if v, ok := StaticBodyMapping["Payloads"]; ok {
			payload, err = json.Marshal(v)
		} else {
			payload, err = json.Marshal(StaticBodyMapping)
		}

		if err != nil {
			return err
		}
	} else {
		joltSpec, err := json.Marshal(restConfig.BodyMapping)
		if err != nil {
			return err
		}

		payload, err = ParseJoltMapping(joltSpec, joltInput)
		if err != nil {
			return err
		}
	}

	_, body, err := GorequestHTTPClient(restConfig.Method, callbackURL, callbackHeaders, string(payload), false, 0)

	fmt.Printf("Response Auth REST from: %v is: %v \n", callbackURL, body)

	if err != nil {
		fmt.Printf("Response Auth REST got error from: %v. Error is: %s \n", callbackURL, err.Error())
		return err
	}

	if config.RestConfig.SuccessResponseBody != nil {
		for bodyKey, bodyValues := range config.RestConfig.SuccessResponseBody {
			if gjson.Get(body, bodyKey).Type == gjson.Null {
				isSuccess = false
				break
			}
			for _, bodyValue := range bodyValues {
				if bodyValue != nil {
					switch bodyTypeVal := bodyValue.(type) {
					case int64:
						respValInt := gjson.Get(body, bodyKey).Int()
						if respValInt != bodyTypeVal {
							isSuccess = false
						}
					case float64: // just to handle 0
						respValInt := int(gjson.Get(body, bodyKey).Float())
						if respValInt != int(bodyTypeVal) {
							isSuccess = false
						}
					case int:
						respValInt := int(gjson.Get(body, bodyKey).Int())
						if respValInt != bodyTypeVal {
							isSuccess = false
						}
					case string:
						if gjson.Get(body, bodyKey).String() != bodyTypeVal {
							isSuccess = false
						}
					}
				}
			}
		}

		if !isSuccess {
			return errors.New("auth rest response body not success")
		}
	}

	// Map the token
	for k, v := range restConfig.ResponseMapping {
		key := fmt.Sprintf("generated_keys.%s.response_mapping.%s", generateKey.Name, k)
		if strings.EqualFold(k, "BEARER_TOKEN") {
			// Concat Bearer to the value
			authToken := gjson.Get(body, v).String()
			if authToken == "" {
				return errors.New("cannot find auth token in the response")
			}
			token := fmt.Sprintf("%s %s", "Bearer", authToken)
			mappedGeneratedKey[key] = token
		} else {
			mappedGeneratedKey[key] = gjson.Get(body, v).String()
		}
	}

	return nil
}

func ParseCallbackConfig(joltSpec, joltInput, callbackConfig []byte) (*gorequest.SuperAgent, error) {
	parsedBody, err := ParseJoltMapping(joltSpec, joltInput)
	if err != nil {
		return nil, err
	}

	// order body payload
	orderedPayload := make(map[string]interface{})
	err = json.Unmarshal(parsedBody, &orderedPayload)
	if err != nil {
		return nil, err
	}

	var parsedPayloadBodyStr string

	// Check if the orderedPayload have Results
	// If so then get from that instead
	if v, ok := orderedPayload["Results"]; ok {
		parsedPayloadBody, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		parsedPayloadBodyStr = string(parsedPayloadBody)
	} else {
		parsedPayloadBody, err := json.Marshal(orderedPayload)
		if err != nil {
			return nil, err
		}
		parsedPayloadBodyStr = string(parsedPayloadBody)
	}

	callbackConf := &schemas.CallbackConfig{}
	err = json.Unmarshal(callbackConfig, callbackConf)
	if err != nil {
		return nil, err
	}

	// Constructing signature and mapping to generate headers in next process
	mappedGeneratedKey := map[string]string{}
	mappedConfigKey := map[string]string{}

	// construct config keys
	for _, configKey := range callbackConf.ConfigKeys {
		mappedConfigKey[fmt.Sprintf("%s%s", "config_keys.", configKey.Name)] = configKey.Key
		// NOTE: will the ssm key already replaced by the actual SSM key value from product service?
	}

	// construct generated keys
	for _, generateKey := range callbackConf.GeneratedKeys {
		generateKeyName := fmt.Sprintf("%s%s", "generated_keys.", generateKey.Name)
		if generateKey.Type == "timestamp" {
			if generateKey.Format != "milliseconds" {
				mappedGeneratedKey[generateKeyName] = time.Now().Format(generateKey.Format)
			} else {
				mappedGeneratedKey[generateKeyName] = strconv.FormatInt(time.Now().Unix(), 10)
			}
		} else if generateKey.Type == "uuid" {
			mappedGeneratedKey[generateKeyName] = uuid.New().String()
		} else if generateKey.Type == "body" {
			mappedGeneratedKey[generateKeyName] = parsedPayloadBodyStr
			if generateKey.AuthConfig.Encoding == "base64" {
				mappedGeneratedKey[generateKeyName] = base64.StdEncoding.EncodeToString([]byte(parsedPayloadBodyStr))
			}
		} else if generateKey.Type == "auth" {
			var generatedSignature string
			secretKey := ReplaceValue(mappedConfigKey, mappedGeneratedKey, generateKey.AuthConfig.Secret)
			switch generateKey.AuthType {
			case "HMAC_SHA256":
				signatureFormat := generateKey.AuthConfig.MessageGeneration.Format
				var signatureParams []interface{}
				for _, paramRawSig := range generateKey.AuthConfig.MessageGeneration.Params {
					signatureParams = append(signatureParams, ReplaceValue(mappedConfigKey, mappedGeneratedKey, paramRawSig))
				}
				finalSignature := fmt.Sprintf(signatureFormat, signatureParams...)
				mac := hmac.New(sha256.New, []byte(secretKey))
				mac.Write([]byte(finalSignature))
				if generateKey.AuthConfig.Encoding == "base64" {
					sign := mac.Sum(nil)
					generatedSignature = base64.StdEncoding.EncodeToString(sign)
				} else {
					sign := mac.Sum(nil)
					generatedSignature = string(sign)
				}
			case "BASIC_AUTH":
				signatureFormat := generateKey.AuthConfig.MessageGeneration.Format
				var signatureParams []interface{}
				for _, paramRawSig := range generateKey.AuthConfig.MessageGeneration.Params {
					signatureParams = append(signatureParams, ReplaceValue(mappedConfigKey, mappedGeneratedKey, paramRawSig))
				}
				finalSignature := fmt.Sprintf(signatureFormat, signatureParams...)
				if generateKey.AuthConfig.Encoding == "base64" {
					generatedSignature = fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(finalSignature)))
				} else {
					generatedSignature = fmt.Sprintf("Basic %s", finalSignature)
				}
			case "REST":
				err := restAuth(joltInput, *generateKey, mappedGeneratedKey)
				if err != nil {
					return nil, err
				}
			case "MD5":
				signatureFormat := generateKey.AuthConfig.MessageGeneration.Format
				var signatureParams []interface{}
				for _, paramRawSig := range generateKey.AuthConfig.MessageGeneration.Params {
					signatureParams = append(signatureParams, ReplaceValue(mappedConfigKey, mappedGeneratedKey, paramRawSig))
				}
				finalSignature := fmt.Sprintf(signatureFormat, signatureParams...)
				hash := md5.Sum([]byte(finalSignature))
				generatedSignature = hex.EncodeToString(hash[:])
			default:
				generatedSignature = secretKey
			}

			if generateKey.AuthType != "REST" {
				mappedGeneratedKey[generateKeyName] = generatedSignature
			}
		}
	}

	constructedHeaders := map[string]string{}
	callbackHeaders := callbackConf.Headers
	for key, val := range callbackHeaders {
		constructedHeaders[key] = ReplaceValue(mappedConfigKey, mappedGeneratedKey, val)
	}

	callbackURL := fmt.Sprintf("%s/%s", callbackConf.BaseUrl, callbackConf.EndPoint)
	req := GorequestSuperAgent(callbackConf.Method, callbackURL, constructedHeaders, parsedPayloadBodyStr, false)

	return req, nil
}

func ReplaceValue(mapConfigKeys, mapGeneratedKeys map[string]string, keyValue string) (result string) {
	if mapConfigKeys[keyValue] != "" {
		result = mapConfigKeys[keyValue]
	} else if mapGeneratedKeys[keyValue] != "" {
		result = mapGeneratedKeys[keyValue]
	} else {
		result = keyValue
	}

	return result
}
