package api

import (
	"github.com/AzfarInan/go-masterclass/simplebank/db/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		/// Check currecny is supported
		return util.IsSUpportedCurrency(currency)
	}
	return false
}
