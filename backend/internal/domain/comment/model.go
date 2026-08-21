package comment

import "time"

// Структура одобренного комментария
type Comment struct {
	ID       int
	Author   string
	Message  string
	Response string
	Date     time.Time
}
