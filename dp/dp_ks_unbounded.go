package dp

import (
	"fmt"
	"go-dsa/util"
	"math"
)


/* set of problems with pattern of unbounded knapsack
 * where we have choice of selecting same item multiple times.
 * * unbounded knapsack
 * * rod cutting problem
 * * coin change I - max/total number of ways to reach k
 * * Coin change II - min number of coins
 * Pattern - where we have set of items and we have to achieve an optimal solumtion
 * by picking/not picking ith item to achieve a target/optimal solution,  there is no limitation on picking ith item
 *
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
 				profit[i][j]=0 // either item size or capacity is zero profit will be be zero
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
					profit[i][j] = 0 // 
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


/*
Given an integer array of coins[ ] of size N representing different types of denominations and an integer sum, 
the task is to count the number of coins required to make a given value sum.  

Note: Assume that you have an infinite supply of each type of coin. // order doesn't matter
- idea behind solution, coins[i] is the weight array , no value array, there is a sum array to achieve by using these weight.

lets say T(i,j) is the solution for total num of ways  to pick the coin and reacha sum of sum
recursion can be written as T(i,j) = T(i-1,j) {{dont pic the ith item,  sum j is still intact}} + 
T(i, j-coin(ith)) {{pic the ith coin i can still pick it in future, sum j is reduces by ith coins weight}}
*/

func Noways_CoinChange(coins []int, sum int) int {
	sol := make([][]int, len(coins)+1)
	for i:=0; i<=len(coins); i++ {
		sol[i] = make([]int, sum+1)  // sol[len(coins)][sum] will have the sol/ no of ways to pick all coins and with sum sum.
		for j:=0; j<=sum; j++ {
			if i==0 {
				//base cases trying to achieve a sum without ny element or
				// else if sum j==0 , even if item i is there, i can have 1 ways by not picking this up
				sol[i][j]=0
				if i==0 && j==0 {
					sol[i][j] = 1
				}
			} else {
				// generic cases
				if get_ith_element(coins,i) <= j {
					sol[i][j] = sol[i-1][j] + sol[i][j- get_ith_element(coins, i)]
				} else {
					sol[i][j] = sol[i-1][j]
				}
			}
		}
	}
	//fmt.Println(sol)
	return sol[len(coins)][sum]
}

func Noways_CoinChange_opt(coins []int, sum int) int {
    sol := make([]int, sum+1)
    sol[0] = 1

    for _, coin := range coins {
        for j := coin; j <= sum; j++ {
            sol[j] += sol[j-coin]
        }
    }

    return sol[sum]
}


/*
Given an integer array of coins[ ] of size N representing different types of denominations and an integer sum, 
the task is to get the min number of coins required to make a given value sum.  

Note: Assume that you have an infinite supply of each type of coin. // order doesn't matter
- idea behind solution, coins[i] is the weight array , no value array, there is a sum array to achieve by using these weight.

lets say T(i,j) is the solution for min coins  to pick the coin and reacha sum of sum
recursion can be written as T(i,j) = MIN(T(i-1,j) {{dont pic the ith item,  sum j is still intact}} , 
T(i, j-coin(ith)) {{pic the ith coin i can still pick it in future, sum j is reduces by ith coins weight}})
base case INTMAX

*/
func Min_CoinChange(coins []int, sum int) int {
	sol := make([][]int, len(coins)+1)
	for i:=0; i<=len(coins); i++ {
		sol[i] = make([]int, sum+1)  // sol[len(coins)][sum] will have the sol/ no of ways to pick all coins and with sum sum.
		for j:=0; j<=sum; j++ {
			if i==0 || j==0 {
				//base cases trying to achieve a sum without ny element or
				// else if sum j==0 , even if item i is there, i can have 1 ways by not picking this up
				if i==0 {
					sol[i][j] = math.MaxInt32
				}
			} else if i==1 { // exceptional case
				if j % get_ith_element(coins, i) == 0 { // divide by this coin, mean j/coini numbers to be chosen
					sol[i][j] = j/get_ith_element(coins, i)

				} else {
					sol[i][j] = sol[i-1][j]
				}
		    } else {
				// generic cases
				if get_ith_element(coins,i) <= j {
					sol[i][j] = util.Min(sol[i-1][j], 1+ sol[i][j- get_ith_element(coins, i)])
				} else {
					sol[i][j] = sol[i-1][j]
				}
			}
		}
	}
	fmt.Println(sol)
	return sol[len(coins)][sum]
}
