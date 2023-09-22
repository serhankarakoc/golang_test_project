package notifiers

import "fmt"

type EmailNotifier struct{}

func (n EmailNotifier) Send(content map[string]string) {
	fmt.Printf("Sending Email to %s: %s - %s\n", content["email"], content["title"], content["content"])
	// Add email sending logic here
}
