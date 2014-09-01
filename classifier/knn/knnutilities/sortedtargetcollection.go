package knnutilities

import (
	"math"

	"github.com/amitkgupta/goodlearn/data/target"
)

type SortedTargetCollection interface {
	Insert(target.Target, float64)
	MaxDistance() float64
	Vote() target.Target
}

type kNNTargetCollection struct {
	k                int
	targetCollection []targetWithDistance
}

type targetWithDistance struct {
	target   target.Target
	distance float64
}

func NewKNNTargetCollection(k int) *kNNTargetCollection {
	return &kNNTargetCollection{k, make([]targetWithDistance, 0, k)}
}

func (stc *kNNTargetCollection) Insert(target target.Target, distance float64) {
	newTargetWithDistance := targetWithDistance{target, distance}

	for i, twd := range stc.targetCollection {
		if distance < twd.distance {
			newCollection := []targetWithDistance{}
			newCollection = append(newCollection, stc.targetCollection[0:i]...)
			newCollection = append(newCollection, newTargetWithDistance)
			newCollection = append(newCollection, stc.targetCollection[i:]...)

			stc.targetCollection = newCollection[:int(math.Min(float64(stc.k), float64(len(newCollection))))]
			return
		}
	}

	stc.targetCollection = append(stc.targetCollection, newTargetWithDistance)
}

func (stc *kNNTargetCollection) MaxDistance() float64 {
	if len(stc.targetCollection) < stc.k {
		return math.MaxFloat64
	}

	return stc.targetCollection[stc.k-1].distance
}

func (stc *kNNTargetCollection) Vote() target.Target {
	winner := stc.targetCollection[0].target
	votesForWinner := 0

	for i, candidate := range stc.targetCollection {
		votesForCurrent := 0

		for _, other := range stc.targetCollection[i:] {
			if candidate.target.Equals(other.target) {
				votesForCurrent++
			}
		}

		if votesForCurrent > votesForWinner {
			winner = candidate.target
			votesForWinner = votesForCurrent
		}
	}

	return winner
}
