package asciibox_test

import (
	"testing"

	"github.com/boundedinfinity/asciibox"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLoader(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Caser Suite")
}

var _ = Describe("Smoke Test", func() {
	It("Phrase to Camel", func() {
		input := []string{
			"a b c d e f g h i j k l m n o p q r s t u v w x y z a b c d e f g h i j k l m n o p q r s t u v w x y z",
			"test message",
			"on mulitple",
			"lines",
		}
		expected := ""
		actual := ""

		asciibox.Box(input, asciibox.BoxOptions{
			Alignment: asciibox.Alignment_Left,
			BoxWidth:  15,
		})
		Expect(actual).To(Equal(expected))

		asciibox.Box(input, asciibox.BoxOptions{
			Alignment: asciibox.Alignment_Left,
		})
		Expect(actual).To(Equal(expected))

		asciibox.Box(input, asciibox.BoxOptions{
			Alignment: asciibox.Alignment_Middle,
		})
		Expect(actual).To(Equal(expected))

		asciibox.Box(input, asciibox.BoxOptions{
			Alignment: asciibox.Alignment_Right,
		})
		Expect(actual).To(Equal(expected))
	})

})
