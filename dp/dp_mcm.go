package dp

import (
	_ "fmt"
	"go-dsa/util"
	"math"
)
/* family of problems based on MCM 
1. MCM
2. Print MCM
3. boolean parenthesization
4. Min/Max value of an Expr
5. Palindrome partitioning
6. scamble string
7. Egg Dropping problem
Pattern - partitioning/ breaking a problem/string/array and 
find the optimal solution from all the partition 
by recursively solving these 2.. partitions of bigger problem
*/

/*
Given the dimension of a sequence of matrices in an array arr[], 
where the dimension of the ith matrix is (arr[i-1] * arr[i]), 
the task is to find the most efficient way to multiply these matrices 
together such that the total number of element multiplications is minimum.
*/
// Matrix Ai has dimention pi-1 x pi
// return min cost of matrix multiplication 
func MatrixChainOrder(p []int, i, j int, dp [][]int) int {
	if i>=j {
		return 0
	}
	if dp[i][j] != -1 {
		// already calculated
		return dp[i][j]
	}
	mincost := math.MaxInt32
	for k:=i; k<j;k++ {
		mincost = util.Min(mincost, MatrixChainOrder(p, i, k, dp)+MatrixChainOrder(p, k+1,j, dp)+p[i-1]*p[k]*p[j])
	}
	dp[i][j] = mincost
	return mincost
}

func Ispalindrome(str string, i, j int) bool {
	for i<j {
		if str[i]!=str[j] {
			return false
		}
		i++
		j--
	}
	return true
}

/*
Given a string s, partition s such that every substring of the partition is a palindrome.

Return the minimum cuts needed for a palindrome partitioning of s
logic - lets say k is the partition point such that we get the minimum number of partition on left side and right side to make palindrome
*/
func MinPartitionPalindrome(str string, i, j int, dp [][]int) int {
	if i >= j {
		return 0 // 0 cost if empty or 1 string len
	}

	if dp[i][j] != -1 {
		// already calculated
		return dp[i][j]
	}

	if Ispalindrome(str, i, j) {
		dp[i][j] = 0
		return 0
	}

	min := math.MaxInt32
	for k := i; k < j; k++ {
		left := MinPartitionPalindrome(str, i, k, dp)
		right := MinPartitionPalindrome(str, k+1, j, dp)
		min = util.Min(min, 1+left+right)
	}
	dp[i][j] = min
	return min
}


/*
You are given k identical eggs and you have access to a building with n floors labeled from 1 to n.

You know that there exists a floor f where 0 <= f <= n such that any egg dropped at a floor higher 
than f will break, and any egg dropped at or below floor f will not break.

Each move, you may take an unbroken egg and drop it from any floor x (where 1 <= x <= n).
 If the egg breaks, you can no longer use it. However, if the egg does not break, you may reuse it in future moves.

Return the minimum number of moves that you need to determine with certainty what the value of f is.
logic- lets say at floor f egg breaks then subproblem will be of size (k-1, f-1),  if it doesn't break (k, n-f)
- further optimization to do binary search on f
*/

func minEggDrop(k int, n int, dp [][]int) int {
	if k ==1 || n==0 || n==1 {
		// if number of eggs = 1, try all possible
		return n
	}
	if dp[k][n]!=-1 {
		return dp[k][n]
	}
    minmoves := math.MaxInt32
    i, j := 1, n
    for i<=j {
        f := (i+j)/2 
        movesbreak := dp[k-1][f-1]
        if movesbreak == -1 {
            movesbreak = minEggDrop(k-1, f-1, dp)
            dp[k-1][f-1] = movesbreak
        }
        movesnobreak := dp[k][n-f]
        if movesnobreak == -1 {
            movesnobreak = minEggDrop(k, n-f, dp)
            dp[k][n-f] = movesnobreak
        }
        if movesnobreak>movesbreak {
            // go up as it gives larger answer
            i=f+1
        } else {
            j=f-1
        }
    	minmoves = util.Min(minmoves, 1+util.Max(movesbreak, movesnobreak))
    }
    dp[k][n] = minmoves
    return minmoves
}