package example3

import (
	"testing"

	"github.com/op/go-logging"
)

var (
	log = logging.MustGetLogger("example")
)

func Test_Multiple_1(t *testing.T) {
	result := multiple(1, 2)

	// log.Debug("multiple result: ", result)
	// log.Debug("what is testing.T? ", t)

	if result != 2 {
		t.Error("錯的結果！", result)
	} else {
		t.Log("結果正確！")
	}
}

func Test_multiple_2(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{1, 2}, 2},
		{name: "test2", args: args{2, 2}, want: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiple(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("multiple() = %v, want %v", got, tt.want)
			}
		})
	}
}
