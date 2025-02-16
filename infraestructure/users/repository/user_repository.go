package repository

import (
	"database/sql"
	"pet-info-backend-docker/infraestructure/users/model"
	"pet-info-backend-docker/infraestructure/users/util"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (connection *UserRepository) CreateUser(user *model.User) error {
	transaction, err := connection.DB.Begin()
	if err != nil {
		return err
	}

	query := "INSERT INTO users.users (username, name, email, status) VALUES (:username, :name, :email, :status)"
	params := map[string]interface{}{
		"username": user.Username,
		"name":     user.Name,
		"email":    user.Email,
		"status":   user.Status,
	}
	query, args := util.ReplaceNamedParams(query, params)

	_, err = transaction.Exec(query, args...)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}

	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}

	return nil
}

func (connection *UserRepository) GetUserByEmail(email string) ([]model.User, error) {
	transaction, err := connection.DB.Begin()
	if err != nil {
		return nil, err
	}

	query := `SELECT u.id, u.username, u."name", u.email, u.status FROM users.users u WHERE u.email = :email`
	params := map[string]interface{}{
		"email": email,
	}
	query, args := util.ReplaceNamedParams(query, params)

	rows, err := transaction.Query(query, args...)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return nil, rbErr
		}
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.Username, &user.Name, &user.Email, &user.Status); err != nil {
			if rbErr := transaction.Rollback(); rbErr != nil {
				return nil, rbErr
			}
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return nil, rbErr
		}
		return nil, err
	}

	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return nil, rbErr
		}
		return nil, err
	}

	return users, nil
}

func (connection *UserRepository) GetUserByUsername(username string) ([]model.User, error) {
	transaction, err := connection.DB.Begin()
	if err != nil {
		return nil, err
	}

	query := "SELECT username, name, email, status FROM users.users WHERE username = :username"
	params := map[string]interface{}{
		"username": username,
	}
	query, args := util.ReplaceNamedParams(query, params)

	rows, err := transaction.Query(query, args...)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return nil, rbErr
		}
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.Username, &user.Name, &user.Email, &user.Status); err != nil {
			if rbErr := transaction.Rollback(); rbErr != nil {
				return nil, rbErr
			}
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return nil, rbErr
		}
		return nil, err
	}

	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return nil, rbErr
		}
		return nil, err
	}

	return users, nil
}
