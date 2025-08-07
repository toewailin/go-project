package serialization

import (
	"go-project/dto"
	"go-project/models"
)

// UserToDTO converts a User model into a UserDTO
func UserToDTO(user models.User) dto.UserDTO {
    return dto.UserDTO{
        ID:    user.ID,
        Name:  user.Name,
        Email: user.Email,
    }
}
