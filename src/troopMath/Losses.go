package troopMath

import (
	"log"

	"fyne.io/fyne/v2/widget"
)

// Saturation & Losses values
type PlayerSatLosses struct {
	PercentSat         float32 // Advantaged player's saturation
	Saturation         float32 // MinusLossSat support
	SaturationC        float32
	MinusLossSat       float32 // % of Advantage given by saturation
	AdvantagePlayerSat string
	LossesFWinner      string
	LossesFLoser       string
	LoserFWinner       string  // Unadvantaged player in saturation mechanic
	AdvantageTTroops   uint32  // Implements saturation mechanic in loser's hypothetical n. troops
	UnadvantageTTroops uint32  // Renew the troops weaker
	AdvantagedPlayer   string  // The Advantaged Player
	UnadvantagedPlayer string  // The Unadvantaged Player
	EightyCTTrAttacker float64 // 80% Theoretical Troops of Advantaged Player statistically
	EightyCTTrDefender float64
}

// Init obj(s) for WinnerCalc.go struct
var (
	p *PlayersEntity
	e *EntryTexts
)

// Get WinnerCalc.go struct
func GetStructPointer(PE *PlayersEntity, ET *EntryTexts) {
	p = PE
	e = ET
}

// Calculates the losses of the battle
func LossesCalc(aE *widget.Entry, dE *widget.Entry, hE *widget.Entry, tE *widget.Entry, aDE *widget.Entry, dDE *widget.Entry, hDE *widget.Entry, tDE *widget.Entry) *PlayerSatLosses {
	l := PlayerSatLosses{}

	if e.troopsEntryText == e.troopsDefenderEntryText {
		if p.EqualTroopNeeded == uint64(p.TroopsWeaker) {
			p.LossesAttacker = float64(p.EightyCTrAttacker)
			p.LossesDefender = float64(p.EightyCTrDefender)
			p.Winner = "Defender; 80/80 case triggered"
			p.LossesAttacker = float64(p.EightyCTrAttacker)
			p.LossesDefender = float64(p.EightyCTrDefender)
			p.RWinner = p.Winner
			log.Print(p.Winner)
			log.Fatalln("Exiting...")
		}
	}

	if e.troopsEntryText > e.troopsDefenderEntryText {
		l.PercentSat = float32(e.troopsEntryText) / float32(e.troopsDefenderEntryText) * 100
		l.AdvantagePlayerSat = "Attacker"
		l.LoserFWinner = "Defender"
	} else if e.troopsEntryText < e.troopsDefenderEntryText {
		l.PercentSat = float32(e.troopsDefenderEntryText) / float32(e.troopsEntryText) * 100
		l.AdvantagePlayerSat = "Defender"
		l.LoserFWinner = "Attacker"
	}

	l.SaturationC = l.PercentSat
	l.Saturation = float32(l.PercentSat - 100)
	l.Saturation = l.Saturation / 15
	l.Saturation = l.Saturation * float32(100)
	l.MinusLossSat = (11.532889 / 100) * l.Saturation

	if l.SaturationC > 115 {
		//Max saturation reached
		l.SaturationC = 115
	}

	if l.MinusLossSat > 11.532889 {
		// Max advantage reached
		l.MinusLossSat = 11.532889
	}

	l.UnadvantagedPlayer = p.WeakerString
	if p.Winner == "Attacker" {
		l.LossesFWinner = "Attacker"
		l.LossesFLoser = "Defender"
	}

	if p.Winner == "Defender" {
		l.LossesFWinner = "Defender"
		l.LossesFLoser = "Attacker"
	}

	l.EightyCTTrAttacker = float64(p.EqualTroopNeeded) / 100 * 80
	l.EightyCTTrDefender = float64(p.EqualTroopNeeded) / 100 * 80

	// Both stats are equal
	if p.PowerAttacker == p.PowerDefender {
		if e.troopsEntryText < e.troopsDefenderEntryText {
			p.Winner = "Defender"
			p.RWinner = "Defender"
			p.LossesAttacker = float64(p.EightyCTrAttacker)
			p.LossesDefender = float64(p.EightyCTrDefender) - ((float64(p.EightyCTrDefender)*float64(l.SaturationC))/10000)*float64(l.MinusLossSat)
		} else {
			p.Winner = "Attacker"
			p.RWinner = "Attacker"
			p.LossesDefender = float64(p.EightyCTrDefender)
			p.LossesAttacker = float64(p.EightyCTrDefender) - ((float64(p.EightyCTrDefender)*float64(l.SaturationC))/10000)*float64(l.MinusLossSat)
		}
	}

	if p.Winner == l.AdvantagePlayerSat {
		if l.LossesFWinner == "Attacker" {
			if p.StrongerString == "Attacker" {
				p.LossesAttacker = float64(float32(p.EightyCTrDefender) - ((float32(p.EightyCTrDefender)*l.SaturationC)/10000)*l.MinusLossSat)
				p.LossesAttacker = p.LossesAttacker * float64(p.PowerDefender) / float64(p.PowerAttacker)
				p.LossesDefender = float64(p.EightyCTrDefender)
				p.RWinner = "Attacker"
			} else {
				// If the Attacker is not advantaged statistically the calculation will be the normal one
				p.LossesAttacker = float64(float32(p.EightyCTrAttacker) - ((float32(p.EightyCTrAttacker)*l.SaturationC)/10000)*l.MinusLossSat)
				p.LossesDefender = float64(p.EightyCTrDefender)
				p.RWinner = "Attacker"
			}
		} else if p.StrongerString == "Defender" {
			p.LossesDefender = float64(float32(p.EightyCTrDefender) - ((float32(p.EightyCTrDefender)*l.SaturationC)/10000)*l.MinusLossSat)
			p.LossesDefender = p.LossesDefender * float64(p.PowerAttacker) / float64(p.PowerDefender)
			p.LossesAttacker = float64(p.EightyCTrAttacker)
			p.RWinner = "Defender"
		} else {
			// If the Defender is not advantaged statistically the calculation will be the normal one
			p.LossesAttacker = float64(p.EightyCTrAttacker)
			p.LossesDefender = float64(float32(p.EightyCTrAttacker) - ((float32(p.EightyCTrAttacker)*l.SaturationC)/10000)*l.MinusLossSat)
			p.RWinner = "Defender"
		}
	} else if l.LossesFLoser == "Attacker" {
		if p.StrongerString == "Attacker" {
			l.AdvantageTTroops = uint32(float64(p.EqualTroopNeeded) + ((float64(p.EqualTroopNeeded) * float64(l.SaturationC) / 10000) * float64(l.MinusLossSat)))
			l.AdvantagedPlayer = "Attacker"
		} else {
			l.UnadvantageTTroops = uint32(float32(p.TroopsWeaker) + (((float32(p.TroopsWeaker) * float32(l.SaturationC)) / 10000) * l.MinusLossSat))
			l.AdvantagedPlayer = "Defender"
		}
	} else if l.LossesFLoser == "Defender" {
		if p.StrongerString == "Defender" {
			l.AdvantageTTroops = uint32(float64(p.EqualTroopNeeded) + (((float64(p.EqualTroopNeeded) * float64(l.SaturationC)) / 10000) * float64(l.MinusLossSat)))
			l.UnadvantageTTroops = p.TroopsWeaker
			l.AdvantagedPlayer = "Defender"
		} else {
			l.UnadvantageTTroops = uint32(float32(p.TroopsWeaker) + (((float32(p.TroopsWeaker) * l.SaturationC) / 10000) * l.MinusLossSat))
			l.AdvantageTTroops = uint32(p.EqualTroopNeeded)
			l.AdvantagedPlayer = "Attacker"
		}
		if l.AdvantageTTroops < l.UnadvantageTTroops {
			p.RWinner = l.UnadvantagedPlayer
		} else {
			p.RWinner = l.AdvantagedPlayer
		}

		if p.RWinner == "Attacker" {
			if p.StrongerString == "Attacker" {
				p.LossesDefender = float64(p.EightyCTrDefender)
				p.LossesAttacker = float64(p.EightyCTrDefender)
				p.LossesAttacker = (p.LossesAttacker * float64(p.PowerDefender)) / float64(p.PowerAttacker)
			} else {
				p.LossesDefender = (float64(l.EightyCTTrDefender) * float64(p.PowerAttacker)) / float64(p.PowerDefender)
				p.LossesAttacker = float64(l.EightyCTTrDefender)
			}
		} else if p.RWinner == "Defender" {
			if p.StrongerString == "Defender" {
				p.LossesAttacker = float64(p.EightyCTrAttacker)
				p.LossesDefender = float64(p.EightyCTrAttacker)
				p.LossesDefender = (p.LossesDefender * float64(p.PowerAttacker)) / float64(p.PowerDefender)
			} else {
				p.LossesAttacker = (float64(l.EightyCTTrAttacker) * float64(p.PowerDefender)) / float64(p.PowerAttacker)
				p.LossesDefender = float64(l.EightyCTTrAttacker)
			}
		}
	}

	return &l
}
