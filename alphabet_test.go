package alphabet

import (
	"testing"
	"testing/quick"
)

const checkMark = "\u2714"
const ballotX = "\u2718"

// reverse a string
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func TestMakeNewAlphabetFromEmptyString(t *testing.T) {
	s := MakeNewAlphabet("")

	switch len(s) {
	case 0:
		t.Logf("Should get an empty string: Got it. %v", checkMark)

	default:
		t.Errorf("Expected an empty slice in output, but got '%v'. %v", s, ballotX)
	}

}

func TestMakeNewAlphabetFromString_01(t *testing.T) {
	ins := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	s := MakeNewAlphabet(ins)

	switch len(s) {

	case 0:
		t.Errorf("Expected a non-empty slice in output, but got an empty one '%v'. %v", s, ballotX)

	default:
		if len(s) == len(ins) {
			t.Logf("Got a slice '%v' from input '%s' and their lengths are the same: OK. %v", s, ins, checkMark)
		} else {
			t.Errorf("Got a slice '%v' from input '%s' but their lengths are not the same: NOT OK. %v", s, ins, ballotX)
		}
	}
}

func TestMakeNewAlphabetFromString_02(t *testing.T) {
	ins := reverse("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	s := MakeNewAlphabet(ins)

	switch len(s) {

	case 0:
		t.Errorf("Expected a non-empty slice in output, but got an empty one '%v'. %v", s, ballotX)

	default:
		if len(s) == len(ins) {
			t.Logf("Got a slice '%v' from input '%s' and their lengths are the same: OK. %v", s, ins, checkMark)
		} else {
			t.Errorf("Got a slice '%v' from input '%s' but their lengths are not the same: NOT OK. %v", s, ins, ballotX)
		}
	}
}

// generative test using the package testing/quick
func TestMakeNewAlphabetGenerative(t *testing.T) {
	attempts := 1000
	fn := func(s string) bool {
		a := MakeNewAlphabet(s)
		return len(s) == len(a)
	}

	if err := quick.Check(fn, &quick.Config{MaxCount: attempts}); err != nil {
		t.Error(err, ballotX)
	} else {
		t.Logf("Generative test OK. %v", checkMark)
	}
}
