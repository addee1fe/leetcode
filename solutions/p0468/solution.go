package solution

import (
	"strconv"
	"strings"
)

func validIPAddress(IP string) string {
	if s := strings.Split(IP, "."); len(s) > 1 {
		return validateIPv4(IP)
	}
	return validateIPv6(IP)
}

func validateIPv4(IP string) string {
	t := strings.Split(IP, ".")
	if len(t) != 4 {
		return "Neither"
	}
	for _, v := range t {
		if len(v) == 0 || len(v) > 3 {
			return "Neither"
		}
		if len(v) > 1 && v[0] == '0' {
			return "Neither"
		}
		if i, err := strconv.Atoi(v); err != nil || i < 0 || i > 255 {
			return "Neither"
		}
	}
	return "IPv4"
}

func validateIPv6(IP string) string {
	t := strings.Split(IP, ":")
	if len(t) != 8 {
		return "Neither"
	}
	hex := "0123456789abcdefABCDEF"
	for _, v := range t {
		if len(v) == 0 || len(v) > 4 {
			return "Neither"
		}
		for _, r := range v {
			if strings.IndexRune(hex, r) == -1 {
				return "Neither"
			}
		}
	}
	return "IPv6"
}

// func validIPAddress(IP string) string {
// 	ip := net.ParseIP(IP)
// 	if ip == nil {
// 		return "Neither"
// 	}
// 	for _, v := range IP {
// 		switch v {
// 		case '.':
// 			return "IPv4"
// 		case ':':
// 			return "IPv6"
// 		}
// 	}
// 	return "Neither"
// }
