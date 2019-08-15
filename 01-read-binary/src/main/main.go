package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

type ThreeIntegers struct {
	First  uint8
	Second uint8
	Third  uint8
}

func main() {
	readBinaryArray("files/OneTwoThree")
}

func readBinaryArray(path string) {
	fmt.Println("readBinaryArray()")

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("error occured 'os.Open()'")
		panic(err)
	}

	var threeIntegers ThreeIntegers
	errb := binary.Read(file, binary.LittleEndian, &threeIntegers)
	if errb != nil {
		fmt.Println("error occured 'binary.Read()'")
		panic(errb)
	}

	fmt.Println("first  : ", threeIntegers.First)
	fmt.Println("second : ", threeIntegers.Second)
	fmt.Println("third  : ", threeIntegers.Third)
}
