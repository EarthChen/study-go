package shapes

type Shape interface {
	Area() float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}

func (rectangle Rectangle) Perimeter() float64 {
	return Perimeter(rectangle.Width, rectangle.Height)
}

func Perimeter(width float64, height float64) float64 {
	return 2 * (width + height)
}

func (rectangle Rectangle) Area() float64 {
	return Area(rectangle.Width, rectangle.Height)
}

func Area(width float64, height float64) float64 {
	return width * height
}
