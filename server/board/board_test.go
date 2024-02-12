package board

import "testing"

func TestBoard_CalculateDepth(t *testing.T) {
	type fields struct {
		Grid [3][3]int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := &Board{
				Grid: tt.fields.Grid,
			}
			if got := board.CalculateDepth(); got != tt.want {
				t.Errorf("Board.CalculateDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
