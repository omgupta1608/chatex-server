package validation

import "github.com/go-playground/validator/v10"

/**
 * validate data
 * @param reqData - must be a pointer to a struct
 * @returns errFields - list of fields in reqData struct that fail validation
 * @returns invalidValidationError - *validator.InvalidValidationError
 */
func ValidateReqData(reqData interface{}) (errFields []string, invalidValidationError error) {
	validate := validator.New()
	err := validate.Struct(reqData)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			panic(err)
		}

		// TODO: register custom validator translations for better error reporting
		for _, err := range err.(validator.ValidationErrors) {
			errFields = append(errFields, err.StructField())
		}
	}

	return errFields, nil
}
