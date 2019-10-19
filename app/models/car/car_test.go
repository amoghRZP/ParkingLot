package car

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		number string
		color  string
	}
	tests := []struct {
		name string
		args args
		want *Car
	}{
		{
			"TestCase 1",
			args{number: "BE4508GE", color: "Red"},
			&Car{Number: "BE4508GE", Color: "Red"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.number, tt.args.color); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
