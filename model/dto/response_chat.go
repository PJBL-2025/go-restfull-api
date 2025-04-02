package dto

type ResponseChat struct {
	Id        int     `json:"id"`
	UserId    int     `json:"user_id"`
	Message   *string `json:"message" gorm:"column:message"`
	RoleMessage string `json:"role_message" gorm:"column:role_message"`
	ImagePath *string `json:"image_path" gorm:"column:image_path"`
	Name      *string `json:"name" gorm:"column:name"`
	Price     *int    `json:"price" gorm:"column:price"`
}
