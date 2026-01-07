package entity

import "gorm.io/gorm"

type Books struct {
	gorm.Model
	Title string  `valid:"stringlength(3|100)~Title is not length 3 between 100"`
	Price float64 `valid:"range(50,5000)~Price must be between 50 and 5000"`
	Code  string  `valid:"matches(^BK//d{6}$)~Price is not Range 50 between 5000"`
}
