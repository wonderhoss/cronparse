package cron

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Expression Parser", func() {
	It("parses cron expressions", func() {
		for _, c := range TestCases {
			e, err := Parse(c.CronString)
			fmt.Printf("%+v\n\n", e)
			fmt.Printf("%+v\n", &c.Exp)
			Expect(err).NotTo(HaveOccurred())
			Expect(e.DeepEquals(&c.Exp)).Should(BeTrue(), "parsed expression should equal sample")
		}
	})
})
