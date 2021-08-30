package math

import (
	"strconv"

	"fyne.io/fyne/v2/widget"
)

var (
	weaker   uint16
	stronger uint16
)

func troopPower(aE *widget.Entry, dE *widget.Entry, hE *widget.Entry, tE *widget.Entry, aDE *widget.Entry, dDE *widget.Entry, hDE *widget.Entry, tDE *widget.Entry) uint16 {
	power1 := aE.Text + dE.Text + hE.Text
	newae, _ := strconv.Atoi(aE.Text)

	powert1 := power1 * tE.Text
	power2 := aDE.Text + dDE.Text + hDE.Text
	powert2 := power2 * tDE.Text

	if power1 < power2 {
		weaker = power1
	} else {
		stronger = power2
	}

	attackEntryTextI, defenceEntryTextI, healthEntryTextI, troopsEntryTextI, attackDefenderEntryTextI, defenceDefenderEntryTextI, healthDefenderEntryTextI, troopsDefenderEntryTextI := utils.strToInt(aE, dE, hE, tE, aDE, dDE, hDE, tDE)

}
