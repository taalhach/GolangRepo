
package main
import "fmt"
func main(){
	a:=make([]int, 5)
	for i,_:=range a{
		a[i]=i
	}
	fmt.Println(a)
	a=delete(a,2)
	fmt.Println(a)
}
func delete(slc[]int,b int) []int{

	for i,_:=range slc{
		if slc[i]==b{
			a:=i
			slc=append(slc[:a],slc[a+1:]...)
			break
		}
	}


	return slc
}