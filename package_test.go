// Gostwire (The Sequel)
//
// Written by Harald Albrecht

package deferrer

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestNetworking(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "gostwire/network/test/deferrer package")
}
