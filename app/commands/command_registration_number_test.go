package commands

import (
	"testing"
)

func TestCommandRegistrationNumber_ParseArguments(t *testing.T) {
	type fields struct {
		Command  Command
		CarColor string
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
			fields{CarColor: "white"},
			args{args: "white"},
			false,
		},
		{
			"TestCase 2: input multiple arguments, is invalid and error",
			fields{CarColor: "white"},
			args{args: "white hello"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &CommandRegistrationNumber{
				Command:  tt.fields.Command,
				CarColor: tt.fields.CarColor,
			}
			if err := this.ParseArguments(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("CommandRegistrationNumber.ParseArguments() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCommandRegistrationNumber_ValidateParams(t *testing.T) {
	type fields struct {
		Command  Command
		CarColor string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"TestCase 1: Only one argument (valid)",
			fields{CarColor: "white", Command: Command{Arguments: []string{"1", "white"}}},
			true,
		},
		{
			"TestCase 2: Less arguments (invalid)",
			fields{CarColor: "white", Command: Command{Arguments: []string{"1"}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &CommandRegistrationNumber{
				Command:  tt.fields.Command,
				CarColor: tt.fields.CarColor,
			}
			if got := this.ValidateInput(); got != tt.want {
				t.Errorf("CommandRegistrationNumber.ValidateInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
