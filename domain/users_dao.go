package domain

import (
	"bookstore/src/github.com/luckyparakh/bookstore_users-api/datasources/mysql/users_db"
	"bookstore/src/github.com/luckyparakh/bookstore_users-api/logger"
	"bookstore/src/github.com/luckyparakh/bookstore_users-api/utils/errors"
	"bookstore/src/github.com/luckyparakh/bookstore_users-api/utils/mysql_utils"

	"fmt"
)

var (
	usersDB             = make(map[int64]*User)
	updateStmt          = "UPDATE users SET first_name=?,last_name=?,email=? WHERE id=?;"
	insertStmt          = "INSERT INTO users (first_name,last_name,email,date_created,status,password) VALUES (?,?,?,?,?,?);"
	selectStmt          = "SELECT id, first_name,last_name,email,date_created,status from users where id=?;"
	delStmt             = "DELETE FROM users WHERE id=?;"
	getUserbyStatusStmt = "SELECT id,first_name,last_name,email,date_created,status from users where status=?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(selectStmt)

	if err != nil {
		logger.Error("error while running prepared stmt for select", err)
		return errors.InternalServerError("error while preparing stmt ")
	}
	defer stmt.Close()
	row := stmt.QueryRow(user.Id)
	if err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		logger.Error("error while running select stmt", err)
		return mysql_utils.ParseError(err)
	}
	return nil
	// This was used when DB was not there
	// result := usersDB[user.Id]
	// if result == nil {
	// 	return errors.NewNotFoundError(fmt.Sprintf("User with id %d not found.", user.Id))
	// }
	// user.Id = result.Id
	// user.FirstName = result.FirstName
	// user.LastName = result.LastName
	// user.DateCreated = result.DateCreated
	// user.Email = result.Email
	// return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(insertStmt)
	if err != nil {
		return errors.InternalServerError(fmt.Sprintf("error while preparing stmt %s", err.Error()))
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	uid, err := result.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	user.Id = uid
	return nil
	// This was used when DB was not there
	// current := usersDB[user.Id]
	// if current != nil {
	// 	if current.Email == user.Email {
	// 		return errors.NewBadRequestError(fmt.Sprintf("User with email %s already present.", user.Email))
	// 	}
	// 	return errors.NewBadRequestError(fmt.Sprintf("User with id %d already present.", user.Id))
	// }
	// user.DateCreated = dateUtils.GetNowString()
	// usersDB[user.Id] = user

}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(updateStmt)
	if err != nil {
		return errors.InternalServerError(fmt.Sprintf("error while preparing stmt %s", err.Error()))
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(delStmt)
	if err != nil {
		return errors.InternalServerError(fmt.Sprintf("error while preparing stmt %s", err.Error()))
	}
	defer stmt.Close()
	if _, delerr := stmt.Exec(user.Id); delerr != nil {
		return mysql_utils.ParseError(err)
	}
	return nil

}

func (user *User) GetUserByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(getUserbyStatusStmt)
	if err != nil {
		return nil, errors.InternalServerError(fmt.Sprintf("error while preparing stmt %s", err.Error()))
	}
	defer stmt.Close()
	rows, rowsErr := stmt.Query(status)
	if rowsErr != nil {
		return nil, mysql_utils.ParseError(err)
	}
	defer rows.Close()
	users := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		users = append(users, user)
	}
	if len(users) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no user with status %s found", status))
	}
	return users, nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(delStmt)
	if err != nil {
		return errors.InternalServerError(fmt.Sprintf("error while preparing stmt %s", err.Error()))
	}
	defer stmt.Close()
	if _, delerr := stmt.Exec(user.Id); delerr != nil {
		return mysql_utils.ParseError(err)
	}
	return nil

}

func (user *User) GetUserByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(getUserbyStatusStmt)
	if err != nil {
		return nil, errors.InternalServerError(fmt.Sprintf("error while preparing stmt %s", err.Error()))
	}
	defer stmt.Close()
	rows, rowsErr := stmt.Query(status)
	if rowsErr != nil {
		return nil, mysql_utils.ParseError(err)
	}
	defer rows.Close()
	users := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		users = append(users, user)
	}
	if len(users) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no user with status %s found", status))
	}
	return users, nil
}
