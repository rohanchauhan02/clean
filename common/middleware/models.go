package middleware

import (
	"encoding/json"
	"net/http"
	"time"

	CommonErrors "github.com/rohanchauhan02/clean/common/error"
	"github.com/rohanchauhan02/clean/common/util"
)

const (
	ErrMessageUnauthorizedToken = "Unauthorized token"

	AuthorizationHeaderEmpty       = "authorization header was empty"
	MalformedBearerTokenFormatting = "malformed bearer token formatting"
	MalformedJWTToken              = "JWT Parse failed, malformed JWT token received"
	MalformedExpiredToken          = "JWT Parse failed, expired JWT token received"
	UnmarshallUserServiceError     = "Failed to unmarshall user check endpoint"

	ContextUserKey          = "UserData"
	ContextHistoryUserKey   = "HistoryUserData"
	ContextHistorySourceKey = "HistoySourceData"

	ApiKeyEmpty = "x-api-key header was empty"
)

type (
	// Config defines the config for the JWTAuthMiddleware
	// Currently all fields are required
	Config struct {
		// Auth endpoint to use for HTTP request call
		// Required.
		JWTKeyEndpoint     string
		PartnerKeyEndpoint string
		Caller             string
		Timeout            time.Duration
	}

	CheckUserV1Request struct {
		Token string `json:"token"`
	}

	CheckUserV1ErrorResponse struct {
		Message string `json:"message"`
	}

	CheckUserV1Response struct {
		Data struct {
			User struct {
				ID                 int           `json:"id"`
				UUID               string        `json:"uuid"`
				FirstName          string        `json:"firstName"`
				LastName           string        `json:"lastName"`
				FullName           string        `json:"fullName"`
				Lang               string        `json:"lang"`
				Email              string        `json:"email"`
				PhoneNumber        string        `json:"phoneNumber"`
				ChattingID         string        `json:"chattingId"`
				LastOnline         interface{}   `json:"lastOnline"`
				IsActive           int           `json:"isActive"`
				IsIdentityVerified string        `json:"isIdentityVerified"`
				IsBankVerified     string        `json:"isBankVerified"`
				IsEmailVerified    string        `json:"isEmailVerified"`
				IsPhoneVerified    string        `json:"isPhoneVerified"`
				CreatedAt          time.Time     `json:"createdAt"`
				Role               string        `json:"role"`
				Groups             []interface{} `json:"groups"`
				BankAccounts       struct {
					ID                int         `json:"id"`
					UUID              string      `json:"uuid"`
					Entity            string      `json:"entity"`
					EntityID          int         `json:"entityId"`
					BankID            interface{} `json:"bankId"`
					BankName          interface{} `json:"bankName"`
					BranchName        interface{} `json:"branchName"`
					AccountNumber     interface{} `json:"accountNumber"`
					AccountHolderName interface{} `json:"accountHolderName"`
					IsPrimary         int         `json:"isPrimary"`
					CreatedAt         time.Time   `json:"createdAt"`
					UpdatedAt         time.Time   `json:"updatedAt"`
					Bank              interface{} `json:"bank"`
					XenditCode        interface{} `json:"xenditCode"`
				} `json:"bankAccounts"`
				VirtualAccount interface{}   `json:"virtualAccount"`
				Agent          interface{}   `json:"agent"`
				Partner        interface{}   `json:"partner"`
				Insurance      interface{}   `json:"insurance"`
				Permissions    []string      `json:"permissions"`
				Companies      []interface{} `json:"companies"`
				IsPassWeak     bool          `json:"isPassWeak"`
			} `json:"user"`
		} `json:"data"`
	}

	CheckUserV2Request struct {
		Token string `json:"token"`
	}

	CheckUserV2ErrorResponse struct {
		Message string `json:"message"`
	}

	CheckUserV2Response struct {
		Data struct {
			User struct {
				ID                 int         `json:"id"`
				UUID               string      `json:"uuid"`
				FirstName          string      `json:"firstName"`
				LastName           string      `json:"lastName"`
				FullName           string      `json:"fullName"`
				Gender             interface{} `json:"gender"`
				Lang               string      `json:"lang"`
				Nationality        interface{} `json:"nationality"`
				Email              string      `json:"email"`
				PhoneNumber        string      `json:"phoneNumber"`
				IsActive           int         `json:"isActive"`
				Locked             int         `json:"locked"`
				IsIdentityVerified string      `json:"isIdentityVerified"`
				IsBankVerified     string      `json:"isBankVerified"`
				IsEmailVerified    string      `json:"isEmailVerified"`
				IsPhoneVerified    string      `json:"isPhoneVerified"`
				CreatedAt          time.Time   `json:"createdAt"`
				Organization       struct {
					ID          string      `json:"id"`
					Name        string      `json:"name"`
					Code        string      `json:"code"`
					Country     string      `json:"country"`
					PartnerID   interface{} `json:"partnerId"`
					Alternative string      `json:"alternative"`
					Description string      `json:"description"`
					Type        string      `json:"type"`
					Logo        string      `json:"logo"`
					LogoSmall   interface{} `json:"logoSmall"`
					ColorTheme  struct {
						Primary   string `json:"primary"`
						Secondary string `json:"secondary"`
					} `json:"colorTheme"`
					Config struct {
						HasLevel      bool `json:"hasLevel"`
						HasAgent      bool `json:"hasAgent"`
						HasCommission bool `json:"hasCommission"`
					} `json:"config"`
					Contact struct {
					} `json:"contact"`
					Address           interface{} `json:"address"`
					DeletedAt         interface{} `json:"deletedAt"`
					SectionLevelID    interface{} `json:"sectionLevelId"`
					ProductCategories []struct {
						ID          string      `json:"id"`
						Name        string      `json:"name"`
						Slug        string      `json:"slug"`
						Description interface{} `json:"description"`
						Industry    string      `json:"industry"`
						Icon        string      `json:"icon"`
						IsDraft     int         `json:"isDraft"`
					} `json:"productCategories"`
					Industries []string      `json:"industries"`
					Relations  []interface{} `json:"relations"`
					Bank       interface{}   `json:"bank"`
				} `json:"organization"`
				Role       string `json:"role"`
				RoleDetail struct {
					ID          string `json:"id"`
					Name        string `json:"name"`
					Rank        int    `json:"rank"`
					Description string `json:"description"`
					IsLock      int    `json:"isLock"`
				} `json:"roleDetail"`
				Agent struct {
					Code string `json:"code"`
				} `json:"agent"`
				BankAccount struct {
				} `json:"bankAccount"`
				Section struct {
					ID    string `json:"id"`
					Name  string `json:"name"`
					Level struct {
						ID          string `json:"id"`
						Name        string `json:"name"`
						Description string `json:"description"`
						Level       int    `json:"level"`
					} `json:"level"`
					Description string `json:"description"`
					Parent      struct {
						ID          string        `json:"id"`
						Name        string        `json:"name"`
						Description string        `json:"description"`
						Levels      []interface{} `json:"levels"`
					} `json:"parent"`
					Levels []interface{} `json:"levels"`
					Code   string        `json:"code"`
				} `json:"section"`
				Level struct {
					ID          string `json:"id"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Level       int    `json:"level"`
				} `json:"level"`
				Groups       []interface{} `json:"groups"`
				Permissions  []string      `json:"permissions"`
				ProfileImage interface{}   `json:"profileImage"`
				LastLogin    time.Time     `json:"lastLogin"`
			} `json:"user"`
		} `json:"data"`
	}

	PartnerKeyErrorResponse struct {
		Message string `json:"message"`
	}

	PartnerKeyResponse struct {
		PartnerCode string `json:"partner_code"`
	}
)

// Unauthorized response is a method that returns an unauthorized response
// back to client. This is mostly used for auth middlewares currently
// such as JWTMiddlewares and KeyAuthMiddleware
func unauthorizedResponse(wrappedErr error, errorCode string, respErrorMessage string, w http.ResponseWriter) {

	resp, _ := CommonErrors.NewClientError(
		wrappedErr,
		http.StatusUnauthorized,
		errorCode,
		respErrorMessage,
		"")

	respStr, _ := json.Marshal(&resp)
	if respStr != nil {
		logger.Debugf("%s -- Response: %s", util.GetCallerMethod(), string(respStr))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	_, _ = w.Write(respStr)
}

func badRequestResponse(wrappedErr error, errorCode string, respErrorMessage string, w http.ResponseWriter) {

	resp, _ := CommonErrors.NewClientError(
		wrappedErr,
		http.StatusBadRequest,
		errorCode,
		respErrorMessage,
		"")

	respStr, _ := json.Marshal(&resp)
	if respStr != nil {
		logger.Debugf("%s -- Response: %s", util.GetCallerMethod(), string(respStr))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write(respStr)
}

func internalServerResponse(wrappedErr error, errorCode string, respErrorMessage string, w http.ResponseWriter) {
	resp, _ := CommonErrors.NewServerError(wrappedErr, http.StatusInternalServerError,
		errorCode, respErrorMessage, "")
	respStr, _ := json.Marshal(&resp)
	if respStr != nil {
		logger.Debugf("%s -- Response: %s", util.GetCallerMethod(), string(respStr))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write(respStr)

}
