package main

import "fmt"

func main() {
	fmt.Println("*********Integers*********")

	var NumberInt8Start = int8(-128)
	var NumberInt8End = int8(127)
	var NumberInt16Start = int16(-32768)
	var NumberInt16End = int16(32767)
	var NumberInt32Start = int32(-2147483648)
	var NumberInt32End = int32(2147483647)
	var NumberInt64Start = int64(-9223372036854775808)
	var NumberInt64End = int64(9223372036854775807)

	fmt.Printf("Type: %T Start: %d End: %d\n", NumberInt8Start, NumberInt8Start, NumberInt8End)
	fmt.Printf("Type: %T Start: %d End: %d\n", NumberInt16Start, NumberInt16Start, NumberInt16End)
	fmt.Printf("Type: %T Start: %d End: %d\n", NumberInt32Start, NumberInt32Start, NumberInt32End)
	fmt.Printf("Type: %T Start: %d End: %d\n", NumberInt64Start, NumberInt64Start, NumberInt64End)

	var NumberUint8Start = uint8(0)
	var NumberUint8End = uint8(255)
	var NumberUint16Start = uint8(0)
	var NumberUint16End = uint16(65535)
	var NumberUint32Start = uint32(0)
	var NumberUint32End = uint32(4294967295)
	var NumberUint64Start = uint64(0)
	var NumberUint64End = uint64(4294967295)

	fmt.Printf("Type: %T Start: %d End: %d\n", NumberUint8Start, NumberUint8Start, NumberUint8End)
	fmt.Printf("Type: %T Start: %d End: %d\n", NumberUint16Start, NumberUint16Start, NumberUint16End)
	fmt.Printf("Type: %T Start: %d End: %d\n", NumberUint32Start, NumberUint32Start, NumberUint32End)
	fmt.Printf("Type: %T Start: %d End: %d\n", NumberUint64Start, NumberUint64Start, NumberUint64End)

	fmt.Println("************************************")
	fmt.Println()

	fmt.Println("*********Float*********")

	var NumberFloat32 = float32(100.555559)

	fmt.Printf("Type %T Value %f\n", NumberFloat32, NumberFloat32)
	fmt.Println(NumberFloat32)

	fmt.Println("************************************")
	fmt.Println()
}
