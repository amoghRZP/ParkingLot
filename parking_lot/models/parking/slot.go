package parking

import (
	"fmt"

	car "../car"
)

// A slot in parking lot
type Slot struct {
	Number int
	Car    *car.Car
}

func NewSlot() *Slot {
	return &Slot{}
}

func (s *Slot) Occupy(cr car.Car) error {
	if s.Car != nil {
		return fmt.Errorf("slot already occupied")
	}
	s.Car = &cr
	return nil
}

func (s *Slot) IsFree() bool {
	return s.Car == nil
}

// Free a slot from vehicle
func (s *Slot) Free() {
	s.Car = nil
}

func (s *Slot) GetCarNumberInSlot() string {
	return s.Car.Number
}
