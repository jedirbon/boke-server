package enum

type LoginType int8

const (
	PwdType LoginType = 1
	QQ      LoginType = 2
	Email   LoginType = 3
)
