package db

import (
	"database/sql"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/elEdupown/planetagoluser/models"
	"github.com/elEdupown/planetagoluser/secretm"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error{
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Connected!")
	return nil
}

func ConnStr(clave models.SecretRDSJson) string{
	var dbUser, authToken, dbEndPoint, dbName string
	dbUser = clave.Username
	authToken = clave.Password
	dbEndPoint = clave.Host
	dbName = "planetagol"
	dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndPoint, dbName)
	fmt.Println(dsn)
	return dsn
}