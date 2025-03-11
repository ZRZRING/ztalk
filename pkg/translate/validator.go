package translate

import (
	"fmt"
	"reflect"
	"strings"
	"ztalk/internal/models"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// Trans 定义一个响应的翻译器
var Trans ut.Translator

// Init 初始化翻译器
func Init(locale string) (err error) {
	// 修改 gin 框架中的 Validator 引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 【强迫症】注册一个获取 json tag 的自定义方法
		// v.RegisterTagNameFunc(RegisterJsonTag)

		// 【强迫症】为 SignUpParam 注册自定义校验方法
		// v.RegisterStructValidation(SignUpParamStructLevelValidation, models.SignUpParam{})

		// 注册中英文
		zhT := zh.New()
		enT := en.New()

		// 第一个参数是备用(fallback)的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)

		// locale 通常取决于 http 请求头的 'Accept-Language'
		// 也可以使用 uni.FindTranslator(...) 传入多个 locale 进行查找
		var ok bool
		Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// 注册翻译器
		if locale == "en" {
			err = enTranslations.RegisterDefaultTranslations(v, Trans)
		} else if locale == "zh" {
			err = zhTranslations.RegisterDefaultTranslations(v, Trans)
		} else {
			err = enTranslations.RegisterDefaultTranslations(v, Trans)
		}

		return
	}
	return
}

// RemoveTopStruct 去除结构体名前缀
// 实现方法：找到最后一个出现的 '.' 截取之后的所有字符
func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

// RegisterJsonTag 获取 json tag
func RegisterJsonTag(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	if name == "-" {
		return ""
	}
	return name
}

// SignUpParamStructLevelValidation 自定义 SignUpParam 结构体校验函数
func SignUpParamStructLevelValidation(sl validator.StructLevel) {
	su := sl.Current().Interface().(models.SignUpParam)
	if su.Password != su.RePassword {
		sl.ReportError(su.RePassword, "re_password", "RePassword", "eqfield", "password")
	}
}
