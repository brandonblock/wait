package wait

import (
	"testing"
	"time"
)

func TestWaiter_StoreEvent(t *testing.T) {
	type fields struct {
		requests <-chan string
		limiter  <-chan time.Time
	}
	type args struct {
		payload string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "test",
			args: args{
				payload: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := NewWaiter()
			for i := 0; i < 10; i++ {
				w.StoreEvent(tt.args.payload)
			}
		})
	}
}
