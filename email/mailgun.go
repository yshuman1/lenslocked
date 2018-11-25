package email

import (
	"fmt"

	mailgun "gopkg.in/mailgun/mailgun-go.v1"
)

func WithMailgun(domain, apiKey, publicKey string) ClientConfig {
	return func(c *Client) {
		mg := mailgun.NewMailgun(domain, apiKey, publicKey)
		c.mg = mg
	}
}

func WithSender(name, email string) ClientConfig {
	return func(c *Client) {
		c.from = buildEmail(name, email)
	}
}

type ClientConfig func(*Client)

func NewClient(opts ...ClientConfig) *Client {
	client := Client{
		from: "support@lenslocked.com",
	}
	for _, opt := range opts {
		opt(&client)
	}
	return &client
}

type Client struct {
	from string
	mg   mailgun.Mailgun
}

func (c *Client) Welcome(toName, toEmail string) error {
	message := mailgun.NewMessage(c.from, welcomeSubject, welcomeText, buildEmail(toName, toEmail))
	message.SetHtml(welcomeHTML)
	_, _, err := c.mg.Send(message)
	return err
}

func buildEmail(name, email string) string {
	if name == "" {
		return email
	}
	return fmt.Sprintf("%s <%s>", name, email)
}

const welcomeSubject = "welcome to lenslocked.com!"
const welcomeText = `Hi there!
Welcome to our website!

Thanks for signing up!
Yasin`

const welcomeHTML = `Hi there!<br/> Welcome to <a href=“https://lenslocked.com”>Lenslocked.com</a>!<br/>Best,</br>Yasin`
