package validator

import (
	"fmt"
	"reflect"

	"ginblog/pkg/helper/errmsg"
	"github.com/go-playground/locales/zh_Hans"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/ar"
)

func Validate(date interface{}) (string, int) {
	validate := validator.New()
	uni := ut.New(zh_Hans.New())
	trans, _ := uni.GetTranslator("zh_Hans_CN")
	err := ar.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println("err", err)
	}
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})
	err = validate.Struct(date)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), errmsg.ERROR
		}
	}
	return "", errmsg.SUCCESS
}
