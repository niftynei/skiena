package four

import (
	"sort"
)

// a lot of this was originally 
// copied from https://play.golang.org/p/N6GbEgBffd

// alias for string slice
type ByLength []string

// implement sort.Interface
func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func Sort(list []string) {
	sort.Sort(ByLength(list))
}
