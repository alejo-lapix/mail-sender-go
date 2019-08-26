package pkg

type MailSender interface {
	Send(...Message) error
}

type Message struct {
	from        string
	to          []string
	subject     string
	body        *body
	attachments []string
}

type body struct {
	contentType string
	content     string
}
