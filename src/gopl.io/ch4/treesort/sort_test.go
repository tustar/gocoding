package treesort_test

import (
	"testing"
	"math/rand"
	"gopl.io/ch4/treesort"
	"sort"
)

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i:= range data {
		data[i] = rand.Int() % 50
	}
	treesort.Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}
