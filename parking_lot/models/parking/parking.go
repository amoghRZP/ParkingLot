package parking

import (
	"fmt"

	"../car"
)

// Parking is main component, for data structure a parking lot
type Parking struct {
	Capacity int
	Slots    []*Slot
}

var (
	// In memory DB to get savedParking parking data
	savedParking *Parking
	startNumber  = 1
)

func New(capacity int) *Parking {
	parking := new(Parking)
	parking.Capacity = capacity
	parking.Slots = make([]*Slot, capacity)
	idx := startNumber
	for i := range parking.Slots {
		parking.Slots[i] = &Slot{Number: int(idx)}
		idx++
	}
	parking.Save()
	return parking
}

// Get for savedParking data
func Get() *Parking {
	return savedParking
}

func (p *Parking) Save() {
	savedParking = p
}

func (p *Parking) FindNearestSlot() (*Slot, error) {
	for _, sl := range p.Slots {
		if sl.IsFree() {
			return sl, nil
		}
	}
	return nil, fmt.Errorf("No Space Available")
}

func (p *Parking) AddCar(cr car.Car) (*Slot, error) {
	sl, err := p.FindNearestSlot()
	if err != nil {
		return nil, err
	}
	if err := sl.Occupy(cr); err != nil {
		return nil, err
	}
	return sl, nil
}

func (p *Parking) GetSlotByCarNumber(carNumber string) (slots *Slot) {
	for _, sl := range p.Slots {
		if !sl.IsFree() {
			if sl.GetCarNumberInSlot() == carNumber {
				slots = sl
				return
			}
		}
	}
	return
}

func (p *Parking) RemoveCar(cr car.Car) {
	for i, sl := range p.Slots {
		if !sl.IsFree() {
			p.Slots[i].Free()
		}
	}
}

func (p *Parking) RemoveCarFromSlot(slotNumber int) error {
	for i, sl := range p.Slots {
		if sl.Number == slotNumber {
			p.Slots[i].Car = nil
			return nil
		}
	}
	return fmt.Errorf("can't find %d slot", slotNumber)
}

func (p *Parking) GetOccupiedSlots() (filledSlots []*Slot) {
	for _, sl := range p.Slots {
		if !sl.IsFree() {
			filledSlots = append(filledSlots, sl)
		}
	}
	return
}
