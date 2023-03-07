package validation

import (
	"encoding/json"
	"errors"

	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/iagoid/HunCoding-Primeiro-CRUD/src/configuration/rest_err"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var (
	Validade = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(val, transl)
	}
}

// ValidateError faz a validação dos erros gerados pela requisição, em caso de erro  de validação também traduz os campos inválidos
func ValidateError(validation_err error) *rest_err.RestErr {
	var jsorErr *json.UnmarshalTypeError               // Erro de tipagem (passar int no lugar de string...)
	var jsonValidationError validator.ValidationErrors // Erro de validação

	if errors.As(validation_err, &jsorErr) {
		rest_err.NewBadRequestError("Invalid field type")
	}

	if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []rest_err.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	}

	return rest_err.NewBadRequestError("Error trying to convert fields")
}
