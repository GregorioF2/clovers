package riddles_test

import (
	"testing"

	controller "github.com/gregorioF2/clovers/controllers/riddles"
	"github.com/gregorioF2/clovers/models/riddles/jug"
)

func TestJugRiddleWrongParameters(t *testing.T) {
	var tests = []struct {
		title   string
		x, y, z int
	}{
		{"X jug negative value", -1, 1, 1},
		{"X jug zero value", 0, 1, 1},
		{"Y jug negative value", 1, -1, 1},
		{"Y jug zero value", 1, 0, 1},
		{"Z greater than X and Y", 1, 1, 3},
		{"Z goal negative value", 1, 1, -1},
	}
	for _, test := range tests {

		testname := test.title
		t.Run(testname, func(t *testing.T) {

			_, err := controller.JugRiddle(test.x, test.y, test.z)
			if err == nil {
				t.Errorf("Got no error, when expecting one.")
			}
		})
	}
}

func TestJugRiddleWithNoSolution(t *testing.T) {
	x, y, z := 3, 6, 2
	solution, err := controller.JugRiddle(x, y, z)

	if err != nil {
		t.Errorf("Error solving riddle: %s.", err.Error())
	}
	if solution != nil {
		t.Errorf("Expected solution to be nil.")
	}
}

func TestJugRiddleWithSolution(t *testing.T) {

	var tests = []struct {
		title    string
		x, y, z  int
		expected []jug.Step
	}{
		{
			title: "X greater than Y",
			x:     5, y: 7, z: 3,
			expected: []jug.Step{{X: 5, Y: 0}, {X: 0, Y: 5}, {X: 5, Y: 5}, {X: 3, Y: 7}},
		},
		{
			title: "Y greater than X",
			x:     7, y: 5, z: 3,
			expected: []jug.Step{{X: 0, Y: 5}, {X: 5, Y: 0}, {X: 5, Y: 5}, {X: 7, Y: 3}},
		},
		{title: "X equal to Y",
			x: 7, y: 7, z: 7,
			expected: []jug.Step{{X: 7, Y: 0}},
		},
		{
			title: "Z equal to 0",
			x:     7, y: 7, z: 0,
			expected: []jug.Step{},
		},
	}
	for _, test := range tests {

		testname := test.title
		t.Run(testname, func(t *testing.T) {
			sol, err := controller.JugRiddle(test.x, test.y, test.z)
			if err != nil {
				t.Errorf("Got error, solving jug riddle.")
			}
			if sol == nil {
				t.Errorf("Got no solution on jug riddle.")
			}
			expectedResultOk := true
			for i, step := range *sol {
				expectedResultOk = expectedResultOk && step == test.expected[i]
			}
			if !expectedResultOk {
				t.Error("Got unexpected result for jug riddle. Result: ", *sol, ". Expected: ", test.expected)
			}
		})
	}
}
