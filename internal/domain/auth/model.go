package auth

type User struct {
	ID       uint
	Email    string
	Password string
	Username string
	Avatar   string
	Rank     int32
}
