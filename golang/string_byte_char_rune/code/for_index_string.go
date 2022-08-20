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
