package knnutilities

import (
	"errors"
	"fmt"
	"math"

	"github.com/amitkgupta/goodlearn/data/target"
)

type SortedTargetCollection interface {
	Insert(target.Target, float64) error
	MaxDistance() float64
	Vote() (target.Target, error)
}

type kNNTargetCollection struct {
	k                int
	targetCollection []targetWithDistance
}

type targetWithDistance struct {
	target   target.Target
	distance float64
}

func NewKNNTargetCollection(k int) (*kNNTargetCollection, error) {
	if k < 1 {
		return nil, newInvalidCapError(k)
	}

	return &kNNTargetCollection{k, make([]targetWithDistance, 0, k)}, nil
}

func (stc *kNNTargetCollection) Insert(target target.Target, distance float64) error {
	if distance >= stc.MaxDistance() {
		return newDistanceTooLargeError(distance, stc.MaxDistance())
	}

	newTargetWithDistance := targetWithDistance{target, distance}

	for i, twd := range stc.targetCollection {
		if distance < twd.distance {
			newCollection := []targetWithDistance{}
			newCollection = append(newCollection, stc.targetCollection[0:i]...)
			newCollection = append(newCollection, newTargetWithDistance)
			newCollection = append(newCollection, stc.targetCollection[i:]...)

			stc.targetCollection = newCollection[:int(math.Min(float64(stc.k), float64(len(newCollection))))]
			return nil
		}
	}

	stc.targetCollection = append(stc.targetCollection, newTargetWithDistance)
	return nil
}

func (stc *kNNTargetCollection) MaxDistance() float64 {
	if len(stc.targetCollection) < stc.k {
		return math.MaxFloat64
	}

	return stc.targetCollection[stc.k-1].distance
}

func (stc *kNNTargetCollection) Vote() (target.Target, error) {
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
