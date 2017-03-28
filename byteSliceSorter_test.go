package alphabet

import (
	"testing"
)

func TestByteSliceSort_01(t *testing.T) {
	er := "!$,.0123456789?abcdefghijklmnopqrstuvwxyz"
	w1 := "zyxwvutsrqponmlkjihgfedcba9876543210.?$!,"

	w2 := SortString(w1)
	t.Log("Unsorted:", w1)
	t.Log("Sorted:  ", w2)

	if w2 == er {
		t.Log("Sorted OK")
	} else {
		t.Error("Error sorting")
	}
}

func TestByteSliceSort_02(t *testing.T) {
	er := "!$,.0123456789?abcdefghijklmnopqrstuvwxyz"
	w1 := "zyxwvutsrqponmlkjihgfedcba9876543210.?$!,"

	t.Log("Unsorted:", w1)
	w1 = SortString(w1)
	t.Log("Sorted:  ", w1)

	if w1 == er {
		t.Log("Sorted OK")
	} else {
		t.Error("Error sorting")
	}
}
