package cron

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCron(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cron Test Suite")
}
