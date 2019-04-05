package arrays

import "fmt"

func ExampleUnique() {
	fmt.Println(Unique("sarath"))
	fmt.Println(Unique("abcdefg"))
	fmt.Println(Unique("ABCDabcd"))
	fmt.Println(Unique(""))
	fmt.Println(Unique("SHARAgfhsjakal"))
	// Output:
	//false
	//true
	//true
	//true
	//false
}

func ExampleUniqueNoExtra() {
	Unique := UniqueNoExtra
	fmt.Println(Unique("sarath"))
	fmt.Println(Unique("abcdefg"))
	fmt.Println(Unique("ABCDabcd"))
	fmt.Println(Unique(""))
	fmt.Println(Unique("SHARAgfhsjakal"))
	// Output:
	//false
	//true
	//true
	//true
	//false
}
