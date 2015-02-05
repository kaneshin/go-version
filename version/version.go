package version

import (
	"strconv"
	"strings"
)

const (
	Ascending  = -1
	Same       = 0
	Descending = 1
)

type Version string

func NewVersion(ver string) Version {
	return (Version)(ver)
}

func (v Version) explode() []string {
	return strings.Split(string(v), ".")
}

func (v Version) element(idx int) int {
	if elms := v.explode(); len(elms) > idx {
		num, err := strconv.Atoi(elms[idx])
		if err != nil {
			panic(err)
		}
		return num
	}
	return 0
}

func (v Version) String() string {
	return string(v)
}

func (v Version) Number() float64 {
	dst := 0.0
	fnd := 100.0
	mul := fnd
	for _, v := range v.explode() {
		val, err := strconv.ParseFloat(v, 0)
		if err != nil {
			val = 0.0
		}
		dst += mul * val
		mul /= fnd
	}
	return dst
}

func (v Version) MajorNumber() int {
	return v.element(0)
}

func (v Version) MinorNumber() int {
	return v.element(1)
}

func (v Version) Compare(o Version) int {
	vn := v.Number()
	on := o.Number()
	if vn < on {
		return Ascending
	} else if vn > on {
		return Descending
	}
	return Same
}

func (v Version) IsGreaterThan(o Version) bool {
	return v.Compare(o) == Descending
}

func (v Version) IsLessThan(o Version) bool {
	return v.Compare(o) == Ascending
}
