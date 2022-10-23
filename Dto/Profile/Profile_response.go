package profileDto

type ProfileResponse struct {
	Gender   string `json:"gender" form:"gender" validation:"required"`
	Phone    string `json:"phone" form:"phone" validation:"required"`
	Image    string `json:"image" form:"image" validation:"required"`
	Address  string `json:"address" from:"address" validate:"required"`
	Location string `json:"location" from:"location" validate:"required"`
	UserID   int    `json:"user_id" from:"user_id" validate:"required"`
}
