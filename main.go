package main

import (
	"fmt"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/Sparpvp/VikingsStatsCalc/src/entriesUtils"
	"github.com/Sparpvp/VikingsStatsCalc/src/troopMath"
)

func addNew(isDecrease bool, dEntries *troopMath.DecreaseEntries, pEntry *widget.Entry, iEntry *widget.Entry) {
	if isDecrease { // Decrease
		entriesUtils.ShowEntries(dEntries)
		pEntry.Hide()
		iEntry.Hide()
	} else { // PalaceLevel
		entriesUtils.HideEntries(dEntries)
		pEntry.Show()
		iEntry.Show()
	}
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Vikings Stat")

	// Attacker Part

	// Attacker Entries Init
	d := &troopMath.DecreaseEntries{}
	dEntries := entriesUtils.InitEntries(d)
	entriesUtils.HideEntries(dEntries)
	palacelvlEntry := widget.NewEntry()
	palacelvlEntry.Hide()
	palacelvlEntry.PlaceHolder = "Palace Level"
	influenceEntry := widget.NewEntry()
	influenceEntry.Hide()
	influenceEntry.PlaceHolder = "Influence"

	attackEntry := widget.NewEntry()
	defenceEntry := widget.NewEntry()
	healthEntry := widget.NewEntry()
	troopsEntry := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{
				Text: "Attack: ", Widget: attackEntry,
			},
		},
	}

	form.Append("Defence:", defenceEntry)
	form.Append("Health:", healthEntry)
	form.Append("Troops:", troopsEntry)

	element := widget.NewSelect([]string{"Base Stats", "Scout Rapport"}, func(value string) {
		log.Println("Select set to", value)
	})
	element.PlaceHolder = "Element Calculation"
	decrease := widget.NewSelect([]string{"Manual Decrease Set", "AI Decrease Calc"}, func(value string) {
		switch value {
		case "Manual Decrease Set":
			addNew(true, dEntries, palacelvlEntry, influenceEntry)

		case "AI Decrease Calc":
			addNew(false, dEntries, palacelvlEntry, influenceEntry)
		}
	})
	decrease.PlaceHolder = "Decrease"

	// Defender Part

	// Defender Entries Init
	palacelvlDefenderEntry := widget.NewEntry()
	influenceDefenderEntry := widget.NewEntry()
	palacelvlDefenderEntry.PlaceHolder = "Palace Level"
	influenceDefenderEntry.PlaceHolder = "Influence"
	palacelvlDefenderEntry.Hide()
	influenceDefenderEntry.Hide()
	attackDefenderEntry := widget.NewEntry()
	defenceDefenderEntry := widget.NewEntry()
	healthDefenderEntry := widget.NewEntry()
	troopsDefenderEntry := widget.NewEntry()

	formDefender := &widget.Form{
		Items: []*widget.FormItem{
			{
				Text: "Attack: ", Widget: attackDefenderEntry,
			},
		},
	}

	formDefender.Append("Defence:", defenceDefenderEntry)
	formDefender.Append("Health:", healthDefenderEntry)
	formDefender.Append("Troops:", troopsDefenderEntry)

	elementDefender := widget.NewSelect([]string{"Base Stats", "Scout Rapport"}, func(value string) {
		log.Println("Select set to", value)
	})
	elementDefender.PlaceHolder = "Element Calculation"

	AIDecrease := widget.NewSelect([]string{"AI Decrease Calc"}, func(value string) {
		palacelvlDefenderEntry.Show()
		influenceDefenderEntry.Show()
	})
	AIDecrease.PlaceHolder = "Decrease"

	// Calculate button

	generalButton := &widget.Form{
		SubmitText: "Calculate",
		OnSubmit: func() {
			pEntity := troopMath.WinnerCalc(attackEntry, defenceEntry, healthEntry, troopsEntry, attackDefenderEntry, defenceDefenderEntry, healthDefenderEntry, troopsDefenderEntry)
			losses := troopMath.LossesCalc(attackEntry, defenceEntry, healthEntry, troopsEntry, attackDefenderEntry, defenceDefenderEntry, healthDefenderEntry, troopsDefenderEntry)
			fmt.Println("No-Saturation Winner (this result doesn't matter in the game): ", pEntity.Winner)
			fmt.Println("Final Winner: ", pEntity.RWinner)
			fmt.Println("Equal Troop Needed: ", pEntity.EqualTroopNeeded)
			fmt.Println("Attacker's Losses", pEntity.LossesAttacker)
			fmt.Println("Defender's Losses", pEntity.LossesDefender)
			fmt.Println("Raw Saturation: ", losses.PercentSat)
		},
	}

	// Text

	attackerText := canvas.NewText("Attacker Stats", color.RGBA{0, 21, 255, 1})
	defenderText := canvas.NewText("Defender Stats", color.RGBA{0, 21, 255, 1})
	attackerText.TextStyle = fyne.TextStyle{Bold: true}
	attackerText.TextSize = 23
	defenderText.TextStyle = fyne.TextStyle{Bold: true}
	defenderText.TextSize = 23

	// Container Part

	attackerBox := container.NewVBox(attackerText, form, element, decrease, d.DecBefAtk, d.DecBefDef, d.DecBefHea, d.DecAftAtk, d.DecAftkDef, d.DecAftHea, palacelvlEntry, influenceEntry)
	defenderBox := container.NewVBox(defenderText, formDefender, elementDefender, AIDecrease, palacelvlDefenderEntry, influenceDefenderEntry)
	generalButtonBox := container.NewCenter(generalButton)
	containerBox := container.New(layout.NewHBoxLayout(), attackerBox, layout.NewSpacer(), generalButtonBox, layout.NewSpacer(), defenderBox)

	// Render

	myWindow.SetContent(containerBox)
	fixedSize := fyne.NewSize(700, 400)
	myWindow.Resize(fixedSize) // Switch between Tiling Mode -> Floating Mode waste this a bit
	myWindow.ShowAndRun()
}
