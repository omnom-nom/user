package handlers

import (
	"sync"

        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/kms"
)

const (
	//DbIP	= "192.168.1.101"
        //DbPort	= "8000"
        dbZone	= "us-west-1"
	dbTableName = "user"
)

var (
        env  *EnvSingleton
        once sync.Once
)


func GetEnvInstance() *EnvSingleton {

        once.Do(func() {
		//dbUrl := fmt.Sprintf("http://%s:%s", DbIP, DbPort)
		config := &aws.Config{
			Region:   aws.String(dbZone),
			//Endpoint: aws.String(dbUrl),
		}
		sess := session.Must(session.NewSession(config))

                env = &EnvSingleton{
			tableName: dbTableName,
                        Db: dynamodb.New(sess),
			Kms: kms.New(sess),
                }
        })

        return env
}
