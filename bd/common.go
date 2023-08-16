package bd

import (
	"os"

	"github.com/elEdupown/planetagoluser/models"
	"github.com/elEdupown/planetagoluser/secretm"
)

var SecretModel models.SecretRDSJson
var err error

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}
