package main1

import (
	"math"
	"my_project/算法/utils"
	"sort"
	"strings"
)

func main1() {
	utils.Info(letterCombinations("234"))
}

var letterMap = map[byte][]byte{
	'2': []byte{'a', 'b', 'c'},
	'3': []byte{'d', 'e', 'f'},
	'4': []byte{'g', 'h', 'i'},
	'5': []byte{'j', 'k', 'l'},
	'6': []byte{'m', 'n', 'o'},
	'7': []byte{'p', 'q', 'r', 's'},
	'8': []byte{'t', 'u', 'v'},
	'9': []byte{'w', 'x', 'y', 'z'},
}

// q17:排列组合
func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	bss := [][]byte{}
	d := []byte(digits)
	for i, v := range d {
		utils.Info("111>>", i, ">>", string(v))
		if m, ok := letterMap[v]; ok {
			bb := [][]byte{}
			for _, letter := range m {
				utils.Info("222>>", string(letter))
				if i == 0 {
					bb = append(bb, []byte{letter})
				} else {
					for _, vv := range bss {
						vv = append(vv, letter)
						bb = append(bb, vv)
						utils.Info(len(vv))
						utils.Info(cap(vv))
						utils.Info("33>>", string(vv))
					}
				}
			}
			bss = bb
			sss := []string{}
			for _, v := range bss {
				sss = append(sss, string(v))
			}
			utils.Info(sss)
		} else {
			return nil
		}
	}
	ss := []string{}
	for _, v := range bss {
		ss = append(ss, string(v))
	}
	return ss
}

// func main() {
// 	utils.Info(threeSumClosest([]int{-1, 0, 1, 2, -1, -4}, -7))
// }

// q16:查找三个元素和最接近目标值
func threeSumClosest(nums []int, target int) int {
	// utils.Info(target)
	sort.Ints(nums)
	l := len(nums)
	if l < 3 {
		return 0
	}
	minDiff := -1
	closest := 0
	for i := 0; i < l-2; i++ {
		// utils.Info(i, ">>>>>", nums[i])
		// utils.Info("minDiff", minDiff)
		// utils.Info("closest", closest)
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := l - 1
		// utils.Info(left)
		// utils.Info(right)
		for left < right {
			if left > i+1 && nums[left] == nums[left+1] {
				left++
				continue
			}
			w := nums[left] + nums[right] + nums[i]
			// utils.Info("w >>", w)
			switch {
			case w == target:
				return target
			case w > target:
				right--
			case w < target:
				left++
			}
			diff := getDiff(w, target)
			// utils.Info("left >>>", nums[left])
			// utils.Info("right >>>", nums[right])
			if minDiff < 0 {
				minDiff = diff
				closest = w
			} else if diff < minDiff {
				minDiff = diff
				closest = w
			}
			// utils.Info(diff)
			// utils.Info(minDiff)
		}
	}
	if minDiff != -1 {
		return closest
	}
	return 0
}

func getDiff(a, b int) int {
	n := a - b
	if n < 0 {
		n = 0 - n
	}
	return n
}

// q15:求和为0的三元组
// S = [-1, 0, 1, 2, -1, -4]
// [
//   [-1, 0, 1],
//   [-1, -1, 2]
// ]
// [-2,0,0,2,2]
// [-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6]
// func main() {
// 	utils.Info(threeSum([]int{-2, 0, 0, 2, 2}))
// }

// q15:求和为0的三元组
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	l := len(nums)
	res := [][]int{}
	for i := 0; i < l-2; i++ {

		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		u := 0 - nums[i]
		left := i + 1
		right := l - 1
		for left < right {
			if left > i+1 && nums[left] == nums[left-1] {
				left++
				continue
			}
			w := nums[left] + nums[right]
			switch {
			case w == u:
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
			case w > u:
				right--
			case w < u:
				left++
			}
		}
	}
	return res
}

//q12:阿拉伯转罗马
// I（1）、V（5）、X（10）、L（50）、C（100）、D（500）和M（1000）
// func main() {
// 	utils.Info(intToRoman(1))
// }

//q12:阿拉伯转罗马
func intToRoman(num int) string {
	if num > 3999 || num < 1 {
		return ""
	}
	m := map[int]byte{1000: 'M', 500: 'D', 100: 'C', 50: 'L', 10: 'X', 5: 'V', 1: 'I'}
	n := []int{1000, 500, 100, 50, 10, 5, 1}
	r := []byte{}
	for i := 0; i < len(n); i += 2 {
		a := num / n[i]
		num = num % n[i]
		utils.Info(a)
		if a > 0 {
			mm, ok := m[n[i]]
			if !ok {
				return ""
			}
			if a < 4 {
				for j := 0; j < a; j++ {
					r = append(r, mm)
				}
			} else {
				kk, ok := m[n[i-1]]
				if !ok {
					return ""
				}
				switch {
				case a == 4:
					r = append(r, mm)
					r = append(r, kk)
				case a == 9:
					r = append(r, mm)
					tt, ok := m[n[i-2]]
					if !ok {
						return ""
					}
					r = append(r, tt)
				default:
					r = append(r, kk)
					for j := 0; j < a-5; j++ {
						r = append(r, mm)
					}
				}
			}
		}
		if num == 0 {
			break
		}
	}
	return string(r)
}

//q11:求最大面积
// func main() {
// 	i := []int{2, 3, 4, 5, 6, 7}
// 	utils.Info(maxArea(i))
// }

//q11:求最大面积
func maxArea(height []int) int {
	l := len(height)
	if l < 2 {
		return 0
	}
	maxArea := 0
	left := 0
	right := l - 1
	for left < right {
		width := right - left
		minHeight := 0
		if height[left] < height[right] {
			minHeight = height[left]
			left++
		} else {
			minHeight = height[right]
			right--
		}
		area := width * minHeight
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

//q10:正则匹配(. *)
// func main() {
// 	utils.Info(isMatch("aaa", "."))
// }

//q10:正则匹配(. *) todo
func isMatch(s string, p string) bool {
	if s == p {
		return true
	}
	if s == "" || p == "" {
		return false
	}
	b := []byte(s)
	var lastLetter byte
	for i := 0; i < len(p); i++ {
		if len(b) == 0 && (p[i] != '*' || i != len(p)-1) {
			return false
		}
		switch p[i] {
		case '.':
			b = b[1:]
		case '*':
			if i == len(p)-1 { //最后一个
				if lastLetter != '.' {
					for _, v := range b {
						if v != lastLetter {
							return false
						}
					}
				}
				return true
			} else { //中间的*

			}
		default:
			if b[0] == p[i] {
				b = b[1:]
			} else {
				return false
			}
		}
		lastLetter = p[i]
	}
	return false
}

/*
aaa*aabbb

*/

//q8:字符串转数字
// func main() {
// 	utils.Info(myAtoi(" -0  "))
// }

//q8:字符串转数字
func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}
	sign := 1
	switch str[0] {
	case '+':
		str = str[1:]
	case '-':
		sign = -1
		str = str[1:]
	}
	d := "0123456789"
	res := 0
	for _, v := range str {
		if index := strings.IndexRune(d, v); index >= 0 {
			res = 10*res + sign*index
			if res > math.MaxInt32 {
				return math.MaxInt32
			}
			if res < math.MinInt32 {
				return math.MinInt32
			}
			continue
		}
		return res
	}
	return res
}

//PAYPALISHIRING
// P   A   H   N
// A P L S I I G
// Y   I   R
//PAHNAPLSIIGYIR
// A     B     C     D     E
// F   G H   I J   K L   M N
// O P   Q R   S T   U V
// W     X     Y     Z

//q6:解析z字形数组
// func main() {
// 	str := "AFOWPGBHQXRICJSYTKDLUZVMEN"
// 	s := convert(str, 2)
// 	utils.Info(s)
// }

func convert(s string, numRows int) string {
	l := len(s)
	if numRows < 2 || l <= numRows || l < 3 {
		return s
	}
	b1 := []byte(s)
	b2 := []byte{}
	bytes := make([][]byte, numRows)
	unit := 2*numRows - 2
	for i, v := range b1 {
		extra := i % unit
		if extra >= numRows {
			extra = 2*numRows - extra - 2
		}
		bytes[extra] = append(bytes[extra], v)
	}
	for _, v := range bytes {
		b2 = append(b2, v[:]...)
	}
	return string(b2)
}
