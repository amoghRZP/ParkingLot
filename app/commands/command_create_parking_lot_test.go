package commands

import "testing"

func TestCommandCreateParkingLot_ParseArguments(t *testing.T) {
	type fields struct {
		Command  Command
		Capacity int
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
			"TestCase 1: one argument in cmd. Input [cmd] 3 produce commandparking object with capacity=3",
			fields{Capacity: 3},
			args{args: "3"},
			false,
		},
		{
			"TestCase 2: input multiple arguments, is invalid and error",
			fields{Capacity: 3},
			args{args: "3 7"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &CommandCreateParkingLot{
				Command:          tt.fields.Command,
				CreationCapacity: tt.fields.Capacity,
			}
			if err := this.ParseArguments(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("CommandCreateParkingLot.ParseArguments() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCommandCreateParkingLot_ValidateParams(t *testing.T) {
	type fields struct {
		Command  Command
		Capacity int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"TestCase 1: Only one argument (actualy valid)",
			fields{Capacity: 3, Command: Command{Arguments: []string{"3"}}},
			true,
		},
		{
			"TestCase 2: Multiple arguments (actualy invalid)",
			fields{Capacity: 3, Command: Command{Arguments: []string{"3", "7"}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &CommandCreateParkingLot{
				Command:          tt.fields.Command,
				CreationCapacity: tt.fields.Capacity,
			}
			if got := this.ValidateInput(); got != tt.want {
				t.Errorf("CommandCreateParkingLot.ValidateInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommandCreateParkingLot_Run(t *testing.T) {
	type fields struct {
		Command  Command
		Capacity int
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			"TestCase 1",
			fields{Capacity: 3},
			"Created a parking lot with 3 slots",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &CommandCreateParkingLot{
				Command:          tt.fields.Command,
				CreationCapacity: tt.fields.Capacity,
			}
			got, err := this.Run()
			if (err != nil) != tt.wantErr {
				t.Errorf("CommandCreateParkingLot.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CommandCreateParkingLot.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
