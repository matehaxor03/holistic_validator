package validate

import (
	"fmt"

	validation_constants "github.com/matehaxor03/holistic_validator/validation_constants"
	validation_functions "github.com/matehaxor03/holistic_validator/validation_functions"
)

type IPAddressWhitelist struct {
	ValidateIPAddress        func(IPAddress string) []error
	GetValidateIPAddressFunc func() *func(string) []error
}

func NewIPAddressWhitelist() *IPAddressWhitelist {
	valid_characters := validation_constants.GetValidIPAddressCharacters()
	cache := make(map[string]interface{})

	validateIPAddress := func(IPAddress string) []error {
		if _, found := cache[IPAddress]; found {
			return nil
		}

		var errors []error
		if IPAddress == "" {
			errors = append(errors, fmt.Errorf("IP Address is empty"))
		}

		whitelist_errors := validation_functions.WhitelistCharacters(valid_characters, IPAddress, "Validator.ValidateIPAddress", "host.IPAddress")
		if whitelist_errors != nil {
			errors = append(errors, whitelist_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		cache[IPAddress] = nil
		return nil
	}

	x := IPAddressWhitelist{
		ValidateIPAddress: func(IPAddress string) []error {
			return validateIPAddress(IPAddress)
		},
		GetValidateIPAddressFunc: func() *func(IPAddress string) []error {
			function := validateIPAddress
			return &function
		},
	}

	return &x
}
