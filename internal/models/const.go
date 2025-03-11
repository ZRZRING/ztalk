package models

const (
	OrderTime        = "time"
	OrderScore       = "score"
	ContextUserIDKey = "userID"
	SecondsInOneDay  = 60 * 60 * 24
	SecondsInOneWeek = SecondsInOneDay * 7
	ScorePerVote     = SecondsInOneDay / 200
)
