package validation

import (
	"fmt"
	"net"
	"regexp"
	"strings"
)

var (
	NotValidEmail     = fmt.Errorf("email is not valid")
	NotValidFQDN      = fmt.Errorf("fqdn is not valid")
	NotValidGroupname = fmt.Errorf("groupname is not valid")
	NotValidHostname  = fmt.Errorf("hostname is not valid")
	NotValidIPAddr    = fmt.Errorf("ip address is not valid")
	NotValidName      = fmt.Errorf("name is not valid")
	NotValidPort      = fmt.Errorf("port is not valid")
	NotValidUppercase = fmt.Errorf("character is not uppercase")
	NotValidUsername  = fmt.Errorf("username is not valid")
)

func IsValidEmail(name string) error {
	reName := regexp.MustCompile(`^[a-zA-Z0-9\-\_\.]+$`)

	if strings.Count(name, "@") != 1 {
		return NotValidEmail
	}

	parts := strings.SplitN(name, "@", 2)

	if !reName.MatchString(parts[0]) {
		return NotValidEmail
	}

	if err := IsValidFQDN(parts[1]); err != nil {
		return NotValidEmail
	}

	return nil
}

func IsValidFQDN(s string) error {
	if len(s) == 0 || len(s) > 254 {
		return NotValidFQDN
	}

	parts := strings.Split(s, ".")

	for i, p := range parts {
		rePart := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9\-]{1,63}$`)

		if i == len(parts)-1 {
			rePart = regexp.MustCompile(`^[a-zA-Z]{2,63}$`)
		}

		if !rePart.MatchString(p) {
			return NotValidFQDN
		}
	}

	return nil
}

func IsValidGroupname(name string) error {
	reName := regexp.MustCompile(`^[a-zA-Z0-9\-\_]+$`)

	if !reName.MatchString(name) {
		return NotValidGroupname
	}

	return nil
}

func IsValidHostname(name string) error {
	reName := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9\-]{1,63}$`)

	if !reName.MatchString(name) {
		return NotValidHostname
	}

	return nil
}

func IsValidIP(ip string) error {
	if net.ParseIP(ip) == nil {
		return NotValidIPAddr
	}

	return nil
}

func IsValidName(name string) error {
	reName := regexp.MustCompile(`^[a-zA-Z0-9\-\_]+$`)

	if !reName.MatchString(name) {
		return NotValidName
	}

	return nil
}

func IsValidPort(port int) error {
	if port < 0 && port > 65535 {
		return NotValidPort
	}

	return nil
}

func IsValidUsername(name string) error {
	reName := regexp.MustCompile(`^[a-zA-Z0-9\-\_\.]+$`)

	if !reName.MatchString(name) {
		return NotValidUsername
	}

	return nil
}

func IsUpper(char string) error {
	reUpper := regexp.MustCompile(`^[A-Z]$`)

	if !reUpper.MatchString(char) {
		return NotValidUppercase
	}

	return nil
}
