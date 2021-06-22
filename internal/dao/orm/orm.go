package orm

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/gradelyng/testRepo/repo/config"
	"github.com/gradelyng/testRepo/repo/internal/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//Connect returns a dao
func Connect() *MySQL {

	c := config.Config
	conStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", c.DBUser, c.DBPassWord, c.DBHost, c.DBName) //
	db, er := gorm.Open("mysql", conStr)
	if er != nil {
		log.Fatal(er)
	}

	db.AutoMigrate(&models.User{})

	return &MySQL{DB: db}
}

//MySQL ...
type MySQL struct {
	DB *gorm.DB
}

func createRandUsers() (users []models.User) {

	rand.Seed(time.Now().Unix())
	for i := 0; i < 100; i++ {
		users = append(users, models.User{FirstName: strconv.Itoa(i + rand.Intn(20)), LastName: strconv.Itoa(i * rand.Intn(5)), Age: rand.Intn(30)})
	}
	return
}

//CreateUsers ...
func (m *MySQL) CreateUsers(users []models.User, userids *[]uint) ([]models.User, error) {

	for i, _ := range users {
		// if m.DB.Table("users").Where("first_name=? AND last_name=?", v.FirstName, v.LastName).First(&users[i]).RowsAffected > 0 {
		// 	// w.Header().Add("content-type", "application/json")
		// 	// w.WriteHeader(http.StatusUnprocessableEntity)
		// 	// json.NewEncoder(w).Encode(map[string]interface{}{"message": "record already exist", "data": v})
		// 	continue
		// }

		result := m.DB.Table("users").Create(&users[i])
		*userids = append(*userids, users[i].ID)
		if er := result.Error; er != nil {
			// w.WriteHeader(http.StatusUnprocessableEntity)
			// w.Write([]byte(er.Error()))
			log.Println(er)
			continue
		}
		// id, _ := result.Value
		// users[i].ID =
	}

	return users, nil
}

//UpdateUsers ...
func (m *MySQL) UpdateUsers(users []models.User, userids *[]uint) ([]models.User, error) {
	for i, _ := range users {
		// if m.DB.Table("users").Where("first_name=? AND last_name=?", v.FirstName, v.LastName).First(&users[i]).RowsAffected > 0 {
		// 	// w.Header().Add("content-type", "application/json")
		// 	// w.WriteHeader(http.StatusUnprocessableEntity)
		// 	// json.NewEncoder(w).Encode(map[string]interface{}{"message": "record already exist", "data": v})
		// 	continue
		// }

		result := m.DB.Table("users").Create(&users[i])
		*userids = append(*userids, users[i].ID)
		if er := result.Error; er != nil {
			// w.WriteHeader(http.StatusUnprocessableEntity)
			// w.Write([]byte(er.Error()))
			log.Println(er)
			continue
		}
	}

	return users, nil

}

//DeleteUsers ...
func (m *MySQL) DeleteUsers(userids *[]uint) error {
	var users []models.User
	result := m.DB.Table("users").Delete(&users)
	if result.RowsAffected < 1 {
		return errors.New("records do not exist")
	}

	if er := result.Error; er != nil {
		return er
	}

	userids = nil

	return nil
}

//GetAllUsers ...
func (m *MySQL) GetAllUsers() ([]models.User, error) {

	var users []models.User

	result := m.DB.Find(&users)

	if result.RowsAffected < 1 {
		return nil, errors.New("no record found")
	}

	if er := result.Error; er != nil {
		return nil, er
	}

	return users, nil
}
