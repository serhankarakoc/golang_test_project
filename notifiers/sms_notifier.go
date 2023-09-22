package notifiers

import "fmt"

type SMSNotifier struct{}

func (n SMSNotifier) Send(content map[string]string) {
	fmt.Printf("Sending SMS to %s: %s\n", content["phone"], content["content"])
	// Add SMS sending logic here
}
