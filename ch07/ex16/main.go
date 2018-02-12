package main

import (
	"html/template"
	"log"
	"net/http"

	"gopl.io/ch07/ex16/eval"
)

type Calculator struct {
	Expr   string
	Result float64
}

var calc = template.Must(template.New("Calculator").Parse(`
<h1>電卓プログラム　でんた君</h1>
<form method="get" action="/">
  <input type="text" name="expr" value="{{.Expr}}"/>
  <input type="submit" value="calc"/>
</form>
<p>result {{.Result}}</p>
`))

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	data := &Calculator{Expr: "", Result: 0}
	for k, v := range r.URL.Query() {
		if k == "expr" {
			data.Expr = v[0]
			expr, err := eval.Parse(v[0])
			if err == nil {
				data.Result = expr.Eval(eval.Env{})
			}
		}
	}
	calc.Execute(w, data)
}
