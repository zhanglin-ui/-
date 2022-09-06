# Slice底层实现
## 切片和数组
~~~go
var num = [2]int{100, 200}
sli := num[:]
fmt.Println(len(sli), cap(sli), sli)
fmt.Printf("%p\n%p", &num, sli)

//输出
2 2 [100 200]
0xc0000180a0
0xc0000180a0

var num = [10]int{100, 200}
sli := num[0:2]
fmt.Println(len(sli), cap(sli), sli)
fmt.Printf("%p\n%p", &num, sli)

//输出
2 10 [100 200]
0xc00001a230
0xc00001a230
~~~

~~~go
func main() {
	arrayA := [2]int{100, 200}
	var arrayB [2]int

	arrayB = arrayA

	fmt.Printf("arrayA : %p, %v\n", &arrayA, arrayA)
	fmt.Printf("arrayB : %p, %v\n", &arrayB, arrayB)

	testArry(arrayA)
}

func testArry(x [2]int) {
	fmt.Printf("func Arrry : %p, %v\n", &x, x)
}

result:
arrayA : 0xc0000120a0, [100 200]
arrayB : 0xc0000120b0, [100 200]
func Arrry : 0xc000012100, [100 200]

~~~

三个地址互不相同，go数组赋值和函数传参都是值复制的

~~~go
func main() {
	arrayA := [2]int{100, 200}
	// testArryPoint(&arrayA)

	arrayB := arrayA[:]
	testArryPoint(&arrayB)

	fmt.Printf("arrayA : %p, %v\n", &arrayA, arrayA)
	fmt.Printf("arrayB : %p, %v\n", &arrayB, arrayB)
}

func testArryPoint(x *[]int) {
	fmt.Printf("func Array : %p, %v\n", x, *x)
	(*x)[1] += 100
}

result:
func Array : 0xc000004078, [100 200]
arrayA : 0xc0000120a0, [100 300]
arrayB : 0xc000004078, [100 300]
~~~

使用指针，传进去的是一个8字节的内存指针，这样可以高效的利用内存

## 切片的数据结构
在go的reflect中存在一个与之对应的数据结构SliceHeader
~~~go
type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

var ptr unsafe.Pointer
length := 10

//创建一个slice
var o []byte
SliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&o))
SliceHeader.Len = length
SliceHeader.Cap = length
SliceHeader.Data = uintptr(ptr)
~~~

## 创建切片

创建切片有两种形式，make 创建切片，空切片。
### make 和切片字面量
