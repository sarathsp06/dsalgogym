package arrays

import "strings"

func URLify(str string) string {
	space := []byte(`%20`)
	s := []byte(str)
	spaceCount := 0
	//TODO: can optimise by storing the indices
	for _, b := range s {
		if b == ' ' {
			spaceCount++
		}
	}
	s = append(s, make([]byte, spaceCount*2)...)
	for i, j := len(str)-1, len(s)-1; i >= 0; i, j = i-1, j-1 {
		if str[i] != ' ' {
			s[j] = str[i]
			continue
		}
		j = j - 2
		copy(s[j:j+4], space)
	}
	return string(s)
}

// URLifyBuiltin has time complexity O(n)...?
func URLifyBuiltin(input string) string {
	// Does it make sense to implement a solution conforming to constraints
	// that are not imposed by the current language you use?
	// Should I artifically limit the slice and resort to manual rune shifting
	// just to solve the problem in its intended form? ¯\_(ツ)_/¯
	return strings.ReplaceAll(input, " ", "%20")
}
