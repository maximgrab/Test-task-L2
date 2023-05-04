
type strategyAlgoritm interface {
	alg()
}

type strategy1 struct {
}

func (strategy1) alg() {

}

type strategy2 struct {
}

func (strategy2) alg() {

}

type algContext struct {
	algMethod strategyAlgoritm
}

func (s *algContext) alg() {
	s.algMethod.alg()
}