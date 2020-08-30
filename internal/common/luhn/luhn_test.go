package luhn

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var validNumbers = []int{
	000000000000000,
	79927398713,
	4929972884676289,
	4532733309529845,
	4539088167512356,
	5577189519503182,
	5499078785968242,
	5236582963742210,
	379537021417898,
	373494930335082,
	379203612454689,
	6011223604226714,
	6011625707082028,
	6011964086036747,
	3544936439662067,
	3533841638640315,
	3536137811022331,
	5460262971178544,
	5493663189154782,
	5469544329973911,
	30154976989581,
	30117694974441,
	30561074627774,
	36687623048586,
	36698326962197,
	36387123823311,
	5038727783656872,
	6763018270225762,
	6763677229110829,
	6706203675790103,
	6709504120728607,
	6771712353138831,
	4913488277530387,
	4913768884688532,
	4844723110962866,
	6386556471849523,
	6387065788050980,
	6388464094939979,
}

func TestLuhnNumbers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LUHN")
}

var _ = Describe("LUHN", func() {
	Describe("Valid Luhn Numbers()", func() {
		for _, number := range validNumbers {
			Describe(fmt.Sprintf("Given the number %d", number), func() {
				It("should be a valid Luhn Number", func() {
					Expect(Valid(number)).To(BeTrue())
				})

				It(fmt.Sprintf("%v's check number should be %v",
					number,
					number%10),
					func() {

						Expect(Number(number / 10)).To(Equal(number % 10))
					})

			})
		}
	})
})
