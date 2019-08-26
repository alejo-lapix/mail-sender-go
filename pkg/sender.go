package pkg

type MailSender interface {
	Send(...message) error
}

type message struct {
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

func NewBody(contentType, content string) *body {
	return &body{
		contentType: contentType,
		content:     content,
	}
}

func NewMessage(from string, to []string, subject string, body *body, attachments []string) (*message, error) {
	// TODO validate input

	return &message{
		from:        from,
		to:          to,
		subject:     subject,
		body:        body,
		attachments: attachments,
	}, nil
}
