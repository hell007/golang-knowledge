package FacadePattern

type ShapeMaker struct {
	circle    Shape
	rectangle Shape
	square    Shape
}

func (s *ShapeMaker) ShapeMaker() {
	s.circle = new(Circle)
	s.rectangle = new(Rectangle)
	s.square = new(Square)
}

func (s *ShapeMaker) DrawCircle() {
	s.circle.Draw2()
}

func (s *ShapeMaker) DrawRectangle() {
	s.rectangle.Draw2()
}

func (s *ShapeMaker) DrawSquare() {
	s.square.Draw2()
}
