package main

type stack struct{ vec []string }

func (s stack) Empty() bool   { return len(s.vec) == 0 }
func (s stack) Peek() string  { return s.vec[len(s.vec)-1] }
func (s stack) Len() int      { return len(s.vec) }
func (s *stack) Put(i string) { s.vec = append(s.vec, i) }
func (s *stack) Pop() string {
	d := s.vec[len(s.vec)-1]
	s.vec = s.vec[:len(s.vec)-1]
	return d
}
