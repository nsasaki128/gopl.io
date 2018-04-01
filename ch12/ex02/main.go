package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

const maxNest = 10

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x), 0)
}

func display(path string, v reflect.Value, nestLevel int) {
	nestLevel++
	if nestLevel > maxNest {
		return
	}
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i), nestLevel)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			filePath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(filePath, v.Field(i), nestLevel)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path,
				formatAtom(key)), v.MapIndex(key), nestLevel)
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem(), nestLevel)
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem(), nestLevel)
		}
	default:
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}
func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	case reflect.Struct:
		var b bytes.Buffer

		b.WriteString(v.Type().String())
		b.WriteRune('{')
		for i := 0; i < v.NumField(); i++ {
			b.WriteString(fmt.Sprintf("%s:%s", v.Type().Field(i).Name, formatAtom(v.Field(i))))
			if i < v.NumField()-1 {
				b.WriteString(", ")
			}
		}
		b.WriteRune('}')
		return b.String()
	case reflect.Array:
		var b bytes.Buffer

		b.WriteString(v.Type().String())
		b.WriteRune('{')
		for i := 0; i < v.Len(); i++ {
			b.WriteString(formatAtom(v.Index(i)))
			if i < v.Len()-1 {
				b.WriteString(", ")
			}
		}
		b.WriteRune('}')
		return b.String()
	default:
		return v.Type().String() + " value"
	}
}

func main() {
	type Sample struct {
		Id   int
		Name string
	}

	type Test struct {
		SampleMap map[Sample]int
	}
	sample := Test{
		SampleMap: map[Sample]int{
			Sample{Id: 1, Name: "01"}: 2,
			Sample{Id: 2, Name: "02"}: 3,
			Sample{Id: 3, Name: "03"}: 4,
			Sample{Id: 4, Name: "04"}: 5,
			Sample{Id: 5, Name: "05"}: 6,
		},
	}
	test := map[string]int{"01": 1, "02": 2, "03": 3}
	arrayMapTest := map[[2]string]int{[2]string{"00", "01"}: 1, [2]string{"01", "02"}: 2}
	Display("sample", sample)
	Display("test", test)
	Display("arrayMap", arrayMapTest)
	type Cycle struct {
		Value int
		Tail  *Cycle
	}
	var c Cycle
	c = Cycle{42, &c}
	Display("c", c)
}
