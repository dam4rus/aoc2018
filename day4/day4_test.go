package main

import (
	day4 "rkalmar/aoc2020/day4/internal/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

var INPUT = []string{
	"[1518-11-01 00:00] Guard #10 begins shift",
	"[1518-11-01 00:05] falls asleep",
	"[1518-11-01 00:25] wakes up",
	"[1518-11-01 00:30] falls asleep",
	"[1518-11-01 00:55] wakes up",
	"[1518-11-01 23:58] Guard #99 begins shift",
	"[1518-11-02 00:40] falls asleep",
	"[1518-11-02 00:50] wakes up",
	"[1518-11-03 00:05] Guard #10 begins shift",
	"[1518-11-03 00:24] falls asleep",
	"[1518-11-03 00:29] wakes up",
	"[1518-11-04 00:02] Guard #99 begins shift",
	"[1518-11-04 00:36] falls asleep",
	"[1518-11-04 00:46] wakes up",
	"[1518-11-05 00:03] Guard #99 begins shift",
	"[1518-11-05 00:45] falls asleep",
	"[1518-11-05 00:55] wakes up",
}

func TestPart1(t *testing.T) {
	entries, err := day4.ParseInput(INPUT)
	if err != nil {
		t.Fatal("Failed to parse input", err)
	}

	guardSleepPeriods, err := day4.ProcessEntries(entries)
	if err != nil {
		t.Fatal("Failed to process entries", err)
	}
	sleepiestGuardId := day4.FindSleepiestGuardId(guardSleepPeriods)
	assert.Equal(t, 10, sleepiestGuardId)

	mostAsleep := day4.FindMostAsleepMinute(guardSleepPeriods[sleepiestGuardId])
	assert.Equal(t, 24, mostAsleep.Minute)
	assert.Equal(t, 10*24, sleepiestGuardId*mostAsleep.Minute)
}

func TestPart2(t *testing.T) {
	entries, err := day4.ParseInput(INPUT)
	if err != nil {
		t.Fatal("Failed to parse input", err)
	}

	guardSleepPeriods, err := day4.ProcessEntries(entries)
	if err != nil {
		t.Fatal("Failed to process entries", err)
	}
	guardId, mostAsleep := day4.FindGuardMostAsleepAtCertainMinute(guardSleepPeriods)
	assert.Equal(t, 99*45, guardId*mostAsleep.Minute)
}
