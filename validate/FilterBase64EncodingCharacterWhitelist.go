package validate

import (
	validation_constants "github.com/matehaxor03/holistic_validator/validation_constants"
	validation_functions "github.com/matehaxor03/holistic_validator/validation_functions"
	"fmt"
)

type Base64EncodingCharacterWhitelist struct {
	ValidateBase64Encoding func(base64 string) ([]error)
	GetValidateBase64EncodingFunc func() (*func(string) []error)
}

func NewBase64EncodingCharacterWhitelist() (*Base64EncodingCharacterWhitelist) {
	valid_characters := validation_constants.GetValidBase64EncodingCharacters()
	cache := make(map[string]interface{})

	validateBase64Encoding := func(base64 string) ([]error) {
		if _, found := cache[base64]; found {
			return nil
		}
		
		var errors []error
		if base64 == "" {
			errors = append(errors, fmt.Errorf("base64 is empty"))
		}

		whitelist_errors := validation_functions.WhitelistCharacters(valid_characters, base64, "Validator.ValidateBase64Encoding", "text.base64")
		if whitelist_errors != nil {
			errors = append(errors, whitelist_errors...)
		}

		//todo: check ends with =

		if len(errors) > 0 {
			return errors
		}

		cache[base64] = nil
		return nil
	}

	x := Base64EncodingCharacterWhitelist {
		ValidateBase64Encoding: func(base64 string) ([]error) {
			return validateBase64Encoding(base64)
		},
		GetValidateBase64EncodingFunc: func() (*func(base64 string) []error) {
			function := validateBase64Encoding
			return &function
		},
	}

	return &x
}
