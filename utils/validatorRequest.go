package utils

import (
	"go-server/global"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

func InitTrans(local string) (err error) {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器
		uni := ut.New(enT, zhT)
		global.Trans, _ = uni.GetTranslator(local)
		switch local {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, global.Trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, global.Trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, global.Trans)
		}
		return
	}
	return
}
