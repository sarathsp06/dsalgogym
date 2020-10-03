package arrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestIsUnique(t *testing.T) {
	testCases := []struct {
		input  string
		output bool
	}{
		{"", true},
		{"a", true},
		{"abcA", false},
		{"heo", true},
		{"hello", false},
		{"Lorem ipsum dolor sit amet.", false},
	}
	testFuncs := []func(string) bool{isUniqueBrute, isUniqueSorted, isUniqueSimple, isUniqueOptimized}
	for _, fn := range testFuncs {
		for _, tc := range testCases {
			t.Run(fmt.Sprintf(tc.input), func(t *testing.T) {
				assert.Equal(t, tc.output, fn(tc.input))
			})
		}
	}
}
