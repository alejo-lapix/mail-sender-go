package pkg

import "gopkg.in/gomail.v2"

type goMailSender struct {
	dialer *gomail.Dialer
}

func NewGoMailer(dialer *gomail.Dialer) *goMailSender {
	return &goMailSender{dialer: dialer}
}

func (mailer *goMailSender) Send(messages ...message) error {
	format := gomail.NewMessage()
	stack := make([]*gomail.Message, len(messages))

	for index, message := range messages {
		format.SetHeader("From", message.from)
		format.SetHeader("To", message.to...)
		format.SetHeader("Subject", message.subject)
		format.SetBody(message.body.contentType, message.body.content)

		for _, path := range message.attachments {
			format.Attach(path)
		}

		stack[index] = format
	}

	return mailer.dialer.DialAndSend(stack...)
}
