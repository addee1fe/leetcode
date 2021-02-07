package solution

import "testing"

func Test_canFinish(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{2, [][]int{{1, 0}}, true},
		{2, [][]int{{1, 0}}, true},
	}
	for _, tt := range tests {
		if got := canFinish(tt.numCourses, tt.prerequisites); got != tt.want {
			t.Errorf("canFinish() = %v, want %v", got, tt.want)
		}
	}
}
