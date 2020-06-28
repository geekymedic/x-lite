package xvalidator

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
)

type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var _ binding.StructValidator = &defaultValidator{}

func (v *defaultValidator) ValidateStruct(obj interface{}) error {

	if kindOfData(obj) == reflect.Struct {

		v.lazyinit()

		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}

	return nil
}

func (v *defaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *defaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")

		// add any custom validations etc. here
		v.validate.RegisterValidation("et_contains", func(fl validator.FieldLevel) bool {
			return EtContains(fl)
		})

		v.validate.RegisterValidation("et_identity", func(fl validator.FieldLevel) bool {
			return Identity(fl)
		})

		v.validate.RegisterValidation("et_phone", func(fl validator.FieldLevel) bool {
			return PhoneNumber(fl)
		})

		v.validate.RegisterValidation("et_chinese", func(fl validator.FieldLevel) bool {
			return Chinese(fl)
		})

		v.validate.RegisterValidation("et_nickname", func(fl validator.FieldLevel) bool {
			return NickName(fl)
		})

		v.validate.RegisterValidation("et_long_time", func(fl validator.FieldLevel) bool {
			return EtLongTimeFormat(fl)
		})

		v.validate.RegisterValidation("et_short_time", func(fl validator.FieldLevel) bool {
			return EtShortTimeFormat(fl)
		})

		v.validate.RegisterValidation("et_json", func(fl validator.FieldLevel) bool {
			return Json(fl)
		})
	})
}

func kindOfData(data interface{}) reflect.Kind {

	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

func EtContains(fl validator.FieldLevel) bool {
	field := fl.Field()
	kind := field.Kind()

	var value string
	switch kind {
	case reflect.String:
		value = field.String()

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value = fmt.Sprintf("%d", field.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		value = fmt.Sprintf("%d", field.Uint())

	case reflect.Float32, reflect.Float64:
		value = fmt.Sprintf("%f", field.Float())

	case reflect.Bool:
		value = fmt.Sprintf("%v", reflect.Bool)

	case reflect.Slice, reflect.Map, reflect.Array: // TODO
		return false

	case reflect.Struct: // TODO
		return false
	}
	for _, s := range strings.Split(fl.Param(), "-") {
		if s == value {
			return true
		}
	}
	return false
}

func EtLongTimeFormat(fl validator.FieldLevel) bool {
	field := fl.Field()
	kind := field.Kind()

	var value string
	switch kind {
	case reflect.String:
		value = field.String()
	default:
		return false
	}
	_, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
	return err == nil
}

func EtShortTimeFormat(fl validator.FieldLevel) bool {
	field := fl.Field()
	kind := field.Kind()

	var value string
	switch kind {
	case reflect.String:
		value = field.String()
	default:
		return false
	}
	_, err := time.ParseInLocation("2006-01-02", value, time.Local)
	return err == nil
}

var identityRegexp = regexp.MustCompile("^(\\d{6})(\\d{8})(.*)")

func Identity(fl validator.FieldLevel) bool {
	if fl.Field().Kind() != reflect.String {
		return false
	}
	return identityRegexp.MatchString(fl.Field().String())
}

var phoneNumberRegexp = regexp.MustCompile(`^1([3578][0-9]|14[57]|5[^4])\d{8}$`)

func PhoneNumber(fl validator.FieldLevel) bool {
	if fl.Field().Kind() != reflect.String {
		return false
	}
	return phoneNumberRegexp.MatchString(fl.Field().String())
}

func Chinese(fl validator.FieldLevel) bool {
	value := fl.Field()
	if value.Kind() != reflect.String {
		return false
	}
	for _, runeValue := range value.String() {
		if !unicode.Is(unicode.Han, runeValue) {
			return false
		}
	}
	return true
}

func NickName(fl validator.FieldLevel) bool {
	value := fl.Field()
	if value.Kind() != reflect.String {
		return false
	}
	strValue := value.String()
	for i, w := 0, 0; i < len(strValue); i += w {
		runeValue, width := utf8.DecodeRuneInString(strValue[i:])
		w = width
		if unicode.Is(unicode.Han, runeValue) {
			continue
		}
		if runeValue == rune('_') {
			continue
		}

		if (runeValue >= rune('A') && runeValue <= rune('Z')) || (runeValue >= rune('a') && runeValue <= rune('z')) {
			continue
		}

		if runeValue >= rune('0') && runeValue <= rune('9') {
			continue
		}

		return false
	}
	return true
}

func Json(fl validator.FieldLevel) bool {
	var i interface{}
	value := fl.Field()
	if value.Kind() != reflect.String {
		return false
	}
	return nil == json.NewDecoder(strings.NewReader(value.String())).Decode(&i)
}
