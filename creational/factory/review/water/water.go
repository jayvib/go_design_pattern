package water

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
