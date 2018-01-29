package eval

import (
	"fmt"
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		name string
		expr string
	}{
		{"plus", "1 + 2"},
		{"plus", "(1 + 2)"},
		{"input pow", "pow(x, 3) + pow(y, 3)"},
		{"input temp", " 5 / 9 * (F - 32)"},
	}
	var prevExpr string
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.expr != prevExpr {
				fmt.Printf("\n%s\n", test.expr)
				prevExpr = test.expr
			}
			expr, err := Parse(test.expr)
			if err != nil {
				t.Error(err)
			}
			exprAgain, err := Parse(expr.String())
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(expr, exprAgain) {
				t.Errorf("%v != Parse(%v.String())\tresult %v", expr, expr, exprAgain)
			}
		})
	}

}
