package main

import (
	"fmt"
	"time"
)

func putEmails(channel chan string, emails []string) {
	for _, email := range (emails) {
		channel <- email
		fmt.Println("email: "+email)
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

	// To check if channel is still open
	// _, ok := <-channel
	
	// if ok { 
	// 	fmt.Println("Channel still open") 
	// 	return
	// }

	// Select
	emailChannel := make(chan string)
	smsChannel := make(chan string)

	go func() {
		emailChannel<-"email"
		close(emailChannel)
	}()

	go func() {
		smsChannel<-"sms"
		close(smsChannel)
	}()

	for {
		select {
		case email, ok := <-emailChannel:
			if !ok {
				return 
			}
			fmt.Println(email)
		case sms, ok := <-smsChannel:
			if !ok {
				return
			}
			fmt.Println(sms)
		}
	}
}