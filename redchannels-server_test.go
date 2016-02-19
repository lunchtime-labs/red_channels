package redchannels_test

import (
  "github.com/sclevine/agouti"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Redchannels", func() {
	var page *agouti.Page

	BeforeEach(func() {
		var err error
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})

  Describe("basic server", func() {
    It("should respond successfully to root requests", func() {
      By("fetching the root url", func() {
        Expect(page.Navigate("http://localhost:5000")).To(Succeed())
      })
    })
  })
})
