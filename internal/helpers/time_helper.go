package helpers

import (
	"time"
)

func CheckIfToday(timeEval time.Time) bool {
	currentTime := time.Now().UTC()
	simpleCurrent := currentTime.Format("2011/01/01")
	timeEvalSimple := timeEval.Format("2011/01/01")
	if simpleCurrent != timeEvalSimple {
		return false
	}

	return true
}
