package utility

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Add() int {
	return 3 + 2
}

func Wind() {
	fmt.Println("Enter the temperature(less than 50) in Farenheat")
	var t, v float64
	fmt.Scanf("%f", &t)
	fmt.Println("Enter the speed (More than 3 and less 120) in miles per hour ")
	fmt.Scanf("%f", &v)
	if (v < 120 && v > 3) && t < 50 {
		w := 35.74 + 0.6215*t + (0.4275*t-35.75)*math.Pow(v, 0.16)
		fmt.Printf("windchill = %f \n", w)
	} else {
		fmt.Println("in valid in put")
		Wind()
	}
}

func Triplet() {
	fmt.Println("Enter the size of the array")
	var n int

	fmt.Scanf("%d", &n)
	arr := make([]int, n)
	fmt.Println("Enter array elements are\n")
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Println("Entered array elements are \n")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i])
	}

	// for i := 0; i < n; i++ {
	// 	for j := i + 1; j < n; j++ {
	// 		if arr[i] > arr[j] {
	// 			temp := arr[i]
	// 			arr[i] = arr[j]
	// 			arr[j] = temp
	// 		}
	// 	}
	// }
	m := 1
	fmt.Println(arr)
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			for k := j + 1; k < len(arr); k++ {
				if arr[i]+arr[j]+arr[k] == 0 {
					fmt.Println("triplet found the elements are,", arr[i], arr[j], arr[k])
					m = 0
				}
			}
		}
	}

	if m == 1 {
		fmt.Println("triplet not found \n")
	}
}

func Euclidean() {
	fmt.Println("Enter the coordinates")

	var x float64
	fmt.Println("Enter X coordinate value :")
	fmt.Scanf("%f", &x)
	fmt.Println("Enter Y coordinate value :")
	var y float64
	fmt.Scanf("%f", &y)
	var result float64
	result = float64(math.Pow(x, 2)) + float64(math.Pow(y, 2))
	var sqrrootresult float64
	sqrrootresult = (math.Sqrt(result))
	fmt.Println("euclidean distance is", sqrrootresult)
}

func Strpermut() {
	fmt.Printf("Enter the String :")
	var str string
	fmt.Scanf("%s", &str)
	//str1 := []rune(str)
	len := len(str)
	l := 0
	r := len - 1
	perm([]rune(str), l, r)

}
func perm(str []rune, l int, r int) {
	if l == r {
		fmt.Println("String=", string(str))
	} else {
		for i := l; i <= r; i++ {
			str[l], str[i] = str[i], str[l]
			perm([]rune(str), l+1, r)
			str[l], str[i] = str[i], str[l]
		}
	}
}
func Quadratic() {
	var a, b, c float64
	fmt.Println("Give values for a,b,c ")
	fmt.Print("A value :")
	fmt.Scanf("%f", &a)
	fmt.Print("B value :")
	fmt.Scanf("%f", &b)
	fmt.Print("C value :")
	fmt.Scanf("%f", &c)
	delta := math.Pow(b, 2) - (4 * a * c)
	if delta < 0 {
		fmt.Println("Roots are imaginary")
		first := (-b) / (2 * a)
		second := (math.Sqrt(-delta)) / (2 * a)
		fmt.Println("Root1 =", first, "+i", second)
		fmt.Println("Root2 =", first, "-i", second)
	} else if delta == 0 {
		first := (-b) / (2 * a)
		fmt.Println("Root1 =", first)
	} else {
		fmt.Println("Roots are real numbers")
		first := (-b) / (2 * a)
		second := float64(math.Sqrt(delta)) / (2 * a)
		root1 := first + second
		root2 := first - second
		fmt.Println("Root1 =", root1)
		fmt.Println("Root2 =", root2)
	}

}
func Anagram() {
	fmt.Printf("Enter the First String.:")
	var str1 string
	fmt.Scanf("%s", &str1)
	fmt.Printf("Enter the Second String.:")
	var str2 string
	fmt.Scanf("%s", &str2)
	var str3 string
	var str4 string
	str3 = strings.ToLower(str1)
	str4 = strings.ToLower(str2)
	var l1 int
	var l2 int
	l1 = len(str3)
	l2 = len(str4)

	runes1 := []rune(str3)
	runes2 := []rune(str4)

	for i := 0; i < len(runes1); i++ {
		for j := 0; j < len(runes1); j++ {
			if runes1[i] < runes1[j] {
				tem := runes1[i]
				runes1[i] = runes1[j]
				runes1[j] = tem
			}
		}
	}
	for i := 0; i < len(runes2); i++ {
		for j := 0; j < len(runes2); j++ {
			if runes2[i] < runes2[j] {
				tem := runes2[i]
				runes2[i] = runes2[j]
				runes2[j] = tem
			}
		}
	}

	if l1 == l2 {
		for i := 0; i < len(runes1); i++ {
			if runes1[i] != runes2[i] {
				fmt.Printf("Not anagram")
				break
			}
		}
		fmt.Printf(" anagram\n")
	} else {
		fmt.Printf("Not anagram\n")
	}
}

func PrimeNumbers() {
	var arr [168]int
	k := 0
	arr[0] = 0
	count := 0
	for i := 3; i <= 1000; i++ {
		prime := 1
		for j := 2; j < i/2; j++ {
			if i%j == 0 {
				prime = 0
				break
			}
		}
		if prime == 1 {
			arr[k] = i
			k++
			count++
		}
	}
	for i := range arr {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Printf("\n ")
	//fmt.Printf("\n%d\n ", count)
}
func AnagramandPal() {
	var arr [168]int
	k := 0
	arr[0] = 0
	count := 0
	for i := 3; i <= 1000; i++ {
		prime := 1
		for j := 2; j < i/2; j++ {
			if i%j == 0 {
				prime = 0
				break
			}
		}
		if prime == 1 {
			arr[k] = i
			k++
			count++
		}
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if checkanagram(arr[i], arr[j]) {
				fmt.Printf("the number %d and %d are annagrams\n", arr[i], arr[j])
			}
		}

	}
	for i := 0; i < len(arr); i++ {
		if Palindrome(arr[i]) {
			fmt.Printf("the number %d is palindrome \n", arr[i])
		}
	}

}
func checkanagram(number int, number2 int) bool {
	A := strconv.Itoa(number)
	str1 := string(sort([]rune(A)))
	B := strconv.Itoa(number2)
	str2 := string(sort([]rune(B)))
	if str1 == str2 {
		return true
	} else {
		return false
	}
}
func sort(s []rune) []rune {
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if s[i] < s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	return s
}
func Palindrome(number int) bool {
	remainder := 0
	reverse := 0
	temp := number
	for {
		remainder = number % 10
		reverse = reverse*10 + remainder
		number /= 10
		if number == 0 {
			break
		}
	}
	if temp == reverse {
		return true
	}
	return false
}
func Insertion() {
	fmt.Printf("Enter the size of the array\n")
	var size int
	fmt.Scanf("%d ", &size)
	arr := make([]string, size)
	fmt.Printf("Enter the string array elements\n")
	for i := range arr {
		fmt.Scanf("%s", &arr[i])
	}

	for i := 0; i < size; i++ {
		j := i
		temp := arr[i]
		for j > 0 && temp < arr[j-1] {
			arr[j] = arr[j-1]
			j = j - 1
		}
		arr[j] = temp
	}
	fmt.Printf("The sorted array elements\n")
	for i := range arr {
		fmt.Printf("%s ", arr[i])
	}
}
func Bubblesort() {
	fmt.Printf("Enter the size of the array\n")
	var size int
	fmt.Scanf("%d ", &size)
	arr := make([]int, size)
	fmt.Printf("Enter the array elements\n")
	for i := range arr {
		fmt.Scanf("%d", &arr[i])
	}

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	fmt.Printf("The sorted array elements\n")
	for i := range arr {
		fmt.Print(arr[i], " ")
	}
}
func Merge(arr []string, l int, m int, r int) {
	n1 := m - l + 1
	n2 := r - m
	Left := make([]string, n1)
	Right := make([]string, n2)
	for i := 0; i < n1; i++ {
		Left[i] = arr[l+i]
	}
	for j := 0; j < n2; j++ {
		Right[j] = arr[m+1+j]
	}
	i := 0
	j := 0
	k := l
	for i < n1 && j < n2 {
		if Left[i] <= Right[j] {
			arr[k] = Left[i]
			i++
		} else {
			arr[k] = Right[j]
			j++
		}
		k++
	}
	for i < n1 {
		arr[k] = Left[i]
		k++
		i++
	}
	for j < n2 {
		arr[k] = Right[j]
		k++
		j++
	}
}

func MergeSort(arr []string, l int, r int) {
	if l < r {
		m := l + (r-l)/2
		MergeSort(arr, l, m)
		MergeSort(arr, m+1, r)
		Merge(arr, l, m, r)
	}
}
