package profileDto

type ProfileRequest struct {
	Gender   string `gorm:"type: varchar(255)" json:"gender" validate:"required"`
	Phone    string `gorm:"type: varchar(255)" json:"phone" validate:"required"`
	Image    string `gorm:"type: varchar(255)" json:"image" validate:"required"`
	Address  string `gorm:"type:varchar(255)" json:"address" validate:"required"`
	Location string `gorm:"type:varchar(255)" json:"location" validate:"required"`
}

type UpdateProfileRequest struct {
	Gender   string `json:"gender" form:"fullname"`
	Phone    string `json:"phone" form:"phone"`
	Image    string `json:"image" form:"image"`
	Address  string `json:"address" form:"address"`
	Location string `json:"location" form:"location"`
}
