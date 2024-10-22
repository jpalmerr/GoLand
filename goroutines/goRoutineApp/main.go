package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type DummyBreakData struct {
	startTime   time.Time
	triggerTime time.Time
	name        string
}

type EndpointResponse struct {
	event    DummyBreakData
	status   int
	error    error
	duration time.Duration
	firedAt  time.Time
}

// simulate polling data with compressed timeline
func generateMockData(simulationStart time.Time) []DummyBreakData {
	var data []DummyBreakData

	fmt.Println("\nSimulating poll for upcoming events...")

	// in real system: events would start 4-5 mins from now
	// demo purposes: events starting 40-55 seconds from now
	for i := 0; i < 6; i++ {
		// random start time between 40-55 seconds from now
		randomFutureOffset := time.Duration(40+rand.Intn(15)) * time.Second
		startTime := simulationStart.Add(randomFutureOffset)

		// in real system: trigger = startTime - 4min
		// demo purposes: trigger = startTime - 30sec
		triggerTime := startTime.Add(-30 * time.Second)

		event := DummyBreakData{
			startTime:   startTime,
			triggerTime: triggerTime,
			name:        fmt.Sprintf("event%d", i+1),
		}
		data = append(data, event)
	}

	// generated events, now select two random events to match times
	idx1 := rand.Intn(len(data))
	idx2 := idx1
	for idx2 == idx1 {
		idx2 = rand.Intn(len(data))
	}

	// Make the second event match the first event's times
	data[idx2].startTime = data[idx1].startTime
	data[idx2].triggerTime = data[idx1].triggerTime

	for _, event := range data {
		fmt.Printf("Found event: %s\n", formatEventTimes(event))
	}

	return data
}

// sortByTriggerTime orders the events by trigger time (soonest first)
func sortByTriggerTime(data []DummyBreakData) []DummyBreakData {
	sort.Slice(data, func(i, j int) bool {
		return data[i].triggerTime.Before(data[j].triggerTime)
	})
	return data
}

/**
weird go formating times thing

Month = 1 (January)
Day = 2
Hour = 15 (3 PM)
Minute = 4
Second = 5
Year = 2006
Timezone = 7 (MST)
*/

// helper to format event times
func formatEventTimes(event DummyBreakData) string {
	return fmt.Sprintf("Name: %s, Start: %s (in %s), Trigger: %s (in %s)",
		event.name,
		event.startTime.Format("15:04:05"),
		time.Until(event.startTime).Round(time.Second),
		event.triggerTime.Format("15:04:05"),
		time.Until(event.triggerTime).Round(time.Second),
	)
}

// mocking our fire to yospace
// we want to build in both some latency and some error responses to demo non blocking async
func simulateEndpointCall(event DummyBreakData) EndpointResponse {
	firedAt := time.Now()

	// random delay between 2 and 4 seconds (a bit much but i want to guarantee some overlap for demo purposes)
	randomDelay := time.Duration(2000+rand.Intn(4000)) * time.Millisecond

	// 40% chance of error (way too high but exaggerated for demo purposes)
	var status int
	var err error
	if rand.Float32() < 0.4 {
		errorCodes := []int{400, 429, 500, 503}
		status = errorCodes[rand.Intn(len(errorCodes))]
		err = errors.New(fmt.Sprintf("HTTP %d error", status))
	} else {
		status = 200
	}

	time.Sleep(randomDelay)

	return EndpointResponse{
		event:    event,
		status:   status,
		error:    err,
		duration: randomDelay,
		firedAt:  firedAt,
	}
}

func scheduleEvent(data DummyBreakData, now time.Time, responses chan<- EndpointResponse) {
	timeUntilTrigger := data.triggerTime.Sub(now)

	if timeUntilTrigger > 0 {
		fmt.Printf("Waiting %s to trigger %s\n",
			timeUntilTrigger.Round(time.Second),
			data.name,
		)
		time.Sleep(timeUntilTrigger)
	}

	fmt.Printf("⚡ Triggering %s endpoint call at %s\n",
		data.name,
		time.Now().Format("15:04:05"),
	)

	response := simulateEndpointCall(data)
	responses <- response
}

func processEvents(data []DummyBreakData) {
	now := time.Now()
	responses := make(chan EndpointResponse, len(data))

	fmt.Println("\nScheduling events...")
	// launch all goroutines
	for _, event := range data {
		fmt.Printf("Scheduled: %s\n", formatEventTimes(event))
		go scheduleEvent(event, now, responses)
	}

	fmt.Println("\nAwaiting responses...")
	// process responses as they arrive
	for i := 0; i < len(data); i++ {
		response := <-responses
		if response.error != nil {
			fmt.Printf("❌ ERROR - %s\n  Status: %d, Error: %s\n  Duration: %.2fs\n  Fired At: %s\n",
				formatEventTimes(response.event),
				response.status,
				response.error,
				response.duration.Seconds(),
				response.firedAt.Format("15:04:05"),
			)
		} else {
			fmt.Printf("✅ SUCCESS - %s\n  Status: %d\n  Duration: %.2fs\n  Fired At: %s\n",
				formatEventTimes(response.event),
				response.status,
				response.duration.Seconds(),
				response.firedAt.Format("15:04:05"),
			)
		}
	}
}

func main() {
	simulationStart := time.Now()
	fmt.Printf("Simulation starting at: %s\n", simulationStart.Format("15:04:05"))

	// generate mock data (demo equivalent to polling)
	mockData := generateMockData(simulationStart)

	// sort by trigger time (priority queue)
	sortedData := sortByTriggerTime(mockData)

	fmt.Println("\nPriority queue (sorted by trigger time):")
	for _, event := range sortedData {
		fmt.Printf("Priority event: %s\n", formatEventTimes(event))
	}

	// process events
	processEvents(sortedData)
}
