//go:build unit
// +build unit

package geo

import "testing"

func TestCircle_Area(t *testing.T) {
	type fields struct {
		R float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
		{
			name:   "Case 1: Find Area",
			fields: fields{R: 7},
			want:   153.93804002589985,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Circle{
				R: tt.fields.R,
			}
			if got := c.Area(); got != tt.want {
				t.Errorf("Circle.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_Perimeter(t *testing.T) {
	type fields struct {
		R float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
		{
			name:   "Case 1: Find Area",
			fields: fields{R: 7},
			want:   43.982297150257104,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Circle{
				R: tt.fields.R,
			}
			if got := c.Perimeter(); got != tt.want {
				t.Errorf("Circle.Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
