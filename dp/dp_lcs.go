package dp

import (
	_ "fmt"
	"go-dsa/util"
)

/*
  set of DPs based on finding Longest common subsequence(LCS)
  * LCS
  * longest common substring
  * Print LCS
  * shortest common supersequence (SCS)
  * Print SCS
  * min edit distance
  * Longest Palindromic subsequence
  pattern - two/one strings given, we have to find optimal solution from both of the string. based on condition.
  - optimal solution for ith charachter depends on previous charachters (pick or not pick)
*/

func get_ith_char(V string, i int) rune {
	if i==0 {
		return ' '
	}
	return rune(V[i-1])
}

/*
given two strings, find length of the longest common subsequence
*/
func Get_LCS(string1 string, string2 string) int {
	lcs := make([][]int, len(string1)+1)
	for i:=0;i<=len(string1); i++ {
		lcs[i] = make([]int, len(string2)+1)
		for j:=0;j<=len(string2);j++ {
			// base case
			if i==0 || j==0 {
				lcs[i][j] = 0
			} else {
				// general case
				if get_ith_char(string1,i) == get_ith_char(string2, j) {
					lcs[i][j]= lcs[i-1][j-1]+1
				} else {
					lcs[i][j] = util.Max(lcs[i-1][j], lcs[i][j-1])
				}
			}
		}
	}
	return lcs[len(string1)][len(string2)]
}


/*
given two strings, print the longest common subsequence
*/
func Print_LCS(string1 string, string2 string) string {
	lcs := make([][]int, len(string1)+1)
	for i:=0;i<=len(string1); i++ {
		lcs[i] = make([]int, len(string2)+1)
		for j:=0;j<=len(string2);j++ {
			// base case
			if i==0 || j==0 {
				lcs[i][j] = 0
			} else {
				// general case
				if get_ith_char(string1,i) == get_ith_char(string2, j) {
					lcs[i][j]= lcs[i-1][j-1]+1
				} else {
					lcs[i][j] = util.Max(lcs[i-1][j], lcs[i][j-1])
				}
			}
		}
	}

	var lcs_str []rune
	i:=len(string1)
	j:=len(string2)
	for i>0&&j>0 { // keep going till any of the substring size becomes 0
		if get_ith_char(string1,i) == get_ith_char(string2, j) {
			// print
			lcs_str = append(lcs_str,get_ith_char(string1,i))
			i--
			j--
			// decrease the problem size after printing
		} else {
			// not equal, detect the max side and loop
			if lcs[i-1][j]>lcs[i][j-1] {
				i--
			} else {
				// string 2, j side is max
				j--
			}
		}
	}
	return util.ReverseString(string(lcs_str))
}


/*
given two strings, find length of the longest common substring:

- idea behind the solution is if the substring discontinues, reinitialize the solution from zero,
i.e if get_ith_char(string1,i) != get_ith_char(string2, j) lcs[i][j] =0
*/
func Get_LongestCommonSubstring(string1 string, string2 string) int {
	longest := 0
	lcs := make([][]int, len(string1)+1)
	for i:=0;i<=len(string1); i++ {
		lcs[i] = make([]int, len(string2)+1)
		for j:=0;j<=len(string2);j++ {
			// base case
			if i==0 || j==0 {
				lcs[i][j] = 0
			} else {
				// general case
				if get_ith_char(string1,i) == get_ith_char(string2, j) {
					lcs[i][j]= lcs[i-1][j-1]+1
					if lcs[i][j]>longest {
						longest = lcs[i][j]
					}
				} else {
					lcs[i][j] = 0 // reinitialize from 0
				}
			}
		}
	}
	return longest
}

/*
Given two sequences X = < x1,...,xm > and Y = < y1,...,yn >, 
a sequence U = < u1,...,uk > is a common supersequence of X and Y if items can be removed from U to produce X and Y.

A shortest common supersequence (SCS) is a common supersequence of minimal length.
- supersequence contains characters from both the string X and Y, 
- idea behind solution is concatenation of the two strings is definitely a supersequence, but it is not the minimal one
- can we optimize it further??
- yes, what if some elements are repeated in this superseq. what ?
- it should be longest common subsequence of the two.
- so len of SCS is SUM(X+Y)-len(LCS)
*/
func Get_ShortestCommonSupersequence(string1 string, string2 string) int {
	lcs := Get_LCS(string1, string2)
	sum := len(string1)+len(string2)
	return sum-lcs
}

/*
Longest Palindromic subsequence:
given two strings, find length of the Longest Palindromic subsequence
- idea behind solution is a palindrome is a string which is same as its reverse.
- can we achieve this  using LCS with its reverse ??
*/
func Get_LongestPalindromicSubseq(string1 string) int {
	rev := util.ReverseString(string1)
	return Get_LCS(string1, rev)
}
