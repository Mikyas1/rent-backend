package messeging

type MessagingAPI interface {
	Notify(userId string,  eventName string, data interface{}) error
	BulkNotify(userIds []string,  eventName string, data interface{}) error
}
