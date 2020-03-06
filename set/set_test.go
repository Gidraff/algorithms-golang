package set

import (
	"fmt"
	"testing"
)

func populateSet(count int, start int) *ItemSet {
	set := ItemSet{}
	for i := start; i < (start + count); i++ {
		set.Add(fmt.Sprintf("item%d", i))
	}
	return &set
}

func TestAdd(t *testing.T) {
	set := populateSet(3, 0)
	if size := set.Size(); size != 3 {
		t.Errorf("Wrong count, expected 3 and got %d", size)
	}

	set.Add("item1") // should not add it, already there
	if size := set.Size(); size != 3 {
		t.Errorf("Wrong count, expected 3 and got %d", size)
	}
	set.Add("item4") // should not add it, already there
	if size := set.Size(); size != 4 {
		t.Errorf("Wrong count, expected 4 and got %d", size)
	}
}
