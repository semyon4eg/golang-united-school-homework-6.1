package golang_united_school_homework

import (
	"fmt"
	"reflect"
)

var errOutOfRange = fmt.Errorf("out of range")

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {

	if len(b.shapes) + 1 > b.shapesCapacity {
		return errOutOfRange
	}

	b.shapes = append(b.shapes, shape)

	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	
	if i >= len(b.shapes) {
		return nil, errOutOfRange
	}

	shape := b.shapes[i]

	return shape, nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {

	if i >= len(b.shapes) {
		return nil, errOutOfRange
	}

	shape := b.shapes[i]
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)

	return shape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	
	if i >= len(b.shapes) {
		return nil, fmt.Errorf("out of range")
	}

	old := b.shapes[i]
	b.shapes[i] = shape

	return old, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	
	sum := 0.0

	for _, shape := range b.shapes {
		sum += shape.CalcPerimeter()
	}

	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	
	sum := 0.0

	for _, shape := range b.shapes {
		sum += shape.CalcArea()
	}

	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	
	n := 0

	for i, shape := range b.shapes {
		if _, ok := shape.(Circle); ok {
			b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
			n++
		}
	}

	if n == 0 {
		return fmt.Errorf("no circles found error")
	}

	return nil
}
