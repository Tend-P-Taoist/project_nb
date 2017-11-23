package db
//
import "testing"
import (
	//"../models"
	"fmt"
)

func TestFindUserAndRole(t *testing.T) {

	user := FindUserAndRole()
	fmt.Print(user)

}