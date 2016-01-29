package golanglearn
import "fmt"



type FatherType struct {
	data int
}

type SonType struct {
	FatherType
	data int
}

func (ft *FatherType)print(){
	fmt.Printf("father data:%d\n",ft.data)
}

func (st *SonType)print(){
	fmt.Printf("son's data:%d\n",st.data)
}
