package parallelloop

import (
	"github.com/opctl/opctl/sdks/go/model"
)

// IsIterationComplete tests if an index is within range of the loop
func IsIterationComplete(
	index int,
	dcgParallelLoop model.DCGParallelLoopCall,
) bool {
	loopRange := dcgParallelLoop.Range
	if nil != loopRange {
		if nil != loopRange.Array {
			return index == len(*loopRange.Array)
		}
		if nil != loopRange.Object {
			return index == len(*loopRange.Object)
		}

		// empty array or object
		return true
	}

	return false
}
