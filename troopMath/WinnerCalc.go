package troopMath

import (
	"fyne.io/fyne/v2/widget"
	"github.com/Sparpvp/VikingsStatsCalc/strUtils"
)

// All the values needed to calculate the verdict of the battle
type PlayersEntity struct {
	powerWeaker       uint32
	powerStronger     uint32
	powerAttacker     uint32
	powerDefender     uint32
	troopsStronger    uint32
	troopsWeaker      uint32
	EqualTroopNeeded  uint64
	powerTrAttacker   uint64
	powerTrDefender   uint64
	powerTrStronger   uint64
	powerTrWeaker     uint64
	strongerString    string
	weakerString      string
	Winner            string
	EightyCTrAttacker uint32
	EightyCTrDefender uint32
	PercentSat        uint32
	LossesAttacker    uint32
	LossesDefender    uint32
}

// GUI Objects with int value
type EntryTexts struct {
	attackEntryText          int
	defenceEntryText         int
	healthEntryText          int
	troopsEntryText          int
	attackDefenderEntryText  int
	defenceDefenderEntryText int
	healthDefenderEntryText  int
	troopsDefenderEntryText  int
}

// Calculate the winner in an attack, returns a PlayersEntity struct which contains the necessary values
func WinnerCalc(aE *widget.Entry, dE *widget.Entry, hE *widget.Entry, tE *widget.Entry, aDE *widget.Entry, dDE *widget.Entry, hDE *widget.Entry, tDE *widget.Entry) *PlayersEntity {
	p := PlayersEntity{}
	e := EntryTexts{}
	e.attackEntryText, e.defenceEntryText, e.healthEntryText, e.troopsEntryText, e.attackDefenderEntryText, e.defenceDefenderEntryText, e.healthDefenderEntryText, e.troopsDefenderEntryText = strUtils.StrToInt(aE, dE, hE, tE, aDE, dDE, hDE, tDE)

	p.powerAttacker = uint32(e.attackEntryText) + uint32(e.defenceEntryText) + uint32(e.healthEntryText)
	p.powerTrAttacker = uint64(p.powerAttacker) * uint64(e.troopsEntryText)
	p.powerDefender = uint32(e.attackDefenderEntryText) + uint32(e.defenceDefenderEntryText) + uint32(e.healthDefenderEntryText)
	p.powerTrDefender = uint64(p.powerDefender) * uint64(e.troopsDefenderEntryText)

	if p.powerAttacker < p.powerDefender {
		p.powerWeaker = p.powerAttacker
		p.troopsWeaker = uint32(e.troopsEntryText)
		p.powerTrWeaker = p.powerTrAttacker
		p.powerStronger = p.powerDefender
		p.troopsStronger = uint32(e.troopsDefenderEntryText)
		p.powerTrStronger = p.powerTrDefender
		p.strongerString = "Defender"
		p.weakerString = "Attacker"
	} else {
		p.powerWeaker = p.powerDefender
		p.troopsWeaker = uint32(e.troopsDefenderEntryText)
		p.powerTrWeaker = p.powerTrDefender
		p.powerStronger = p.powerAttacker
		p.troopsStronger = uint32(e.troopsEntryText)
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

	p.EightyCTrAttacker = uint32(((e.troopsEntryText) * 80) / 100)
	p.EightyCTrDefender = uint32(((e.troopsDefenderEntryText) * 80) / 100)

	if e.troopsEntryText == e.troopsDefenderEntryText {
		if p.EqualTroopNeeded == uint64(p.troopsWeaker) {
			p.LossesAttacker = p.EightyCTrAttacker
			p.LossesDefender = p.EightyCTrDefender
			p.Winner = "Defender; 80/80 case triggered"
		}
	}

	return &p
}
