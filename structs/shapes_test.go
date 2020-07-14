package structs

import "testing"

func TestPerimeter(t *testing.T) {
	r := Rectangle{
		Width:  10.0,
		Height: 10.0,
	}
	got := r.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("\ngot:  %.2f\nwant: %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{
			"Rectangle",
			Rectangle{12.0, 6.0},
			72.0,
		},
		{
			"Circle",
			Circle{10.0},
			314.1592653589793,
		},
		{"Triangle", Triangle{12.0, 6.0}, 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("\n%#v\ngot:  %g\nwant: %g", tt.shape, got, tt.want)
			}
		})
	}
}
