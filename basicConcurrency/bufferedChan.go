package main

import (
	"fmt"
	"time"
)

func putEmails(channel chan string, emails []string) {
	for _, email := range (emails) {
		channel <- email
	}
	close(channel) // Close the channel after all emails are sent
}

func sendEmails(channel chan string) {
	for email := range channel {
		fmt.Println("Sending email: ", email)
		time.Sleep(1 * time.Second) // Simulate email sending delay
	}
}

func main() {
	emails := []string{"efef", "hello", "efwef", "hi", "bye"}

	channel := make(chan string, 5)
	go putEmails(channel, emails)
	
	sendEmails(channel)
}
