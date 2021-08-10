package rules

import (
	//"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"

	//enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// 定义一个全局翻译器T
var trans ut.Translator

//初始化翻译器
func Init() {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		zhT := zh.New() // 中文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(zhT, zhT)

		// locale 通常取决于 http 请求头的 'Accept-Language'
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		trans, _ = uni.GetTranslator("zh")

		// 注册翻译器
		err := zhTranslations.RegisterDefaultTranslations(v, trans)
		if err != nil {
			panic("validate init failed, err:" + err.Error())
		}
	}
}

func Translate(errs validator.ValidationErrors) string {
	var errList []string
	for _, err := range errs {
		errList = append(errList, err.Translate(trans))
	}
	return strings.Join(errList, "|")
}
