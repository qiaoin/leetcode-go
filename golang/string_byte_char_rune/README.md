[back to contents](https://github.com/qiaoin/leetcode-go#leetcode-go)

# String

## 如何遍历 string？

### for-range 遍历

当使用 `for-range` 遍历 string 值时，char 的类型为 `rune` (alias for `int32`)，代表一个 Unicode 的码点（a Unicode code point），，等同于将 `string` 转换为 `[]rune` 进行处理。Go 中所有的 string 字面量均为 UTF8 编码。[代码示例](code/for_range_string.go)。

```go
package main

import "fmt"

func main()  {
	str := "hello"
	for _, c := range str {
		fmt.Printf("%U, %q = %v, type = %T\n", c, c, c, c)
	}

	fmt.Println()

	runeArr := []rune(str)
	for _, c := range runeArr {
		fmt.Printf("%U, %q = %v, type = %T\n", c, c, c, c)
	}

	//Output:
	//U+0068, 'h' = 104, type = int32
	//U+0065, 'e' = 101, type = int32
	//U+006C, 'l' = 108, type = int32
	//U+006C, 'l' = 108, type = int32
	//U+006F, 'o' = 111, type = int32
	//
	//U+0068, 'h' = 104, type = int32
	//U+0065, 'e' = 101, type = int32
	//U+006C, 'l' = 108, type = int32
	//U+006C, 'l' = 108, type = int32
	//U+006F, 'o' = 111, type = int32
}
```

### for index 遍历

当使用 `for-index` 遍历 string 值时，char 的类型为 `byte` (alias for `uint8`)，等同于将 `string` 转换为 `[]byte` 进行处理。[代码示例](code/for_index_string.go)。

```go
package main

import "fmt"

func main() {
	str := "hello"

	for i := 0; i < len(str); i++ {
		c := str[i]
		fmt.Printf("%U, %q = %v, type = %T\n", c, c, c, c)
	}

	fmt.Println()

	byteArr := []byte("hello")
	for _, c := range byteArr {
		fmt.Printf("%U, %q = %v, type = %T\n", c, c, c, c)
	}

	//Output:
	//U+0068, 'h' = 104, type = uint8
	//U+0065, 'e' = 101, type = uint8
	//U+006C, 'l' = 108, type = uint8
	//U+006C, 'l' = 108, type = uint8
	//U+006F, 'o' = 111, type = uint8
	//
	//U+0068, 'h' = 104, type = uint8
	//U+0065, 'e' = 101, type = uint8
	//U+006C, 'l' = 108, type = uint8
	//U+006C, 'l' = 108, type = uint8
	//U+006F, 'o' = 111, type = uint8
}
```

### 处理中文字符串遍历

在为中文字符串时需要特别注意

- `for-range` 将中文每一个汉字作为一个 c 输出，表示为 `rune` 类型
- `for-index` 将中文每一个汉字包含的三个字节分别进行输出，表示为 `byte` 类型

## 修改字符串

[5730. 将所有数字用字符替换](https://leetcode-cn.com/problems/replace-all-digits-with-characters/)

编译错误，`Cannot assign to s[i],  (strings are immutable)`，为什么呢？

```go
func ReplaceDigits(s string) string {
	for i := 1; i < len(s); i += 2 {
		fmt.Printf("i = %v, arr[i] = %v, type = %T\n", i, s[i], s[i])
		s[i] = s[i-1] + byte(int(s[i])-48)
	}

	return s
}
```

正确做法：

- 将字符串转换为 `[]byte` (alias for `uint8`)
- 在与 `int` 做加减运算时，将 `byte` 转换为 `int`
- 在与 `byte` 做加减运算时，将 `int` 转换为 `byte`
- 将 `[]byte` 转换为字符串

```go
func ReplaceDigits(s string) string {
	arr := []byte(s)
	for i := 1; i < len(arr); i += 2 {
		arr[i] = arr[i-1] + byte(int(arr[i])-48)
	}

	return string(arr)
}
```

