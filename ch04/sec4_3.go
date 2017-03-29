package main

import (
	"fmt"
	"sort"
)

func main() {

	// 4.3. マップ
	{
		ages := make(map[string]int)
		ages["alice"] = 34
		ages["john"] = 20

		ages2 := map[string]int{
			"alice" : 34,
			"john" : 20,
		}

		fmt.Printf("%d\n", ages2["alice"])

		for name, age := range ages {
			fmt.Printf("%s is %d years old.\n", name, age)
		}

		//var names []string
		names := make([]string, 0, len(ages))
		println("names:", len(names), cap(names))
		for name := range ages {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			fmt.Printf("%s is %d\n", name, ages[name])
		}

		var types map[string]int
		println(types == nil)
		println(len(types) == 0)
		// panic.
		//types["aaa"] = 1

		if age, ok := ages["bob"]; !ok {
			println("bob is not found.")
		} else {
			println(age)
		}

		// Set としての利用.
		set := make(map[string]bool)
		set["a"] = true
		set["b"] = true
		if _, ok := set["c"]; !ok {
			println("'c' is not found.")
		}
	}
}
