package eval

type Var string

type literal float64

type Expr interface {
	Eval(env Env) float64
	// Checkは、この Expr 内のエラーを報告し、セットにその Var を追加します。
	Check(vars map[Var]bool) error

	String() string
}

type unary struct {
	op rune
	x  Expr
}

type binary struct {
	op   rune
	x, y Expr
}

type call struct {
	fn   string
	args []Expr
}
