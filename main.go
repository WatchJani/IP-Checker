package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func init() {
	loadEnv()
}

func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println(err)
	}
}

func main() {
	v1, err := exec.Command("curl", "ifconfig.me").Output()
	if err != nil {
		log.Fatal(err)
	}

	// Sender's email address and password
	from := os.Getenv("FROM")
	password := os.Getenv("PASSWORD")

	// Recipient's email address
	to := "jankokondic@gmail.com"

	// SMTP server configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Email subject and message
	subject := "My IP address"
	body := string(v1)

	// Compose the email
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	// Connect to the SMTP server
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}

	fmt.Println("Email sent successfully!")
}
