package book

import (
	"errors"
	B "group-project2/entities/book"
	"time"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (ur *BookRepository) Insert(NewBook B.Books) (B.Books, error) {
	NewBook.Status = "booked"
	if err := ur.db.Create(&NewBook).Error; err != nil {
		return B.Books{}, err
	}
	return NewBook, nil
}

func (repo *BookRepository) SetPaid(BookID uint) (B.Books, error) {
	Book := B.Books{}
	if RowsAffected := repo.db.Table("books").Where("id = ?", BookID).Update("status", "paid").RowsAffected; RowsAffected == 0 {
		return B.Books{}, errors.New("gagal mengubah status booking menjadi paid")
	}
	repo.db.First(&Book, BookID)
	return Book, nil
}

func (repo *BookRepository) SetCancel(BookID uint) (B.Books, error) {
	Book := B.Books{}
	if RowsAffected := repo.db.Table("books").Where("id = ?", BookID).Update("status", "cancel").RowsAffected; RowsAffected == 0 {
		return B.Books{}, errors.New("gagal mengubah status booking menjadi cancel")
	}
	repo.db.First(&Book, BookID)
	return Book, nil
}

func (repo *BookRepository) SetCheckInTime(BookID uint, CheckInTime time.Time) (B.Books, error) {
	Book := B.Books{}
	if RowsAffected := repo.db.Table("books").Where("id = ?", BookID).Update("check_in_time", CheckInTime).RowsAffected; RowsAffected == 0 {
		return B.Books{}, errors.New("gagal mengatur check_in_time")
	}
	repo.db.First(&Book, BookID)
	return Book, nil
}

func (repo *BookRepository) SetCheckOutTime(BookID uint, CheckOutTime time.Time) (B.Books, error) {
	Book := B.Books{}
	if RowsAffected := repo.db.Table("books").Where("id = ?", BookID).Update("check_out_time", CheckOutTime).RowsAffected; RowsAffected == 0 {
		return B.Books{}, errors.New("gagal mengatur check_out_time")
	}
	repo.db.First(&Book, BookID)
	return Book, nil
}
