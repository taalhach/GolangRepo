package int64arr

import "fmt"

type Int64arr []int64


func (a Int64arr) Len() int           { return len(a) }
func (a Int64arr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Int64arr) Less(i, j int) bool { return a[i] < a[j] }
func (a Int64arr) String() (s string) {
	sep := "" // for printing separating commas
	for _, el := range a {
		s += sep
		sep = ", "
		s += fmt.Sprintf("%d", el)
	}
	return
}

