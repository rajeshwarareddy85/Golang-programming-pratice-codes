//package main

// // // // import "fmt"

// // // // func main() {

// // // // 	score := 98
// // // // 	switch {
// // // // 	case score >= 90:
// // // // 		fmt.Println("good")
// // // // 	case score >= 80:
// // // // 		fmt.Println("avg")
// // // // 	case score >= 70:
// // // // 		fmt.Println("loss")
// // // // 	}
// // // // }

// // // // package main

// // // // import "fmt"

// // // // func main() {
// // // // 	ages := make(map[string]int)
// // // // 	ages["raju"] = 23
// // // // 	ages["ram"] = 34
// // // // 	ages["chaithu"] = 33
// // // // 	ages["ramesh"] = 24

// // // // 	for name, ages := range ages {
// // // // 		fmt.Println(name, "->", ages)
// // // // 	}

// // // // 	for name := range ages {
// // // // 		fmt.Println(name)
// // // // 	}

// // // // 	for _, age := range ages {
// // // // 		fmt.Println(age)
// // // // 	}

// // // // }

// // // package main

// // // import "fmt"

// // // func main() {

// // // 	ages := map[string]int{
// // // 		"raju":      23,
// // // 		"ram":       24,
// // // 		"chaitanya": 20,
// // // 		"hari":      23,
// // // 	}

// // // 	for name, age := range ages {
// // // 		fmt.Println(name, "->", age)
// // // 	}

// // // 	for _, age := range ages {
// // // 		fmt.Println(age)
// // // 	}

// // // 	for name := range ages {
// // // 		fmt.Println(name)
// // // 	}

// // // }

// // //

// // package main

// // import "fmt"

// // func main() {
// // 	ages := make(map[string]int)

// // 	var n int

// // 	fmt.Println("enter the length of map:")
// // 	fmt.Scan(&n)

// // 	for i := 0; i < n; i++ {
// // 		var name string
// // 		var age int

// // 		fmt.Println("enter the name:")
// // 		fmt.Scan(&name)

// // 		fmt.Println("enter the age:")
// // 		fmt.Scan(&age)

// // 		ages[name] = age
// // 	}
// // 	for name, age := range ages {
// // 		fmt.Println(name, "->", age)
// // 	}
// // }

// package main

// import (
// 	"fmt"
// 	"os"
// 	"bufio"
// 	"strings"
// )

// func main(){
// 	reader:= bufio.NewReader(os.Stdin)
// 	details := make(map[string]int)

// 	var n int
// 	fmt.Println("enter the count:")
// 	fmt.Scan(&n)
//     reader.ReadString("\n")

// 	for i:=0;i<n;i++{
//       var name string
// 	  var age int
// 	  fmt.Println("enter the name:")
// 	  name,_ := reader.ReadString("\n")
// 	  name=strings.TrimSpace(name)
// 	}

// }

// package main

// import "fmt"

// func addsub(a, b int) (int, int) {
// 	return a + b, a - b
// }

// func main() {
// 	add, sub := addsub(12, 79)
// 	fmt.Println(add)
// 	fmt.Println(sub)
// }

// package main

// import "fmt"

// func sum(nums ...int) int {
// 	total := 0
// 	for _, n := range nums {
// 		total += n
// 	}
// 	return total
// }

// func main() {
// 	fmt.Println(sum(99, 90, 87, 66, 63, 60))
// }

// package main

// import "fmt"

// func main() {
// 	result := func(a, b int) int {
// 		return a + b
// 	}

// 	fmt.Println(result(22, 54))
// }

// package main

// import "fmt"

// func main() {
// 	func() {
// 		fmt.Println("cool morning")
// 	}()
// }

// closure is a function value that references variables from outside its body.
// The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
// package main

// import "fmt"

// func main() {
// 	count := 0

// 	increament := func() {
// 		count++
// 		fmt.Println(count)
// 	}

// 	increament()
// 	increament()
// 	increament()

// }

// package main

// import "fmt"

// func main() {
// 	fmt.Println("start")
// 	defer fmt.Println("middle")
// 	fmt.Println("end")
// }

// package main

// import "fmt"

// func main() {
// 	var n int
// 	fmt.Println("enter the array size:")
// 	fmt.Scan(&n)

// 	var index int
// 	var value int

// 	var arr []int
// 	//var s []int

// 	fmt.Println("enter the index:")
// 	fmt.Scan(&index)

// 	fmt.Println("Enther the value:")
// 	fmt.Scan(&value)

// 	for i := 0; i < n; i++ {
// 		var ele int
// 		fmt.Scan(&ele)
// 		arr = append(arr, ele)
// 	}

// 	// for i := 0; i < n; i++ {
// 	// 	var m int
// 	// 	fmt.Scan(&m)
// 	// 	s = append(s, m)
// 	// }

// 	//arr = append(arr, s...)

// 	arr = append(arr[:index], append([]int{value}, arr[index:]...)...)

// 	for i := 0; i < len(arr); i++ {
// 		fmt.Print(arr[i], ",")
// 	}
// 	fmt.Println()
// 	arr = append(arr[:index], arr[index+1:]...)
// 	for i := 0; i < len(arr); i++ {
// 		fmt.Print(arr[i], ",")
// 	}

// }

package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr []int
	var n int
	fmt.Println("enter the size:")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)
		arr = append(arr, m)
	}

	sort.Ints(arr)
	fmt.Println(arr)

	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	fmt.Println(arr)

}
