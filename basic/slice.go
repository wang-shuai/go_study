package main
import "fmt"

func printSlice(ss []int)  {
	fmt.Printf("len=%d capicity=%d slice=%v \n",len(ss),cap(ss),ss)
}

func main(){
	var nums []int
	printSlice(nums)

	nums = append(nums,1,2,3,4,5,6,7,8,9)
	printSlice(nums)

	nums1 := append(nums,1,2)
	printSlice(nums1)

	nums2 := []int{1}
	printSlice(nums2)
	nums3 := append(nums2,6,7,3,5)
	printSlice(nums3)

	nums4 := nums[1:4]
	printSlice(nums4)
	nums[1] = 100
	printSlice(nums4)
}