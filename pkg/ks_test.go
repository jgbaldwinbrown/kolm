package kolm

import (
	"testing"
	"slices"
	"fmt"
)

func TestKolmogorovSmirnovChi2(t *testing.T) {
	data := []float64{3, 4, 5, 6, 7, 7, 7, 7}
	res, e := KolmogorovSmirnovChi2(slices.Values(data))
	if e != nil {
		panic(e)
	}
	fmt.Printf("%#v\n", res)
}
