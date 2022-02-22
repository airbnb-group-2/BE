package auth

import U "group-project2/entities/user"

type Auth interface {
	Login(email, password string) (U.Users, error)
}
