package utils

import (
	"log"
	"strconv"

	"fyne.io/fyne/v2/widget"
)

func StrToInt(aE *widget.Entry, dE *widget.Entry, hE *widget.Entry, tE *widget.Entry, aDE *widget.Entry, dDE *widget.Entry, hDE *widget.Entry, tDE *widget.Entry) (int, int, int, int, int, int, int, int) {
	attackEntryTextI, err := strconv.Atoi(aE.Text)
	if err != nil {
		log.Fatalln(err)
	}
	defenceEntryTextI, err := strconv.Atoi(dE.Text)
	if err != nil {
		log.Fatalln(err)
	}
	healthEntryTextI, err := strconv.Atoi(hE.Text)
	if err != nil {
		log.Fatalln(err)
	}
	troopsEntryTextI, err := strconv.Atoi(tE.Text)
	if err != nil {
		log.Fatalln(err)
	}

	attackDefenderEntryTextI, err := strconv.Atoi(aDE.Text)
	if err != nil {
		log.Fatalln(err)
	}
	defenceDefenderEntryTextI, err := strconv.Atoi(dDE.Text)
	if err != nil {
		log.Fatalln(err)
	}
	healthDefenderEntryTextI, err := strconv.Atoi(hDE.Text)
	if err != nil {
		log.Fatalln(err)
	}
	troopsDefenderEntryTextI, err := strconv.Atoi(tDE.Text)
	if err != nil {
		log.Fatalln(err)
	}

	return attackEntryTextI, defenceEntryTextI, healthEntryTextI, troopsEntryTextI, attackDefenderEntryTextI, defenceDefenderEntryTextI, healthDefenderEntryTextI, troopsDefenderEntryTextI
}
