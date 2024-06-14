package repository

import (
	"database/sql"
	"errors"
	"gts-dry/model"
	"time"
)

type UserRepository interface {
	GetByUserEmail(email string) (*model.User, error)
	ChangePassword(email, password, id string) (bool, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByUserEmail(email string) (*model.User, error) {

	rows := r.db.QueryRow("SELECT email, password, is_active, Unit FROM user_login WHERE email = ? and is_active=1", email)
	user := model.User{}
	err := rows.Scan(&user.Email, &user.Password, &user.IsActive, &user.Unit)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("akun tidak ditemukan")
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) ChangePassword(email, password, id string) (bool, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return false, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	dateUpdate := time.Now()
	formattedDate := dateUpdate.Format("2006-01-02 15:04:05")

	updateHeaderQuery := `
		UPDATE user_login set Password=?, UserUpdate=?, DateTimeUpdate=?
		where Email=? and DateTimeDelete is null
	`
	_, err = tx.Exec(updateHeaderQuery, password, id, formattedDate, email)
	if err != nil {
		return false, err
	}

	if err := tx.Commit(); err != nil {
		return false, err
	}

	return true, nil
}
