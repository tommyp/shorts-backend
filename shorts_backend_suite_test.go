package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestShortsBackend(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ShortsBackend Suite")
}
