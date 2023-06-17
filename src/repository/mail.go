package repository

type Mail struct{}

type SendData struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
	MailTo  string `json:"mailto"`
}