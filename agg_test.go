package snarksnarf

import (
	"testing"

	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/test"
)

type Circuit struct {
	Base          frontend.Variable `gnark:",public`
	End           frontend.Variable `gnark:",public`
	nbConstraints int
}

func (c *Circuit) Define(api frontend.API) error {
	output := c.Base
	for i := 0; i < c.nbConstraints; i++ {
		output = api.Mul(output, c.Base)
	}
	api.AssertIsEqual(output, c.End)
	return nil
}

func NewTestCircuit(nbConstraints int) *Circuit {
	base := 3
	out := base
	for i := 0; i < nbConstraints; i++ {
		out = out * base
	}
	return &Circuit{
		Base:          base,
		End:           out,
		nbConstraints: nbConstraints,
	}
}

func TestCircuitGroth16(t *testing.T) {
	assert := test.NewAssert(t)
	assert.ProverSucceeded(&Circuit{}, NewTestCircuit(3))
}
