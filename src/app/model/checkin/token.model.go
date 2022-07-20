package checkin

type CheckinToken struct {
	Token       string
	UserId      string
	CheckinType int32
}

type TokenInfo struct {
	Id          string
	CheckinType int32
	EventType   int32
}
