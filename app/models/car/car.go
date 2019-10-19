package car

// TODO:: TO create a vehicle interface where different kind of vehicles will be implementing this (car, bike, bus etc)
// car model is used for car entity  which will be inside parking lot and occupy space
type Car struct {
	Number string
	Color  string
}

// Constructor for car model
func New(number, color string) *Car {
	return &Car{
		Number: number,
		Color:  color,
	}
}
