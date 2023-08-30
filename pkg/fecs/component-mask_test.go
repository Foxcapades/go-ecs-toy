package fecs

import (
	"fmt"
"testing"
)

func TestComponentMask(t *testing.T) {
	var mask ComponentMask

	mask.Add(1)
	if !mask.Has(1) {
		fmt.Print("x")
		t.Error("expected mask to have 1")
	} else {
		fmt.Print(".")
	}

	mask.Remove(1)
	if mask.Has(1) {
		fmt.Print("x")
		t.Error("expected mask to not have 1")
	} else {
		fmt.Print(".")
	}

	mask.Add(64)
	if !mask.Has(64) {
		fmt.Print("x")
		t.Error("expected mask to have 64")
	} else {
		fmt.Print(".")
	}

	mask.Remove(64)
	if mask.Has(64) {
		fmt.Print("x")
		t.Error("expected mask to not have 64")
	} else {
		fmt.Print(".")
	}

	mask.Add(65)
	if !mask.Has(65) {
		fmt.Print("x")
		t.Error("expected mask to have 64")
	} else {
		fmt.Print(".")
	}

	mask.Remove(65)
	if mask.Has(65) {
		fmt.Print("x")
		t.Error("expected mask to not have 65")
	} else {
		fmt.Print(".")
	}

	fmt.Println()
}
