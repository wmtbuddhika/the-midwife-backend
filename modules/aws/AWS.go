package aws

import (
	"back-end/modules/logger"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/sns"
	"os"
)

func authSNS() *sns.SNS {
	region := os.Getenv("AWS_REGION")
	svc := sns.New(session.New(&aws.Config {
		Region: &region,
	}))
	return svc
}

func authSES() *ses.SES {
	region := os.Getenv("AWS_REGION")
	svc := ses.New(session.New(&aws.Config {
		Region: &region,
	}))
	return svc
}

func SendSMS(phoneNo string, message string)  {
	svc := authSNS()

	msg := &sns.PublishInput{
		Message: aws.String(message),
		PhoneNumber: aws.String(phoneNo),
		MessageAttributes: map[string]*sns.MessageAttributeValue{
			"AWS.SNS.SMS.SenderID": {
				DataType: aws.String("String"),
				StringValue: aws.String("MidWife"),
			},
		},
	}

	result, err := svc.Publish(msg)
	if err == nil {
		logger.AddLogger(logger.INFO, "SMS sent - " + *result.MessageId)
	}
}

func SendEmail(e Email) {
	svc := authSES()

	if e.HTML == "" {
		e.HTML = e.Text
	}

	msg := &ses.Message{
		Subject: &ses.Content{
			Charset: aws.String("utf-8"),
			Data:    &e.Subject,
		},
		Body: &ses.Body{
			Html: &ses.Content{
				Charset: aws.String("utf-8"),
				Data:    &e.HTML,
			},
			Text: &ses.Content{
				Charset: aws.String("utf-8"),
				Data:    &e.Text,
			},
		},
	}

	dest := &ses.Destination{
		ToAddresses: aws.StringSlice(e.To),
	}

	email := &ses.SendEmailInput{
		Source:           &e.From,
		Destination:      dest,
		Message:          msg,
		ReplyToAddresses: aws.StringSlice(e.ReplyTo),
	}

	result, err := svc.SendEmail(email)

	if err == nil {
		logger.AddLogger(logger.INFO, "Email sent - " + *result.MessageId)
	}
}

func SendVerificationEmail(Recipient string) {
	svc := authSES()

	_, err := svc.VerifyEmailAddress(&ses.VerifyEmailAddressInput{EmailAddress: aws.String(Recipient)})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}

		return
	}

	fmt.Println("Verification sent to address: " + Recipient)
}