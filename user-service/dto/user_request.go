package dto

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	ImageUrl string `json:"imageUrl"`
}
