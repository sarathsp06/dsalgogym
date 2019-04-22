package arrays

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
