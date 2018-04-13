package params

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var emailPattern = regexp.MustCompile(`^[a-zA-Z0-9\-._]+@[a-zA-Z0-9\-._]+$`)
var creditPattern = regexp.MustCompile(`^[0-9]{16}$`)
var zipPattern = regexp.MustCompile(`^[0-9]{7}$`)

func isConstraintPattern(value, constraint string) bool {
	switch constraint {
	case "email":
		return emailPattern.MatchString(value)
	case "credit":
		return creditPattern.MatchString(value)
	case "zip":
		return zipPattern.MatchString(value)
	default:
		return false
	}
}

// Unpackは、req内のHTTPリクエストパラメータから
// ptrが指す構造体のフィールドに値を移し替えます。
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	//実効的な名前をキーとするフィールドのマップを構築する
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem() //構造体変数
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
	}

	//リクエストの中の個々のパラメータに対する構造体のフィールドを更新
	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue //認識されなかったHTTPパラメータを無視
		}
		for _, value := range values {
			if !isConstraintPattern(value, name) {
				return fmt.Errorf("%s invalid for %s", value, name)
			}
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}

		}
	}
	return nil
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}

		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}

	return nil
}
