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
	"github.com/Sparpvp/VikingsStatsCalc/math"
)

var (
	bDecrease          bool = false
	bPalacelvl         bool = false
	bDefenderDecrease  bool = false
	bDefenderPalacelvl bool = false
)

func addNew(bDec bool, bPal bool, dEntry *widget.Entry, pEntry *widget.Entry, iEntry *widget.Entry) {
	for {
		if bDec {
			dEntry.Show()
			pEntry.Hide()
			iEntry.Hide()
			break
		}
		if bPal {
			dEntry.Hide()
			pEntry.Show()
			iEntry.Show()
			break
		}
	}
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Vikings Stat")

	// Attacker Part

	// Attacker Entries Init
	decreaseEntry := widget.NewEntry()
	decreaseEntry.Hide()
	palacelvlEntry := widget.NewEntry()
	palacelvlEntry.Hide()
	influenceEntry := widget.NewEntry()
	influenceEntry.Hide()

	attackEntry := widget.NewEntry()
	defenceEntry := widget.NewEntry()
	healthEntry := widget.NewEntry()
	troopsEntry := widget.NewEntry()
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Attack: ", Widget: attackEntry}},
	}
	form.Append("Defence:", defenceEntry)
	form.Append("Health:", healthEntry)
	form.Append("Troops:", troopsEntry)

	element := widget.NewSelect([]string{"Base Stats", "Attack VS"}, func(value string) {
		log.Println("Select set to", value)
	})
	element.PlaceHolder = "Element Calculation"
	decrease := widget.NewSelect([]string{"Manual Decrease Set", "AI Decrease Calc"}, func(value string) {
		if value == "Manual Decrease Set" {
			bDecrease = true
			bPalacelvl = false
			go addNew(bDecrease, bPalacelvl, decreaseEntry, palacelvlEntry, influenceEntry)
		}
		if value == "AI Decrease Calc" {
			bPalacelvl = true
			bDecrease = false
			go addNew(bDecrease, bPalacelvl, decreaseEntry, palacelvlEntry, influenceEntry)
		}
	})
	decrease.PlaceHolder = "Decrease"

	// Defender Part

	// Defender Entries Init
	decreaseDefenderEntry := widget.NewEntry()
	decreaseDefenderEntry.Hide()
	palacelvlDefenderEntry := widget.NewEntry()
	palacelvlDefenderEntry.Hide()
	influenceDefenderEntry := widget.NewEntry()
	influenceDefenderEntry.Hide()

	attackDefenderEntry := widget.NewEntry()
	defenceDefenderEntry := widget.NewEntry()
	healthDefenderEntry := widget.NewEntry()
	troopsDefenderEntry := widget.NewEntry()
	formDefender := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Attack: ", Widget: attackDefenderEntry}},
	}
	formDefender.Append("Defence:", defenceDefenderEntry)
	formDefender.Append("Health:", healthDefenderEntry)
	formDefender.Append("Troops:", troopsDefenderEntry)

	elementDefender := widget.NewSelect([]string{"Base Stats", "Attack VS"}, func(value string) {
		log.Println("Select set to", value)
	})
	elementDefender.PlaceHolder = "Element Calculation"
	decreaseDefender := widget.NewSelect([]string{"Manual Decrease Set", "AI Decrease Calc"}, func(value string) {
		if value == "Manual Decrease Set" {
			bDefenderDecrease = true
			bDefenderPalacelvl = false
			go addNew(bDefenderDecrease, bDefenderPalacelvl, decreaseDefenderEntry, palacelvlDefenderEntry, influenceDefenderEntry)
		}
		if value == "AI Decrease Calc" {
			bDefenderPalacelvl = true
			bDefenderDecrease = false
			go addNew(bDefenderDecrease, bDefenderPalacelvl, decreaseDefenderEntry, palacelvlDefenderEntry, influenceDefenderEntry)
		}
	})
	decreaseDefender.PlaceHolder = "Decrease"

	// Calculate button

	generalButton := &widget.Form{
		SubmitText: "Calculate",
		OnSubmit: func() {
			winner := math.WinnerCalc(attackEntry, defenceEntry, healthEntry, troopsEntry, attackDefenderEntry, defenceDefenderEntry, healthDefenderEntry, troopsDefenderEntry)
			fmt.Println("the winner is: ", winner)
			fmt.Println("(dbg) equal troop needed:", math.EqualTroopNeeded)
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

	attackerBox := container.NewVBox(attackerText, form, element, decrease, decreaseEntry, palacelvlEntry, influenceEntry)
	defenderBox := container.NewVBox(defenderText, formDefender, elementDefender, decreaseDefender, decreaseDefenderEntry, palacelvlDefenderEntry, influenceDefenderEntry)
	generalButtonBox := container.NewCenter(generalButton)
	containerBox := container.New(layout.NewHBoxLayout(), attackerBox, layout.NewSpacer(), generalButtonBox, layout.NewSpacer(), defenderBox)

	// Render

	myWindow.SetContent(containerBox)
	fixedSize := fyne.NewSize(700, 400)
	myWindow.Resize(fixedSize) // Switch between Tiling Mode -> Floating Mode waste this a bit
	myWindow.ShowAndRun()
}
