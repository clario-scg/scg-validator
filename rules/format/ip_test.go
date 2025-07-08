package format_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/format"
)

type ipTestCase struct {
	name       string
	ipType     string
	value      any
	shouldPass bool
}

func TestIPRule(t *testing.T) {
	testCases := []ipTestCase{
		// --- IP: Accepts both IPv4 & IPv6 & MAC
		{"IP - valid IPv4", format.IP, "192.168.1.1", true},
		{"IP - valid IPv6", format.IP, "2001:db8::ff00:42:8329", true},
		{"IP - valid MAC", format.IP, "01:23:45:67:89:ab", true},
		{"IP - invalid string", format.IP, "not-an-ip", false},

		// --- IPv4
		{"IPv4 - valid", format.IPV4, "10.0.0.1", true},
		{"IPv4 - invalid (IPv6)", format.IPV4, "2001:db8::1", false},
		{"IPv4 - invalid MAC", format.IPV4, "01:23:45:67:89:ab", false},

		// --- IPv6
		{"IPv6 - valid", format.IPV6, "2001:db8::ff00:42:8329", true},
		{"IPv6 - invalid (IPv4)", format.IPV6, "127.0.0.1", false},
		{"IPv6 - invalid MAC", format.IPV6, "01-23-45-67-89-ab", false},

		// --- MAC
		{"MAC - valid colon", format.MAC, "01:23:45:67:89:ab", true},
		{"MAC - valid dash", format.MAC, "01-23-45-67-89-ab", true},
		{"MAC - invalid", format.MAC, "not-a-mac", false},
		{"MAC - invalid IPv4", format.MAC, "192.168.1.1", false},

		// --- Non-string input
		{"Non-string - integer", format.IP, 12345, false},
		{"Non-string - nil", format.IP, nil, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			rule, err := format.NewIPRule([]string{tc.ipType})
			if err != nil {
				t.Fatalf("failed to create IPRule: %v", err)
			}

			ctx := contract.NewValidationContext("ip_field", tc.value, nil, nil)
			err = rule.Validate(ctx)

			if tc.shouldPass && err != nil {
				t.Errorf("expected validation to pass, but got error: %v", err)
			}
			if !tc.shouldPass && err == nil {
				t.Errorf("expected validation to fail for value: %#v", tc.value)
			}
		})
	}
}
