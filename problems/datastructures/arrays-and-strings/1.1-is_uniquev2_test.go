package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsUnique2(t *testing.T) {
	testFns := []func(string) bool{isUniqueBrute, isUniqueSorted, isUniqueSimple, isUniqueOptimized}
	for _, fn := range testFns {
		assert.Equal(t, true, fn("heo"), "CASE heo")
		assert.Equal(t, false, fn("hello"), "CASE hello")
		assert.Equal(t, true, fn(""), "CASE empty")
		assert.Equal(t, false, fn("abcA"), "CASE abcA")
		assert.Equal(t, true, fn("a"), "CASE a")
		loremIpsum := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris sem diam, iaculis et consectetur cursus, vehicula id eros. Sed feugiat in risus at fringilla. Vestibulum ipsum dui, suscipit sit amet risus sit amet, rutrum posuere leo. Sed sodales rutrum scelerisque. Fusce viverra id tortor vel pharetra. Morbi aliquam venenatis odio. Nulla maximus imperdiet pulvinar. Praesent ornare massa eget elit condimentum sollicitudin. Suspendisse sed pellentesque nunc. Vivamus vehicula venenatis nisi, nec fringilla ipsum consectetur quis. "
		assert.Equal(t, false, fn(loremIpsum), "CASE lorem ipsum")
	}
}
