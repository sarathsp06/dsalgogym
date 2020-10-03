package lastn

// LastN stores last n entries
// and provides interface to add and query
type LastN struct {
	items []int
	last  int
	cap   int
}

// Record Records the val as the most recent entry
// it may delete the oldest entry if need  be
func (l *LastN) Record(val int) {
	l.last = (l.last + 1) % l.cap
	l.items[l.last] = val
}

// GetLast returns `idx`th element from last
// returns if index out of range
// counting from 0
func (l *LastN) GetLast(idx int) int {
	if l.last == -1 || idx >= l.cap {
		return 0
	}
	return l.items[(l.last-idx+l.cap)%l.cap]
}

// New creates an instance of LastN
func New(cap uint) *LastN {
	return &LastN{cap: int(cap), items: make([]int, cap)}
}
