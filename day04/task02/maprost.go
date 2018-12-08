package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/maprost/adventofcode2018/golib"
)

type Action struct {
	time   time.Time
	action string
}

type DurationInfo struct {
	duration    time.Duration
	markedTimes []int
}

func main() {
	actions := golib.Read("day04/task02/input_102095.txt")
	actionList := make([]Action, 0, len(actions))

	// prepare
	for _, action := range actions {
		actionArr := strings.Split(action, "] ")

		t, err := time.Parse("[2006-01-02 15:04", actionArr[0])
		if err != nil {
			panic(err)
		}

		actionList = append(actionList, Action{
			time:   t,
			action: actionArr[1],
		})
	}

	// sort
	sort.Slice(actionList, func(i, j int) bool {
		return actionList[i].time.Before(actionList[j].time)
	})

	// calculate
	var guard int
	var startTime time.Time
	guardMap := make(map[int]DurationInfo)

	for _, action := range actionList {
		fmt.Println(action.time, " - ", action.action)

		if strings.HasPrefix(action.action, "Guard") {
			idStr := strings.TrimPrefix(strings.TrimSuffix(action.action, " begins shift"), "Guard #")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				panic(err)
			}

			guard = id

		} else if action.action == "falls asleep" {
			startTime = action.time

		} else if action.action == "wakes up" {
			duration := action.time.Sub(startTime)

			dInfo, ok := guardMap[guard]
			if !ok {
				dInfo = DurationInfo{
					duration:    duration,
					markedTimes: make([]int, 60),
				}
			} else {
				dInfo.duration += duration
			}

			// mark times
			for i := startTime.Minute(); i < action.time.Minute(); i++ {
				dInfo.markedTimes[i] = dInfo.markedTimes[i] + 1
			}

			guardMap[guard] = dInfo

		} else {
			panic("No action found")
		}
	}

	// find guard with most sleep at a time
	maxTimes := 0
	minute := 0
	bestSleeperGuard := 0
	for g, dInfo := range guardMap {
		for m, t := range dInfo.markedTimes {
			if t > maxTimes {
				bestSleeperGuard = g
				minute = m
				maxTimes = t
			}
		}
	}

	fmt.Printf("Best Sleeper: %d, %d min (%dx) (%d)\n",
		bestSleeperGuard, minute, maxTimes, bestSleeperGuard*minute)
}
