package user

import U "group-project2/entities/user"

type User interface {
	Insert(NewUser U.Users) (U.Users, error)
	GetUserByID(UserID uint) (U.Users, error)
	Update(UpdatedUser U.Users) (U.Users, error)
	SetRenter(UserID uint) (U.Users, error)
	DeleteByID(UserID uint) error
}
