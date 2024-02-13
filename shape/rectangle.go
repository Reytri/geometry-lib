package shape

type Rectangle struct {
	Width  float32
	Length float32
}

func (rectange *Rectangle) Area() float32 {
	return rectange.Width * rectange.Length
}

func (rectangle *Rectangle) Perimeter() float32 {
	return 2 * (rectangle.Width + rectangle.Length)
}
