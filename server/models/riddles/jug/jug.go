package jug

type Step struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (s *Step) InvertJugs() {
	tmp := s.X
	s.X = s.Y
	s.Y = tmp
}
