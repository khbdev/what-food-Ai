package models



// CreateUserRequest — yangi user yaratish uchun
type CreateUserRequest struct {
	Name  string
	Phone string
	Email string
}

// UpdateUserRequest — user yangilash uchun
type UpdateUserRequest struct {
	Name  string
	Phone string
	Email string
}

// GetAllUsersRequest — pagination bilan userlar ro'yxati olish uchun
type GetAllUsersRequest struct {
	Limit  int
	Offset int
}

// GetAllUsersResponse — pagination meta + userlar ro'yxati
type GetAllUsersResponse struct {
	Users  []User
	Total  int
	Limit  int
	Offset int
}