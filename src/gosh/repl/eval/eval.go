package eval

import . "gosh/repl/read"

func Eval(e Expr) (int, error) {
	return e.Eval()
}
