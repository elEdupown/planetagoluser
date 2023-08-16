package bd

import (
	"os"

	"github.com/planetagoluser/models"
	"github.com/planetagoluser/secretm"
)

var SecretModel models.SecretRDSJson
var err error

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}
