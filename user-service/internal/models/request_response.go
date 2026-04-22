package models




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