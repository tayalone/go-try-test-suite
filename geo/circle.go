package geo

import "math"

/*Circle is Cirlcle */
type Circle struct {
	R float64
}

/*Area is Area Of Reactangle */
func (c *Circle) Area() float64 {
	return math.Pi * c.R * c.R
}

/*Perimeter is Perimeter Of Circle*/
func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.R
}
