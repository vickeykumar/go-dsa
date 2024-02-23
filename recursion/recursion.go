package recursion

import (
	_ "fmt"
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

