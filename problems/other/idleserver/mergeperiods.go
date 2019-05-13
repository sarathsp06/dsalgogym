package idleserver

import (
	"sort"
)

type period struct {
	Start int
	End   int
}

// upsert tries to insert a period to an slice of periods
// the function assumes the inputs as non-overlapping periods and a new priod in asceding order of time
// if the new period overlaps with the last period in the array it will be merged with the last
// if the new period does not overlap it is added as another period in the list of non-overlapping period
func upsert(merged []period, item period) []period {
	if len(merged) == 0 {
		return append(merged, item)
	}
	last := merged[len(merged)-1]
	if last.Start <= item.Start && last.End >= item.Start {
		if last.End < item.End {
			merged[len(merged)-1].End = item.End
		}
		return merged
	}

	merged = append(merged, item)
	return merged
}

// mergePeriods merges two non-overlapping period slices
// it merges them to  to one non-ovelapping perodm slice
func mergePeriods(p1, p2 []period) []period {
	sort.Slice(p1, func(i, j int) bool { return p1[i].Start < p1[j].Start })
	sort.Slice(p2, func(i, j int) bool { return p2[i].Start < p2[j].Start })
	var merged []period

	var i, j int
	for ; i < len(p1); i++ {
		for ; j < len(p2); j++ {
			if p1[i].Start < p2[j].Start {
				merged = upsert(merged, p1[i])
				break
			}
			merged = upsert(merged, p2[j])
		}
	}
	for ; j < len(p2); j++ {
		merged = upsert(merged, p2[j])
	}

	return merged
}
