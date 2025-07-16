package enum

type LogType int8

const (
	LoginLogType   LogType = 1 //登录日志
	ActionLogType  LogType = 2 //运行日志
	RuntimeLogTYpe LogType = 3 //操作日志
)
