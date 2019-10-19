package commands

import "testing"

func TestCommand_Run(t *testing.T) {
	type fields struct {
		Args []string
		Menu map[string]ICommand
	}
	type args struct {
		command string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"TestCase 1: Test create parking lot with capacity 6 (create_parking_lot 6)",
			fields{
				Menu: map[string]ICommand{"create_parking_lot": &CommandCreateParkingLot{}},
			},
			args{command: "create_parking_lot 6"},
			"Created a parking lot with 6 slots",
		},
		{
			"TestCase 2: Test undefined command (e.g foo)",
			fields{
				Menu: map[string]ICommand{"foo": nil},
			},
			args{command: "foo"},
			"foo: command not found\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Command{
				Arguments:   tt.fields.Args,
				CommandList: tt.fields.Menu,
			}
			if got := this.Run(tt.args.command); got != tt.want {
				t.Errorf("Command.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
