package shapes

import "testing"

func TestPerimeter(t *testing.T) {

	perimeterTests := []struct {
		name         string
		shape        Shape
		hasPerimeter float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10.0, Height: 10.0}, hasPerimeter: 40.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasPerimeter: 62.83185307179586},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			if shapeWithPerimeter, ok := tt.shape.(Options); ok {
				got := shapeWithPerimeter.Perimeter()
				if got != tt.hasPerimeter {
					t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasPerimeter)
				}
			} else {
				println("not implemented")
			}
		})
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10.0, Height: 10.0}, hasArea: 100.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
