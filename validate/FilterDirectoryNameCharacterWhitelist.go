package validate

import(
	validation_constants "github.com/matehaxor03/holistic_validator/validation_constants"
	validation_functions "github.com/matehaxor03/holistic_validator/validation_functions"
	"fmt"
	"strings"
)

type DirectoryNameCharacterWhitelist struct {
	ValidateDirectoryName func(directory_name string) ([]error)
}

func NewDirectoryNameCharacterWhitelist() (*DirectoryNameCharacterWhitelist) {
	whiltelist := validation_constants.GetValidDirectoryNameCharacters()

	validateDirectoryName := func(directory_name string) ([]error) {
		var errors []error
		
		if directory_name == "" {
			errors = append(errors, fmt.Errorf("is empty"))
		}

		if strings.Contains(directory_name, "..") {
			errors = append(errors, fmt.Errorf("is unable to contain .."))
		}

		whitelist_errors := validation_functions.WhitelistCharacters(whiltelist, directory_name, "Validator.ValidateDirectoryName", "file.directory_name")
		if whitelist_errors != nil {
			errors = append(errors, whitelist_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
	}


	x := DirectoryNameCharacterWhitelist {
		ValidateDirectoryName: func(directory_name string) ([]error) {
			return validateDirectoryName(directory_name)
		},
	}

	return &x
}
