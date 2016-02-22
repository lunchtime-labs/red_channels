package service_test

import (
	"testing"
	"github.com/sclevine/agouti"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRedchannels(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Redchannels Service Test Suite")
}

var agoutiDriver *agouti.WebDriver

var _ = BeforeSuite(func() {
	// Choose a WebDriver:
	agoutiDriver = agouti.PhantomJS()
	// agoutiDriver = agouti.Selenium()
	// agoutiDriver = agouti.ChromeDriver()

	Expect(agoutiDriver.Start()).To(Succeed())
})

var _ = AfterSuite(func() {
	Expect(agoutiDriver.Stop()).To(Succeed())
})
