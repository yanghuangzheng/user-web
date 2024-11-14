package initialize
import(
	"strings"
	"reflect"
	"fmt"

	"user-web/global"

	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin/binding"
    "github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
    ut "github.com/go-playground/universal-translator"
)
////////////////////////////////////////////////////////gin错误转换

func Inittrans(local string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册一个获取json的tag的自定义方法
		v.RegisterTagNameFunc(func (fld reflect.StructField)string{
			name:=strings.SplitN(fld.Tag.Get("json"),",",2)[0]
			if name=="-"{
				return""
			}
			return name
		})
		zh := zh.New() // 中文翻译器
		en := en.New()
		// 第一个是备用的语言环境，后面是应该支持的语言环境
		uni := ut.New(en, zh, en)
		global.Trans, ok = uni.GetTranslator(local)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", local)
		}
		switch local {
		case "en":
			en_translations.RegisterDefaultTranslations(v, global.Trans)
		case "zh":
			zh_translations.RegisterDefaultTranslations(v, global.Trans)
		default:
			en_translations.RegisterDefaultTranslations(v, global.Trans)
		}
		return
	}
	return
}