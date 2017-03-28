package alphabet

import (
	"sort"
)

// An alphabet is a slice of bytes (single characters)
type Alphabet []byte

type Alphabeter interface {
	MakeAlphabet()
}

func (a *Alphabet) MakeAlphabet(alphabetAsAString string) {
	*a = Alphabet(SortString(alphabetAsAString))
	return
}

func MakeNewAlphabet(alphabetAsAString string) Alphabet {
	var a Alphabet
	a.MakeAlphabet(alphabetAsAString)

	// An alternative way to create a new alphabet
	// a := Alphabet("")
	// a.MakeAlphabet(alphabetAsAString)

	return a
}

///////////////////////////////////////////////////////////////////////////////
// Byte slice sorter, string sorter

type sortBytes []byte

func (s sortBytes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortBytes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortBytes) Len() int {
	return len(s)
}

type sortedstring string

type stringsorter interface {
	sortString()
}

func (s *sortedstring) sortString() {
	sb := sortBytes(*s)
	sort.Sort(sb)
	*s = sortedstring(sb)
	return
}

func SortString(s string) string {
	ss := sortedstring(s)
	ss.sortString()
	return string(ss)
}
