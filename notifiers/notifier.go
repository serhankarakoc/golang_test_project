package notifiers

type Notifier interface {
	Send(content map[string]string)
}
