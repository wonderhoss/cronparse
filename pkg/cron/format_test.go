package cron

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Formatter Suite", func() {
	It("generates expected string representation", func() {
		s := ExampleExpression1.AsString()

		Expect(s).Should(Equal(ExampleString1))
	})

})
