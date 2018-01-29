package eval

import (
	"math"
	"testing"
)

func TestMin_String(t *testing.T) {
	tests := []struct {
		name     string
		expr     Expr
		expected string
	}{
		{"2 plus value", &min{literal(1), literal(128)}, "min(1, 128)"},
		{"minus and plus value", &min{literal(-30), literal(20)}, "min(-30, 20)"},
		{"2 minus value", &min{literal(-30), literal(-200)}, "min(-30, -200)"},
		{"0 and plus value", &min{literal(0), literal(200)}, "min(0, 200)"},
		{"0 and minus value", &min{literal(-30), literal(0)}, "min(-30, 0)"},
		{"2 value 0", &min{literal(0), literal(0)}, "min(0, 0)"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.expr.String()
			if actual != test.expected {
				t.Errorf("%v expceted %s but actual is %s\n", test.expr, test.expected, actual)
			}

		})
	}
}

func TestMin_Eval(t *testing.T) {
	tests := []struct {
		name     string
		expr     Expr
		env      Env
		expected float64
	}{
		{"0 var", &min{literal(1.0), literal(2.0)}, Env{}, 1.0},
		{"1 var", &min{Var("x"), literal(2.0)}, Env{"x": 1.0}, 1.0},
		{"2 var", &min{Var("x"), Var("y")}, Env{"x": 1.0, "y": -2.0}, -2.0},
		{"2 plus value", &min{literal(1), literal(128)}, Env{}, 1},
		{"minus and plus value", &min{literal(-30), literal(20)}, Env{}, -30},
		{"2 minus value", &min{literal(-30), literal(-200)}, Env{}, -200},
		{"0 and plus value", &min{literal(0), literal(200)}, Env{}, 0},
		{"0 and minus value", &min{literal(-30), literal(0)}, Env{}, -30},
		{"2 value 0", &min{literal(0), Var("y")}, Env{"y": 0}, 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.expr.Eval(test.env)
			const eps = 1e-10

			if math.Abs(actual-test.expected) > eps {
				t.Errorf("")
			}

		})
	}
}

func TestMin_Check(t *testing.T) {
	tests := []struct {
		name     string
		expr     Expr
		expected bool
	}{
		{"2 literal ok", &min{literal(1), literal(2)}, true},
		{"1 literal and 1 var ok", &min{literal(1), Var("x")}, true},
		{"unknown function ng", &min{literal(1), Expr(call{})}, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			vars := make(map[Var]bool)
			err := test.expr.Check(vars)
			if (err == nil) != test.expected {
				t.Errorf("%s Check expected %t but actually %t\n", test.expr, test.expected, err != nil)
			}
		})
	}
}
