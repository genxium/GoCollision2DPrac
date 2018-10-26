package main

import (
	"fmt"
  "github.com/Tarliton/collision2d"
)

func main() {
  pos1 := collision2d.NewVector(100.0, -100.0)
  pos2 := collision2d.NewVector(100.0, 0.0)
  pos3 := collision2d.NewVector(100.0, 300.0)
  pos4 := collision2d.NewVector(100.0, 100.0)

  var polygon collision2d.Polygon
  {
    polygonCorners := [...]float64{
      0, 100,
      -100, 0,
      -100, -100,
      100, -100,
      100, 0,
    }
    pos := collision2d.NewVector(pos1.X, pos1.Y)
    offset := collision2d.NewVector(0.0, 0.0)
    angle := 0.0
    polygon = collision2d.NewPolygon(pos, offset, angle, polygonCorners[:])
  }

  var circle collision2d.Circle
  {
    circle = collision2d.Circle{collision2d.Vector{100, 100}, 32}
  }

  {
    result, response := collision2d.TestPolygonCircle(polygon, circle)
    fmt.Printf("\n#1 result == %v, response == %v.\n", result, response)
  }

  {
    polygon.Pos = pos2
    result, response := collision2d.TestPolygonCircle(polygon, circle)
    fmt.Printf("\n#2 result == %v, response == %v.\n", result, response)
  }

  {
    polygon.Pos = pos3
    result, response := collision2d.TestPolygonCircle(polygon, circle)
    fmt.Printf("\n#3 result == %v, response == %v.\n", result, response)
  }

  {
    polygon.Pos = pos4
    result, response := collision2d.TestPolygonCircle(polygon, circle)
    fmt.Printf("\n#4 result == %v, response == %v.\n", result, response)
  }
}
