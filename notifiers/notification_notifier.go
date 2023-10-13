package notifiers

import "fmt"

type NotificationNotifier struct{}

func (n NotificationNotifier) Send(content map[string]string) {
	fmt.Printf("Sending Notification: %s - %s\n", content["title"], content["content"])
	// Bildirim gönderim işlemi
}
