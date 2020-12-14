package cron

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Expression Formatter", func() {
	It("generates expected string representation", func() {
		for _, c := range TestCases {
			k := c.Exp
			Expect(k.AsString()).Should(Equal(c.OutString))
		}
	})
})
