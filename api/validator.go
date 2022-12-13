package api

import (
	"simplebank/utils"
	"github.com/go-playground/validator/v10"
)

// validate the currency implementation validator...
var validCurrency validator.Func = func(fl validator.FieldLevel) bool {

    if currency, ok :=	fl.Field().Interface().(string); ok {
		return utils.IsSuppertedCurrency( currency )
	}

	return false
}