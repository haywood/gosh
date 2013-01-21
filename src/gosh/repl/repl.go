package repl

import (
	. "gosh/repl/read"
	. "gosh/repl/eval"
	. "gosh/repl/print"
)

func Repl() {
	for {
		Print(Eval(Read()))
	}
}
