package day4

import (
	"errors"
	"regexp"
	"sort"
	"strconv"
)

type Entry struct {
	Month  int
	Day    int
	Hour   int
	Minute int
	Log    string
}

type EntrySlice []*Entry

func (entries EntrySlice) Len() int {
	return len(entries)
}

func (entries EntrySlice) Less(i, j int) bool {
	if entries[i].Month < entries[j].Month {
		return true
	}
	if entries[i].Month > entries[j].Month {
		return false
	}
	if entries[i].Day < entries[j].Day {
		return true
	}
	if entries[i].Day > entries[j].Day {
		return false
	}
	if entries[i].Hour < entries[j].Hour {
		return true
	}
	if entries[i].Hour > entries[j].Hour {
		return false
	}
	return entries[i].Minute < entries[j].Minute
}

func (entries EntrySlice) Swap(i, j int) {
	tmp := entries[i]
	entries[i] = entries[j]
	entries[j] = tmp
}

func (entries EntrySlice) Sort() {
	sort.Sort(entries)
}

type SleepPeriod struct {
	StartMinute int
	EndMinute   int
}

func (period SleepPeriod) IsAsleepAtMinute(minute int) bool {
	return minute >= period.StartMinute && minute < period.EndMinute
}

func (period SleepPeriod) SleepInterval() int {
	return period.EndMinute - period.StartMinute
}

type MostAsleep struct {
	Minute int
	Count  int
}

func ParseInput(input []string) (entries []*Entry, err error) {
	regex, err := regexp.Compile(`\[[0-9]{4}-([0-9]{2})-([0-9]{2}) ([0-9]{2}):([0-5][0-9])\] (.*)`)
	if err != nil {
		return entries, err
	}

	for _, line := range input {
		entry, err := parseEntry(line, regex)
		if err != nil {
			return entries, err
		}
		entries = append(entries, entry)
	}
	EntrySlice(entries).Sort()
	return
}

func ProcessEntries(entries []*Entry) (map[int][]SleepPeriod, error) {
	shiftBeginsRegex, err := regexp.Compile(`Guard #([0-9]+) begins shift`)
	if err != nil {
		return nil, err
	}
	guardSleepPeriods := make(map[int][]SleepPeriod)
	var currentGuardId *int
	var currentStartMinute *int
	for _, entry := range entries {
		guardShiftSubmatches := shiftBeginsRegex.FindStringSubmatch(entry.Log)
		if guardShiftSubmatches != nil {
			guardId, err := strconv.Atoi(guardShiftSubmatches[1])
			if err != nil {
				return nil, err
			}
			currentGuardId = &guardId
			continue
		}
		if entry.Log == "falls asleep" {
			currentStartMinute = &entry.Minute
			continue
		}
		if entry.Log == "wakes up" {
			guardSleepPeriods[*currentGuardId] = append(guardSleepPeriods[*currentGuardId], SleepPeriod{
				StartMinute: *currentStartMinute,
				EndMinute:   entry.Minute,
			})
			continue
		}
		return nil, errors.New("invalid entry log")
	}
	return guardSleepPeriods, nil
}

func FindSleepiestGuardId(guardSleepPeriods map[int][]SleepPeriod) (sleepiestGuardId int) {
	guardTotalSleepTime := make(map[int]int)
	for guardId, sleepPeriods := range guardSleepPeriods {
		for _, sleepPeriod := range sleepPeriods {
			guardTotalSleepTime[guardId] += sleepPeriod.SleepInterval()
		}
	}
	var sleepiestGuardTotalSleepTime int
	for guardId, totalSleepTime := range guardTotalSleepTime {
		if totalSleepTime > sleepiestGuardTotalSleepTime {
			sleepiestGuardId = guardId
			sleepiestGuardTotalSleepTime = totalSleepTime
		}
	}
	return
}

func FindMostAsleepMinute(sleepPeriods []SleepPeriod) (mostAsleep MostAsleep) {
	for minute := 0; minute < 60; minute++ {
		var totalAsleepCount int
		for _, schedule := range sleepPeriods {
			if schedule.IsAsleepAtMinute(minute) {
				totalAsleepCount += 1
			}
		}
		if totalAsleepCount > mostAsleep.Count {
			mostAsleep.Minute = minute
			mostAsleep.Count = totalAsleepCount
		}
	}
	return
}

func FindGuardMostAsleepAtCertainMinute(guardSleepPeriods map[int][]SleepPeriod) (guardId int, mostAsleep MostAsleep) {
	for currentGuardId, sleepPeriods := range guardSleepPeriods {
		mostAsleepForCurrentGuard := FindMostAsleepMinute(sleepPeriods)
		if mostAsleepForCurrentGuard.Count > mostAsleep.Count {
			guardId = currentGuardId
			mostAsleep = mostAsleepForCurrentGuard
		}
	}
	return
}

func parseEntry(line string, entryRegex *regexp.Regexp) (*Entry, error) {
	submatches := entryRegex.FindStringSubmatch(line)
	month, err := strconv.Atoi(submatches[1])
	if err != nil {
		return nil, err
	}
	day, err := strconv.Atoi(submatches[2])
	if err != nil {
		return nil, err
	}
	hour, err := strconv.Atoi(submatches[3])
	if err != nil {
		return nil, err
	}
	minute, err := strconv.Atoi(submatches[4])
	if err != nil {
		return nil, err
	}
	return &Entry{
		Month:  month,
		Day:    day,
		Hour:   hour,
		Minute: minute,
		Log:    submatches[5],
	}, nil
}
