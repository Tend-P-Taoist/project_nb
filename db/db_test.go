package db
//
import "testing"
import (
	"fmt"
	"project_nb/models"
)



func TestCreateUser(t *testing.T) {
	user := models.User{}

	user.PhoneNumber = "23423423423"
	CreateUser(&user)
}

