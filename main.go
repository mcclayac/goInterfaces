package main

import (
	"fmt"
	"math/rand"
)

type shuffler interface {
	Len() int
	Swap(i, j int)
}

func shuffle(s shuffler) {

	for i:= 0; i < s.Len() ; i++ {
		j := rand.Intn(s.Len()-i)
		s.Swap(i, j)
	}

}

type intSlice []int

func (is intSlice) Len() int  {
	return len(is)
}

func (is intSlice) Swap(i,j int)  {
/*	Swap code old style
	temp := is[i]
	is[i] = is[j]
	is[j] = temp*/

	//concurrent swapping in go
	is[i] , is[j] = is[j], is[i]
}

type stringSlice []string

func (ss stringSlice) Length() int  {
	return len(ss)
}

func (ss stringSlice) Len() int  {
	return len(ss)
}


func (ss stringSlice) Swap(i,j int)  {
	/*	Swap code old style
		temp := is[i]
		is[i] = is[j]
		is[j] = temp*/

	//concurrent swapping in go
	ss[i] , ss[j] = ss[j], ss[i]
}

func main() {

	is := intSlice{1,2,3,4,5,6}
	shuffle(is)
	fmt.Printf("%q\n\n", is)
	fmt.Printf("%v\n\n", is)


	fmt.Printf("=======================\n\n")
	fmt.Printf("String Slices \n")


	ss := stringSlice{"1","2","3","4","5","6"}
	ss2 := stringSlice{"The", "quick","brown","fox"}
	shuffle(ss)
	shuffle(ss2)
	fmt.Printf("%s\n\n", ss)
	fmt.Printf("%s\n\n", ss2)




}
