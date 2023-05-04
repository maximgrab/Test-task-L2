import "fmt"

type Facade struct {
	str1 *struct1
	str2 *struct2
}

type struct1 struct {
	msg string
}
type struct2 struct {
	msg string
}

func (f *Facade) str1f() {
	fmt.Println(f.str1.msg)
}
func (f *Facade) str2f() {
	fmt.Println(f.str2.msg)
}
