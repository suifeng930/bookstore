package model

// session 结构体
type Session struct {
	SessionId string //sessionId
	Username  string //username
	UserId    int    //userId
	Cart      *Cart
}
