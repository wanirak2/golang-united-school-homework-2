package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type intCustomeType int

const (
	SidesTriangle intCustomeType = 3
	SidesSquare                  = 4
	SidesCircle                  = 0
)

func CalcSquare(sideLen float64, sidesNum intCustomeType) float64 {
	if sidesNum == SidesTriangle {
		return math.Pow(3, 0.5) * sideLen * sideLen / 4
	} else if sidesNum == SidesSquare {
		return sideLen * sideLen
	} else if sidesNum == SidesCircle {
		return math.Pi * sideLen * sideLen
	} else {
		return 0
	}
}
