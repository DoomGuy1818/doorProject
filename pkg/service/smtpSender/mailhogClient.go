package smtpSender

type MailhogClient struct {
	SmtpDSN  string
	SmtpFrom string
}

func NewMailhogClient(smtpDSN, smtpFrom string) *MailhogClient {
	return &MailhogClient{
		SmtpDSN:  smtpDSN,
		SmtpFrom: smtpFrom,
	}
}
