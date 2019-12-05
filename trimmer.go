package moco

import (
	"regexp"
	"strings"
)

func WhitespaceSplit(s string) []string {
	r := regexp.MustCompile("[^\\s]+")

	return r.FindAllString(s, -1)
}

func BlacklistTrim(ar []string, blacklist []string) []string {

	var trimmed []string
	var found = false

	for _, v := range ar {
		found = false
		for _, v2 := range blacklist {
			if strings.ToLower(v) == strings.ToLower(v2) {
				found = true
				break
			}
		}
		if !found {
			trimmed = append(trimmed, v)
		}
	}

	return trimmed
}

func NumericTrim(s string) (string, error) {
	s = strings.Replace(s, ",00", "", -1)
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		return "", err
	}
	return reg.ReplaceAllString(s, ""), nil
}
