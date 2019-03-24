package short

import "testing"

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
