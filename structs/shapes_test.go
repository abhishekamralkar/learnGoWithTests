package shapes

import "testing"

/*
type Rectangle struct {
	Width  float64
	Height float64
        }
*/

func TestShapes(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

areaTests := []struct{
	name string
	got  Area
	want float64
}{
	{name: "Rectangle", got: rectangle.Area(), want:= 72.0}
}

func TestArea(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		got := rectangle.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %.2f but want %.2f", got, want)
		}
	})

	t.Run("cirlces", func(t *testing.T) {
		circle := Circle.Area()
		got := Area(circle)
		want := 314.16

		if got != want {
			t.Errorf("got %.2f but want %.2f", got, want)
		}
	})
}
