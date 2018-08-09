package types

import "fmt"

type Str struct {
	Name string
	Agee int
}
func (s *Str)Print(){
	fmt.Println(s.Name)
}
