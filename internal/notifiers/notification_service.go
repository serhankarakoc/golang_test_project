package notifiers

import "fmt"

type NotificationService struct {
	SMSNotifier
	EmailNotifier
	NotificationNotifier
}

func (ns NotificationService) Send(notificationType string, content map[string]string) {
	var notifier Notifier

	switch notificationType {
	case "sms":
		notifier = ns.SMSNotifier
	case "email":
		notifier = ns.EmailNotifier
	case "notification":
		notifier = ns.NotificationNotifier
	default:
		fmt.Println("Unsupported notification type")
		return
	}

	notifier.Send(content)
}
