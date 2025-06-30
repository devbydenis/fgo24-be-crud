package models

import (
	u "be_crud/utils"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type User struct {
	ID   			int				`json:"id"`
	Name 			string		`json:"name"`
	Email 		string		`json:"email"`
	Password	string		`json:"password"`
	CreatedAt *time.Time	`json:"created_at,omitempty"`
	UpdatedAt *time.Time	`json:"updated_at,omitempty"`
}

type users []User

type CreateUser struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserType struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func FindAllUser(query string) []User {
	// connect ke db dulu
	conn, err := u.ConnectDB()
	if err != nil {
		fmt.Println("failed to connect to database", err)
	}
	defer func(){
		conn.Conn().Close(context.Background())
	}()

	// get row of database
	rows, err := conn.Query(
		context.Background(),
		`
			SELECT id, name, email, password, created_at, updated_at FROM users
			WHERE name = $1
		`,
		query,
	)
	if err != nil {
		fmt.Println("failed to query rows:", err)
	}

	// get collect row to mapping into struct
	users, err := pgx.CollectRows[User](rows, pgx.RowToStructByName)
	if err != nil {
		fmt.Println("failed to collect rows:", err)
		return []User{}
	}

	return users
}

func FindUserById(id int) User {
	// connect ke db dulu
	conn, err := u.ConnectDB()
	if err != nil {
		fmt.Println("failed to connect to database", err)
	}
	defer func(){
		conn.Conn().Close(context.Background())
	}()

	// get row of database
	row, err := conn.Query(
		context.Background(),
		`
			SELECT * FROM users WHERE id = $1
		`,
		id,
	)
	if err != nil {
		fmt.Println("failed to query row:", err)
	}

	// get collect row to mapping into struct
	user, err := pgx.CollectOneRow[User](row, pgx.RowToStructByName)
	if err != nil {
		fmt.Println("failed to collect row:", err)
		return User{}
	}

	return user
}

func AddingNewUSer(user CreateUser) {
	// connect ke db dulu
	conn, err := u.ConnectDB()
	if err != nil {
		fmt.Println("failed to connect to database", err)
	}
	defer func(){
		conn.Conn().Close(context.Background())
	}()

	// get row of database
	_, err = conn.Exec(
		context.Background(),
		`
			INSERT INTO users(name, email, password) VALUES ($1, $2, $3)
		`,
		user.Name,
		user.Email,
		user.Password,
	)

	if err != nil {
		fmt.Println("failed to update row:", err)
	}
}

func UpdateUser(user *UpdateUserType) {
	// connect ke db dulu
	conn, err := u.ConnectDB()
	if err != nil {
		fmt.Println("failed to connect to database", err)
	}
	defer func(){
		conn.Conn().Close(context.Background())
	}()

	queryChangeName := `
	UPDATE users 
	SET name = $1, updated_at = NOW() 
	WHERE id = $2
	`
	queryChangeEmail := `
	UPDATE users 
	SET email = $1, updated_at = NOW() 
	WHERE id = $2
	`

	// get row of database
	if user.Name == "" {
		_, err = conn.Exec(context.Background(), queryChangeEmail, user.Email, user.ID)
		if err != nil {
			fmt.Println("failed to update row with column email:", err)
		}
		return
	}
	_, err = conn.Exec(context.Background(), queryChangeName, user.Name, user.ID)
	
	
	if err != nil {
		fmt.Println("failed to update row with column name:", err)
	}

}