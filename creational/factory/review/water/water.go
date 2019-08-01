package water

type from int

const (
	// The constant will be available to the user.
	// in behalf for the concrete type
	Faucet from = iota
	Well
)

// Available from values are:
// water.Faucet
// water.Well
func GetWater(f from) WaterGetter {
	var wg WaterGetter
	switch f {
	case Faucet:
		wg = faucet{}
	case Well:
		wg = well{}
	}
	return wg
}

// Define the commont method interface

// WaterGetter is a general method
// for getting a water.
//
// The Implementation will vary
// depending on the structs
// that will implement it.
type WaterGetter interface {
	Get() string
}

type faucet struct{}

func (faucet) Get() string {
	return "Got water from faucet"
}

type well struct{}

func (well) Get() string {
	return "Got water from well"
}
