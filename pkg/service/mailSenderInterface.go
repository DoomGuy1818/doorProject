package service

type MailSenderInterface interface {
	SendMessage(receiver string) error
}
