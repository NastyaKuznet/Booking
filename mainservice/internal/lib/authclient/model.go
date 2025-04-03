package authclient

type ValidateToken struct {
	Valid  bool
	Error  string
	Login  string
	UserId int64
}
