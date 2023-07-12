package models 

type Testimonials struct {
	Id 	  		int 		`json:"id" gorm:"primaryKey;autoIncrement"`
	Image 		string 		`json:"image"`
	Text  		string 		`json:"text" gorm:"not null"`
	Authors 	string 		`json:"authors" gorm:"not null"`
	Rating		int 		`json:"rating" gorm:"not null"`
	IsCompany 	bool 		`json:"is_company" gorm:"not null"`
}

func (Testimonials) TableName() string {
	return "testimonials"
}