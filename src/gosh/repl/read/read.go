package read

import (
	. "bytes"
	. "os"
	. "os/exec"
	. "gosh/repl/print"
)

var buffer []byte = make([]byte, 1000)
var ifs = []byte{' '}
var In = Stdin

var echo = []byte{'e', 'c', 'h', 'o'}
var exit = []byte{'e', 'x', 'i', 't'}

type Expr interface {
	Eval() (int, error)
}

type NoOp struct {}

func (n NoOp) Eval() (int, error) {
	return 0, nil
}

type EchoExpr struct {
	s []byte
}

func (e EchoExpr) Eval() (int, error) {
	Out.Write(e.s)
	return 0, nil
}

type CmdExpr struct {
	c *Cmd
}

func (c CmdExpr) Eval() (int, error) {
	output, err := c.c.Output()
	Out.Write(output)
	if err != nil {
		return 1, err
	}
	return 0, err
}

type ExitExpr struct {
	code int
}

func (e ExitExpr) Eval() (int, error) {
	return e.code, nil
}

func Read() Expr {
	In.Read(buffer)
	tokens := Split(TrimRight(buffer, string([]byte{0, '\n', ' ', '\t'})), ifs) // TODO: write a parser...
	switch {
	default:
		return CmdExpr{Command(string(tokens[0]))}
	case Equal(tokens[0], echo):
		return EchoExpr{Join(tokens[1:], ifs)}
	case Equal(tokens[0], exit):
		return ExitExpr{0}
	}
	panic("this can't happen")
}
