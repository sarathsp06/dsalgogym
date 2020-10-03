package minsteps

import "testing"

func TestMinSteps(t *testing.T) {
	type Scenario struct {
		input                  [][]bool
		srcX, srcY, dstX, dstY int
		expectedOutput         int
	}
	table := map[string]Scenario{
		"route exists": Scenario{
			input: [][]bool{
				{false, false, false, false},
				{true, true, false, true},
				{false, false, false, false},
				{false, false, false, false},
			},
			srcX: 3, srcY: 0,
			dstX: 0, dstY: 0,
			expectedOutput: 7,
		},
		"no route": Scenario{
			input: [][]bool{
				{false, false, false, false},
				{true, true, false, true},
				{false, true, false, false},
				{false, false, true, false},
			},
			srcX: 3, srcY: 0,
			dstX: 0, dstY: 0,
			expectedOutput: 0,
		},
	}

	for name, s := range table {
		if m := MinSteps(s.input,
			s.srcX, s.srcY,
			s.dstX, s.dstY,
		); m != s.expectedOutput {
			t.Errorf("\n\t%s: expected:%d got:%d", name, s.expectedOutput, m)
		}
	}
}
