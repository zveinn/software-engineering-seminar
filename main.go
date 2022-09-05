package main

import (
	"fmt"
	"math"
)

// GLOBAL VARIABLES CAN BE ACCESSED ANYWHERE
var GlobalA string = "This is a global variable"

const GlobalConstantA string = "This is a global constant"

func main() {

	// BASIC VARIABLE TYPES
	var LocalA string = "this is a local variable"
	fmt.Println(GlobalA)
	fmt.Println(LocalA)

	const LocalConstantA string = "this is a local constant"
	fmt.Println(GlobalConstantA)
	fmt.Println(LocalConstantA)

	fmt.Println("Bytes to Binary:")
	var Byte1 byte = 1
	var Byte2 byte = 2
	var Byte4 byte = 4
	var Byte8 byte = 8
	var Byte16 byte = 16
	var Byte32 byte = 32
	var Byte64 byte = 64
	var Byte128 byte = 128
	fmt.Printf("%08b \n", Byte1)
	fmt.Printf("%08b \n", Byte2)
	fmt.Printf("%08b \n", Byte4)
	fmt.Printf("%08b \n", Byte8)
	fmt.Printf("%08b \n", Byte16)
	fmt.Printf("%08b \n", Byte32)
	fmt.Printf("%08b \n", Byte64)
	fmt.Printf("%08b \n", Byte128)

	fmt.Println("Combining bits to represent numbers:")
	var Byte57 byte = 57
	fmt.Printf("%08b \n", Byte57)

	fmt.Println("Shifting Byte1:")
	var Byte1Shifted byte = Byte1 << 1
	fmt.Printf("%08b \n", Byte1Shifted)
	var Byte1ShiftedR byte = Byte1 >> 1
	fmt.Printf("%08b \n", Byte1ShiftedR)

	// SIGNED INTEGERS CAN HAVE NEGATIVE VALUES
	var IntX int8 = 127
	fmt.Printf("Integer:%d Bytes:1 Binary:%08b \n", IntX, IntX)
	var IntXNegative int8 = -127
	fmt.Printf("Integer:%d Bytes:1 Binary:%08b \n", IntXNegative, IntXNegative)

	// UNSIGED INTEGERS CAN NOT HAVE NEGATIVE VALUES
	// INSTEAD THEY CAN HAVE HIGHER POSITIVE VALUES
	var UIntX uint8 = 255
	fmt.Printf("UInt8:%d Bytes:1 Binary:%08b \n", UIntX, UIntX)

	var UInt16 uint16 = math.MaxUint16
	fmt.Printf("UInt16:%d Bytes:2 Binary:%08b \n", UInt16, UInt16)

	var UInt32 uint32 = math.MaxUint32
	fmt.Printf("UInt32:%d Bytes:3 Binary:%08b \n", UInt32, UInt32)

	var UInt64 uint64 = math.MaxUint64
	fmt.Printf("UInt64:%d Bytes:3 Binary:%08b \n", UInt64, UInt64)

}
