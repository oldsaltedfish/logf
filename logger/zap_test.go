package logger

import (
	"go.uber.org/zap"
	"reflect"
	"testing"
)

func TestLogger_Debug(t *testing.T) {
	type fields struct {
		sugar *zap.SugaredLogger
	}
	type args struct {
		i interface{}
	}
	l, _ := zap.NewDevelopment()
	sugar := l.Sugar()
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "test",
			fields: fields{sugar: sugar},
			args: args{
				i: "test debug log",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := ZapLogger{
				logger: tt.fields.sugar,
			}
			l.Debug(tt.args.i)
		})
	}
}

func TestLogger_Debugf(t *testing.T) {
	type fields struct {
		sugar *zap.SugaredLogger
	}
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := ZapLogger{
				logger: tt.fields.sugar,
			}
			l.Debugf(tt.args.format, tt.args.a...)
		})
	}
}

func TestLogger_Error(t *testing.T) {
	type fields struct {
		sugar *zap.SugaredLogger
	}
	type args struct {
		i interface{}
	}
	l, _ := zap.NewProduction()
	sugar := l.Sugar()
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "error log",
			fields: fields{sugar: sugar},
			args: args{
				i: "error log",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := ZapLogger{
				logger: tt.fields.sugar,
			}
			l.Error(tt.args.i)
		})
	}
}

func TestLogger_Errorf(t *testing.T) {
	type fields struct {
		sugar *zap.SugaredLogger
	}
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := ZapLogger{
				logger: tt.fields.sugar,
			}
			l.Errorf(tt.args.format, tt.args.a...)
		})
	}
}

func TestLogger_Info(t *testing.T) {
	type fields struct {
		sugar *zap.SugaredLogger
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := ZapLogger{
				logger: tt.fields.sugar,
			}
			l.Info(tt.args.i)
		})
	}
}

func TestLogger_Infof(t *testing.T) {
	type fields struct {
		sugar *zap.SugaredLogger
	}
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := ZapLogger{
				logger: tt.fields.sugar,
			}
			l.Infof(tt.args.format, tt.args.a...)
		})
	}
}

func TestLogger_Warn(t *testing.T) {
	type fields struct {
		sugar *zap.SugaredLogger
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := ZapLogger{
				logger: tt.fields.sugar,
			}
			l.Warn(tt.args.i)
		})
	}
}

func TestLogger_Warnf(t *testing.T) {
	type fields struct {
		sugar *zap.SugaredLogger
	}
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := ZapLogger{
				logger: tt.fields.sugar,
			}
			l.Warnf(tt.args.format, tt.args.a...)
		})
	}
}

func TestNewLogger(t *testing.T) {
	type args struct {
		sugar *zap.SugaredLogger
	}
	tests := []struct {
		name string
		args args
		want *ZapLogger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewZapLogger(tt.args.sugar); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}
