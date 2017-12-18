package huffman

import (
	_ "fmt"
	"testing"
)

func TestWriteBits(t *testing.T) {
	packer := NewBitPacker()
	packer.WriteBits(0x1, 1)
	packer.WriteBits(0x2, 2)
	packer.WriteBits(0x3, 2)
	packer.WriteBits(0x4, 3)
	packer.WriteBits(0x5, 3)
	packer.WriteBits(0x6, 3)
	expect := []byte{0xdc, 0xb8}
	// Now the bits should be 11011100 10111000
	// nbytes = 1, nbit = 6
	if packer.nbyte != 1 {
		t.Errorf("Wrong number of bytes, expected %d, got %d", 1, packer.nbyte)
	}
	if packer.nbit != 6 {
		t.Errorf("Wrong number of bits, expected %d, got %d", 6, packer.nbit)
	}
	if packer.Size() != 14 {
		t.Errorf("Wrong size: expected %d, got %d", 14, packer.Size())
	}
	data := packer.Data()
	for i := 0; i <= int(packer.nbyte); i++ {
		if data[i] != expect[i] {
			t.Errorf("Wrong byte at %d: expected %x, got %x", i, expect[i], packer.data[i])
		}
	}
}

func TestWriteRead(t *testing.T) {
	data := [][]uint32{
		{0x1, 1}, {0x2, 2}, {0x3, 2}, {0x4, 3},
		{0x8, 4}, {0x7, 3}, {0x6, 3}, {0x5, 3},
		{0x9, 4}, {0xa, 4}, {0xb, 4}, {0xc, 4},
		{0x10, 5}, {0xf, 4}, {0xe, 4}, {0xd, 4},
		{0x18, 5}, {0x1f, 5}, {0x42, 7}, {0x17, 7},
		{0x12, 5}, {0x31, 6}, {0x23, 7}, {0x9, 7},
		{0x23, 6}, {0x5f, 7}, {0x14, 7}, {0x5f, 7},
		{0x20, 6}, {0x1a, 6}, {0x45, 7}, {0x5d, 7},
		{0x34, 7}, {0x4f, 7}, {0x46, 7}, {0x5e, 7},
	}
	packer := NewBitPacker()
	for i := 0; i < len(data); i++ {
		packer.WriteBits(uint64(data[i][0]), data[i][1])
	}
	/*
		d := packer.Data()
		for i := 0; i < len(data); i++ {
			fmt.Printf("%d: %02x\n", i, d[i])
		}*/
	unpacker := NewBitUnpacker(packer.Data(), packer.Size())
	if unpacker == nil {
		t.Errorf("Error in creating unpacker")
		return
	}
	for i := 0; i < len(data); i++ {
		bits, err := unpacker.ReadBits(data[i][1])
		if err != nil {
			t.Errorf("Error in reading bits: %s", err.Error())
		}
		if bits != uint64(data[i][0]) {
			t.Errorf("Mismatched unpacked result at %d: expected %d, got %d",
				i, data[i][0], bits)
		}
	}
}
