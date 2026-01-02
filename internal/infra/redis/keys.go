package redis

func SocketOwner(userID string) string {
	return "socket:user:" + userID
}

func WSEventsChannel() string {
	return "ws:events"
}
