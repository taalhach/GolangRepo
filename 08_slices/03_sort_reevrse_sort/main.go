package main
import (
	"sort"
	"fmt"
)


func main() {
	myInts := []int{ 8, 6, 7, 5, 3, 0, 9 }
	sort.Ints(myInts)
	fmt.Printf("Ints %v reversed: %v\n", myInts, reverseInts(myInts))
}

func reverseInts(input []int) []int {
	sort.Reverse(sort.IntSlice(input))

	if len(input) == 0 {
		return input
	}

	return append(reverseInts(input[1:]), input[0])
	return input
}

