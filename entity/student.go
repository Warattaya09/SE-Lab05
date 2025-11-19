package entity

import (
	
	"gorm.io/gorm"
	
)

type Student struct {
	gorm.Model
	FullName string `valid:"required~Fullname is required"`
	Age      int	`valid:"range(18|120)~Age must be at least 18"`
	Email	 string `valid:"required~Email is required,email~Email is Invalid"`
	GPA		 float32`valid:"range(0|4)~GPA must be between 0.0 to 4.0"`
}