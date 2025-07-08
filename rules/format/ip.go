package format

import (
	"errors"
	"net"
	"strings"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	IP   = "ip"
	IPV4 = "ipv4"
	IPV6 = "ipv6"
	MAC  = "mac"
	ANY  = "any" // accepts ipv4, ipv6, mac

	ipRuleName = "ip"

	ipRuleDefaultMsg         = "the :attribute must be a valid IP address"
	ipRuleDataNotProvidedMsg = "the :attribute must provide a valid IP address, but it is empty or not provided"
	ipRuleInvalidFormatMsg   = "the :attribute must be a valid :param0 address"
	ipRuleUnsupportedTypeMsg = "unsupported IP rule type: :param0"
)

// IPRule is a validation rule for IP and MAC addresses.
type IPRule struct {
	common.BaseRule
	ipType string
}

// NewIPRule creates a new IPRule instance.
func NewIPRule(parameters []string) (contract.Rule, error) {
	ipType := ANY
	if len(parameters) > 0 {
		ipType = strings.ToLower(parameters[0])
	}
	return &IPRule{
		BaseRule: common.NewBaseRule(ipRuleName, ipRuleDefaultMsg, parameters),
		ipType:   ipType,
	}, nil
}

// Validate performs the IP/MAC format validation.
func (r *IPRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	val, ok := ctx.Value().(string)
	if !ok || strings.TrimSpace(val) == "" {
		return errors.New(ipRuleDataNotProvidedMsg)
	}

	switch r.ipType {
	case IPV4:
		if !isValidIPv4(val) {
			return r.newError(ipRuleInvalidFormatMsg)
		}
	case IPV6:
		if !isValidIPv6(val) {
			return r.newError(ipRuleInvalidFormatMsg)
		}
	case MAC:
		if !isValidMAC(val) {
			return r.newError(ipRuleInvalidFormatMsg)
		}
	case IP, ANY:
		if isValidIPv4(val) || isValidIPv6(val) || isValidMAC(val) {
			return nil
		}
		return r.newError(ipRuleInvalidFormatMsg)
	default:
		return r.newError(ipRuleUnsupportedTypeMsg)
	}

	return nil
}

// isValidIPv4 checks for a valid IPv4 address.
func isValidIPv4(ip string) bool {
	parsed := net.ParseIP(ip)
	return parsed != nil && parsed.To4() != nil
}

// isValidIPv6 checks for a valid IPv6 address.
func isValidIPv6(ip string) bool {
	parsed := net.ParseIP(ip)
	return parsed != nil && parsed.To4() == nil
}

// isValidMAC checks for a valid MAC address.
func isValidMAC(addr string) bool {
	_, err := net.ParseMAC(addr)
	return err == nil
}

// newError formats error with dynamic :param0 replacement.
func (r *IPRule) newError(msg string) error {
	resolved := strings.ReplaceAll(msg, ":param0", r.ipType)
	return errors.New(resolved)
}

func (r *IPRule) Name() string {
	return ipRuleName
}
