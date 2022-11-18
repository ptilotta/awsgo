package awsgo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context
var Cfg aws.Config
var err error

// InicializoAWS es la función inicial para inicializar AWS
func InicializoAWS() {
	Ctx = context.TODO()
	//Para Lambda
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1"))
	//Para prueba en local
	//Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1"), config.WithSharedConfigProfile("tiendaclic-dba"))
	if err != nil {
		panic("Error al cargar la configuracions .aws/config " + err.Error())
	}
}
