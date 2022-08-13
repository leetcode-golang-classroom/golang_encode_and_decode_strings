package sol

import (
	"reflect"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	input := []string{"lint", "code", "love", "you"}
	for idx := 0; idx < b.N; idx++ {
		RunTest(input)
	}
}
func TestRunTest(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "[\"lint\",\"code\",\"love\",\"you\"]",
			args: args{input: []string{"lint", "code", "love", "you"}},
			want: []string{"lint", "code", "love", "you"},
		},
		{
			name: "[\"we\",\"say\",\":\",\"yes\"]",
			args: args{input: []string{"we", "say", ":", "yes"}},
			want: []string{"we", "say", ":", "yes"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RunTest(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RunTest() = %v, want %v", got, tt.want)
			}
		})
	}
}
