package smtpSender

import "net/smtp"

type MailSenderService struct {
	client *MailhogClient
}

func NewSenderService(client *MailhogClient) *MailSenderService {
	return &MailSenderService{
		client: client,
	}
}

func (m *MailSenderService) SendMessage(receiver string) error {
	subject := "Subject: Аутентификация"
	body := "Здравствуйте! Перейдите по ссылке, чтобы подтвердить аккаунт"
	message := []byte(subject + "\n" + body)
	to := []string{receiver}

	err := smtp.SendMail(m.client.SmtpDSN, nil, m.client.SmtpFrom, to, message)

	if err != nil {
		return err
	}

	return nil
}
