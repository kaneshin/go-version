package main

import (
	"reflect"
	"testing"
)

func TestVersion_1(t *testing.T) {
	var version Version
	version = "2.1.5"

	expected1 := (Version)("2.1.5")
	if !reflect.DeepEqual(version, expected1) {
		t.Fatalf("Expected %v, but %v:", expected1, version)
	}

	expected2 := []string{"2", "1", "5"}
	if !reflect.DeepEqual(version.explode(), expected2) {
		t.Fatalf("Expected %v, but %v:", expected2, version)
	}

	expected3 := []int{2, 1, 5}
	for idx, val := range expected3 {
		if !reflect.DeepEqual(version.element(idx), val) {
			t.Fatalf("Expected %v, but %v:", val, version)
		}
	}

}
