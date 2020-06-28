package xvalidator

import "github.com/gin-gonic/gin/binding"

func ValidateStruct(obj interface{}) error {
	var validate = defaultValidator{}
	return validate.ValidateStruct(obj)
}

func init() {
	binding.Validator = new(defaultValidator)
}
