package validate

import(
	validation_constants "github.com/matehaxor03/holistic_validator/validation_constants"
	validation_functions "github.com/matehaxor03/holistic_validator/validation_functions"
	"fmt"
	"strings"
)

type FileNameCharacterWhitelist struct {
	ValidateFileName func(file_name string) ([]error)
}

func NewFileNameCharacterWhitelist() (*FileNameCharacterWhitelist) {
	whiltelist := validation_constants.GetValidFilenameCharacters()

	validateFilenName := func(file_name string) ([]error) {
		var errors []error
		
		if file_name == "" {
			errors = append(errors, fmt.Errorf("is empty"))
		}

		if strings.Contains(file_name, "..") {
			errors = append(errors, fmt.Errorf("is unable to contain .."))
		}

		whitelist_errors := validation_functions.WhitelistCharacters(whiltelist, file_name, "Validator.ValidateFileName", "file.file_name")
		if whitelist_errors != nil {
			errors = append(errors, whitelist_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
	}


	x := FileNameCharacterWhitelist {
		ValidateFileName: func(file_name string) ([]error) {
			return validateFilenName(file_name)
		},
	}

	return &x
}
