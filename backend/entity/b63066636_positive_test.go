package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func Test_positive(t *testing.T) {
	g := NewGomegaWithT(t)

	e := Employee{
		Name:       "ASDF",
		Email:      "teppharit@gmail.com",
		EmployeeID: "M12345678",
	}
	ok, err := govalidator.ValidateStruct(e)

	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
}
