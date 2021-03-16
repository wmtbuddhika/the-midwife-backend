package aws

type Email struct {
	// From is the source email.
	From string
	// To is a set of destination emails.
	To []string
	// ReplyTo is a set of reply to emails.
	ReplyTo []string
	// Subject is the email subject text.
	Subject string
	// Text is the plain text representation of the body.
	Text string
	// HTML is the HTML representation of the body.
	HTML string
}