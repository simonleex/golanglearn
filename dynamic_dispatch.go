package golanglearn
import "fmt"


type TypeA struct {

}

type TypeB struct {

}

func MyTestDynamicDispatch(typ Printer){
	typ.print()
}

func (typ TypeA)print(){
	fmt.Printf("typea\n")
}

func (typ TypeB)print(){
	fmt.Printf("typeb\n")
}


type Printer interface {
	print()
}