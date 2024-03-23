package dp

import (
	"testing"
)

func TestMatrixChainOrder(t *testing.T) {
	tests := []struct {
		p        []int
		expected int
	}{
		// Test cases with different sizes of input matrices
		{[]int{10, 20, 30}, 6000},          // (A1 A2)A3 = (10x20)*(20x30) = 6000
		{[]int{40, 20, 30, 10, 30}, 26000}, // ((A1(A2A3))A4) = ((40x20)*(20x30))*(30x10) = 26000
		{[]int{10, 20, 30, 40, 30}, 30000}, // (A1(A2(A3A4))) = (10x20)*((20x30)*(30x40)) = 30000
	}

	for _, test := range tests {
		dp := make([][]int, len(test.p))
		for i,_:=range dp {
			dp[i] = make([]int, len(test.p))
			for j,_ := range dp[i] {
				dp[i][j] = -1
			}
		}
		actual := MatrixChainOrder(test.p, 1, len(test.p)-1, dp)
		if actual != test.expected {
			t.Errorf("For p=%v, expected=%d, but got=%d", test.p, test.expected, actual)
		} else {
			t.Logf("For p=%v, expected=%d, got=%d (Success)", test.p, test.expected, actual)
		}
	}
}

func TestMinPartitionPalindrome(t *testing.T) {
	tests := []struct {
		str      string
		expected int
	}{
		{"aab", 1},    // "aa"|"b"
		{"abbab", 1},  // "a"|"bb"|"a"|"b"
		{"abcba", 0},  // No partition needed
		{"banana", 1}, // "b"|"anana"
	}

	for _, test := range tests {
		dp := make([][]int, len(test.str))
		for i := range dp {
			dp[i] = make([]int, len(test.str))
			for j := range dp[i] {
				dp[i][j] = -1
			}
		}
		actual := MinPartitionPalindrome(test.str, 0, len(test.str)-1, dp)
		if actual != test.expected {
			t.Errorf("For str=%s, expected=%d, but got=%d", test.str, test.expected, actual)
		} else {
			t.Logf("For str=%s, expected=%d, got=%d (Success)", test.str, test.expected, actual)
		}
	}
}

func TestMinEggDrop(t *testing.T) {
	tests := []struct {
		k        int
		n        int
		expected int
	}{
		{1, 10, 10},   // 1 egg, 10 floors -> Try all possible floors
		{2, 6, 3},     // 2 eggs, 6 floors -> Optimal strategy: Drop from floor 3, if broken, drop from 1st or 2nd floor; if not, drop from 4th, 5th, or 6th
		{3, 14, 4},    // 3 eggs, 14 floors -> Optimal strategy: Drop from floor 4, 7, 11, if broken, use binary search between previous floor and current one.
		{9, 10000, 14}, // 4 eggs, 10000 floors -> Testing for large values
	}

	for _, test := range tests {
		dp := make([][]int, test.k+1)
		for i := range dp {
			dp[i] = make([]int, test.n+1)
			for j := range dp[i] {
				dp[i][j] = -1
			}
		}
		actual := minEggDrop(test.k, test.n, dp)
		if actual != test.expected {
			t.Errorf("For k=%d, n=%d, expected=%d, but got=%d", test.k, test.n, test.expected, actual)
		} else {
			t.Logf("For k=%d, n=%d, expected=%d, got=%d (Success)", test.k, test.n, test.expected, actual)
		}
	}
}