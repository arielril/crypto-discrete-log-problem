package discrete

import (
	"fmt"

	"github.com/ericlagergren/decimal"
	"github.com/projectdiscovery/gologger"
)

/**
 * Compute computes the value of `x=x0*B+x1`
 */
func Compute(ps, gs, hs string) string {
	p, g, h := decimal.New(0,0),decimal.New(0,0),decimal.New(0,0)
	p, _ = p.SetString(ps)
	g, _ = g.SetString(gs)
	h, _ = h.SetString(hs)

	gologger.Info().Msgf("using h => %v", h)
	gologger.Info().Msgf("using p => %v", p)
	gologger.Info().Msgf("using g => %v", g)

	B := new(decimal.Big)

	B = B.Context.Pow(B, decimal.New(2, 0), decimal.New(20, 0))
	hashTableSize, _ := B.Int64()
	
	gologger.Info().Msgf("hash table size -> %v", hashTableSize)

	hTable := make(map[int64]int64, hashTableSize)

	// creates the hash table for the left hand side of the equation `h/(g^x1) = (g^B)^x0`
	for x1 := int64(0); x1 <= hashTableSize; x1 += 1 {
		x1Big := decimal.New(x1, 0)

		leftHandSide := decimal.New(0,0)
		
		// (g^B)^x0
		leftHandSideQuocient := decimal.New(0,0)
		_ = leftHandSideQuocient.Context.Pow(leftHandSideQuocient, g, x1Big)

		_ = leftHandSide.Quo(h, leftHandSideQuocient)
		gologger.Debug().Msgf("executing the division %v / %v == %v \n", h, leftHandSideQuocient, leftHandSide)
		
		lhsInt64, _ := leftHandSide.Int64()
		x1BigInt64, _ := x1Big.Int64()

		hTable[lhsInt64] = x1BigInt64
		gologger.Debug().Msgf("[.] h table [%v] ==> %v", leftHandSide, x1Big)
	}

	// fmt.Printf("h table --> %v\n", hTable)

	foundX0 := int64(0);
	foundX1 := int64(0);

	// computes the right hand side of the equation `h/(g^x1) = (g^B)^x0`
	for x0 := int64(0); x0 <= hashTableSize; x0 += 1 {
		x0Big := decimal.New(x0, 0)

		// (g^B)^x0
		rightHandSide := decimal.New(0,0)
		_ = rightHandSide.Context.Pow(rightHandSide, g, B)
		_ = rightHandSide.Context.Pow(rightHandSide, rightHandSide, x0Big)

		gologger.Debug().Msgf("right hand side result --> %v", rightHandSide)
		rhsInt64, _ := rightHandSide.Int64()
		x1, ok := hTable[rhsInt64]
		if ok {
			gologger.Info().Msgf("found the right hand side in the hash table --> %v", x1)
			foundX0 = x0
			foundX1 = x1

			gologger.Info().Msgf("(x0, x1) <-> (%v, %v) ==> valid? [%v]", foundX0, foundX1, verifyResult(h, g, B, x0Big, decimal.New(x1, 0)))

			// x=x0*B+x1
			x := decimal.New(0,0)
			x.Mul(x0Big, B)
			x.Add(x, decimal.New(x1, 0))

			// hVerifier := decimal.New(0,0)
			// hVerifier.Context.Pow(hVerifier, g, x)

			
			// fmt.Printf("g [%v], x [%v] -> hverifier [%v]\n", g, x, hVerifier)
			// return fmt.Sprintf("h = g^x = %v ==> in Zp | h = g^x? [%v]", hVerifier, hVerifier.Cmp(h))
			return fmt.Sprintf("x = %v", x)
		}
	}

	return "x value not found"
}

func verifyResult(h, g, B, x0, x1 *decimal.Big) bool {
	// h/(g^x1) = (g^B)^x0
	
	leftQuocient := decimal.New(1,0)
	leftQuocient.Context.Pow(leftQuocient, g, x1)

	left := decimal.New(0,0)
	left.Quo(h, leftQuocient)

	right := decimal.New(0,0)
	right.Context.Pow(right, g, B)
	right.Context.Pow(right, right, x0)

	leftInt, _ := left.Int64()
	rightInt, _ := right.Int64()
	return leftInt == rightInt
}
