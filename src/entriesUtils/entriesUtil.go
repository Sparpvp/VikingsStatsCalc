package entriesUtils

import (
	"fyne.io/fyne/v2/widget"
	"github.com/Sparpvp/VikingsStatsCalc/src/troopMath"
)

func InitEntries(d *troopMath.DecreaseEntries) *troopMath.DecreaseEntries {
	// Create Entries & Set Placeholder

	d.DecBefAtk = widget.NewEntry()
	d.DecBefDef = widget.NewEntry()
	d.DecBefHea = widget.NewEntry()
	d.DecAftAtk = widget.NewEntry()
	d.DecAftkDef = widget.NewEntry()
	d.DecAftHea = widget.NewEntry()

	d.DecBefAtk.PlaceHolder = "Normal Scout Attack"
	d.DecBefDef.PlaceHolder = "Normal Scout Defence"
	d.DecBefHea.PlaceHolder = "Normal Scout Health"
	d.DecAftAtk.PlaceHolder = "Decreased Scout Attack"
	d.DecAftkDef.PlaceHolder = "Decreased Scout Defence"
	d.DecAftHea.PlaceHolder = "Decreased Scout Health"

	return d
}

func HideEntries(d *troopMath.DecreaseEntries) {
	d.DecBefAtk.Hide()
	d.DecBefDef.Hide()
	d.DecBefHea.Hide()
	d.DecAftAtk.Hide()
	d.DecAftkDef.Hide()
	d.DecAftHea.Hide()
}

func ShowEntries(d *troopMath.DecreaseEntries) {
	d.DecBefAtk.Show()
	d.DecBefDef.Show()
	d.DecBefHea.Show()
	d.DecAftAtk.Show()
	d.DecAftkDef.Show()
	d.DecAftHea.Show()
}
