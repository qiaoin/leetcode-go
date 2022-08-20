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
