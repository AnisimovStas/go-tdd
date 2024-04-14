package geom

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.3, 15.5}
	got := Perimeter(rectangle)
	want := 51.6
	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.name, got, tt.want)
			}
		})
	}

}
