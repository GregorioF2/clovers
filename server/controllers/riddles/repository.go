package riddles

import (
	. "github.com/gregorioF2/clovers/lib/errors"
	. "github.com/gregorioF2/clovers/lib/utils"
	"github.com/gregorioF2/clovers/models/riddles/jug"
)

func thereIsSolutionJugRiddle(x int, y int, z int) bool {
	return (z % GCD(x, y)) == 0
}

func newStep(x, y int, invertJugs bool) jug.Step {
	if invertJugs {
		return jug.Step{Y: x, X: y}
	}
	return jug.Step{X: x, Y: y}
}

func pourJuggle(xCap, yCap, z int, invertJugs bool) []jug.Step {
	xCurr := xCap
	yCurr := 0

	var steps []jug.Step
	steps = append(steps, newStep(xCurr, yCurr, invertJugs))

	for xCurr != z && yCurr != z {
		maxPourAmmount := MinInt(xCurr, yCap-yCurr)

		xCurr -= maxPourAmmount
		yCurr += maxPourAmmount
		steps = append(steps, newStep(xCurr, yCurr, invertJugs))

		if xCurr == z || yCurr == z {
			return steps
		}

		if xCurr == 0 {
			xCurr = xCap
			steps = append(steps, newStep(xCurr, yCurr, invertJugs))
		}

		if yCurr == yCap {
			yCurr = 0
			steps = append(steps, newStep(xCurr, yCurr, invertJugs))
		}
	}
	return steps
}

func JugRiddle(x int, y int, z int) ([]jug.Step, *Error) {
	if !thereIsSolutionJugRiddle(x, y, z) {
		return []jug.Step{}, nil
	}

	pourXtoYSol := pourJuggle(x, y, z, false)
	pourYtoXSol := pourJuggle(y, x, z, true)

	if len(pourXtoYSol) <= len(pourYtoXSol) {
		return pourXtoYSol, nil
	} else {
		return pourYtoXSol, nil
	}
}
