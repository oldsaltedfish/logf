package logger

import (
	"github.com/rs/zerolog"
	"os"
	"testing"
)

func TestZeroLogger_Error(t *testing.T) {
	type fields struct {
		logger *zerolog.Logger
	}
	type args struct {
		i interface{}
	}
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "test",
			fields: fields{
				logger: &logger,
			},
			args: args{
				i: struct {
					A string
					B int
				}{
					A: "test",
					B: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := ZeroLogger{
				logger: tt.fields.logger,
			}
			z.Errorf("aa %d", 1)
		})
	}
}
