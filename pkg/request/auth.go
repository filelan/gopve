package request

type AuthenticationMethod int

const (
	AuthenticationMethodCookie AuthenticationMethod = iota
	AuthenticationMethodHeader
)
