package assumeRole

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Service interface {
	Connect() error
}

type AssumeRole struct {
	Sess  *session.Session
	Creds *credentials.Credentials
	Svc   *s3.S3
}

func (a *AssumeRole) Connect() error {

	a.Sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	}))

	a.Creds = stscreds.NewCredentials(a.Sess, "arn:aws:iam::512933419476:role/stsRole")

	a.Svc = s3.New(a.Sess, &aws.Config{Credentials: a.Creds})
	fmt.Println(a.Creds.Get())

	return nil
}

func NewAssumeRole() *AssumeRole {
	return &AssumeRole{}
}
