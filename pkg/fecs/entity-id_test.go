package fecs

import (
	"fmt"
	"testing"
)

func TestEntityID_String(t *testing.T) {
	tests := []struct {
		expect string
		id     EntityID
	}{
		{"10000000000000001", newEntityID(1, 1)},
		{"a00000000000000ff", newEntityID(255, 10)},
		{"ff000000000000000a", newEntityID(10, 255)},
	}

	for i := range tests {
		if tests[i].expect != tests[i].id.String() {
			fmt.Print("x")
			t.Error("expected", tests[i].expect, "got", tests[i].id.String())
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func TestEntityID_Equals(t *testing.T) {
	tests := []struct {
		expect bool
		id1    EntityID
		id2    EntityID
	}{
		{true, newEntityID(1, 1), newEntityID(1, 1)},
		{false, newEntityID(1, 2), newEntityID(1, 1)},
		{false, newEntityID(2, 1), newEntityID(1, 1)},
	}

	for i := range tests {
		if tests[i].expect != tests[i].id1.Equals(&tests[i].id2) {
			fmt.Print("x")
			t.Error("expected", tests[i].expect, "got", tests[i].id1.Equals(&tests[i].id2))
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}
