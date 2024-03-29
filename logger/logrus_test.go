package logger

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLogrusLogger_Debug(t *testing.T) {
	type fields struct {
		logrus *logrus.Logger
	}
	type args struct {
		i interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			"1",
			fields{
				logrus: logrus.New(),
			},
			args{
				i: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LogrusLogger{
				logger: tt.fields.logrus,
			}
			l.Error(tt.args.i)
		})
	}
}
