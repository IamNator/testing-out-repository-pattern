package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/IamNator/testRepo/repo/config"
	"github.com/IamNator/testRepo/repo/http/handlers"
	"github.com/IamNator/testRepo/repo/internal/dao/orm"
	"github.com/IamNator/testRepo/repo/internal/dao/raw"
	"github.com/IamNator/testRepo/repo/internal/models"
)

func createRandUsers() (users []models.User) {
	for i := 0; i < 50; i++ {
		users = append(users, models.User{FirstName: strconv.Itoa(i + rand.Intn(20)), LastName: strconv.Itoa(i * rand.Intn(5)), Age: rand.Intn(30)})
	}
	return
}

func main() {
	ormDao := orm.Connect()
	rawDao := raw.Connect()

	rand.Seed(time.Now().Unix())
	user1, user2 := createRandUsers(), createRandUsers()

	handlersOrm := handlers.New(ormDao, user1, user2)
	handlersRaw := handlers.New(rawDao, user1, user2)

	http.HandleFunc("/orm", handlersOrm.Handle)
	http.HandleFunc("/raw", handlersRaw.Handle)

	port := config.Config.PORT
	fmt.Printf("server running on port %s \n/[raw|orm]\nPOST to create users\nPUT to update user details\nGET to get all users\nDELETE to delete all users\n", port)
	http.ListenAndServe(":"+port, nil)
}
