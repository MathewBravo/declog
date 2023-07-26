package decisions

import "time"

type Decision struct {
	Title       string
	Date        time.Time
	Description string
}

type Decisions struct {
	Decisions []Decision
}

func NewDecision(title string, description string) Decision {
	return Decision{
		Title:       title,
		Date:        time.Now(),
		Description: description,
	}
}
