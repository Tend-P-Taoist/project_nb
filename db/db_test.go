package db
//
import "testing"
import (
	"fmt"
	"project_nb/models"
)

func TestFindUserAndRole(t *testing.T) {

	user := FindUserAndRole()
	fmt.Print(user)

}

func TestCreateUser(t *testing.T) {
	user := models.User{}

	user.PhoneNumber = "23423423423"
	CreateUser(&user)
}