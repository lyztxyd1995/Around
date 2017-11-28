package learn_and_test

import "fmt"

func printmap(m map[string]int){
for key,value:=range m{
	fmt.Println("key: ",key , "Value: ", value)
}
}
func printarray(array []int ){
	for temp :=range array{
		fmt.Println(temp)
	}
}
func slice() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	m:=make(map[string]int)
	m["route"]=66
	m["userid"]=1111
	m["itemid"]=1234
	printmap(m)
}
