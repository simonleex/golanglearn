package golanglearn

import "testing"
import "unsafe"

func TestImplictInherit(t *testing.T){
	tpyf := FatherType{1}
	tpys := SonType{tpyf,2}

	tpyf.print()
	tpys.print()

	((*FatherType)((unsafe.Pointer(&tpys)))).print()
}