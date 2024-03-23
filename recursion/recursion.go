package recursion

import (
	_ "fmt"
	"sort"
)

/*
the append function creates a new slice, and the modified slice header 
inside the myFunc function doesn't affect the original slice header in 
the main function. Therefore, the original mySlice slice remains unchanged,
and its length and capacity are still the same.
*/

func Powersets(nums []int, currpos,last int) [][]int {
    //base case when currpos>last
    if currpos>last {
        return [][]int{
            []int{},
        }
    }
    result_powerset := make([][]int,0)
    // lets say recursion is giving me sets of all the subsets (powerset) till currpos+1th index to last index
    // include or exclude ith eleent to generate possible power set
    powersets := Powersets(nums, currpos+1, last)
    //add to result without and with currpos
    result_powerset = append(result_powerset, powersets...)
    // now add with the num[currpos] to each elemnt form previously calculated powerset
    for i:=0;i<len(powersets);i++ {
        set := append([]int{nums[currpos]},powersets[i]...)
        result_powerset = append(result_powerset, set)
    }
    return result_powerset
}

func Powersets2(nums []int, currpos, last int) [][]int {
    // Sort nums to group duplicate elements together
    sort.Ints(nums)
    // Base case when currpos > last
    if currpos > last {
        return [][]int{{}} // Return an empty set
    }
    resultPowerset := make([][]int, 0)
    // Count the number of occurrences of the current element
    count := 1
    for currpos+count <= last && nums[currpos+count] == nums[currpos] {
        count++
    }
    // Recursively generate subsets from currpos+count to last
    powersets := Powersets(nums, currpos+count, last)
    // Include or exclude the current element to generate subsets
    for _, subset := range powersets {
        resultPowerset = append(resultPowerset, subset) // Add subset without current element
        // Add all possible combinations of the current element to the subset
        for i := 1; i <= count; i++ {
            set := make([]int, i)
            copy(set, nums[currpos:currpos+i])
            resultPowerset = append(resultPowerset, append(set, subset...))
        }
    }
    return resultPowerset
}

func Permutehelper(nums []int, pos int) (results [][]int) {
    // Base case when pos becomes equal to the last index
    if pos == len(nums)-1 {
        // Append a copy of nums to results
        results = append(results, append([]int{}, nums...))
        return
    }
    // Swap currpos to each place in nums[pos+1:] and then permute the rest of the places
    // This way currpos will get a chance to swap with each individual character
    for i := pos; i < len(nums); i++ {
        // Swap
        nums[pos], nums[i] = nums[i], nums[pos]
        // Recursively permute the rest of the array
        temp := Permutehelper(nums, pos+1)
        // Append the permutations generated from the rest of the array to results
        results = append(results, temp...)
        // Backtrack: restore the original order of nums
        nums[pos], nums[i] = nums[i], nums[pos]
    }
    return
}

// lookup , and insert if char not found
func lookup_and_insert(set []int, c int) ([]int,bool) {
    //fmt.Println(set)
    for i:=0; i<len(set); i++ {
        if (set[i]==c) {
            return set, true;
        }
    }    
    // not found, insert
    set=append(set,c)
    return set, false;
}


func Permutehelper2(nums []int, pos int) (results [][]int) {
    // Base case when pos becomes equal to the last index
    if pos == len(nums)-1 {
        // Append a copy of nums to results
        results = append(results, append([]int{}, nums...))
        return
    }
    set := make([]int,0)
    var ok bool
    // Swap currpos to each place in nums[pos+1:] and then permute the rest of the places
    // This way currpos will get a chance to swap with each individual character
    for i := pos; i < len(nums); i++ {
        // if this element is already traversed previously skip
        set, ok = lookup_and_insert(set, nums[i]) 
        if ok {
            continue
        }
        // Swap
        nums[pos], nums[i] = nums[i], nums[pos]
        // Recursively permute the rest of the array
        temp := Permutehelper2(nums, pos+1)
        // Append the permutations generated from the rest of the array to results
        results = append(results, temp...)
        // Backtrack: restore the original order of nums
        nums[pos], nums[i] = nums[i], nums[pos]
    }
    return
}

//input/output method, input- array
// output- subset
func GenSubsets(nums []int, subset []int, pos int) (powerset [][]int) {
    if pos == len(nums) {
        // base case
        subsetCopy := make([]int, len(subset))
        copy(subsetCopy, subset)
        powerset = append(powerset, subsetCopy)
        return 
    } 
    tmpsubsets := GenSubsets(nums, subset, pos+1) // without nums[pos] 
    powerset = append(powerset, tmpsubsets...)

    subset = append(subset, nums[pos])
    tmpsubsets = GenSubsets(nums, subset, pos+1) // with nums[pos]
    powerset = append(powerset, tmpsubsets...)
    return
}

//input/output method, input- number
// output- subset
func Combinations(n, index, k int, subset []int) (result [][]int) {
    if len(subset)==k {
        // base case 
        subsetCopy := make([]int, len(subset))
        copy(subsetCopy, subset)
        result = append(result, subsetCopy)
        return
    }
    if index == n+1 {
        return
    }
    tmpsubsets := Combinations(n, index+1, k, subset) // without
    result = append(result, tmpsubsets...)

    subset = append(subset, index) // include current number
    tmpsubsets = Combinations(n, index+1, k, subset) // with current num
    result = append(result, tmpsubsets...)
    return
}


/* selection of k items from n items
input/output method, input- items
// output- subset */
func Combinations2(items []int, index, k int, subset []int) (result [][]int) {
    if len(subset)==k {
        // base case 
        subsetCopy := make([]int, len(subset))
        copy(subsetCopy, subset)
        result = append(result, subsetCopy)
        return
    }
    if index == len(items) {
        return
    }
    tmpsubsets := Combinations2(items, index+1, k, subset) // without current item
    result = append(result, tmpsubsets...)

    subset = append(subset, items[index]) // include current item 
    tmpsubsets = Combinations2(items, index+1, k, subset) // with current num
    result = append(result, tmpsubsets...)
    return
}

/*
Given an array of distinct integers candidates and a target integer target, 
return a list of all unique combinations of candidates where the chosen numbers sum to target. 
You may return the combinations in any order.
The same number may be chosen from candidates an unlimited number of times.
Two combinations are unique if the frequency of at least one of the chosen numbers is different.

The test cases are generated such that the number of unique combinations that sum up to target 
is less than 150 combinations for the given input.

*/
func CombinationsSum(items []int, index, target int, subset []int) (result [][]int) {
    if target<0 {
        return
    }
    if target==0 {
        // base case 
        subsetCopy := make([]int, len(subset))
        copy(subsetCopy, subset)
        result = append(result, subsetCopy)
        return
    }
    if index == len(items) {
        return
    }
    tmpsubsets := CombinationsSum(items, index+1, target, subset) // without current item
    result = append(result, tmpsubsets...)

    subset = append(subset, items[index]) // include current item and dont move on
    tmpsubsets = CombinationsSum(items, index, target-items[index], subset) // with current num
    result = append(result, tmpsubsets...)
    
    return
}

/*
Given a collection of candidate numbers (candidates) and a target number (target), 
find all unique combinations in candidates where the candidate numbers sum to target.

Each number in candidates may only be used once in the combination.

Note: The solution set must not contain duplicate combinations.
*/
func CombinationsSum2(items []int, index, target int, subset []int) (result [][]int) {
    if target<0 {
        return
    }
    if target==0 {
        // base case 
        subsetCopy := make([]int, len(subset))
        copy(subsetCopy, subset)
        result = append(result, subsetCopy)
        return
    }
    if index == len(items) {
        return
    }
    subset = append(subset, items[index]) // include current item and move on/ cz only once
    tmpsubsets := CombinationsSum(items, index+1, target-items[index], subset) // with current num
    result = append(result, tmpsubsets...)
    
    // backtrack vvi ?? why keep picking till u want to pick and ignore once saturated
    subset = subset[:len(subset)-1]
    // Skip duplicates at the same level of recursion
    for index+1 < len(items) && items[index+1] == items[index] {
        index++
    }
    tmpsubsets = CombinationsSum(items, index+1, target, subset) // move on without current item if not picking up
    result = append(result, tmpsubsets...)
    
    return
}


/*
The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.

Given an integer n, return all distinct solutions to the n-queens puzzle. You may return the answer in any order.
*/
func isSafe(chessboard [][]bool, row, col int) bool {
    // up
    r := row-1
    for r>=0 {
        if chessboard[r][col] {
            // already a queen in this col, not safe
            return false
        }
        r--;
    }
    
    // left diag, (r-1,c-1)-> (r-2,c-2)
    r = row-1
    c := col-1
    for r>=0 && c>=0 {
        if chessboard[r][c] {
            // already a queen in this col, not safe
            return false
        }
        r--;
        c--;
    }
    
    // right diag, (r-1,c+1)-> (r-2,c+2)
    r = row-1
    c = col+1
    for r>=0 && c<len(chessboard) {
        if chessboard[r][c] {
            // already a queen in this col, not safe
            return false
        }
        r--;
        c++;
    }
    return true   
}

func NQueens(chessboard [][]bool, row int, ans *[][]string) {
    // base case when all queens are placed
    if row == len(chessboard) {
        // print the solutions
        //fmt.Println(chessboard)
        chess_str := make([]string, len(chessboard))
        for i:=0; i<len(chessboard); i++ {
            bytes := make([]byte, len(chessboard)) // byte array for string
            for j:=0; j<len(chessboard); j++ {
                if chessboard[i][j] {
                    bytes[j] = byte('Q')
                } else {
                    bytes[j] = byte('.')
                }
            }
            chess_str[i] = string(bytes) // push ith row string
        }
        *ans = append(*ans, chess_str)
    }
    for col:=0;col<len(chessboard);col++ {
        if isSafe(chessboard, row, col) {
            // place the queen here if is safe, then move on
            chessboard[row][col] = true
            NQueens(chessboard, row+1, ans) // place next queen in next row
            // backtrack to find next valid solution
            chessboard[row][col] = false
        }
    }
} 

func SolveNQueens(n int) [][]string {
    var ans [][]string
    chessboard := make([][]bool, n)
    for i:=0;i<n;i++ {
        chessboard[i] = make([]bool, n)
    }
    //fmt.Println(chessboard)
    NQueens(chessboard, 0, &ans) // start placing from 1st row , 0
    return ans
}
