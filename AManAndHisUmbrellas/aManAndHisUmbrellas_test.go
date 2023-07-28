package kata

import "testing"

func TestMinUmbrellas(t *testing.T) {
	type args struct {
		weather []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test-1", args{[]string{"cloudy"}}, 0},
		{"test-2", args{[]string{"rainy", "rainy", "rainy", "rainy"}}, 1},
		{"test-3", args{[]string{"overcast", "rainy", "clear", "thunderstorms"}}, 2},
		{"test-4", args{weather: []string{"rainy", "sunny", "thunderstorms", "rainy", "thunderstorms", "rainy", "sunny", "cloudy", "rainy", "thunderstorms"}}, 2},
		{"test-5", args{weather: []string{"clear", "thunderstorms", "sunny", "clear", "clear", "sunny", "cloudy", "thunderstorms", "clear", "windy", "thunderstorms", "sunny", "thunderstorms", "thunderstorms", "sunny", "sunny", "clear", "thunderstorms", "rainy", "clear", "clear", "sunny", "sunny", "rainy", "clear", "sunny", "windy", "thunderstorms", "windy", "clear", "sunny", "clear", "sunny", "clear", "clear", "clear", "sunny", "sunny", "sunny", "sunny", "rainy", "sunny", "cloudy", "thunderstorms", "thunderstorms", "sunny", "clear", "cloudy", "thunderstorms", "clear", "clear", "clear", "clear", "windy", "clear", "clear", "windy", "sunny", "clear", "clear", "clear", "clear", "sunny", "clear", "clear", "clear", "sunny", "clear", "clear", "clear", "thunderstorms", "windy", "clear", "sunny", "rainy", "clear", "clear", "clear", "sunny", "sunny", "sunny", "windy", "sunny", "windy", "windy", "thunderstorms", "windy", "clear", "clear", "thunderstorms", "sunny", "sunny", "sunny"}}, 4},
		{"test-6", args{weather: []string{"rainy", "rainy", "rainy", "rainy", "thunderstorms", "rainy"}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinUmbrellas(tt.args.weather); got != tt.want {
				t.Errorf("MinUmbrellas() = %v, want %v", got, tt.want)
			}
		})
	}
}
