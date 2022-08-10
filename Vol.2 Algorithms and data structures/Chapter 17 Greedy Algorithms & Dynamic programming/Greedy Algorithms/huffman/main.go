package main

import "fmt"

func main() {
	str := "East or West home is best"
	fmt.Printf("The original string is: %s\n", str)
	ht := CreateHuffnamTree([]byte(str))
	ht.PrintHuffmanCode()
	encodedStr := ht.GetEncodedString()
	fmt.Printf("The encoded string is: %s\n", encodedStr)
	decodedStr := ht.GetDecodedString(encodedStr)
	fmt.Printf("The decoded string is: %s\n", decodedStr)
}

// func main() {
// 	str := "AAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
// 	fmt.Printf("The original string is: %s\n", str)
// 	ht := CreateHuffnamTree([]byte(str))
// 	ht.PrintHuffmanCode()
// 	encodedStr := ht.GetEncodedString()
// 	fmt.Printf("The encoded string is: %s\n", encodedStr)
// 	decodedStr := ht.GetDecodedString(encodedStr)
// 	fmt.Printf("The decoded string is: %s\n", decodedStr)
// }
