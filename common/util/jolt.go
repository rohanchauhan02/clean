package util

import (
	"github.com/qntfy/kazaam/v4"
)

func ParseJoltMapping(spec, input []byte) ([]byte, error) {
	k, err := kazaam.NewKazaam(string(spec))
	if err != nil {
		return nil, err
	}

	parsedStr, err := k.TransformJSONStringToString(string(input))
	if err != nil {
		return nil, err
	}
	return []byte(parsedStr), nil
}
