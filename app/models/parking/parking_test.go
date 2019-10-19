package parking

import (
	"reflect"
	"testing"

	"github.com/amogmish/parkingLot/app/models/car"
	"github.com/amogmish/parkingLot/app/models/slot"
)

func TestNew(t *testing.T) {
	type args struct {
		capacity int
	}

	tests := []struct {
		name string
		args args
		want *Parking
	}{
		{
			"TestCase 1: Create with capacity=3 produce 3 slots with Number 1 to 3 and null Car",
			args{capacity: 3},
			&Parking{
				Capacity: 3,
				Slots: []*slot.Slot{
					{
						Number: 1,
						Car:    nil,
					},
					{
						Number: 2,
						Car:    nil,
					},
					{
						Number: 3,
						Car:    nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.capacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\x1b[31;1mNew() = %v, want %v\x1b[0m", got, tt.want)
			}
		})
	}
}
func TestParking_FindNearestSlot(t *testing.T) {
	type fields struct {
		Capacity int
		Slots    []*slot.Slot
	}
	tests := []struct {
		name    string
		fields  fields
		want    *slot.Slot
		wantErr bool
	}{
		{
			"TestCase 1. Find nearest free slot with nearest filled slot in Number 2. Want slot with Number=1",
			fields{
				Capacity: 3,
				Slots: []*slot.Slot{
					{
						Number: 1,
						Car:    nil,
					},
					{
						Number: 2,
						Car:    &car.Car{Number: "BE4508GE", Color: "Red"},
					},
					{
						Number: 3,
						Car:    nil,
					},
				},
			},
			&slot.Slot{
				Number: 1,
				Car:    nil,
			},
			false,
		},
		{
			"TestCase 2",
			fields{
				Capacity: 3,
				Slots: []*slot.Slot{
					{
						Number: 1,
						Car:    &car.Car{Number: "BE4508GE", Color: "Red"},
					},
					{
						Number: 2,
						Car:    nil,
					},
					{
						Number: 3,
						Car:    nil,
					},
				},
			},
			&slot.Slot{
				Number: 2,
				Car:    nil,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Parking{
				Capacity: tt.fields.Capacity,
				Slots:    tt.fields.Slots,
			}
			got, err := this.FindNearestSlot()
			if (err != nil) != tt.wantErr {
				t.Errorf("\x1b[31;1mParking.FindNearestSlot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\x1b[31;1mParking.FindNearestSlot() = %v, want %v\x1b[0m", got, tt.want)
			}
		})
	}
}

func TestParking_AddCar(t *testing.T) {
	type fields struct {
		Capacity int
		Slots    []*slot.Slot
	}
	type args struct {
		cr car.Car
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *slot.Slot
		wantErr bool
	}{
		{
			"TestCase 1: Parking slots is not full",
			fields{
				Capacity: 3,
				Slots: []*slot.Slot{
					{
						Number: 1,
						Car:    nil,
					},
					{
						Number: 2,
						Car:    nil,
					},
					{
						Number: 3,
						Car:    nil,
					},
				},
			},
			args{cr: car.Car{Number: "BE4508GE", Color: "Red"}},
			&slot.Slot{
				Number: 1,
				Car:    &car.Car{Number: "BE4508GE", Color: "Red"},
			},
			false,
		},
		{
			"TestCase 2: Parking slots is full",
			fields{
				Capacity: 3,
				Slots: []*slot.Slot{
					{
						Number: 1,
						Car:    &car.Car{Number: "BE1000GE", Color: "Red"},
					},
					{
						Number: 2,
						Car:    &car.Car{Number: "BE2000GE", Color: "Red"},
					},
					{
						Number: 3,
						Car:    &car.Car{Number: "BE3000GE", Color: "Red"},
					},
				},
			},
			args{cr: car.Car{Number: "BE4508GE", Color: "Red"}},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Parking{
				Capacity: tt.fields.Capacity,
				Slots:    tt.fields.Slots,
			}
			got, err := this.AddCar(tt.args.cr)
			if (err != nil) != tt.wantErr {
				t.Errorf("\x1b[31;1mParking.AddCar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\x1b[31;1mParking.AddCar() = %v, want %v\x1b[0m", got, tt.want)
			}
		})
	}
}

func TestParking_GetFilledSlots(t *testing.T) {
	type fields struct {
		Capacity int
		Slots    []*slot.Slot
	}
	tests := []struct {
		name            string
		fields          fields
		wantFilledSlots []*slot.Slot
	}{
		{
			"TestCase 1: ",
			fields{
				Capacity: 3,
				Slots: []*slot.Slot{
					{
						Number: 1,
						Car:    &car.Car{Number: "BE1000GE", Color: "Red"},
					},
					{
						Number: 2,
						Car:    nil,
					},
					{
						Number: 3,
						Car:    &car.Car{Number: "BE3000GE", Color: "Red"},
					},
				},
			},
			[]*slot.Slot{
				{
					Number: 1,
					Car:    &car.Car{Number: "BE1000GE", Color: "Red"},
				},
				{
					Number: 3,
					Car:    &car.Car{Number: "BE3000GE", Color: "Red"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Parking{
				Capacity: tt.fields.Capacity,
				Slots:    tt.fields.Slots,
			}
			if gotFilledSlots := this.GetOccupiedSlots(); !reflect.DeepEqual(gotFilledSlots, tt.wantFilledSlots) {
				t.Errorf("\x1b[31;1mParking.GetOccupiedSlots() = %v, want %v\x1b[0m", gotFilledSlots, tt.wantFilledSlots)
			}
		})
	}
}

func TestParking_GetSlotsByCarColor(t *testing.T) {
	type fields struct {
		Capacity int
		Slots    []*slot.Slot
	}
	type args struct {
		carColor string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantSlots []*slot.Slot
	}{
		{
			"TestCase 1:",
			fields{
				Capacity: 3,
				Slots: []*slot.Slot{
					{
						Number: 1,
						Car:    &car.Car{Number: "BE1000GE", Color: "Red"},
					},
					{
						Number: 2,
						Car:    &car.Car{Number: "BE2000GE", Color: "Blue"},
					},
					{
						Number: 3,
						Car:    &car.Car{Number: "BE3000GE", Color: "red"},
					},
				},
			},
			args{carColor: "red"},
			[]*slot.Slot{
				{
					Number: 1,
					Car:    &car.Car{Number: "BE1000GE", Color: "Red"},
				},
				{
					Number: 3,
					Car:    &car.Car{Number: "BE3000GE", Color: "red"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Parking{
				Capacity: tt.fields.Capacity,
				Slots:    tt.fields.Slots,
			}
			if gotSlots := this.GetSlotsByCarColor(tt.args.carColor); !reflect.DeepEqual(gotSlots, tt.wantSlots) {
				t.Errorf("\x1b[31;1mParking.GetSlotsByCarColor() = %v, want %v\x1b[0m", gotSlots, tt.wantSlots)
			}
		})
	}
}
