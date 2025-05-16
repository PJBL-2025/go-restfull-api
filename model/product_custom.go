package model

type ProductCustom struct {
	FrontImagePath string `json:"font_image_path" gorm:"column:font_image_path"`
	BackImagePath  string `json:"back_image_path" gorm:"column:back_image_path"`
	FrontWidth     string `json:"front_width" gorm:"column:front_width"`
	BackWidth      string `json:"back_width" gorm:"column:back_width"`
	ProductId      int    `json:"product_id" gorm:"column:product_id"`
}
