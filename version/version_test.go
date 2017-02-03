package version

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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
	expected5 := 2000001.000005
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
	expected8 := 5
	if !reflect.DeepEqual(version.RevisionNumber(), expected8) {
		t.Fatalf("Expected %v, but %v:", expected8, version.RevisionNumber())
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

func TestVersion_2(t *testing.T) {
	assert := assert.New(t)

	version_1 := NewVersion("3.2.1")
	version_10 := NewVersion("3.2.10")
	version_100 := NewVersion("3.2.100")
	version_1000 := NewVersion("3.2.1000")
	version_x := NewVersion("3434.20941.141592")

	assert.Equal(3000002.000001, version_1.Number())
	assert.Equal(3000002.000010, version_10.Number())
	assert.Equal(3000002.000100, version_100.Number())
	assert.Equal(3000002.001000, version_1000.Number())
	assert.Equal(3434020941.141592, version_x.Number())

	assert.NotEqual(Same, version_1.Compare(version_10))
	assert.False(version_1.IsGreaterThan(version_10))
	assert.True(version_1.IsLessThan(version_10))

	assert.NotEqual(Same, version_10.Compare(version_100))
	assert.False(version_10.IsGreaterThan(version_100))
	assert.True(version_10.IsLessThan(version_100))

	assert.NotEqual(Same, version_100.Compare(version_1000))
	assert.False(version_100.IsGreaterThan(version_1000))
	assert.True(version_100.IsLessThan(version_1000))

	assert.NotEqual(Same, version_1000.Compare(version_1))
	assert.True(version_1000.IsGreaterThan(version_1))
	assert.False(version_1000.IsLessThan(version_1))
}

func TestVersion_3(t *testing.T) {
	var version Version
	version = NewVersion("3.5")

	expected1 := NewVersion("3.5")
	if !reflect.DeepEqual(version, expected1) {
		t.Fatalf("Expected %v, but %v:", expected1, version)
	}

	expected2 := []string{"3", "5"}
	if !reflect.DeepEqual(version.explode(), expected2) {
		t.Fatalf("Expected %v, but %v:", expected2, version.explode())
	}

	expected3 := []int{3, 5}
	for idx, val := range expected3 {
		if !reflect.DeepEqual(version.element(idx), val) {
			t.Fatalf("Expected %v, but %v:", val, version.element(idx))
		}
	}

	expected4 := "3.5"
	if !reflect.DeepEqual(version.String(), expected4) {
		t.Fatalf("Expected %v, but %v:", expected4, version.String())
	}
	expected5 := 3000005.0
	if !reflect.DeepEqual(version.Number(), expected5) {
		t.Fatalf("Expected %v, but %v:", expected5, version.Number())
	}
	expected6 := 3
	if !reflect.DeepEqual(version.MajorNumber(), expected6) {
		t.Fatalf("Expected %v, but %v:", expected6, version.MajorNumber())
	}
	expected7 := 5
	if !reflect.DeepEqual(version.MinorNumber(), expected7) {
		t.Fatalf("Expected %v, but %v:", expected7, version.MinorNumber())
	}
	expected8 := 0
	if !reflect.DeepEqual(version.RevisionNumber(), expected8) {
		t.Fatalf("Expected %v, but %v:", expected8, version.MinorNumber())
	}
}

func TestIsGreaterThanOrEqualTo(t *testing.T) {
	testVersion := NewVersion("2.1.5")
	equalVersion := NewVersion("2.1.5")
	greaterVersion := NewVersion("2.1.6")
	lessVersion := NewVersion("2.1.4")
	assert.Equal(t, true, testVersion.IsGreaterThanOrEqualTo(equalVersion))
	assert.Equal(t, true, testVersion.IsGreaterThanOrEqualTo(lessVersion))
	assert.Equal(t, false, testVersion.IsGreaterThanOrEqualTo(greaterVersion))
}

func TestIsLessThanOrEqualTo(t *testing.T) {
	testVersion := NewVersion("2.1.5")
	equalVersion := NewVersion("2.1.5")
	greaterVersion := NewVersion("2.1.6")
	lessVersion := NewVersion("2.1.4")
	assert.Equal(t, true, testVersion.IsLessThanOrEqualTo(equalVersion))
	assert.Equal(t, true, testVersion.IsLessThanOrEqualTo(greaterVersion))
	assert.Equal(t, false, testVersion.IsLessThanOrEqualTo(lessVersion))
}
