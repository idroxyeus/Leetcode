package main

func resultArray(nums []int) []int {
    // Initialize both arrays with their first respective elements
    arr1 := []int{nums[0]}
    arr2 := []int{nums[1]}
    
    // Simulate the distribution for the remaining elements
    for i := 2; i < len(nums); i++ {
        if arr1[len(arr1)-1] > arr2[len(arr2)-1] {
            arr1 = append(arr1, nums[i])
        } else {
            arr2 = append(arr2, nums[i])
        }
    }
    
    // Concatenate arr2 to the end of arr1 and return the result
    return append(arr1, arr2...)
}
