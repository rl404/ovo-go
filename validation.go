package ovo

import (
	"context"
	"reflect"
	"strings"

	"github.com/go-playground/mold/v4"
	"github.com/go-playground/mold/v4/modifiers"
	"github.com/go-playground/validator/v10"
)


var mod *mold.Transformer
var val *validator.Validate

func init() {
	val = validator.New()

	mod = modifiers.New()
	mod.Register("no_space", modNoSpace)
}


func modNoSpace(ctx context.Context, fl mold.FieldLevel) error {
	switch fl.Field().Kind() {
	case reflect.String:
		fl.Field().SetString(strings.Replace(fl.Field().String(), " ", "", -1))
	}
	return nil
}

func validate(data interface{}) error {
	if err := mod.Struct(context.Background(), data); err != nil {
		return err
	}
	if err := val.Struct(data); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return err
		}
		for _, e := range errs {
			switch e.Tag() {
			case "required":
				return errRequiredField(e.Field())
			case "gt":
				return errGTField(e.Field(), e.Param())
			case "lt":
				return errLTField(e.Field(), e.Param())
			// case "gte":
			// 	return errGTEField(e.Field(), e.Param())
			// case "max":
			// 	return errMaxField(e.Field(), e.Param())
			// case "numeric":
			// 	return errNumericField(e.Field())
			// case "url":
			// 	return errURLField(e.Field())
			default:
				return errInvalidValueField(e.Field())
			}
		}
		return err
	}
	return nil
}
