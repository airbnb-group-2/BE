package user

import (
	U "group-project2/entities/user"

	"gorm.io/gorm"
)

type RequestCreateUser struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	IsRenter bool   `json:"is_renter" form:"is_renter"`
}

type ResponseCreateUser struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsRenter bool   `json:"is_renter"`
}

func (Req RequestCreateUser) ToEntityUser() U.Users {
	return U.Users{
		Name:     Req.Name,
		Email:    Req.Email,
		Password: Req.Password,
		IsRenter: Req.IsRenter,
	}
}

func ToResponseCreateUser(User U.Users) ResponseCreateUser {
	return ResponseCreateUser{
		ID:       User.ID,
		Name:     User.Name,
		Email:    User.Email,
		IsRenter: User.IsRenter,
	}
}

type ResponseGetByID struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsRenter bool   `json:"is_renter"`
}

func ToResponseGetByID(User U.Users) ResponseGetByID {
	return ResponseGetByID{
		ID:       User.ID,
		Name:     User.Name,
		Email:    User.Email,
		IsRenter: User.IsRenter,
	}
}

type RequestUpdateUser struct {
	ID       uint
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	IsRenter bool   `json:"is_renter" form:"is_renter"`
}

type ResponseUpdateUser struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsRenter bool   `json:"is_renter"`
}

func (Req RequestUpdateUser) ToEntityUser(UserID uint) U.Users {
	return U.Users{
		Model:    gorm.Model{ID: UserID},
		Name:     Req.Name,
		Email:    Req.Email,
		Password: Req.Password,
		IsRenter: Req.IsRenter,
	}
}

func ToResponseUpdate(User U.Users) ResponseUpdateUser {
	return ResponseUpdateUser{
		ID:       User.ID,
		Name:     User.Name,
		Email:    User.Email,
		IsRenter: User.IsRenter,
	}
}
