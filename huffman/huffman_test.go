package huffman

import (
	_ "fmt"
	"testing"
)

func TestHuffmanEncodeDecode(t *testing.T) {
	data := []int{
		3, 14, 15, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9, 3, 2, 3, 8, 4, 6, 2, 6,
	}
	code := &HuffmanCode{
		[]Pair{
			Pair{25, 5}, /*   0: (0,-1) 11001 */
			Pair{0, 1},  /*   1: (0, 0) 0 */
			Pair{24, 5}, /*   2: (0, 1) 11000 */

			Pair{52, 6}, /*   3: (1,-1) 110100 */
			Pair{2, 2},  /*   4: (1, 0) 10 */
			Pair{27, 5}, /*   5: (1, 1) 11011 */

			Pair{428, 9}, /*   6: (2,-1) 110101100 */
			Pair{7, 3},   /*   7: (2, 0) 111 */
			Pair{215, 8}, /*   8: (2, 1) 11010111 */

			Pair{3432, 12}, /*   9: (3,-1) 110101101000 */
			Pair{106, 7},   /*  10: (3, 0) 1101010 */
			Pair{1717, 11}, /*  11: (3, 1) 11010110101 */

			Pair{13732, 14}, /*  12: (4,-1) 11010110100100 */
			Pair{859, 10},   /*  13: (4, 0) 1101011011 */
			Pair{6867, 13},  /*  14: (4, 1) 1101011010011 */

			Pair{109868, 17}, /*  15: (5,-1) 11010110100101100 */
			Pair{27466, 15},  /*  16: (5, 0) 110101101001010 */
			Pair{54935, 16},  /*  17: (5, 1) 1101011010010111 */

			Pair{439479, 19}, /*  18: (6,-1) 1101011010010110111 */
			Pair{219738, 18}, /*  19: (6, 0) 110101101001011010 */
			Pair{439478, 19}, /*  20: (6, 1) 1101011010010110110 */
		},
		[]Triple{
			Triple{1, 2, -1},   /*   0: */
			Triple{-1, -1, 1},  /*   1: (0, 0)  1 bit  */
			Triple{3, 4, -1},   /*   2: */
			Triple{-1, -1, 4},  /*   3: (1, 0)  2 bits */
			Triple{5, 40, -1},  /*   4: */
			Triple{6, 9, -1},   /*   5: */
			Triple{7, 8, -1},   /*   6: */
			Triple{-1, -1, 2},  /*   7: (0, 1)  5 bits */
			Triple{-1, -1, 0},  /*   8: (0,-1)  5 bits */
			Triple{10, 39, -1}, /*   9: */
			Triple{11, 12, -1}, /*  10: */
			Triple{-1, -1, 3},  /*  11: (1,-1)  6 bits */
			Triple{13, 14, -1}, /*  12: */
			Triple{-1, -1, 10}, /*  13: (3, 0)  7 bits */
			Triple{15, 38, -1}, /*  14: */
			Triple{16, 17, -1}, /*  15: */
			Triple{-1, -1, 6},  /*  16: (2,-1)  9 bits */
			Triple{18, 37, -1}, /*  17: */
			Triple{19, 36, -1}, /*  18: */
			Triple{20, 21, -1}, /*  19: */
			Triple{-1, -1, 9},  /*  20: (3,-1) 12 bits */
			Triple{22, 35, -1}, /*  21: */
			Triple{23, 24, -1}, /*  22: */
			Triple{-1, -1, 12}, /*  23: (4,-1) 14 bits */
			Triple{25, 26, -1}, /*  24: */
			Triple{-1, -1, 16}, /*  25: (5, 0) 15 bits */
			Triple{27, 34, -1}, /*  26: */
			Triple{28, 29, -1}, /*  27: */
			Triple{-1, -1, 15}, /*  28: (5,-1) 17 bits */
			Triple{30, 31, -1}, /*  29: */
			Triple{-1, -1, 19}, /*  30: (6, 0) 18 bits */
			Triple{32, 33, -1}, /*  31: */
			Triple{-1, -1, 20}, /*  32: (6, 1) 19 bits */
			Triple{-1, -1, 18}, /*  33: (6,-1) 19 bits */
			Triple{-1, -1, 17}, /*  34: (5, 1) 16 bits */
			Triple{-1, -1, 14}, /*  35: (4, 1) 13 bits */
			Triple{-1, -1, 11}, /*  36: (3, 1) 11 bits */
			Triple{-1, -1, 13}, /*  37: (4, 0) 10 bits */
			Triple{-1, -1, 8},  /*  38: (2, 1)  8 bits */
			Triple{-1, -1, 5},  /*  39: (1, 1)  5 bits */
			Triple{-1, -1, 7},  /*  40: (2, 0)  3 bits */
		},
	}
	encoder := NewHuffmanEncoder(code)
	for i := 0; i < len(data); i++ {
		encoder.Update(data[i])
	}
	result := encoder.Digest()
	// for i := 0; i < len(result); i++ {
	// 	fmt.Printf("%02x ", result[i])
	// }
	// fmt.Println()
	decoder := NewHuffmanDecoder(code, result)
	for i := 0; i < len(data); i++ {
		next, err := decoder.Next()
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
			return
		}
		if next != data[i] {
			t.Errorf("Wrong result at %d: expected %d, got %d", i, data[i], next)
		}
	}
}
