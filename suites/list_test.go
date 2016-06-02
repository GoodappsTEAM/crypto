package suites

import (
	"gopkg.in/dedis/crypto.v0/test"
	"testing"
)

func TestSuites(t *testing.T) {
	s := All()
	for _, suite := range s {
		test.TestSuite(suite)
	}
}

func TestString(t *testing.T) {
	_, err := StringToSuite("unknown")
	if err == nil {
		t.Fatal("Shouldn't find suite 'unknown'")
	}
	s := All()
	for n := range s {
		search, _ := StringToSuite(n)
		if n != search.String() {
			t.Fatal("Suite", n, "returned", search.String())
		}
	}
}
