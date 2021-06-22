package raw

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/IamNator/testRepo/repo/config"
	"github.com/IamNator/testRepo/repo/internal/models"
)

//Raw takes a pointer to sql connection
type Raw struct {
	db *sql.DB
}

//Connect returns a Raw object
func Connect() *Raw {

	c := config.Config
	conStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", c.DBUser, c.DBPassWord, c.DBHost, c.DBName)
	db, er := sql.Open("mysql", conStr)
	if er != nil {
		log.Fatal(er)
	}

	return &Raw{db}
}

//CreateUsers adds a new user to database
func (r *Raw) CreateUsers(users []models.User, userids *[]uint) ([]models.User, error) {
	for i, user := range users {

		result, err := r.db.Exec(`INSERT INTO users(first_name, last_name, age) VALUES(?, ?, ? )`,
			user.FirstName, user.LastName, user.Age,
		)
		if err != nil {
			return nil, err
		}

		n, _ := result.LastInsertId()
		*userids = append(*userids, uint(n))
		users[i].ID = uint(n)

	}

	return users, nil
}

//UpdateUsers update users information
func (r *Raw) UpdateUsers(users []models.User, userids *[]uint) ([]models.User, error) {

	for i, user := range users {

		if i > len(*userids)-1 {
			return nil, errors.New("no records exist")
		}
		_, err := r.db.Exec(`UPDATE users SET first_name = ?, last_name = ?, age = ? WHERE id = ?`,
			user.FirstName, user.LastName, user.Age, (*userids)[i],
		)

		if err != nil {
			return nil, err
		}

		//	n, _ := result.LastInsertId()
		//idsupdated = append(idsupdated, uint(n))
	}

	return users, nil

}

//DeleteUsers delete users from a database
func (r *Raw) DeleteUsers(userids *[]uint) error {
	result, er := r.db.Exec("DELETE FROM users")
	if er != nil {

		return er
	}

	userids = nil

	if i, _ := result.RowsAffected(); i < int64(1) {

		return errors.New("records do not exist")
	}

	return nil
}

//GetAllUsers gets all users from database
func (r *Raw) GetAllUsers() ([]models.User, error) {
	var users []models.User

	row, err := r.db.Query(`SELECT * FROM users;`)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var user models.User
		er := row.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Age,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if er != nil {
			return nil, er
		}
		if user.FirstName != "" {
			users = append(users, user)
		}

	}

	if len(users) < 1 {
		return nil, errors.New("no record found")
	}

	return users, nil

}
