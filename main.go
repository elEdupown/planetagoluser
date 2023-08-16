package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/elEdupown/planetagoluser/awsgo"
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
}

func ValidateParams() bool {
	var getParam bool
	_, getParam = os.LookupEnv("SecretName")
	return getParam
}
