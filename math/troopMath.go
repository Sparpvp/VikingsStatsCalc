package math

import (
	"fyne.io/fyne/v2/widget"
	"github.com/Sparpvp/VikingsStatsCalc/utils"
)

var (
	powerWeaker      uint32
	powerStronger    uint32
	powerAttacker    uint32
	powerDefender    uint32
	troopsStronger   uint32
	troopsWeaker     uint32
	equalTroopNeeded uint64
	powerTrAttacker  uint64
	powerTrDefender  uint64
	powerTrStronger  uint64
	powerTrWeaker    uint64
	strongerString   string
	weakerString     string
	winner           string
)

func TroopPower(aE *widget.Entry, dE *widget.Entry, hE *widget.Entry, tE *widget.Entry, aDE *widget.Entry, dDE *widget.Entry, hDE *widget.Entry, tDE *widget.Entry) string {
	attackEntryText, defenceEntryText, healthEntryText, troopsEntryText, attackDefenderEntryText, defenceDefenderEntryText, healthDefenderEntryText, troopsDefenderEntryText := utils.StrToInt(aE, dE, hE, tE, aDE, dDE, hDE, tDE)

	powerAttacker = uint32(attackEntryText) + uint32(defenceEntryText) + uint32(healthEntryText)
	powerTrAttacker = uint64(powerAttacker) * uint64(troopsEntryText)
	powerDefender = uint32(attackDefenderEntryText) + uint32(defenceDefenderEntryText) + uint32(healthDefenderEntryText)
	powerTrDefender = uint64(powerDefender) * uint64(troopsDefenderEntryText)

	if powerAttacker < powerDefender {
		powerWeaker = powerAttacker
		troopsWeaker = uint32(troopsEntryText)
		powerTrWeaker = powerTrAttacker
		powerStronger = powerDefender
		troopsStronger = uint32(troopsDefenderEntryText)
		powerTrStronger = powerTrDefender
		strongerString = "Defender"
		weakerString = "Attacker"
	} else {
		powerWeaker = powerDefender
		troopsWeaker = uint32(troopsDefenderEntryText)
		powerTrWeaker = powerTrDefender
		powerStronger = powerAttacker
		troopsStronger = uint32(troopsEntryText)
		powerTrStronger = powerTrAttacker
		strongerString = "Attacker"
		weakerString = "Defender"
	}

	equalTroopNeeded = powerTrStronger / uint64(powerWeaker) // The attacker/defender DOESN'T HAVE THIS AMOUNT OF TROOPS

	if equalTroopNeeded < uint64(troopsWeaker) {
		winner = weakerString
	} else {
		winner = strongerString
	}

	return winner
}
