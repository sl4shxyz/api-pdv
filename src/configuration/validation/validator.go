package validation

import (
  "api_pdv/src/configuration/rest_error"

  "github.com/go-playground/locales/en"
  validator "github.com/go-playground/validator/v10"
  ut "github.com/go-playground/universal-translator"
  "github.com/gin-gonic/gin/binding"
  en_translation "github.com/go-playground/validator/v10/translations/en"
  "encoding/json"
  "errors"
)

var (
  validate = validator.New()
  transl ut.Translator
)

func init() {
  if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
    en := en.New()
    un := ut.New(en, en)
    transl, _ = un.GetTranslator("en")
    en_translation.RegisterDefaultTranslations(val, transl)
  }
}

func ValidateUserError(val_err error) *rest_error.RestError {
  var jsonErr *json.UnmarshalTypeError
  var jsonValError validator.ValidationErrors

  if errors.As(val_err, &jsonErr) {
    return rest_error.NewBadRequestError("invalid field type")
  } else if errors.As(val_err, &jsonValError) {
    errCauses := []rest_error.Causes{}

    for _, v := range val_err.(validator.ValidationErrors) {
      cause := rest_error.Causes {
        Message: v.Translate(transl),
        Field: v.Field(),
      }

      errCauses = append(errCauses, cause)
    }

    return rest_error.NewBadRequestValidationError("invalid field", errCauses)
  }

  return rest_error.NewBadRequestError("error converting field type")
}
