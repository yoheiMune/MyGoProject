package main

func main() {

	// 2.3.2 ポインタ
	{
		x := 1
		p := &x
		println(p)   // ポインタ
		println(*p)  // 値

		*p = 2
		println(*p) // 2
	}

	// ポインタの比較
	{
		var a, b int
		println(&a == &a, &a == &b, &a == nil)
	}

	// 関数がポインタを返す.
	{
		var p = f()
		println(p)
		println(f())
	}

	// 関数へのポインタ渡し.
	{
		v := 1
		incr(&v)
		println(incr(&v))
	}

	// 2.3.3 new関数.
	{
		p := new(int)
		println(*p) // 0
		*p = 2
		println(*p) // 2

		a := new(int)
		b := new(int)
		println(a == b, a, b) // false
	}
}

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}
