package dto

// UserDTO represents the structure of the user response
type UserDTO struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
