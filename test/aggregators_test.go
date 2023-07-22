package test

import (
	"testing"

	"github.com/seedy1/UpfTest/helpers"
)

func TestMinMax(t *testing.T) {
	arr := [...]int64{1, 2, 3, 4, 5}
	min, max := helpers.GetMinMaxTimeStamp(arr[:])
	if min != 1 && max != 5 {
		t.Error("Error getting min and max values")
	}
	// t.Error("Just a test")

}
