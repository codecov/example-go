package awesome

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

var _ = Describe("Awesome", func() {
    Describe("Making an expression", func() {
        Context("with a smile", func() {
            It("should result :)", func() {
                Setup()
                Expect(GetResult()).To(Equal(":)"))
            })
        })
    })
})
