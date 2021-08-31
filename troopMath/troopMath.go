package troopMath

import (
	"fyne.io/fyne/v2/widget"
	"github.com/Sparpvp/VikingsStatsCalc/strUtils"
)

type PlayersEntity struct {
	powerWeaker      uint32
	powerStronger    uint32
	powerAttacker    uint32
	powerDefender    uint32
	troopsStronger   uint32
	troopsWeaker     uint32
	EqualTroopNeeded uint64
	powerTrAttacker  uint64
	powerTrDefender  uint64
	powerTrStronger  uint64
	powerTrWeaker    uint64
	strongerString   string
	weakerString     string
	Winner           string
}

// Calculate the winner in an attack, returns a PlayersEntity struct which contains the necessary values
func WinnerCalc(aE *widget.Entry, dE *widget.Entry, hE *widget.Entry, tE *widget.Entry, aDE *widget.Entry, dDE *widget.Entry, hDE *widget.Entry, tDE *widget.Entry) *PlayersEntity {
	attackEntryText, defenceEntryText, healthEntryText, troopsEntryText, attackDefenderEntryText, defenceDefenderEntryText, healthDefenderEntryText, troopsDefenderEntryText := strUtils.StrToInt(aE, dE, hE, tE, aDE, dDE, hDE, tDE)
	p := PlayersEntity{}

	p.powerAttacker = uint32(attackEntryText) + uint32(defenceEntryText) + uint32(healthEntryText)
	p.powerTrAttacker = uint64(p.powerAttacker) * uint64(troopsEntryText)
	p.powerDefender = uint32(attackDefenderEntryText) + uint32(defenceDefenderEntryText) + uint32(healthDefenderEntryText)
	p.powerTrDefender = uint64(p.powerDefender) * uint64(troopsDefenderEntryText)

	if p.powerAttacker < p.powerDefender {
		p.powerWeaker = p.powerAttacker
		p.troopsWeaker = uint32(troopsEntryText)
		p.powerTrWeaker = p.powerTrAttacker
		p.powerStronger = p.powerDefender
		p.troopsStronger = uint32(troopsDefenderEntryText)
		p.powerTrStronger = p.powerTrDefender
		p.strongerString = "Defender"
		p.weakerString = "Attacker"
	} else {
		p.powerWeaker = p.powerDefender
		p.troopsWeaker = uint32(troopsDefenderEntryText)
		p.powerTrWeaker = p.powerTrDefender
		p.powerStronger = p.powerAttacker
		p.troopsStronger = uint32(troopsEntryText)
		p.powerTrStronger = p.powerTrAttacker
		p.strongerString = "Attacker"
		p.weakerString = "Defender"
	}

	p.EqualTroopNeeded = p.powerTrStronger / uint64(p.powerWeaker) // The attacker/defender DOESN'T HAVE THIS AMOUNT OF TROOPS

	if p.EqualTroopNeeded < uint64(p.troopsWeaker) {
		p.Winner = p.weakerString
	} else {
		p.Winner = p.strongerString
	}

	return &PlayersEntity{}
}
