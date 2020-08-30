package imei

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDecodeSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IMEI")
}

var _ = Describe("IMEI", func() {
	for _, i := range fixtures {
		validateIMEI(i)
	}
})

func validateIMEI(i fixture) {
	code, err := Decode([]byte(i.imei))
	Describe(fmt.Sprintf("When decoding %s", i.imei), func() {
		if i.valid {
			checkValidIMEIExpectation(i, code, err)
		} else {
			checkInvalidIMEIExpectation(i, err)
		}
	})
}

func checkInvalidIMEIExpectation(i fixture, err error) bool {
	return Context("With an invalid IMEI", func() {
		It("returns a non-nil error", func() {
			Expect(err).NotTo(BeNil())
		})
		It("returns appropriate error", func() {
			Expect(err).To(Equal(i.decodeError))
		})
	})
}

func checkValidIMEIExpectation(i fixture, code uint64, err error) bool {
	return Context("With a valid IMEI", func() {
		It("returns nil for error", func() {
			Expect(err).To(BeNil())
		})
		It("correctly decodes", func() {
			Expect(code).To(Equal(i.decoded))
		})
	})
}

