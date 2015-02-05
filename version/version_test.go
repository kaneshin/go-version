package version

import (
	"reflect"
	"testing"
)

func TestVersion_1(t *testing.T) {
	var version Version
	version = NewVersion("2.1.5")

	expected1 := NewVersion("2.1.5")
	if !reflect.DeepEqual(version, expected1) {
		t.Fatalf("Expected %v, but %v:", expected1, version)
	}

	expected2 := []string{"2", "1", "5"}
	if !reflect.DeepEqual(version.explode(), expected2) {
		t.Fatalf("Expected %v, but %v:", expected2, version.explode())
	}

	expected3 := []int{2, 1, 5}
	for idx, val := range expected3 {
		if !reflect.DeepEqual(version.element(idx), val) {
			t.Fatalf("Expected %v, but %v:", val, version.element(idx))
		}
	}

	expected4 := "2.1.5"
	if !reflect.DeepEqual(version.String(), expected4) {
		t.Fatalf("Expected %v, but %v:", expected4, version.String())
	}

	expected5 := 201.05
	if !reflect.DeepEqual(version.Number(), expected5) {
		t.Fatalf("Expected %v, but %v:", expected5, version.Number())
	}

	expected6 := 2
	if !reflect.DeepEqual(version.MajorNumber(), expected6) {
		t.Fatalf("Expected %v, but %v:", expected6, version.MajorNumber())
	}

	expected7 := 1
	if !reflect.DeepEqual(version.MinorNumber(), expected7) {
		t.Fatalf("Expected %v, but %v:", expected7, version.MinorNumber())
	}

	var otherVersion Version
	otherVersion = NewVersion("2.1.5")
	if version.Compare(otherVersion) != Same {
		t.Fatalf("Expected same")
	}
	if version.IsGreaterThan(otherVersion) {
		t.Fatalf("Expected false")
	}
	if version.IsLessThan(otherVersion) {
		t.Fatalf("Expected false")
	}

	otherVersion = NewVersion("2.1.4")
	if version.Compare(otherVersion) == Same {
		t.Fatalf("Expected not same")
	}
	if !version.IsGreaterThan(otherVersion) {
		t.Fatalf("Expected true")
	}
	if version.IsLessThan(otherVersion) {
		t.Fatalf("Expected false")
	}

	otherVersion = NewVersion("2.1.6")
	if version.Compare(otherVersion) == Same {
		t.Fatalf("Expected not same")
	}
	if version.IsGreaterThan(otherVersion) {
		t.Fatalf("Expected false")
	}
	if !version.IsLessThan(otherVersion) {
		t.Fatalf("Expected true")
	}

}
