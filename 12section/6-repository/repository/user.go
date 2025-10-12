package repository

import (
	"context"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"learning-go/12section/6-repository/models"
)

type UserRepository interface {
	CreateUser(name, email, hashedPassword, avatar string) (int64, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUsers() ([]models.User, error)
}

type SQLUserRepository struct {
	db *sql.DB
}

// NewSQLUserRepository create new UserRepository type
func NewSQLUserRepository(db *sql.DB) UserRepository {
	return &SQLUserRepository{
		db: db,
	}
}

func (r *SQLUserRepository) CreateUser(name, email, hashedPassword, avatar string) (int64, error) {
	ctx := context.Background()

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	hp, err := bcrypt.GenerateFromPassword([]byte(hashedPassword), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(name, email, string(hp))
	if err != nil {
		return 0, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	profileStm, err := tx.PrepareContext(ctx, `INSERT INTO profile (user_id, avatar) VALUES( ?, ?)`)
	if err != nil {
		return 0, err
	}

	defer profileStm.Close()
	_, err = profileStm.Exec(userID, avatar)
	if err != nil {
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (r *SQLUserRepository) GetUserByEmail(email string) (*models.User, error) {
	stmt := `SELECT u.id, u.name, u.email,  u.hashed_password, u.created_at, p.avatar FROM users u 
	INNER JOIN profile p ON u.id = p.user_id WHERE u.email = ?`

	row := r.db.QueryRow(stmt, email)
	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword, &user.CreatedAt, &user.Profile.Avatar)
	if err != nil {
		return nil, err
	}
	user.Profile.UserID = user.ID

	return &user, nil
}

func (r *SQLUserRepository) GetUsers() ([]models.User, error) {
	stmt := `SELECT id, name, email,  hashed_password, created_at FROM users`
	rows, err := r.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID,
			&user.Name,
			&user.Email,
			&user.HashedPassword,
			&user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
