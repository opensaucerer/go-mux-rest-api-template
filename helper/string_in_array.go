package helper

import "strings"

func StringInArray(arr []string, val string, special bool) bool {
	if special {
		for _, a := range arr {
			if !strings.Contains(val, ".zip") && strings.Contains(val, a) {
				return true
			}
		}
		return false
	}
	for _, a := range arr {
		if a == val {
			return true
		}
	}
	return false
}
