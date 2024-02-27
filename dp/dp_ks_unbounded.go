package dp

import (
	_ "fmt"
	"go-dsa/util"
	_ "math"
)


/* set of problems with pattern of unbounded knapsack
 * where we have choice of selecting same item multiple times.
 * * unbounded knapsack
 * * rod cutting problem
 * * coin change I
 * * Coin change II
 */

/*
given a value and weight array and a bag of limited capacity
maximize the profit, note: same item can be selected multiple times.
*/
 func Knapsack_unbounded (V, W []int, capacity int) int {
 	profit := make([][]int, len(V)+1)
 	for i:=0;i<len(V)+1;i++ {
 		profit[i]=make([]int, capacity+1)
 		for j:=0;j<capacity+1;j++ {
 			//base case
 			if i==0 || j==0 {
 				if i==0 {
 					profit[i][j]=0
 				}
 				if i==0 && j==0 {
 					profit[i][j]=1
 					// or profit[0][0]=1
 				}

 			} else {
 				if j<get_ith_element(W,i) {
 					profit[i][j]=profit[i-1][j]
 				} else {
 					profit[i][j]=util.Max((get_ith_element(V,i)+profit[i][j- get_ith_element(W,i)]), profit[i-1][j])
 				}
 			}
 		}
 	}
 	return profit[len(V)][capacity]
 }



 /*
Given a rod of length n inches and an array of prices that includes prices of all pieces of size smaller than n. 
Determine the maximum value obtainable by cutting up the rod and selling the pieces. 
For example, if the length of the rod is 8 and the values of different pieces are given as the following, 
then the maximum obtainable value is 22 (by cutting in two pieces of lengths 2 and 6) 

length   | 1   2   3   4   5   6   7   8  
--------------------------------------------
price    | 1   5   8   9  10  17  17  20
And if the prices are as follows, then the maximum obtainable value is 24 (by cutting in eight pieces of length 1) 

length   | 1   2   3   4   5   6   7   8  
--------------------------------------------
price    | 3   5   8   9  10  17  17  20
 */

func Rodcut(price []int, length int) int {
	profit := make([][]int, len(price)+1)
	for i:=0;i<=len(price);i++ {
		profit[i]=make([]int,length+1)
		for j:=0;j<=length;j++ {
			// base case
			if i==0 || j==0 {
				if i==0 {
					profit[i][j] = 0
				}
				if i==0 && j==0 {
					profit[i][j]=1
				}
			} else {
				// generic case
				if j < i {
					profit[i][j]=profit[i-1][j]
				} else {
					profit[i][j]=util.Max(get_ith_element(price,i)+profit[i][j-i], profit[i-1][j])
				}
			}
		}
	}
	return profit[len(price)][length]
}