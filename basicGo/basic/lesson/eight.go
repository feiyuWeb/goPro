package lesson

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"net/url"
	"os"
	"sort"
	"strings"
)

func Eight()  {
	//sum(1,2,3)
	//
	//nums := []int{1,2,3,4}
	//sum(nums...)

	/*--闭包--*/
	//nextInt := intSeq()
	//fmt.Println(nextInt())
	//fmt.Println(nextInt())

	/*--递归--*/
	//fmt.Println(fact(0),"--")
	//fmt.Println(fact(5))

	/*--接口--*/
	//r := rect{width: 3, height: 4}
	//c := circle{radius: 5}
	//
	//measure(r)
	//measure(c)

	/*排序*/
	strs := []string{"c","a","b"}
	sort.Strings(strs)
	fmt.Println("strings",strs)

	ints := []int{7,2,4}
	sort.Ints(ints)
	fmt.Println(ints)

	// rand.Intn 返回一个随机的整数 n
	fmt.Println(rand.Intn(100))

	/*
	URL解析
	我们将解析这个 URL 示例，它包含了一个 scheme， 认证信息，主机名，端口，路径，查询参数和片段
	*/
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	u,err := url.Parse(s)
	if err !=nil{
		panic(err)
	}

	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.Host)
	fmt.Println(u.Path)
	fmt.Println(u.RawQuery)

	// 写文件
	fileURL,_ := os.Getwd()
	path := strings.Replace(fileURL, "\\", "/", -1)	// 替换\\为/
	fmt.Println(path,"--")
	d1 := []byte("hello\ngo\n")
	errs := ioutil.WriteFile(path + "/lesson/file/aa.txt",d1,0644)
	check(errs)

	f,err:=os.Create(path + "/lesson/file/bb.txt")
	check(err)

	defer f.Close()
}

func check(e error)  {
	if e!=nil{
		panic(e)
	}
}

// 变参函数
func sum(nums ...int){
	fmt.Println(nums)
	total := 0
	for _,num := range nums{
		total += num
	}
	fmt.Println(total)
}

// 闭包
func intSeq() func() int{
	i := 0
	return func() int {
		i +=1
		return i
	}
}

// 递归
func fact(n int) int{
	if n==0{
		return 1
	}
	return n*fact(n-1) // 5 * 4 * 3 * 2 * 1
}


// 接口

// 这里是一个几何体的基本接口。
type geometry interface {
	area() float64
	perim() float64
}

// 在我们的例子中，我们将在类型 `rect` 和 `circle` 上实现
// 这个接口
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// 要在 Go 中实现一个接口，我们就需要实现接口中的所有
// 方法。这里我们在 `rect` 上实现了 `geometry` 接口。
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// `circle` 的实现。
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 如果一个变量具有接口类型，那么我们可以调用指定接口中的
// 方法。这里有一个通用的 `measure` 函数，利用它来在任
// 何的 `geometry` 上工作。
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}



