package publicdefine

import (
	"regexp"
	"strconv"
)

func IsNumber(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}
	reg := regexp.MustCompile(`[\d]+`)
	if reg.FindString(s) == s {
		ret, err := strconv.Atoi(s)
		if err == nil {
			return ret, true
		}
	}
	return 0, false
}

func IsUINT64(s string) (uint64, bool) {
	if len(s) == 0 {
		return 0, false
	}
	reg := regexp.MustCompile(`[\d]+`)
	if reg.FindString(s) == s {
		ret, err := strconv.ParseUint(s, 10, 64)
		if err == nil {
			return ret, true
		}
	}
	return 0, false
}

func IsFLOAT64(s string) (float64, bool) {
	if len(s) == 0 {
		return 0, false
	}
	ret, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return ret, true
	}
	return 0, false
}
