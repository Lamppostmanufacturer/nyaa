package search

import (
	"math"
	"strconv"

	humanize "github.com/dustin/go-humanize"
)

// SizeBytes size in bytes
type SizeBytes uint64

func (sz *SizeBytes) Parse(s string, sizeType string) bool {
	if s == "" {
		*sz = 0
		return false
	}
	var multiplier uint64
	switch sizeType {
	case "b":
		multiplier = 1
	case "k":
		multiplier = uint64(math.Exp2(10))
	case "m":
		multiplier = uint64(math.Exp2(20))
	case "g":
		multiplier = uint64(math.Exp2(30))
	}
	size64, err := humanize.ParseBytes(s)
	if err != nil {
		*sz = 0
		return false
	}
	*sz = SizeBytes(size64 * multiplier)
	return true
}

// ToESQuery convert DateFilter type to ES query format
func (sz SizeBytes) ToESQuery() string {
	if sz > 0 {
		return strconv.FormatUint(uint64(sz), 10)
	}
	return "*"
}

// ToDBQuery convert DateFilter type to ES query format
func (sz SizeBytes) ToDBQuery() string {
	if sz > 0 {
		return strconv.FormatUint(uint64(sz), 10)
	}
	return ""
}
