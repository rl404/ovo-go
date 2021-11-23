package ovo

import (
	"context"
	"reflect"
	"regexp"
	"strings"

	"github.com/dongri/phonenumber"
	"github.com/go-playground/mold/v4"
	"github.com/go-playground/mold/v4/modifiers"
	"github.com/go-playground/validator/v10"
)

var mod *mold.Transformer
var val *validator.Validate

func init() {
	val = validator.New()
	val.RegisterValidationCtx("e164", validatePhone)

	mod = modifiers.New()
	mod.Register("no_space", modNoSpace)
	mod.Register("e164", modPhone)
}

func modNoSpace(ctx context.Context, fl mold.FieldLevel) error {
	switch fl.Field().Kind() {
	case reflect.String:
		fl.Field().SetString(strings.Replace(fl.Field().String(), " ", "", -1))
	}
	return nil
}

func modPhone(ctx context.Context, fl mold.FieldLevel) error {
	switch fl.Field().Kind() {
	case reflect.String:
		fl.Field().SetString(phonenumber.ParseWithLandLine(fl.Field().String(), "ID"))
	}
	return nil
}

func validatePhone(ctx context.Context, fl validator.FieldLevel) bool {
	return regexp.MustCompile("^[1-9]?[0-9]{7,14}$").MatchString(fl.Field().String())
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
			case "e164":
				return errPhoneField()
			default:
				return errInvalidValueField(e.Field())
			}
		}
		return err
	}
	return nil
}
