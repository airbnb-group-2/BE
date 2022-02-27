package book

import B "group-project2/entities/book"

type Book interface {
	Insert(NewBook B.Books) (B.Books, error)
	SetPaid(BookID uint) (B.Books, error)
	SetCancel(BookID uint) (B.Books, error)
}
