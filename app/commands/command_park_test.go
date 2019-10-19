package commands

import (
	"testing"

	"github.com/amogmish/parkingLot/app/models/car"
)

func TestCommandPark_ParseArgs(t *testing.T) {
	type fields struct {
		Command Command
		Car     *car.Car
	}
	type args struct {
		args string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"TestCase 1: one argument in cmd. should parse agruments",
			fields{Car: car.New("1", "white")},
			args{args: "1 white"},
			false,
		},
		{
			"TestCase 2: input multiple arguments, is invalid and error",
			fields{Car: car.New("1", "white")},
			args{args: "1 white hello"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &CommandPark{
				Command: tt.fields.Command,
				Car:     tt.fields.Car,
			}
			if err := this.ParseArguments(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("CommandPark.ParseArgs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCommandPark_ValidateParams(t *testing.T) {
	type fields struct {
		Command Command
		Car     *car.Car
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"TestCase 1: Only one argument (valid)",
			fields{Car: car.New("1", "white"), Command: Command{Arguments: []string{"1", "white"}}},
			true,
		},
		{
			"TestCase 2: Less arguments (invalid)",
			fields{Car: car.New("1", "white"), Command: Command{Arguments: []string{"1"}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &CommandPark{
				Command: tt.fields.Command,
				Car:     tt.fields.Car,
			}
			if got := this.ValidateInput(); got != tt.want {
				t.Errorf("CommandPark.ValidateInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
