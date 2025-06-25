package models

import (
	u "be_crud/utils"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type User struct {
	ID   			int				`json:"id"`
	Name 			string		`json:"name"`
	Email 		string		`json:"email"`
	Password	string		`json:"password"`
}

type users []User

func FindAllUser() []User {
	// connect ke db dulu
	conn, err := u.ConnectDB()
	if err != nil {

	}
	defer func(){
		conn.Conn().Close(context.Background())
	}()

	// get row of database
	rows, err := conn.Query(
		context.Background(),
		`
			SELECT * FROM users
		`,
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

