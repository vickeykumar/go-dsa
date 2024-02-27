package dp

import (
	"fmt"
	"testing"
)

func TestGet_LCS(t *testing.T) {
	// Test case 1: Strings have common characters
	string1 := "ABCBDAB"
	string2 := "BDCAB"
	expectedResult1 := 4 // "BCAB"
	if result1 := Get_LCS(string1, string2); result1 != expectedResult1 {
		t.Errorf("Test case 1 failed: string1=%s, string2=%s, expected %d but got %d", string1, string2, expectedResult1, result1)
	} else {
		fmt.Printf("Test case 1 passed: string1=%s, string2=%s, result=%d\n", string1, string2, result1)
	}

	// Test case 2: Strings have no common characters
	string3 := "ABC"
	string4 := "DEF"
	expectedResult2 := 0 // No common substring
	if result2 := Get_LCS(string3, string4); result2 != expectedResult2 {
		t.Errorf("Test case 2 failed: string1=%s, string2=%s, expected %d but got %d", string3, string4, expectedResult2, result2)
	} else {
		fmt.Printf("Test case 2 passed: string1=%s, string2=%s, result=%d\n", string3, string4, result2)
	}

	// Test case 3: Empty strings
	string5 := ""
	string6 := ""
	expectedResult3 := 0 // Empty strings have no common substring
	if result3 := Get_LCS(string5, string6); result3 != expectedResult3 {
		t.Errorf("Test case 3 failed: string1=%s, string2=%s, expected %d but got %d", string5, string6, expectedResult3, result3)
	} else {
		fmt.Printf("Test case 3 passed: string1=%s, string2=%s, result=%d\n", string5, string6, result3)
	}
}


func TestPrintLCS(t *testing.T) {
	testCases := []struct {
		name     string
		string1  string
		string2  string
		expected string
	}{
		{
			name:     "Example 1",
			string1:  "AGGTAB",
			string2:  "GXTXAYB",
			expected: "GTAB",
		},
		{
			name:     "Example 2",
			string1:  "ABCBDAB",
			string2:  "BDCABA",
			expected: "BDAB",
		},
		// Add more test cases as needed
		{
			name:     "Empty Strings",
			string1:  "",
			string2:  "",
			expected: "",
		},
		{
			name:     "No Common Subsequence",
			string1:  "ABC",
			string2:  "DEF",
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Print_LCS(tc.string1, tc.string2)
			if actual != tc.expected {
				t.Errorf("Test case %s failed: expected '%s' but got '%s' (string1: '%s', string2: '%s')",
					tc.name, tc.expected, actual, tc.string1, tc.string2)
			} else {
				fmt.Printf("Test case %s passed: expected '%s' and got '%s' (string1: '%s', string2: '%s')\n",
					tc.name, tc.expected, actual, tc.string1, tc.string2)
			}
		})
	}
}


func TestGet_LongestCommonSubstring(t *testing.T) {
	testCases := []struct {
		string1  string
		string2  string
		expected int
	}{
		// Test cases where the longest common substring is present
		{"ABCBDAB", "BDCABA", 2},  // Expected: 2 (BD)
		{"ABC", "XYZABC", 3},      // Expected: 3 (ABC)
		{"ABAB", "BABA", 3},       // Expected: 3 (BAB or ABA)

		// Test cases where there is no common substring
		{"XYZ", "ABC", 0},     // Expected: 0
		{"12345", "67890", 0}, // Expected: 0
		{"", "", 0},           // Expected: 0 (both empty strings)
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test case %d", i+1), func(t *testing.T) {
			result := Get_LongestCommonSubstring(tc.string1, tc.string2)
			if result != tc.expected {
				t.Errorf("Test case %d failed: expected %d, got %d", i+1, tc.expected, result)
			} else {
				fmt.Printf("Test case %d passed: string1='%s', string2='%s', longest common substring length=%d\n", i+1, tc.string1, tc.string2, result)
			}
		})
	}
}

func TestGet_ShortestCommonSupersequence(t *testing.T) {
	testCases := []struct {
		string1  string
		string2  string
		expected int
	}{
		// Both strings are empty
		{"", "", 0}, // Expected: 0

		// One string is empty and the other is non-empty
		{"", "XYZ", 3},   // Expected: 3 (shortest common supersequence is the non-empty string)
		{"ABC", "", 3},   // Expected: 3 (shortest common supersequence is the non-empty string)
		{"", "ABC", 3},   // Expected: 3 (shortest common supersequence is the non-empty string)
		{"XYZ", "", 3},   // Expected: 3 (shortest common supersequence is the non-empty string)
		{"", "ABCD", 4},  // Expected: 4 (shortest common supersequence is the non-empty string)
		{"ABCD", "", 4},  // Expected: 4 (shortest common supersequence is the non-empty string)

		// Both strings are non-empty
		{"ABC", "XYZ", 6},      // Expected: 6 (length of combined strings)
		{"ABC", "BCD", 4},      // Expected: 4 (shortest common supersequence is "ABCD")
		{"AGGTAB", "GXTXAYB", 9}, // Expected: 9 (shortest common supersequence is "AGGXTXAYB")
		{"ABCBDAB", "BDCABA", 9}, // Expected: 9 (shortest common supersequence is "ABDCBDAB")
		{"ABCD", "ACBCD", 5},   // Expected: 5 (shortest common supersequence is "ACBCD")
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test case %d", i+1), func(t *testing.T) {
			result := Get_ShortestCommonSupersequence(tc.string1, tc.string2)
			if result != tc.expected {
				t.Errorf("Test case %d failed: expected %d, got %d", i+1, tc.expected, result)
			} else {
				fmt.Printf("Test case %d passed: string1='%s', string2='%s', shortest common supersequence length=%d\n", i+1, tc.string1, tc.string2, result)
			}
		})
	}
}

func TestGet_LongestPalindromicSubseq(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		// Test cases where the longest palindromic subsequence is present
		{"BBABCBCAB", 7}, // Expected: 7 (BABCBAB)
		{"cbbd", 2},      // Expected: 2 (bb)
		{"abca", 3},      // Expected: 3 (aba)
		{"abcde", 1},     // Expected: 1 (a)

		// Test cases where the input string is empty
		{"", 0}, // Expected: 0

		// Additional test cases can be added here
	}

	for _, tc := range testCases {
		result := Get_LongestPalindromicSubseq(tc.input)
		if result != tc.expected {
			t.Errorf("Test case failed: input=%s, expected=%d, got=%d", tc.input, tc.expected, result)
		} else {
			t.Logf("Test case passed: input=%s, expected=%d, got=%d", tc.input, tc.expected, result)
		}
	}
}
