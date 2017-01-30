package model

type Product struct {
	ID		int		`gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name		string		`json:"name,omitempty" form:"name"  binding:"required"`
	Price 		int		`json:"price,omitempty" form:"price" binding:"required"`
	Description	string		`json:"description,omitempty" form:"description" binding:"required"`
}