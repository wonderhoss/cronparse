package cron

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Utils", func() {
	Describe("Sequence Generator", func() {
		It("generates full sequences", func() {
			s := fullSequence(60, false)
			Expect(s).Should(Equal([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
				21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38,
				39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56,
				57, 58, 59}))
			s = fullSequence(24, false)
			Expect(s).Should(Equal([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}))
			s = fullSequence(31, true)
			Expect(s).Should(Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}))

		})
		It("generates partial sequences", func() {
			s := sequenceBetween(5, 17, 60, false)
			Expect(s).Should(Equal([]int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}))
			s = sequenceBetween(5, 17, 24, false)
			Expect(s).Should(Equal([]int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}))
			s = sequenceBetween(58, 7, 60, false)
			Expect(s).Should(Equal([]int{58, 59, 0, 1, 2, 3, 4, 5, 6, 7}))
			s = sequenceBetween(3, 6, 24, false)
			Expect(s).Should(Equal([]int{3, 4, 5, 6}))
			s = sequenceBetween(3, 6, 60, false)
			Expect(s).Should(Equal([]int{3, 4, 5, 6}))
			s = sequenceBetween(22, 3, 24, false)
			Expect(s).Should(Equal([]int{22, 23, 0, 1, 2, 3}))
			s = sequenceBetween(28, 2, 31, true)
			Expect(s).Should(Equal([]int{28, 29, 30, 31, 1, 2}))
		})
		It("generates periodic sequences", func() {
			s := periodicSequence(0, 15, 60, false)
			Expect(s).Should(Equal([]int{0, 15, 30, 45}))
			s = periodicSequence(0, 19, 60, false)
			Expect(s).Should(Equal([]int{0, 19, 38, 57}))
			s = periodicSequence(17, 18, 60, false)
			Expect(s).Should(Equal([]int{17, 35, 53}))
			s = periodicSequence(0, 4, 24, false)
			Expect(s).Should(Equal([]int{0, 4, 8, 12, 16, 20}))
			s = periodicSequence(4, 3, 24, false)
			Expect(s).Should(Equal([]int{4, 7, 10, 13, 16, 19, 22}))
			s = periodicSequence(3, 10, 31, true)
			Expect(s).Should(Equal([]int{3, 13, 23}))
		})
	})
})
