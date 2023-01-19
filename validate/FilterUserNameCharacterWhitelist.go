package validate

import (
	validation_constants "github.com/matehaxor03/holistic_validator/validation_constants"
	validation_functions "github.com/matehaxor03/holistic_validator/validation_functions"
	"fmt"
)

type UsernameCharacterWhitelist struct {
	ValidateUsername func(username string) ([]error)
	GetValidateUsernameFunc func() (*func(string) []error)
}

func NewUsernameCharacterWhitelist() (*UsernameCharacterWhitelist) {
	valid_username_characters := validation_constants.GetValidUsernameCharacters()
	valid_usernames_cache := make(map[string]interface{})

	validateUsername := func(username string) ([]error) {
		if _, found := valid_usernames_cache[username]; found {
			return nil
		}
		
		var errors []error
		if username == "" {
			errors = append(errors, fmt.Errorf("username is empty"))
		}

		if len(username) < 2 {
			errors = append(errors, fmt.Errorf("username is too short must be at least 2 characters"))
		}

		whitelist_errors := validation_functions.WhitelistCharacters(valid_username_characters, username, "Validator.ValidateUsername", "dao.User.username")
		if whitelist_errors != nil {
			errors = append(errors, whitelist_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		valid_usernames_cache[username] = nil
		return nil
	}

	x := UsernameCharacterWhitelist {
		ValidateUsername: func(username string) ([]error) {
			return validateUsername(username)
		},
		GetValidateUsernameFunc: func() (*func(string) []error) {
			function := validateUsername
			return &function
		},
	}

	return &x
}
