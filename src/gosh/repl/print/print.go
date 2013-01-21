package print

import . "os"

var Out *File = Stdout
var Err *File = Stderr

func Print(code int, err error) {
	if err != nil {
		Err.WriteString(err.Error())
	} else {
		Out.WriteString(string(code))
	}
}
