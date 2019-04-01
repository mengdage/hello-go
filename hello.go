package main

import (
	"fmt"
)

type ByteSlice []byte

func (b ByteSlice) Append(data []byte) []byte {
	sliceLen := len(b)
	if sliceLen+len(b) > cap(b) {
		newb := make(ByteSlice, (sliceLen+len(b))*2)

		copy(newb, b)
		b = newb
	}

	b = b[:(sliceLen + len(data))]

	copy(b[sliceLen:], data)

	return b
}

func (b *ByteSlice) Append2(data []byte) {
	slice := *b
	sliceLen := len(slice)
	if sliceLen+len(data) > cap(slice) {
		newslice := make(ByteSlice, (sliceLen+len(data))*2)

		copy(newslice, slice)
		slice = newslice
	}

	slice = slice[:(sliceLen + len(data))]
	copy(slice[sliceLen:], data)
	*b = slice
}

func main() {
	// bs := ByteSlice{1, 2, 3,}
	bs := make(ByteSlice, 3, 4)
	nbs := bs.Append([]byte{4, 5, 6})
	fmt.Printf("%v\n", nbs)
	fmt.Println(len(nbs), cap(nbs))

	bs2 := &ByteSlice{1, 2, 3, 4}
	bs2.Append2([]byte{4, 5, 6})
	fmt.Printf("%v\n", bs2)
	fmt.Println(len([]byte(*bs2)), cap([]byte(*bs2)))
}
