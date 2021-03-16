package notification

import "back-end/modules/aws"

func SendSMS(phoneNo string, message string) {
	aws.SendSMS(phoneNo, message)
}
func SendVerificationEmail(to string) {
	aws.SendVerificationEmail(to)
}

func SendEmail(from string, to []string, subject string, message string) {
	email := aws.Email {
		From:		from,
		To: 		to,	//[]string{"tbuddhika99@gmail.com"}
		ReplyTo: 	nil,
		Subject: 	subject,
		HTML:    	message,
	}
	aws.SendEmail(email)
}