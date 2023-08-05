package day7

import (
	"regexp"
	"sort"

	"golang.org/x/exp/slices"
)

type Step struct {
	RequiredBy []rune
	Requires   []rune
}

type Worker struct {
	WorkingOn rune
	DoneIn    int
}

type SortByDoneIn []*Worker

func (w SortByDoneIn) Len() int {
	return len(w)
}

func (w SortByDoneIn) Less(i, j int) bool {
	return w[i].DoneIn < w[j].DoneIn
}

func (w SortByDoneIn) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

type StepIterator struct {
	steps     map[rune]*Step
	stepsDone []rune
	stepCount int
}

func NewStepIterator(steps map[rune]*Step) StepIterator {
	return StepIterator{
		steps:     steps,
		stepCount: len(steps),
	}
}

func (s *StepIterator) iterateSteps(callable func([]rune) rune) {
	for len(s.stepsDone) < s.stepCount {
		availableSteps := s.collectAvailableSteps()
		slices.Sort(availableSteps)
		stepDone := callable(availableSteps)
		s.stepsDone = append(s.stepsDone, stepDone)
	}
}

func (s *StepIterator) collectAvailableSteps() (availableSteps []rune) {
	for stepIndex, step := range s.steps {
		if slices.Contains(s.stepsDone, stepIndex) {
			continue
		}
		if len(step.Requires) == 0 || s.allRequiredStepDone(step) {
			availableSteps = append(availableSteps, stepIndex)
		}
	}
	return
}

func (s *StepIterator) allRequiredStepDone(step *Step) bool {
	for _, requiredStep := range step.Requires {
		if !slices.Contains(s.stepsDone, requiredStep) {
			return false
		}
	}
	return true
}

func ParseInput(lines []string) map[rune]*Step {
	regex := regexp.MustCompile("Step ([A-Z]) must be finished before step ([A-Z]) can begin")
	steps := make(map[rune]*Step)
	for _, line := range lines {
		submatches := regex.FindStringSubmatch(line)
		if submatches == nil {
			continue
		}
		prerequisiteStepIndex := rune(submatches[1][0])
		stepIndex := rune(submatches[2][0])
		prerequisiteStep, found := steps[prerequisiteStepIndex]
		if !found {
			prerequisiteStep = new(Step)
			steps[prerequisiteStepIndex] = prerequisiteStep
		}
		prerequisiteStep.RequiredBy = append(steps[prerequisiteStepIndex].RequiredBy, stepIndex)

		step, found := steps[stepIndex]
		if !found {
			step = new(Step)
			steps[stepIndex] = step
		}
		step.Requires = append(steps[stepIndex].Requires, prerequisiteStepIndex)
	}
	return steps
}

func DetermineStepOrder(steps map[rune]*Step) string {
	iterator := NewStepIterator(steps)
	iterator.iterateSteps(func(availableSteps []rune) rune { return availableSteps[0] })
	return string(iterator.stepsDone)
}

func CalculateTotalTime(steps map[rune]*Step, maxWorkerCount int, extraSecondsToComplete int) int {
	var totalTime int
	var workers []*Worker
	iterator := NewStepIterator(steps)
	iterator.iterateSteps(func(availableSteps []rune) rune {
		for _, availableStep := range availableSteps {
			workers = append(workers, &Worker{
				WorkingOn: availableStep,
				DoneIn:    int(availableStep) - int('A') + 1 + extraSecondsToComplete,
			})
			delete(steps, availableStep)
			if len(workers) == maxWorkerCount {
				break
			}
		}
		sort.Sort(SortByDoneIn(workers))
		firstWorkerToBeDone := workers[0]
		totalTime += firstWorkerToBeDone.DoneIn
		workers = slices.Delete(workers, 0, 1)
		for _, worker := range workers {
			worker.DoneIn -= firstWorkerToBeDone.DoneIn
		}
		return firstWorkerToBeDone.WorkingOn
	})

	return totalTime
}
