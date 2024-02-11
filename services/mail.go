package services

import (
	"fmt"
	"net/smtp"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func sendMail(to []string, cc []string, subject, message string) error {
	body := "From: " + os.Getenv("CONFIG_SENDER_NAME") + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := getAuth()
	smtpAddr := getSmtpAddress()
	email := os.Getenv("CONFIG_AUTH_EMAIL")
	if gin.Mode() == gin.DebugMode {
		email = "from@example.com"
		to = []string{"to@example.com"}
	}
	return smtp.SendMail(smtpAddr, auth, email, append(to, cc...), []byte(body))
}

func SendVerificationEmail(to string, token string) error {
	subject := "Account Verification"
	message := fmt.Sprintf("You've succesfully register your account. All you need now is to verify your account through this link https://digitalent.games.test.shopee.io/vm5/auth/verify/%s", token)
	return sendMail([]string{to}, []string{}, subject, message)
}

func SendResetPasswordEmail(to string, token string) error {
	subject := "Reset Password"
	message := fmt.Sprintf("You have requested to reset your Varmasea Account password. Here is your reset password link https://digitalent.games.test.shopee.io/vm5/auth/reset-password/%s", token)
	return sendMail([]string{to}, []string{}, subject, message)
}

func SendNewAdminPharmacyEmail(to string, pass string) error {
	subject := "New Admin Pharmacy Account"
	message := fmt.Sprintf(`We are delighted to grant you access to our admin panel. As an admin, you will have the privilege to manage and oversee various aspects of our platform.
	Here's the password to access your account using THIS email:
	Password: %s
	If you have any questions or need assistance, please do not hesitate to reach out to our support team.
	Thank you for joining our team of admins and helping us maintain a secure and efficient platform for all users.`, pass)
	return sendMail([]string{to}, []string{}, subject, message)
}

func getAuth() smtp.Auth {
	email := os.Getenv("CONFIG_AUTH_EMAIL")
	pass := os.Getenv("CONFIG_AUTH_PASSWORD")
	host := os.Getenv("CONFIG_SMTP_HOST")
	if gin.Mode() == gin.DebugMode {
		email = os.Getenv("CONFIG_AUTH_EMAIL_DEV")
		pass = os.Getenv("CONFIG_AUTH_PASSWORD_DEV")
		host = os.Getenv("CONFIG_SMTP_HOST_DEV")
	}
	return smtp.PlainAuth("", email, pass, host)
}

func getSmtpAddress() string {
	port, _ := strconv.Atoi(os.Getenv("CONFIG_SMTP_PORT"))
	host := os.Getenv("CONFIG_SMTP_HOST")
	if gin.Mode() == gin.DebugMode {
		host = os.Getenv("CONFIG_SMTP_HOST_DEV")
	}
	return fmt.Sprintf("%s:%d", host, port)
}
