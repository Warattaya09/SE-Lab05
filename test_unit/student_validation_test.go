package unit

import (
	"lab05/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

// func validStudent() Student {
//     return Student{
//         Fullname: "John Doe",
//         Age:      20,
//         Email:    "john.doe@example.com",
//         GPA:      3.50,
//     }
// }

//กรณีข้อมูลครบทุก fieldถูกต้อง
func TestValidateStudent_AllFieldsValid(t *testing.T) {
    g := NewGomegaWithT(t)

    student := entity.Student{
        FullName: "Alice Johnson",
        Age:      22,
        Email:    "alice.johnson@example.com",
        GPA:      3.8,
    }
    ok, err := govalidator.ValidateStruct(student)
    g.Expect(ok).To(BeTrue())
    g.Expect(err).To(BeNil())
}

//กรณี Fullname เป็นค่าว่าง
func TestValidateStudent_EmptyFullname(t *testing.T) {
    g := NewGomegaWithT(t)

    student := entity.Student{
        FullName: "",
        Age:      22,
        Email:    "alice.johnson@example.com",
        GPA:      3.8,
    }
    ok, err := govalidator.ValidateStruct(student)
    g.Expect(ok).ToNot(BeTrue())
    g.Expect(err).ToNot(BeNil())
    g.Expect(err.Error()).To(Equal("Fullname is required")) // ตรวจว่าข้อความข้อผิดพลาดมีคำว่า "FullName"
}

//กรณีอายุน้อยกว่า 18 ปี
func TestValidateStudent_Underage(t *testing.T) {
    g := NewGomegaWithT(t)
    student := entity.Student{
        FullName: "Alice Johnson",
        Age:      17,
        Email:    "alice.johnson@example.com",
        GPA:      3.8,
    }
    ok, err := govalidator.ValidateStruct(student)
    g.Expect(ok).ToNot(BeTrue())
    g.Expect(err).ToNot(BeNil())
    g.Expect(err.Error()).To(Equal("Age must be at least 18")) // ตรวจว่าข้อความข้อผิดพลาดมีคำว่า "Age"
}

//กรณี Email รูปแบบไม่ถูกต้อง
func TestValidateStudent_InvalidEmail(t *testing.T) {
    g := NewGomegaWithT(t)
    student := entity.Student{
        FullName: "Alice Johnson",
        Age:      22,
        Email:    "invalid-email-at-example",
        GPA:      3.8,
    }
    ok, err := govalidator.ValidateStruct(student)
    g.Expect(ok).ToNot(BeTrue())
    g.Expect(err).ToNot(BeNil())
    g.Expect(err.Error()).To(Equal("Email is Invalid")) // ตรวจว่าข้อความข้อผิดพลาดมีคำว่า "Email"
}

//กรณี GPA ไม่อยู่ในช่วง 0.00-4.00
func TestValidateStudent_InvalidGPA_OutOfRange(t *testing.T) {
    g := NewGomegaWithT(t)
    student := entity.Student{
        FullName: "Alice Johnson",
        Age:      22,
        Email:    "alice.johnson@example.com",
        GPA:      4.5,
    }
    ok, err := govalidator.ValidateStruct(student)
    g.Expect(ok).ToNot(BeTrue())
    g.Expect(err).ToNot(BeNil())
    g.Expect(err.Error()).To(Equal("GPA must be between 0.0 to 4.0")) // ตรวจว่าข้อความข้อผิดพลาดมีคำว่า "GPA"
}