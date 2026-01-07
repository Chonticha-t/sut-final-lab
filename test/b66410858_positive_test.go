package test

import (
	"sut-final-lab/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"
)

func TestTask1(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Title is not length 3 between 100", func(t *testing.T) {
		books := entity.Books{
			Title: "Love",
			Price: 100,
			Code:  "BK123456",
		}
		ok, err := govalidator.ValidateStruct(books)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})

}
