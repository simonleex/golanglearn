package golanglearn

import (
	"testing"
)



func TestMyTestDynamicDispatch(t *testing.T){
	typa := TypeA{}
	typb := TypeB{}
	MyTestDynamicDispatch(typa)
	MyTestDynamicDispatch(typb)
}
