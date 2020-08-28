package imei

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type ImeiExpectation struct {
	imei        string
	decoded     uint64
	valid       bool
	decodeError error
}

var imeiValid1 = ImeiExpectation{
	imei:        "123456789012344",
	valid:       true,
	decodeError: nil,
	decoded:     123456789012344,
}

var imeiValid2 = ImeiExpectation{
	imei:        "000000000000000",
	valid:       true,
	decodeError: nil,
	decoded:     0,
}

var imeiInvalid1 = ImeiExpectation{
	imei:        "01234567890XYZA",
	valid:       false,
	decodeError: ErrInvalid,
	decoded:     0,
}

var imeiInvalid2 = ImeiExpectation{
	imei:        "01234",
	valid:       false,
	decodeError: ErrInvalid,
	decoded:     0,
}

var imeis = []ImeiExpectation{
	imeiValid2,
	imeiValid1,
	imeiInvalid1,
	imeiInvalid2,
}

func TestDecode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IMEI")
}

var _ = Describe("IMEI", func() {
	for _, i := range imeis {
		if i.valid {
			validateIMEI(i)
		} else {
			invalidateIMEI(i)
		}

	}
})

func invalidateIMEI(i ImeiExpectation) {
	var err error
	BeforeEach(func() {
		_, err = Decode([]byte(i.imei))
	})

	Describe("Decode()", func() {
		Context("With an invalid IMEI", func() {
			It("returns appropriate error", func() {
				Expect(err).To(Equal(i.decodeError))
			})
		})
	})
}

func validateIMEI(i ImeiExpectation) {
	var code uint64
	var err error
	BeforeEach(func() {
		code, err = Decode([]byte(i.imei))
	})

	Describe("Decode()", func() {
		Context("With a valid IMEI", func() {
			It("correctly decodes", func() {
				Expect(code).To(Equal(i.decoded))
			})
			It("returns nil for error", func() {
				Expect(err).To(BeNil())
			})
		})
	})
}

func TestDecodeAllocations(t *testing.T) {
}

func TestDecodePanics(t *testing.T) {
}

func BenchmarkDecode(b *testing.B) {
}
