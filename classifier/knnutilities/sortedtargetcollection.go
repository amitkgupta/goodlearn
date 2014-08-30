package knnutilities

import (
	"math"

	"github.com/amitkgupta/goodlearn/dataset/target"
)

type sortedTargetCollection struct {
	k                int
	targetCollection []targetWithDistance
}

type targetWithDistance struct {
	target   target.Target
	distance float64
}

func NewSortedTargetCollection(k int) *sortedTargetCollection {
	return &sortedTargetCollection{k, make([]targetWithDistance, 0, k)}
}

func (stc *sortedTargetCollection) Insert(target target.Target, distance float64) {
	if distance >= stc.MaxDistance() {
		return
	}

	newTargetWithDistance := targetWithDistance{target, distance}
	if len(stc.targetCollection) == 0 {
		stc.targetCollection = []targetWithDistance{newTargetWithDistance}
		return
	}

	for i, twd := range stc.targetCollection {
		if distance < twd.distance {
			newCollection := append(stc.targetCollection[0:i], newTargetWithDistance)
			newCollection = append(newCollection, stc.targetCollection[i:]...)
			newCollection = newCollection[:int(math.Min(float64(stc.k), float64(len(newCollection))))]
			stc.targetCollection = newCollection
		}
	}
}

func (stc *sortedTargetCollection) MaxDistance() float64 {
	numTargets := len(stc.targetCollection)

	if numTargets == 0 {
		return math.MaxFloat64
	}

	return stc.targetCollection[numTargets-1].distance
}

// bad, return real error
func (stc *sortedTargetCollection) Vote() (target.Target, error) {
	return stc.targetCollection[0].target, nil
}
