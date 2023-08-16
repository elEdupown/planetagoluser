package db

import (
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/elEdupown/planetagoluser/models"
	"github.com/elEdupown/planetagoluser/tools"
)

func SignUp(sig models.SignUp) error{
	fmt.Println("Start Register")

	err := DbConnect()
	if err != nil {
		return err
	}

	defer.Db.Close()

	sentence := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + sig.UserEmail + "', '" + sig.UserUUID + "', '" + tools.FechaMySQL() + "')"
	fmt.Println(sentence)

	_, err = Db.Exec(sentence)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("End Register")
	return nil
}