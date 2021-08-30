package deferrer

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("outer Deferrer", func() {

	It("defers in an outer function", func() {
		d := Deferrer{}
		resourceRemoved := false

		outer := func() {
			defer d.Cleanup()

			By("Defer'ing resource cleanup", func() {
				d.Defer(func() {
					resourceRemoved = true
				})
				Expect(d).To(HaveLen(1))
				Expect(resourceRemoved).To(BeFalse())
			})

			Expect(d).To(HaveLen(1))
			Expect(resourceRemoved).To(BeFalse())
		}
		outer()

		Expect(d).To(HaveLen(1))
		Expect(resourceRemoved).To(BeTrue())
	})

})
