package commands

import "testing"

func TestCommandLeave_ParseArguments(t *testing.T) {
	type fields struct {
		Command    Command
		SlotNumber int
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
			fields{SlotNumber: 3},
			args{args: "3"},
			false,
		},
		{
			"TestCase 2: input multiple arguments, is invalid and error",
			fields{SlotNumber: 3},
			args{args: "3 7"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &CommandLeave{
				Command:    tt.fields.Command,
				SlotNumber: tt.fields.SlotNumber,
			}
			if err := this.ParseArguments(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("CommandLeave.ParseArguments() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCommandLeave_ValidateParams(t *testing.T) {
	type fields struct {
		Command    Command
		SlotNumber int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"TestCase 1: Only one argument (valid)",
			fields{SlotNumber: 3, Command: Command{Arguments: []string{"3"}}},
			true,
		},
		{
			"TestCase 2: Multiple arguments (invalid)",
			fields{SlotNumber: 3, Command: Command{Arguments: []string{"3", "7"}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &CommandLeave{
				Command:    tt.fields.Command,
				SlotNumber: tt.fields.SlotNumber,
			}
			if got := this.ValidateInput(); got != tt.want {
				t.Errorf("CommandLeave.ValidateInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
