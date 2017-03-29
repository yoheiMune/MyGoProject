package main

import (
	"time"
	"fmt"
)

type Employee struct {
	ID				int
	Name, Address 	string      // まとめることもできる.
	DoB 			time.Time
	Position 		string
	Salary 			int
	ManagerID 		int
}

func main() {

	// 4.4. 構造体

	{
		var dilbert Employee
		println(&dilbert)

		var dilbert2 *Employee
		println(dilbert2)
		println(dilbert2 == nil)
		//dilbert2.Salary = 1000  // panic : nil pointer access.

		println("-----------------------")

		dilbert.Salary -= 5000 // 減給

		dilbert.Position = "Normal Employee"
		position := &dilbert.Position
		println(position, *position)
		*position = "Senior " + *position
		println(position, *position)

		var employeeOfTheMonth *Employee = &dilbert
		println(employeeOfTheMonth, &employeeOfTheMonth.Position) // ポインタ
		employeeOfTheMonth.Position += " (proactive team player)"
		println(employeeOfTheMonth.Position)

		employee := EmployeeById(12)
		println(employee, &employee.ID, employee.ID)

		employee2 := EmployeeById2(21)
		//println(employee2, *employee2.ID)
		println(&employee2, &employee2.ID, employee2.ID)
	}

	main2()
}

// ポインターを返す.
func EmployeeById(id int) *Employee {
	println("------------")
	var employee Employee
	employee.ID = id
	println(&employee, &employee.ID)
	return &employee
}

// 構造体のコピーを返す.
func EmployeeById2(id int) Employee {
	println("------------")
	var employee Employee
	employee.ID = id
	println(&employee, &employee.ID)
	return employee
}

// ======================================================================

type tree struct {
	value 			int
	left, right 	*tree
}


// Sort は values 内の値をその中でソートします.
func Sort(values []int) {
	var root *tree
	for i, v := range values {
		fmt.Printf("%d : %d\n", i, v)
		root = add(root, v)
		fmt.Printf("values1: %v\n", &values)
		appendValues(values[:0], root)
	}
}

// appendValues は t の要素を values の正しい順序に追加し、結果のスライスを返却します.
func appendValues(values []int, t *tree) []int {
	fmt.Printf("values2: %x\n", values)
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main2() {

	nums := []int{5, 2, 7, 8, 1}
	Sort(nums)
	fmt.Printf("%v\n", nums)




}




























