package flow_test

import (
	"testing"

	"github.ncsu.edu/jjuecks/vv8-post-processor/flow"
)

func testFlow(T *testing.T) {
	T.Log("Test flow")
	flowAgg, err := flow.NewAggregator()
	if err != nil {
		T.Fatal(err)
	}

}
