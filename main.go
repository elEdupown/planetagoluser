package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/elEdupown/planetagoluser/awsgo"
	"github.com/elEdupown/planetagoluser/db"
	"github.com/elEdupown/planetagoluser/models"
)

func main() {
	lambda.Start(ExecLambda)
}

func ExecLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InitializeAWS()

	if !ValidateParams() {
		fmt.Println("Error: No se encontraron los parametros necesarios")
		err := errors.New("Error en los parametros, debe enviar SecretName")
		return event, err
	}

	var datos models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("Email: " + datos.UserEmail)
		case "sub":
			datos.UserUUID = att
			fmt.Println("Sub: " + datos.UserUUID)

		}
	}

	err := db.ReadSecret()
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return event, err
	}

	err = db.SignUp(datos)
	return event, err
}

func ValidateParams() bool {
	var getParam bool
	_, getParam = os.LookupEnv("SecretName")
	return getParam
}
