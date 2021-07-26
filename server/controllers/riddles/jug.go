package riddles

import (
	"fmt"

	. "github.com/gregorioF2/clovers/lib/consts"
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

func pourJuggle(fromCapacity, toCapacity, goal int, invertJugs bool) *[]jug.Step {
	fromCurrentVol := fromCapacity
	toCurrentVol := 0

	steps := make([]jug.Step, 0)
	if goal == 0 {
		return &steps
	}
	steps = append(steps, newStep(fromCurrentVol, toCurrentVol, invertJugs))

	for fromCurrentVol != goal && toCurrentVol != goal {
		maxPourAmmount := MinInt(fromCurrentVol, toCapacity-toCurrentVol)

		fromCurrentVol -= maxPourAmmount
		toCurrentVol += maxPourAmmount
		steps = append(steps, newStep(fromCurrentVol, toCurrentVol, invertJugs))

		if fromCurrentVol == goal || toCurrentVol == goal {
			return &steps
		}

		if fromCurrentVol == 0 {
			fromCurrentVol = fromCapacity
			steps = append(steps, newStep(fromCurrentVol, toCurrentVol, invertJugs))
		}

		if toCurrentVol == toCapacity {
			toCurrentVol = 0
			steps = append(steps, newStep(fromCurrentVol, toCurrentVol, invertJugs))
		}
	}
	return &steps
}

func validateJugRiddleParameters(x int, y int, z int) (bool, string) {
	if x <= 0 {
		return false, "'x' jug capacity must be greater than zero."
	}
	if y <= 0 {
		return false, "'y' jug capacity must be greater than zero."
	}
	if z < 0 {
		return false, "'z' goal must be greater or equal than zero."
	}

	if z > x && z > y {
		return false, "'z' goal can't be greater than 'x' jug capacity, and 'y' jug capacity."
	}
	return true, ""
}

func JugRiddle(x int, y int, z int) (*[]jug.Step, *Error) {
	ok, err := validateJugRiddleParameters(x, y, z)
	if !ok {
		return nil, NewError(fmt.Sprintf("Error: %s", err), HttpStatusCode["ClientError"]["BadRequest"])
	}

	if !thereIsSolutionJugRiddle(x, y, z) {
		return nil, nil
	}

	pourXtoYSol := pourJuggle(x, y, z, false)
	pourYtoXSol := pourJuggle(y, x, z, true)

	if len(*pourXtoYSol) <= len(*pourYtoXSol) {
		return pourXtoYSol, nil
	} else {
		return pourYtoXSol, nil
	}
}
