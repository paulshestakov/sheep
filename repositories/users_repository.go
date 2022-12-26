package repositories

import (
	"database/sql"
	"fmt"
	"main/domain"

	_ "github.com/lib/pq"
)



type usersRepository struct {
	db *sql.DB
}

type UsersRepositoryInterface interface {
	Get(string) (*domain.User, error)
	Create(user *domain.User) (*domain.User, error)
}

func NewUsersRepository(db *sql.DB) UsersRepositoryInterface {
	return &usersRepository { db: db }
}

func (r *usersRepository) Get(id string) (*domain.User, error) {
	var user domain.User
	row := r.db.QueryRow("SELECT id, name FROM users WHERE id=$1;", id)
	err := row.Scan(&user.Id, &user.Name)
	if err != nil {
		fmt.Println("this is the error man: ", err)
		return nil, err
	}
	return &user, nil
}

func (r *usersRepository) Create(user *domain.User) (*domain.User, error) {
	fmt.Println(user.Id)
	fmt.Println(user.Name)
	_, err := r.db.Exec("INSERT INTO users (id, name) VALUES($1, $2);", user.Id, user.Name)
	if err != nil {
		fmt.Println("this is the error man: ", err)
		return nil, err
	}
	return user, nil
}

