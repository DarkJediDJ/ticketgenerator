package cloud

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

// AwsServiceInterface describes necessary methods
type AwsServiceInterface interface {
	GetSession() *session.Session
}

type awsService struct {
	Session *session.Session
}

var (
	AwsService AwsServiceInterface = &awsService{}
)

// GetSession creates new S3 connection
func (as *awsService) GetSession() *session.Session {

	S3Region := os.Getenv("REGION")

	if as.Session != nil {
		return as.Session
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      &S3Region,
		Credentials: credentials.NewEnvCredentials(),
	})

	if err != nil {
		log.Fatal(err)
	}

	as.Session = sess

	return as.Session
}
