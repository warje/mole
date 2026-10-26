package mole_test

import (
	"fmt"
	"testing"

	"github.com/davrodpin/mole/mole"
)

func TestAddressInputSet(t *testing.T) {
	tests := []struct {
		user string
		host string
		port string
	}{
		{
			"mole",
			"mole-server",
			"22",
		},
	}

	var ai mole.AddressInput
	for id, test := range tests {
		ai.Set(fmt.Sprintf("%s@%s:%s", test.user, test.host, test.port))

		if test.user != ai.User {
			t.Errorf("user does not match on test %d: expected: %s, value: %s", id, test.user, ai.User)
		}

		if test.host != ai.Host {
			t.Errorf("host does not match on test %d: expected: %s, value: %s", id, test.host, ai.Host)
		}

		if test.port != ai.Port {
			t.Errorf("port does not match on test %d: expected: %s, value: %s", id, test.port, ai.Port)
		}
	}

}

func TestAddressInputListSet(t *testing.T) {

	tests := []struct {
		ai string
	}{
		{
			"mole@mole-server:22",
		},
	}

	for id, test := range tests {
		aiList := mole.AddressInputList{}
		aiList.Set(test.ai)

		if test.ai != aiList.String() {
			t.Errorf("test %d: expected: %s, value: %s", id, test.ai, aiList.String())
		}
	}
}

func TestAddressInputAddress(t *testing.T) {

	tests := []struct {
		host     string
		port     string
		expected string
	}{
		{
			"mole-server",
			"22",
			"mole-server:22",
		},
		{
			"mole-server",
			"",
			"mole-server",
		},
		{
			"::1",
			"4500",
			"[::1]:4500",
		},
	}

	for id, test := range tests {
		ai := mole.AddressInput{Host: test.host, Port: test.port}
		addr := ai.Address()

		if test.expected != addr {
			t.Errorf("address does not match on test %d: expected: %s, value: %s", id, test.expected, addr)
		}
	}
}

func TestAddressInputSetIPv6(t *testing.T) {
	tests := []struct {
		input string
		user  string
		host  string
		port  string
	}{
		{"[::1]:4500", "", "::1", "4500"},
		{"mole@[fe80::1%eth0]:22", "mole", "fe80::1%eth0", "22"},
		{"127.0.0.1:3001", "", "127.0.0.1", "3001"},
		{":3001", "", "", "3001"},
	}

	for id, test := range tests {
		var ai mole.AddressInput
		ai.Set(test.input)

		if ai.User != test.user || ai.Host != test.host || ai.Port != test.port {
			t.Errorf("test %d: expected: %q %q %q, value: %q %q %q", id, test.user, test.host, test.port, ai.User, ai.Host, ai.Port)
		}
	}
}
