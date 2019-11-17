package online_shopping

import "testing"

func TestStructEmbedding(t *testing.T) {
	rice := &Rice{
		Product: Product{
			Name:  "rice",
			Price: 1.000,
		},
	}

	t.Log(rice.GetName())
	t.Log(rice.GetPrice())
}
