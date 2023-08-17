package interview

import (
	"bytes"
	"container/list"
	"fmt"
	"math"
	"sort"
	"strconv"
)

//算法面试题重刷
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func InitTreeNode(data []int, i int) *TreeNode {
	head := &TreeNode{
		Value: data[i],
		Left:  nil,
		Right: nil,
	}

	if 2*i+1 < len(data) {
		head.Left = InitTreeNode(data, 2*i+1)
	}

	if 2*i+2 < len(data) {
		head.Right = InitTreeNode(data, 2*i+2)
	}

	return head
}

func BinartTreeLevel(node *TreeNode) []int {
	var result []int
	if node == nil {
		return result
	}

	queue := new(list.List)
	queue.PushBack(node)
	for queue.Len() > 0 {
		tmpNode := queue.Remove(queue.Front()).(*TreeNode)
		result = append(result, tmpNode.Value)

		if tmpNode.Left != nil {
			queue.PushBack(tmpNode.Left)
		}

		if tmpNode.Right != nil {
			queue.PushBack(tmpNode.Right)
		}
	}

	return result
}

//二叉树先序遍历-迭代-中左右
func BinartTreePre(node *TreeNode) []int {
	var result []int
	if node == nil {
		return result
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, node)
	for len(stack) > 0 {
		topNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, topNode.Value)

		if topNode.Right != nil {
			stack = append(stack, topNode.Right)
		}

		if topNode.Left != nil {
			stack = append(stack, topNode.Left)
		}
	}

	return result
}

//二叉树中序遍历-迭代-左中右
func BinartTreeMiddle(node *TreeNode) []int {
	var result []int
	if node == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	for node != nil || len(stack) != 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		topNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, topNode.Value)
		if topNode.Right != nil {
			node = topNode.Right
		}

	}
	return result
}

//二叉树后序遍历-迭代-左右中
func BinartTreeBack(node *TreeNode) []int {
	var result []int
	if node == nil {
		return result
	}

	//中右左
	stack := make([]*TreeNode, 0)
	stack = append(stack, node)
	for len(stack) > 0 {
		topNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, topNode.Value)

		if topNode.Left != nil {
			stack = append(stack, topNode.Left)
		}

		if topNode.Right != nil {
			stack = append(stack, topNode.Right)
		}
	}

	//左右中
	Reverse(result)
	return result
}

func Reverse(data []int) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

//递归法
func BinartTreeBack_2(node *TreeNode) {
	if node == nil {
		return
	}

	BinartTreeBack_2(node.Left)
	BinartTreeBack_2(node.Right)

	fmt.Println(node.Value)
}

//递归法解决树
func RecursiveBinartTreePre(node *TreeNode) {
	if node == nil {
		return
	}

	print(node.Value, "-")
	RecursiveBinartTreePre(node.Left)
	RecursiveBinartTreePre(node.Right)
}

func RecursiveBinartTreeMiddle(node *TreeNode) {
	if node == nil {
		return
	}

	RecursiveBinartTreeMiddle(node.Left)
	print(node.Value, "-")
	RecursiveBinartTreeMiddle(node.Right)
}

func RecursiveBinartTreeBack(node *TreeNode) {
	if node == nil {
		return
	}

	RecursiveBinartTreeBack(node.Left)
	RecursiveBinartTreeBack(node.Right)
	print(node.Value, "-")
}

var result [][]int

func RecursiveBinartTreeLevel(node *TreeNode) {
	if node == nil {
		return
	}
	RecursiveBinartTreeLevelOrder(node, 0)

	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			print(result[i][j], "-")
		}
	}
}

func RecursiveBinartTreeLevelOrder(node *TreeNode, depth int) {
	if node == nil {
		return
	}

	if len(result) == depth {
		result = append(result, []int{})
	}

	result[depth] = append(result[depth], node.Value)
	RecursiveBinartTreeLevelOrder(node.Left, depth+1)
	RecursiveBinartTreeLevelOrder(node.Right, depth+1)
}

var (
	path       []int
	resultData [][]int
)

//1155. 掷骰子等于目标和的方法数
//可以通过测试用例，会超时
func numRollsToTargetSelf(n int, k int, target int) int {
	//resultData = make([][]int, 0)
	Trace(n, k, target, 0)

	return len(resultData)
}

func Trace(n int, k int, target int, current int) {
	if current > target {
		return
	}

	if current == target && len(path) == n {
		tmp := make([]int, len(path))
		copy(tmp, path)
		resultData = append(resultData, tmp)
		return
	}

	for j := 1; j <= k; j++ {
		path = append(path, j)
		current += j
		Trace(n, k, target, current)
		path = path[:len(path)-1]
		current -= j
	}

}

func Sum(path []int) int {
	sum := 0
	for _, item := range path {
		sum += item
	}

	return sum
}

/*dp[i][j]代表i个骰子凑成target=j的方案数
dp[i][j] +== dp[i-1][j-[1~k]].
第i个骰子的数字是1，当骰子是1时，那么dp[i-1][j-1]就代表i-1个骰子骰出j-1的种类有多少。
第i个骰子的数字是2，当骰子是2时，那么dp[i-1][j-2]就代表i-1个骰子骰出j-2的种类有多少。
一直到k。思路主要是反着来的。最后一颗骰子的范围在1~k,那么当第i颗投出这个结果时，种类数就依赖
前i-1颗能投出j-k的数量了。慢慢品 你可以的。
初始化：dp[0][j]:0颗骰子投不出其他的j，只能dp[0][0]=1,其他的dp[0][j] = 0,不可能抛出来
遍历顺序：背包问题，先遍历物品，再背包

*/
func numRollsToTarget(n int, k int, target int) int {
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, target+1)
	}

	dp[0][0] = 1
	mod := int(1e9 + 7)
	for i := 1; i <= n; i++ {
		for j := 1; j <= target; j++ {
			for z := 1; z <= k && z <= j; z++ {
				dp[i][j] = (dp[i][j] + dp[i-1][j-z]) % mod
			}
		}
	}

	return dp[n][target]
}

/*
# 1768 交替合并字符串
*/
func mergeAlternately(word1 string, word2 string) string {
	len1 := len(word1)
	len2 := len(word2)

	var result bytes.Buffer
	for i, j := 0, 0; i < len1 || j < len2; i, j = i+1, j+1 {
		if i < len1 {
			result.WriteByte(word1[i])
		}
		if j < len2 {
			result.WriteByte(word2[j])
		}
	}

	return result.String()
}

func gcdOfStringsMain(str1 string, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)

	for i := len1; i > 0; i-- {
		for j, k := 0, 0; k < len2; j, k = j+1, k+1 {
			j = j % i
			if str1[j] != str2[k] {
				break
			}
			flag := len1 % i
			flag1 := len2 % i
			if k == len2-1 && j == i-1 && flag == 0 && flag1 == 0 {
				return str1[:i]
			}
		}
	}

	return ""

}

func gcdOfStrings(str1 string, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)

	if len1 > len2 {
		str1, str2 = str2, str1
	}

	data := gcdOfStringsMain(str1, str2)
	data2 := gcdOfStringsMain(data, str1)

	if data == data2 {
		return data
	}

	return ""
}

/*
1431. 拥有最多糖果的孩子
*/
func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	result := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}

	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= max {
			result[i] = true
		}
	}

	return result
}

//605. 种花问题
func canPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed) == 1 {
		if flowerbed[0] == 0 && n <= 1 {
			return true
		}

		if flowerbed[0] == 1 && n == 0 {
			return true
		}
		return false
	}

	if n <= 0 {
		return true
	}

	can := 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			if i >= 1 && i < len(flowerbed)-1 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
				can += 1
				flowerbed[i] = 1
			} else if i-1 < 0 && flowerbed[i+1] == 0 {
				can += 1
				flowerbed[i] = 1
			} else if i+1 == len(flowerbed) && flowerbed[i-1] == 0 {
				can += 1
				flowerbed[i] = 1
			}
		}

	}

	if can >= n {
		return true
	}

	return false
}

func reverseVowels(s string) string {
	letterMap := map[byte]struct{}{
		'A': struct{}{},
		'E': struct{}{},
		'I': struct{}{},
		'O': struct{}{},
		'U': struct{}{},
		'a': struct{}{},
		'e': struct{}{},
		'i': struct{}{},
		'o': struct{}{},
		'u': struct{}{},
	}

	bb := []byte(s)
	i, j := 0, len(s)-1
	for i < j {
		_, ok1 := letterMap[bb[i]]
		for i < j && !ok1 {
			i++
			_, ok1 = letterMap[bb[i]]
		}

		_, ok2 := letterMap[bb[j]]

		for i < j && !ok2 {
			j--
			_, ok2 = letterMap[bb[j]]
		}

		if i >= j {
			break
		}

		bb[i], bb[j] = bb[j], bb[i]
		i += 1
		j -= 1
	}

	return string(bb)
}

//151. 反转字符串中的单词
func reverseWords(s string) string {
	if len(s) <= 1 {
		return s
	}

	rr := []rune(s)
	//全部反转
	for i, j := 0, len(rr)-1; i < len(rr)/2; i, j = i+1, j-1 {
		rr[i], rr[j] = rr[j], rr[i]
	}

	//去除左右，中间多余空格
	i, j := 0, len(rr)-1
	for rr[i] == ' ' {
		i += 1
	}

	for rr[j] == ' ' {
		j -= 1
	}

	rr = rr[i : j+1]
	fmt.Println("-xxxxx", string(rr))

	//去除左右，中间多余空格
	slow, fast := 0, 0
	for fast < len(rr) {
		if rr[fast] == ' ' && rr[fast-1] == ' ' && fast-1 >= 0 {
			fast++
			continue
		}

		rr[slow] = rr[fast]
		fast++
		slow++
	}

	rr = rr[:slow]
	fmt.Println("----", string(rr))

	fmt.Println(",", string(rr))
	//逐个单词反转
	last := 0
	for k := 0; k < len(rr); k++ {
		if rr[k] == ' ' || k == len(rr)-1 {
			if k == len(rr)-1 {
				k += 1
			}
			for i, j := last, k-1; i < (k-last)/2+last; i, j = i+1, j-1 {
				rr[i], rr[j] = rr[j], rr[i]
			}
			last = k + 1
		}
	}

	return string(rr)

}

/*
238. 除自身以外数组的乘积
*/
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	left := make([]int, len(nums))
	right := make([]int, len(nums))

	left[0] = 1
	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	right[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	for i := 0; i < len(nums); i++ {
		result[i] = left[i] * right[i]
	}

	return result
}

func productExceptSelfOp(nums []int) []int {
	result := make([]int, len(nums))

	result[0] = 1
	for i := 1; i < len(nums); i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	R := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = result[i] * R
		R = R * nums[i]
	}

	return result
}

/*
334. 递增的三元子序列

make by myself.test case is ok, but long nums will execute timeout
*/
func increasingTriplet(nums []int) bool {
	path := make([]int, 0)
	return BackTrace(nums, 0, path)
}

func BackTrace(nums []int, index int, path []int) bool {
	if len(path) == 3 {
		return true
	}

	for i := index; i < len(nums); i++ {
		if len(path) == 0 || (path)[len(path)-1] < nums[i] {
			path = append(path, nums[i])
			if BackTrace(nums, i+1, path) {
				return true
			}
			path = (path)[:len(path)-1]
		}
	}

	return false
}

func increasingTripletOp2(nums []int) bool {
	first, second := nums[0], math.MaxInt32
	for i := 0; i < len(nums); i++ {
		t := nums[i]
		if t > second {
			return true
		} else if t > first {
			second = t
		} else {
			first = t
		}
	}

	return false
}
func increasingTripletOp3(nums []int) bool {
	leftMin, rightMax := make([]int, len(nums)), make([]int, len(nums))
	leftMin[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		leftMin[i] = int(math.Min(float64(leftMin[i-1]), float64(nums[i])))
	}

	rightMax[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		rightMax[i] = int(math.Max(float64(rightMax[i+1]), float64(nums[i])))
	}

	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > leftMin[i-1] && nums[i] < rightMax[i+1] {
			return true
		}
	}

	return false
}

/*
443. 压缩字符串
*/
func compress(chars []byte) int {
	write, left := 0, 0
	for read, ch := range chars {
		if read == len(chars)-1 || chars[read+1] != ch {
			chars[write] = ch
			write++
			num := read - left + 1
			if num > 1 {
				counteByte := IntToBytes(num)
				for i := 0; i < len(counteByte); i++ {
					chars[write] = counteByte[i]
					write++
				}

			}

			left = read + 1
		}
	}

	return write
}
func IntToBytes(n int) []byte {
	str := strconv.Itoa(n)
	return []byte(str)
}

/*
283. 移动零

输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]
*/
func moveZeroes(nums []int) {
	slow := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[slow], nums[i] = nums[i], nums[slow]
			slow++
		}
	}
}

/*
392. 判断子序列

*/
func isSubsequence(s string, t string) bool {
	m := len(s)
	n := len(t)

	i, j := 0, 0
	count := 0
	for i < m && j < n {
		if s[i] == t[j] {
			i++
			j++
			count++
		} else {
			j++
		}
	}

	return count == m
}

/*
11. 盛最多水的容器
输入：[1,8,6,2,5,4,8,3,7]
输出：49
*/

func maxArea(height []int) int {
	minInter := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	left, right := 0, len(height)-1
	total := 0
	for left < right {
		//According to math algorithms, we should move the min index. t may be bigger
		t := minInter(height[left], height[right]) * (right - left)
		if t > total {
			total = t
		}

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return total
}

/*
1679. K 和数对的最大数目
输入：nums = [1,2,3,4], k = 5
输出：2
*/
func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	first := 0
	second := len(nums) - 1
	count := 0
	for first < second {
		if nums[first]+nums[second] < k {
			first++
		} else if nums[first]+nums[second] > k {
			second--
		} else {
			count++
			first++
			second--
		}
	}

	return count
}

/*
643. 子数组最大平均数 I
输入：nums = [1,12,-5,-6,50,3], k = 4
输出：12.75
解释：最大平均数 (12-5-6+50)/4 = 51/4 = 12.75
*/

func findMaxAverage(nums []int, k int) float64 {
	start := 0
	total := 0
	result := math.MinInt32
	for index, num := range nums {
		if index-start+1 < k {
			total += num
			continue
		}

		total += num

		if total > result {
			result = total
		}

		total -= nums[start]
		start += 1

	}

	fmt.Println("result", result, float64(result)/float64(k))
	return float64(result) / float64(k)
}

/*
1456. 定长子串中元音的最大数目
输入：s = "abciiidef", k = 3
输出：3
解释：子字符串 "iii" 包含 3 个元音字母
*/
func maxVowels(s string, k int) int {
	start := 0
	count := 0
	result := 0
	uMap := make(map[byte]struct{}, 5)
	uMap['a'] = struct{}{}
	uMap['e'] = struct{}{}
	uMap['i'] = struct{}{}
	uMap['o'] = struct{}{}
	uMap['u'] = struct{}{}

	for i := 0; i < len(s); i++ {
		if i-start+1 < k {
			if _, ok := uMap[s[i]]; ok {
				count++
			}
			continue
		}

		if _, ok := uMap[s[i]]; ok {
			count++
		}

		if count > result {
			result = count
		}

		_, ok := uMap[s[start]]
		if ok {
			count--
		}
		start++
	}

	return result
}

/*
1004. 最大连续1的个数 III
输入：nums = [1,1,1,0,0,0,1,1,1,1,0], K = 2
输出：6
解释：[1,1,1,0,0,1,1,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 6。
*/
func longestOnes(nums []int, k int) int {
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] == 0 {
			k--
		}
		right++

		if k < 0 {
			if nums[left] == 0 {
				k++
			}
			left++
		}
	}

	return right - left
}

func longestOnes2(nums []int, k int) int {
	left, right := 0, 0
	res := 0
	count := 0
	for right < len(nums) {
		if nums[right] == 0 {
			count++
		}

		for count > k {
			if nums[left] == 0 {
				count--
			}
			left++
		}

		if right-left+1 > res {
			res = right - left + 1
		}

		right++

	}

	return res
}

/*
1493. 删掉一个元素以后全为 1 的最长子数组

*/
func longestSubarray(nums []int) int {
	left, right := 0, 0
	result := -1
	count := 0
	for right < len(nums) {
		if nums[right] == 0 {
			count++
		}

		for count > 1 {
			if nums[left] == 0 {
				count--
			}
			left++
		}

		if result < right-left+1 {
			result = right - left
		}

		right++
	}

	return result
}

/*
1732. 找到最高海拔
输入：gain = [-5,1,5,0,-7]
输出：1
解释：海拔高度依次为 [0,-5,-4,1,1,-6] 。最高海拔为 1 。
*/
func largestAltitude(gain []int) int {
	result := 0
	lastNum := 0
	for _, item := range gain {
		lastNum += item
		result = maxInt(result, lastNum)
	}

	return result
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

/*
724. 寻找数组的中心下标
输入：nums = [1, 7, 3, 6, 5, 6]
输出：3
解释：
中心下标是 3 。
*/
func pivotIndex(nums []int) int {
	for i := 0; i < len(nums); i++ {
		leftTotal := getTotalFromSlice(nums, 0, i-1)
		rightTotal := getTotalFromSlice(nums, i+1, len(nums)-1)

		if leftTotal == rightTotal {
			return i
		}
	}

	return -1
}

func getTotalFromSlice(data []int, start, end int) int {
	total := 0

	if end < 0 || start > len(data)-1 {
		return total
	}
	for i := start; i <= end; i++ {
		total += data[i]
	}

	return total
}

/*
2215. 找出两数组的不同
输入：nums1 = [1,2,3], nums2 = [2,4,6]
输出：[[1,3],[4,6]]
*/

func findDifference(nums1 []int, nums2 []int) [][]int {
	nums1Map := make(map[int]struct{}, len(nums1))
	for _, item := range nums1 {
		nums1Map[item] = struct{}{}
	}

	nums2Map := make(map[int]struct{}, len(nums2))
	for _, item := range nums2 {
		nums2Map[item] = struct{}{}
	}

	result := make([][]int, 0, 2)
	first := make([]int, 0)
	for k := range nums1Map {
		if _, ok := nums2Map[k]; !ok {
			first = append(first, k)
		}
	}

	result = append(result, first)

	second := make([]int, 0)
	for k := range nums2Map {
		if _, ok := nums1Map[k]; !ok {
			second = append(second, k)
		}
	}

	result = append(result, second)

	return result
}

func uniqueOccurrences(arr []int) bool {
	showMap := make(map[int]int, 0)
	countMap := make(map[int]int, 0)
	for _, item := range arr {
		count := showMap[item]
		count++
		showMap[item] = count
	}

	for k, v := range showMap {
		if _, ok := countMap[v]; ok {
			return false
		}

		countMap[v] = k
	}

	return true
}

func closeStrings(word1 string, word2 string) bool {
	nums1Map := make(map[rune]int, 0)
	countMap := make(map[int]int, 0)
	for _, item := range word1 {
		count := nums1Map[item]
		count++
		nums1Map[item] = count
	}

	for _, v := range nums1Map {
		cur, ok := countMap[v]
		if !ok {
			countMap[v] = 1
			continue
		}
		cur++
		countMap[v] = cur
	}

	nums2Map := make(map[rune]int, 0)
	for _, item := range word2 {
		count := nums2Map[item]
		count++
		nums2Map[item] = count
	}

	for k, v := range nums2Map {
		cur, ok := countMap[v]
		_, ok2 := nums1Map[k]
		if !ok || !ok2 {
			return false
		}

		cur--
		countMap[v] = cur
	}

	for _, v := range countMap {
		if v != 0 {
			return false
		}
	}

	return true

}

func equalPairs(grid [][]int) int {
	m := make(map[int]int, len(grid[0]))
	n := make(map[int]int, len(grid[0]))
	for i := 0; i < len(grid[0]); i++ {
		m[grid[0][i]] = i
		n[grid[i][0]] = i
	}

	count := 0
	for mk, mv := range m {
		nv, ok := n[mk]
		if !ok {
			continue
		}

		i := 0
		for i < len(grid[0]) {
			if grid[i][mv] != grid[nv][i] {
				break
			}
			i++
		}

		if i == len(grid[0]) {
			count++
		}
	}

	return count
}
