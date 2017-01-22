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

func StringInSlice(a string, list []string) bool {
	for _, v := range list {
		if a == v {
			return true
		}
	}

	return false
}

func Exit(err error) {
	fmt.Println(err)

	os.Exit(1)
}
