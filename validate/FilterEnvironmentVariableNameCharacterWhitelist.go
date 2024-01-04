package validate

import(
	validation_constants "github.com/matehaxor03/holistic_validator/validation_constants"
	validation_functions "github.com/matehaxor03/holistic_validator/validation_functions"
	"fmt"
)

type EnvironmentVariableNameWhitelist struct {
	ValidateEnvironmentVariableName func(environment_variable_name string) ([]error)
	GetValidateEnvironmentVariableNameFunc func() (*func(string) []error)
}

func NewEnvironmentVariableNameWhitelist() (*EnvironmentVariableNameWhitelist) {
	environment_varable_name_character_whitelist := validation_constants.GetValidEnvironmentVariableNameCharacters()
	valid_environment_variable_name_cache := make(map[string]interface{})

	validateEnvironmentVariableName := func(environment_variable_name string) ([]error) {
		if _, found := valid_environment_variable_name_cache[environment_variable_name]; found {
			return nil
		}
		
		var errors []error
		
		if environment_variable_name == "" {
			errors = append(errors, fmt.Errorf("environment_variable_name is empty"))
		}

		whitelist_errors := validation_functions.WhitelistCharacters(environment_varable_name_character_whitelist, environment_variable_name, "Validator.validateEnvironmentVariableName", "sys.env.variable_name")
		if whitelist_errors != nil {
			errors = append(errors, whitelist_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		valid_environment_variable_name_cache[environment_variable_name] = nil
		return nil
	}


	x := EnvironmentVariableNameWhitelist {
		ValidateEnvironmentVariableName: func(environment_variable_name string) ([]error) {
			return validateEnvironmentVariableName(environment_variable_name)
		},
		GetValidateEnvironmentVariableNameFunc: func() (*func(environment_variable_name string) []error) {
			function := validateEnvironmentVariableName
			return &function
		},
	}

	return &x
}
