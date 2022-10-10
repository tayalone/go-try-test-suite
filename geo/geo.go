package geo

/*Geo must have methode Area */
type Geo interface {
	Area() float64
	Perimeter() float64
}

/*Desc present basic Geometric Desctription*/
type Desc struct {
	Area      float64 `json:"area"`
	Perimeter float64 `json:"perimeter"`
}
