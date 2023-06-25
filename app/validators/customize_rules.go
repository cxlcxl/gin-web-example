package validators

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"regexp"
)

var (
	rules = map[string]validator.Func{
		"password": password,
		"username": username,
	}
)

func password(fl validator.FieldLevel) bool {
	_pass := fl.Field().String()
	if _pass == "" {
		return true
	}
	if ok, err := regexp.MatchString(passRule, _pass); err != nil {
		return false
	} else {
		return ok
	}
}

func username(fl validator.FieldLevel) bool {
	_uname := fl.Field().String()
	if _uname == "" {
		return true
	}
	if ok, err := regexp.MatchString(unameRule, _uname); err != nil {
		return false
	} else {
		return ok
	}
}

func RegisterValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		for s, f := range rules {
			_ = v.RegisterValidation(s, f)
		}
	}
}
