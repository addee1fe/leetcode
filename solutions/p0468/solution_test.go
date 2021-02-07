package solution

import (
	"testing"
)

func TestValidIPAddress(t *testing.T) {
	tests := []struct {
		IP   string
		want string
	}{
		{"172.16.254.1", "IPv4"},
		{"01.01.01.01", "Neither"},
		{"172.16.254.01", "Neither"},
		{"256.256.256.256", "Neither"},
		{"2001:0db8:85a3:0:0:8A2E:0370:7334", "IPv6"},
		{"02001:0db8:85a3:0000:0000:8a2e:0370:7334", "Neither"},
	}
	for _, tt := range tests {
		if got := validIPAddress(tt.IP); got != tt.want {
			t.Errorf("validIPAddress(%v) = %v, want %v", tt.IP, got, tt.want)
		}
	}
}

func TestValidateIPv4(t *testing.T) {
	tests := []struct {
		IP   string
		want string
	}{
		{"1.1.1.1", "IPv4"},
		{"1.1.01", "Neither"},
		{"1.2..4", "Neither"},
		{"192.0.0.1", "IPv4"},
		{"172.16.254.1", "IPv4"},
		{"1.2.3.1234", "Neither"},
		{"01.01.01.01", "Neither"},
		{"172.16.254.01", "Neither"},
	}
	for _, tt := range tests {
		if got := validateIPv4(tt.IP); got != tt.want {
			t.Errorf("validateIPv6(%v) = %v, want %v", tt.IP, got, tt.want)
		}
	}
}

func TestValidateIPv6(t *testing.T) {
	tests := []struct {
		IP   string
		want string
	}{
		{"2001:0db8:85a3:0:0:8A2E:0370:7334", "IPv6"},
		{"02001:0db8:85a3:0000:0000:8a2e:0370:7334", "Neither"},
	}
	for _, tt := range tests {
		if got := validateIPv6(tt.IP); got != tt.want {
			t.Errorf("validateIPv6(%v) = %v, want %v", tt.IP, got, tt.want)
		}
	}
}
