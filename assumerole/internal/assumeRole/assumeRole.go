package assumeRole

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

type Service interface {
	Connect() error
}

type AssumeRole struct {
	Sess    *session.Session
	Creds   *credentials.Credentials
	Svc     *s3.S3
	s3Token credentials.Value
}

func (a *AssumeRole) Connect() error {
	var err error
	a.Sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	a.Creds = stscreds.NewCredentials(a.Sess, os.Getenv("ROLE"))

	a.Svc = s3.New(a.Sess, &aws.Config{Credentials: a.Creds})
	a.s3Token, err = a.Creds.Get()
	if err != nil {
		return err
	}
	fmt.Println(a.s3Token.SessionToken)

	return nil
}

func NewAssumeRole() *AssumeRole {
	return &AssumeRole{}
}
