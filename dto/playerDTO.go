package dto

// PlayerDTO represents the structure of the user response
type PlayerDTO struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Nickname  string `json:"nickname"`
    Email string `json:"email"`
}
