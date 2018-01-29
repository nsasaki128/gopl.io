package eval

import "fmt"

type min struct {
	x, y Expr
}

func (m min) Eval(env Env) float64 {
	if m.x.Eval(env) < m.y.Eval(env) {
		return m.x.Eval(env)
	}
	return m.y.Eval(env)
}

func (m min) Check(vars map[Var]bool) error {
	if err := m.x.Check(vars); err != nil {
		return err
	}
	return m.y.Check(vars)
}

func (m min) String() string {
	return fmt.Sprintf("min(%s, %s)", m.x, m.y)
}
