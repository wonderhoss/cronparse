package cron

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Expression Parser", func() {
	It("parses cron expressions", func() {
		for _, c := range TestCases {
			e, err := Parse(c.CronString)
			Expect(err).NotTo(HaveOccurred())
			Expect(e.DeepEquals(&c.Exp)).Should(BeTrue(), "parsed expression %+v should equal sample %+v", e, &c.Exp)
		}
	})
	It("correctly identifies errors", func() {
		for _, c := range BrokenTestCases {
			_, err := Parse(c.CronString)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(c.Err))
		}
	})
})
