package service_test

import (
  . "github.com/lunchtime-labs/redchannels/Godeps/_workspace/src/github.com/onsi/ginkgo"
  . "github.com/lunchtime-labs/redchannels/Godeps/_workspace/src/github.com/onsi/gomega"
  "github.com/lunchtime-labs/redchannels/Godeps/_workspace/src/github.com/sclevine/agouti"
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

  Describe("Blacklist Service", func() {
    It("should respond successfully to requests", func() {
      By("fetching the root url", func() {
        Expect(page.Navigate("http://localhost:5000")).To(Succeed())
      })
    })
  })
})
