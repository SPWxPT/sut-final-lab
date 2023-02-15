package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNameNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)
	customer := Customer{
		Name:       "",
		Email:      "test@gmail.com",
		CustomerID: "L0123456", // รหัสลูกค้าขึ8นต้นด้วย L หรือ M หรือ H แล้วตามด้วยตัวเลขจํานวน 7 ตัว
	}
	ok, err := govalidator.ValidateStruct(customer)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Name cannot be balnk"))
}
