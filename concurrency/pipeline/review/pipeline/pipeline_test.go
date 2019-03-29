package pipeline

import "testing"

func TestLaunchPipeline(t *testing.T) {
	tableTest := [][]int{
		{3, 12},
		{5, 30},
	}

	var res int
	for _, test := range tableTest {
		res = LaunchPipeline(test[0])
		if res != test[1] {
			t.Fail()
		}
		t.Logf("%d == %d\n", res, test[1])
	}
}
