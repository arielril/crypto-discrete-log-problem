package discrete

import (
	"fmt"
	"math/big"
)

func Compute(ps, gs, hs string) string {
	p, g, h := new(big.Int),new(big.Int),new(big.Int)
	
	B := new(big.Int)

	B = B.Exp(big.NewInt(2), big.NewInt(20), nil)
	p.SetString(ps, 10)
	g.SetString(gs, 10)
	h.SetString(hs, 10)
	
	// x_0, x_1 := 0,0

	hTableSize := new(big.Int)
	hTableSize = hTableSize.Exp(big.NewInt(2), big.NewInt(20), nil)
	hTable := make(map[int]*big.Int, hTableSize.Int64()+1)

	for i:=0; i < (int) (hTableSize.Int64()+1); i++ {
		ax := new(big.Int)
		// ! implement the hash table
		hTable[i] = ax.Exp(big.NewInt(2),big.NewInt(int64(i)), nil)
	}

	fmt.Printf("hash table: %#v\n", hTable)

	

	return B.String()
}
