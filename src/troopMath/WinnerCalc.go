package troopMath

import (
	"fyne.io/fyne/v2/widget"
	"github.com/Sparpvp/VikingsStatsCalc/src/strUtils"
)

// All the values needed to calculate the verdict of the battle
type PlayersEntity struct {
	PowerAttacker     uint32 // Hypothetical power of the attacker's troops
	PowerDefender     uint32 // Hypothetical power of the defender's troops
	PowerWeaker       uint32
	PowerStronger     uint32
	TroopsStronger    uint32
	TroopsWeaker      uint32
	EqualTroopNeeded  uint64 // Number of troops needed by the weaker player to equate the stronger opponent
	PowerTrAttacker   uint64
	PowerTrDefender   uint64
	PowerTrStronger   uint64
	PowerTrWeaker     uint64
	StrongerString    string
	WeakerString      string
	EightyCTrAttacker uint32 // Max losses of the attacker
	EightyCTrDefender uint32 // Max losses of the defender
	LossesAttacker    float64
	LossesDefender    float64
	Winner            string
	RWinner           string // The Real Winner counting Saturation
}

// GUI Objects with integer value
type EntryTexts struct {
	attackEntryText          int // Attacker's attack
	defenceEntryText         int // Attacker's defence
	healthEntryText          int // Attacker's health
	troopsEntryText          int // Attacker's number of troops
	attackDefenderEntryText  int // Defender's attack
	defenceDefenderEntryText int // Defender's defence
	healthDefenderEntryText  int // Defender's health
	troopsDefenderEntryText  int // Defender's number of troops
}

// Calculate the winner in an attack, returns a PlayersEntity struct which contains the necessary values
func WinnerCalc(aE *widget.Entry, dE *widget.Entry, hE *widget.Entry, tE *widget.Entry, aDE *widget.Entry, dDE *widget.Entry, hDE *widget.Entry, tDE *widget.Entry) *PlayersEntity {
	p := PlayersEntity{}
	e := EntryTexts{}

	e.attackEntryText, e.defenceEntryText, e.healthEntryText, e.troopsEntryText, e.attackDefenderEntryText, e.defenceDefenderEntryText, e.healthDefenderEntryText, e.troopsDefenderEntryText = strUtils.StrToInt(aE, dE, hE, tE, aDE, dDE, hDE, tDE)

	p.PowerAttacker = uint32(e.attackEntryText) + uint32(e.defenceEntryText) + uint32(e.healthEntryText)
	p.PowerTrAttacker = uint64(p.PowerAttacker) * uint64(e.troopsEntryText)
	p.PowerDefender = uint32(e.attackDefenderEntryText) + uint32(e.defenceDefenderEntryText) + uint32(e.healthDefenderEntryText)
	p.PowerTrDefender = uint64(p.PowerDefender) * uint64(e.troopsDefenderEntryText)

	if p.PowerAttacker < p.PowerDefender {
		p.PowerWeaker = p.PowerAttacker
		p.TroopsWeaker = uint32(e.troopsEntryText)
		p.PowerTrWeaker = p.PowerTrAttacker
		p.PowerStronger = p.PowerDefender
		p.TroopsStronger = uint32(e.troopsDefenderEntryText)
		p.PowerTrStronger = p.PowerTrDefender
		p.StrongerString = "Defender"
		p.WeakerString = "Attacker"
	} else {
		p.PowerWeaker = p.PowerDefender
		p.TroopsWeaker = uint32(e.troopsDefenderEntryText)
		p.PowerTrWeaker = p.PowerTrDefender
		p.PowerStronger = p.PowerAttacker
		p.TroopsStronger = uint32(e.troopsEntryText)
		p.PowerTrStronger = p.PowerTrAttacker
		p.StrongerString = "Attacker"
		p.WeakerString = "Defender"
	}

	p.EqualTroopNeeded = p.PowerTrStronger / uint64(p.PowerWeaker) // The attacker/defender DOESN'T HAVE THIS AMOUNT OF TROOPS

	if p.EqualTroopNeeded < uint64(p.TroopsWeaker) {
		p.Winner = p.WeakerString
	} else {
		p.Winner = p.StrongerString
	}

	p.EightyCTrAttacker = uint32(((e.troopsEntryText) * 80) / 100)
	p.EightyCTrDefender = uint32(((e.troopsDefenderEntryText) * 80) / 100)

	GetStructPointer(&p, &e)
	return &p
}
