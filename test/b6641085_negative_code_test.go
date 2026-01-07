package test

import (
	"sut-final-lab/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"
)
func TestTask3(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Code must start with BK followed by 6 digits (0-9)", func(t *testing.T) {
		books := entity.Books{
			Title: "Love",
			Price: 100,
			Code:  "BK2",
		}
		ok, err := govalidator.ValidateStruct(books)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Code must start with BK followed by 6 digits (0-9)"))
	})

}