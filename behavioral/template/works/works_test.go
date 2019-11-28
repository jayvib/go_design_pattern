package works_test

import (
	"github.com/jayvib/go_design_pattern/behavioral/template/works"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWork(t *testing.T) {
	maidTemplate := works.NewMaidTemplate(&works.MaiderImpl{
		"Jayson",
		"Car Wash",
	})

	result := maidTemplate.Do()
	assert.Equal(t, "Jayson is doing Car Wash", result, "expected result doesn't match with the actual result")
}
