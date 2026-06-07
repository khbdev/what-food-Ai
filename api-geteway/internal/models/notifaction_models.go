package models

type CreateNotificationRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

type UpdateNotificationRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

type GetNotificationsQuery struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}
