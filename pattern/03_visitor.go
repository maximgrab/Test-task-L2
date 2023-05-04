type visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
}

type areaCalcVisitor struct {
	area int
}

func (a *areaCalcVisitor) visitForSquare(s *square) {
	a.area = s.side * s.side
}
func (a *areaCalcVisitor) visitForCircle(c *circle) {
	a.area = int(float64(c.radius) * float64(c.radius) * 3.14)
}

type shape interface {
	getType() string
	accept(visitor)
}

type square struct {
	side int
}

func (s *square) accept(v visitor) {
	v.visitForSquare(s)
}
func (s *square) getType() string {
	return "Square"
}

type circle struct {
	radius int
}

func (c *circle) accept(v visitor) {
	v.visitForCircle(c)
}
func (c *circle) getType() string {
	return "Circle"
}