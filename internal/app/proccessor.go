package app

import (
	"awesomeProject/internal/models"
)

func Proccessor(pchan chan models.AccountUpdate, acct *map[int64]float64) {
	for {
		update := <-pchan
		(*acct)[update.ID] += update.Delta
	}

}
