package util

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
	"log"
	"strings"
)

var trans ut.Translator

// 设置验证器中文翻译
func init() {
	if trans != nil {
		return
	}
	zhLoc := zh.New()
	enLoc := en.New()
	uni := ut.New(enLoc, zhLoc)
	trans, _ = uni.GetTranslator("zh")
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := zhtranslations.RegisterDefaultTranslations(v, trans); err != nil {
			log.Fatalln(err)
		}
	}
}

// TransError 翻译验证器错误提示
func TransError(err error) map[string]string {
	if errs, ok := err.(validator.ValidationErrors); ok {
		return errs.Translate(trans)
	}
	log.Println("[TransError]", err)
	return nil
}

func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
