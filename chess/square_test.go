package chess

import "testing"

func TestSquare_File(t *testing.T) {
	tests := []struct {
		s    Square
		want File
	}{
		{A1, FileA}, {B1, FileB}, {C1, FileC}, {D1, FileD}, {E1, FileE}, {F1, FileF}, {G1, FileG}, {H1, FileH},
		{A2, FileA}, {B2, FileB}, {C2, FileC}, {D2, FileD}, {E2, FileE}, {F2, FileF}, {G2, FileG}, {H2, FileH},
		{A3, FileA}, {B3, FileB}, {C3, FileC}, {D3, FileD}, {E3, FileE}, {F3, FileF}, {G3, FileG}, {H3, FileH},
		{A4, FileA}, {B4, FileB}, {C4, FileC}, {D4, FileD}, {E4, FileE}, {F4, FileF}, {G4, FileG}, {H4, FileH},
		{A5, FileA}, {B5, FileB}, {C5, FileC}, {D5, FileD}, {E5, FileE}, {F5, FileF}, {G5, FileG}, {H5, FileH},
		{A6, FileA}, {B6, FileB}, {C6, FileC}, {D6, FileD}, {E6, FileE}, {F6, FileF}, {G6, FileG}, {H6, FileH},
		{A7, FileA}, {B7, FileB}, {C7, FileC}, {D7, FileD}, {E7, FileE}, {F7, FileF}, {G7, FileG}, {H7, FileH},
		{A8, FileA}, {B8, FileB}, {C8, FileC}, {D8, FileD}, {E8, FileE}, {F8, FileF}, {G8, FileG}, {H8, FileH},
	}

	for _, tc := range tests {
		got := tc.s.File()
		if got != tc.want {
			t.Errorf("File(%v) = %v, want %v", tc.s, got, tc.want)
		}
	}
}

func TestSquare_Rank(t *testing.T) {
	tests := []struct {
		s    Square
		want Rank
	}{
		{A1, Rank1}, {B1, Rank1}, {C1, Rank1}, {D1, Rank1}, {E1, Rank1}, {F1, Rank1}, {G1, Rank1}, {H1, Rank1},
		{A2, Rank2}, {B2, Rank2}, {C2, Rank2}, {D2, Rank2}, {E2, Rank2}, {F2, Rank2}, {G2, Rank2}, {H2, Rank2},
		{A3, Rank3}, {B3, Rank3}, {C3, Rank3}, {D3, Rank3}, {E3, Rank3}, {F3, Rank3}, {G3, Rank3}, {H3, Rank3},
		{A4, Rank4}, {B4, Rank4}, {C4, Rank4}, {D4, Rank4}, {E4, Rank4}, {F4, Rank4}, {G4, Rank4}, {H4, Rank4},
		{A5, Rank5}, {B5, Rank5}, {C5, Rank5}, {D5, Rank5}, {E5, Rank5}, {F5, Rank5}, {G5, Rank5}, {H5, Rank5},
		{A6, Rank6}, {B6, Rank6}, {C6, Rank6}, {D6, Rank6}, {E6, Rank6}, {F6, Rank6}, {G6, Rank6}, {H6, Rank6},
		{A7, Rank7}, {B7, Rank7}, {C7, Rank7}, {D7, Rank7}, {E7, Rank7}, {F7, Rank7}, {G7, Rank7}, {H7, Rank7},
		{A8, Rank8}, {B8, Rank8}, {C8, Rank8}, {D8, Rank8}, {E8, Rank8}, {F8, Rank8}, {G8, Rank8}, {H8, Rank8},
	}

	for _, tc := range tests {
		got := tc.s.Rank()
		if got != tc.want {
			t.Errorf("Rank(%v) = %v, want %v", tc.s, got, tc.want)
		}
	}
}
