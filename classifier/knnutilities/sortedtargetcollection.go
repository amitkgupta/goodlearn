package knnutilities

import (
	"errors"
	"fmt"
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

func NewSortedTargetCollection(k int) (*sortedTargetCollection, error) {
	if k < 1 {
		return nil, newInvalidCapError(k)
	}

	return &sortedTargetCollection{k, make([]targetWithDistance, 0, k)}, nil
}

func (stc *sortedTargetCollection) Insert(target target.Target, distance float64) error {
	if distance >= stc.MaxDistance() {
		return newDistanceTooLargeError(distance, stc.MaxDistance())
	}

	newTargetWithDistance := targetWithDistance{target, distance}
	if len(stc.targetCollection) == 0 {
		stc.targetCollection = []targetWithDistance{newTargetWithDistance}
		return nil
	}

	for i, twd := range stc.targetCollection {
		if distance < twd.distance {
			newCollection := append(stc.targetCollection[0:i], newTargetWithDistance)
			newCollection = append(newCollection, stc.targetCollection[i:]...)
			newCollection = newCollection[:int(math.Min(float64(stc.k), float64(len(newCollection))))]
			stc.targetCollection = newCollection
		}
	}

	return nil
}

func (stc *sortedTargetCollection) MaxDistance() float64 {
	if len(stc.targetCollection) < stc.k {
		return math.MaxFloat64
	}

	return stc.targetCollection[stc.k-1].distance
}

func (stc *sortedTargetCollection) Vote() (target.Target, error) {
	if len(stc.targetCollection) == 0 {
		return nil, newUnpopulatedVoteError()
	}

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

	return winner, nil
}

func newDistanceTooLargeError(distance, maxDistance float64) error {
	return errors.New(fmt.Sprintf("Cannot insert element with distance %.4f into collection with max distance %.4f", distance, maxDistance))
}
func newUnpopulatedVoteError() error {
	return errors.New("Cannot vote with unpopulated collection")
}

func newInvalidCapError(cap int) error {
	return errors.New(fmt.Sprintf("Cannot create collection with cap %d", cap))
}
