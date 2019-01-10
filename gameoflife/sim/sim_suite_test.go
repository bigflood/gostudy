package sim_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSim(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sim Suite")
}
