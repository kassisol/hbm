package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func RecoverFunc() {
	if r := recover(); r != nil {
		log.Println("Recovered:", r)
	}
}

func CreateSlice(input, sep string) []string {
	result := []string{}

	items := strings.Split(input, sep)

	for _, item := range items {
		result = append(result, strings.TrimSpace(item))
	}

	return result
}

func ConvertSliceToMap(sep string, slice []string) map[string]string {
	result := make(map[string]string)

	if len(slice) > 0 {
		for _, s := range slice {
			if ! strings.Contains(s, sep) {
				continue
			}

			split := strings.Split(s, sep)

			result[split[0]] = split[1]
		}
	}

	return result
}

func StringInSlice(a string, list []string, insensitive bool) bool {
	for _, v := range list {
		a1 := a
		v1 := v
		if insensitive {
			a1 = strings.ToLower(a)
			v1 = strings.ToLower(v)
		}

		if a1 == v1 {
			return true
		}
	}

	return false
}

func Exit(err error) {
	fmt.Println(err)

	os.Exit(1)
}

func RemoveLastChar(s string) string {
	strLen := len(s) - 1
	newStr := s
	if strLen > 0 {
		newStr = s[0:strLen]
	}

	return newStr
}
