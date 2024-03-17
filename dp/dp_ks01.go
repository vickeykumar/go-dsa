package dp

import (
	"fmt"
	"go-dsa/util"
	"math"
)

/*
DP variations:
    * 0/1 Knapsack
    * Unbounded Knapsack
    * Shortest Path (eg: Unique Paths I/II)
    * Fibonacci Sequence (eg: House Thief, Jump Game)
    * Longest Common Substring/Subsequeunce
*/

func get_ith_element(V []int, i int) int {
	if i==0 {
		return 0
	}
	return V[i-1]
}

/* dynamic programming problems similar to 0/1 knapsack
 * Knapsack_01(V , W []int, capacity int) int {
 * IsSubsetSumPresent(nums []int, sum int) bool {
 * IsEqualSumPartPresent(nums []int) bool {
 * CountSubsetSum(nums []int, sum int) int { 
 * MinSubsetSumDiff(nums []int) int {
 * CountSubsetSumdiff(nums []int, diff int) int {
 * FindTargetSumWays(nums []int, target int) int {
 * 
 * Pattern - where we have set of items and we have to achieve an optimal solumtion
 * by picking/not picking ith item to achieve a target/optimal solution
 */

// takes value array and weight array, capacity
func Knapsack_01(V , W []int, capacity int) int {
	if len(V) != len(W) {
		panic(fmt.Errorf("Invalid: value array and weight array len is not same"))
	}
	profit := make([][]int, len(V)+1)
	for i, _ :=range profit {
		profit[i] = make([]int, capacity+1)
		for j,_ := range profit[i] {
			 if i==0 || j==0 {
			 	profit[i][j] = 0 
			 	// base case if profit(i) is zero, net profit will be zero
			 	// if capacity is zero, profit is also zero, nothing to put inside knapsack
			 } else {
				 if get_ith_element(W, i) > j {
				 	 profit[i][j] =  profit[i-1][j] // if current weight is more than capacity can't choose anyway
				 } else {
				 	// profit i will make if i choose ith item or not choose
				 	// profit[i][j] gives max profit if i take/not take  this ith element considering j capacity
				 	profit[i][j] = util.Max(get_ith_element(V, i)+profit[i-1][j- get_ith_element(W, i)], profit[i-1][j])
				 }
			}
		}
	}
	return profit[len(V)][capacity]
}


/* given an array of numbers, 
is there a subset with a sum, sum  
subsum[i][j] is the solution if i select/not select ith item, can j be sum
*/
func IsSubsetSumPresent(nums []int, sum int) bool {
	subsum := make([][]bool, len(nums)+1)
	for i:=0; i< len(subsum); i++ {
		subsum[i] = make([]bool, sum+1)
		for j:=0; j<sum+1; j++ {
			if i==0 {
				// base case
				if i==0 {
					// no element in arr, sum can't be possible
					subsum[i][j] = false
				}
				if j==0 {
					// to achieve a sum of 0,its possible by not selecting current element
					subsum[i][j] = true
				}
			} else {
				// general case
				if get_ith_element(nums, i) > j {
				 	 subsum[i][j] =  subsum[i-1][j] // if current value is more than sum itself can't choose anyway
				 } else {
					// can we achieve subset sub by either selecting current element or not
					subsum[i][j] = (subsum[i-1][j- get_ith_element(nums, i)]) || subsum[i-1][j]
				}
			}
		}
	}
	return subsum[len(nums)][sum]
}

/* is there an equal sum partition present in array
 * 2s = sumarr so problem is same as IsSubsetSumPresent(nums, sumarr/2)
*/

func IsEqualSumPartPresent(nums []int) bool {
	if len(nums)%2!=0 {
		// odd, not possible
		return false
	}
	sumarr := 0 
	for _,item := range nums {
		sumarr += item
	}
	return IsSubsetSumPresent(nums, sumarr/2)
}

/* count of all the subsets with a sum in array */
func CountSubsetSum(nums []int, sum int) int {
	count := make([][]int, len(nums)+1) 
	// count[i][j] represents count of subsets if i select/notselect item(i) and getting a sum j
	for i:=0; i<len(count); i++ {
		count[i] = make([]int, sum+1) // to maintain index sum 0...sum
		for j:=0;j<=sum;j++ {
			//base case
			if i==0 {
				if i==0 {
					count[i][j] = 0 // if ther is no items to choose..
				}
				if j==0 {
					count[i][j] = 1 
					// don't pick any item, 1 way, if sum is zero
					// what if some of items are zero, we still can choose thoise items to make sum 0,
					// for item i>0 if they are zero, still we can have multiple ways to select to reach sum j=0
					// so, lets calculate this in else part for those i>0

				}
			} else {
				// general case
				if get_ith_element(nums, i) > j {
					// item is greater than sum j itself
					count[i][j]=count[i-1][j]
				} else {
					// by selecting + by not selecting
					count[i][j] = count[i-1][j-get_ith_element(nums, i)] + count[i-1][j]
				}
			}
		} 
	}
	//fmt.Println(count)
	return count[len(nums)][sum] // count of number of subsets for all items and sum -> sum
}

/*
this can be considered a variation of subset sum problem
given an array, find min possible value of abs(s1-s2)=diff , where s1 is sum of 
first subset and s2 is sum of second subset ( 2 variables unknown -> reduce to single variable)
we know s1+s2 = sumarr,
s2 = sumarr-s1
we have to minimize abs(s2-s1) , abs((sumarr-s1)-s1) ->
minimize sumarr-2s1 ->
it means if we are able to find just optimal s1 such that abs(sumarr-2s1) is minimum, we will find optimal solution
steps in substet sum problem of size sumarr, try to find all the sums that is possible with all the elements in array,
minimize abs(sumarr-2s1) - return min(sumarr-2s1)
soving eq. 2s = sumarr+diff
-diff = sumarr-2s
diff = |sumarr-2s|
so return minimum of sumarr-2s
*/
func MinSubsetSumDiff(nums []int) int {
	sumarr := 0
	for _, num := range nums {
		sumarr += num 
	}
	subsetsum := make([][]bool, len(nums)+1)
	for i:=0; i<len(nums)+1; i++ {
		subsetsum[i]=make([]bool,sumarr+1) // since golang array 0-indexed
		for j:=0;j<sumarr+1;j++ {
			//base case
			if i==0 || j==0 {
				if i==0 {
					subsetsum[i][j]=false
				}
				if j==0 {
					subsetsum[i][j]=true
				}
			} else {
				// generic cases
				if get_ith_element(nums, i)>j {
					subsetsum[i][j]=false
				} else {
					subsetsum[i][j] = subsetsum[i-1][j- get_ith_element(nums, i) ] || subsetsum[i-1][j]
				}
			}
		}
	}

	// after this subsetsum[len(nums)] should give all the possible subset sums i can get using whole array
	// calculate array |sumarr-2j| and miniumize it
	min := math.MaxInt32
	for j, subpresent := range subsetsum[len(nums)] {
		if subpresent {
			min = util.Min(min, int(math.Abs(float64(sumarr-2*j))))
		}
	}
	return min
}


/*
count all the subset sums with given diff
s1-s2 = diff
s1+s2 = sumarr

solving, 2s1=diff+sumarr
s1 = (diff+sumarr)/2
same as CountSubsetSum(nums, (diff+sumarr)/2)
*/

func CountSubsetSumdiff(nums []int, diff int) int {
	sumarr := 0
	for _, num := range nums {
		sumarr += num 
	}
	targetSum := (diff + sumarr) / 2
	if targetSum < 0 || (diff + sumarr)%2 > 0 {
        // If the target sum is odd, return 0
        return 0
    }
	return CountSubsetSum(nums, targetSum)
}



/*
Leetcode 494 targetsum

You are given an integer array nums and an integer target.

You want to build an expression out of nums by adding one of the symbols '+' and '-' before each integer in nums and then concatenate all the integers.

For example, if nums = [2, 1], you can add a '+' before 2 and a '-' before 1 and concatenate them to build the expression "+2-1".
Return the number of different expressions that you can build, which evaluates to target.

solution:

if we assign + and - ... ultimately it will resolve to s1-s2=diff

s1+s2=sumarr
s1=(diff+sumarr)/2

*/

func FindTargetSumWays(nums []int, target int) int {
    sumarr := 0
	for _, num := range nums {
		sumarr += num 
	}
	targetSum := (target + sumarr) / 2
	if targetSum < 0 || (target + sumarr)%2 != 0 {
        // If the target sum is odd, return 0
        return 0
    }
	return CountSubsetSum(nums, targetSum)
}