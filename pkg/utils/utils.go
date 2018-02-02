package utils

import (
	"regexp"
	"strconv"
	"strings"
)

func GetPortRangeFromString(value string) (int, int, error) {
	var startPort int
	var endPort int

	valid, err := regexp.Match(`^[0-9]+\-[0-9]+$`, []byte(value))
	if err != nil {
		return startPort, endPort, err
	}

	if valid {
		pr := strings.Split(value, "-")

		startPort, err = strconv.Atoi(pr[0])
		if err != nil {
			return startPort, endPort, err
		}
		endPort, err = strconv.Atoi(pr[1])
		if err != nil {
			return startPort, endPort, err
		}
	} else {
		port, err := strconv.Atoi(value)
		if err != nil {
			return startPort, endPort, err
		}

		startPort = port
		endPort = port
	}

	return startPort, endPort, nil
}
