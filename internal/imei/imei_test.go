package imei

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
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
	imeiValid1,
	imeiValid2,
	imeiInvalid1,
	imeiInvalid2,
}

func TestThermomatic(t *testing.T) {
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
			It("correctly decode", func() {
				Expect(code).To(Equal(i.decoded))
			})
			It("has a nil error", func() {
				Expect(err).To(BeNil())
			})
		})
	})
}

func TestDecode(t *testing.T) {
	for _, i := range imeis {
		if code, err := Decode([]byte(i.imei)); err != nil {
			t.Run(`invalid-error `, func(t *testing.T) {
				assert.Equal(t, err, i.decodeError)
			})

			t.Run("IMEI is invalid", func(t *testing.T) {
				assert.Equal(t, i.valid, false)
			})
		} else {
			t.Run("code is invalid", func(t *testing.T) {
				fmt.Printf("%d", i.decoded)
				assert.Equal(t, code-i.decoded, uint64(0))
			})
		}
	}
}

func TestDecodeAllocations(t *testing.T) {
}

func TestDecodePanics(t *testing.T) {
}

func BenchmarkDecode(b *testing.B) {
}
