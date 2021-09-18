package troopMath

import "fyne.io/fyne/v2/widget"

type DecreaseEntries struct {
	DecBefAtk         *widget.Entry
	DecBefDef         *widget.Entry
	DecBefHea         *widget.Entry
	DecAftAtk         *widget.Entry
	DecAftkDef        *widget.Entry
	DecAftHea         *widget.Entry
}

func DecreaseCalc() *DecreaseEntries {
	d := DecreaseEntries{}

	return &d
}
