package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// DummyBreakData
type DummyBreakData struct {
	startTime   time.Time
	triggerTime time.Time
	name        string
}

// generateMockData creates a list of DummyBreakData within a 1 min range, spaced 10 secs apart
func generateMockData(adBreakStart time.Time) []DummyBreakData {
	var data []DummyBreakData

	// lets space the start times by 10 seconds within a 1-minute range.
	// => start times will be at :00, :10, :20, :30, :40, :50 seconds
	for i := 0; i < 6; i++ {
		// add i*10 seconds to adBreakStart to create a 10-second gap between events
		startTime := adBreakStart.Add(-time.Minute).Add(time.Duration(i*10) * time.Second)
		// set the triggerTime to 4 minutes before the startTime
		triggerTime := startTime.Add(-4 * time.Minute)

		// add the new event to the slice
		data = append(data, DummyBreakData{
			startTime:   startTime,
			triggerTime: triggerTime,
			name:        fmt.Sprintf("event%d", i+1),
		})
	}

	// Shuffle the data slice randomly
	rand.Shuffle(len(data), func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})

	return data

}

// sortByTriggerTime orders the events by trigger time (soonest first)
func sortByTriggerTime(data []DummyBreakData) []DummyBreakData {
	sort.Slice(data, func(i, j int) bool {
		return data[i].triggerTime.Before(data[j].triggerTime)
	})
	return data
}

// todo^ look at the newer [slices.SortFunc]

// simulateEndpointCall simulates an HTTP endpoint call by printing to the console
// and waiting for a random response time between 0 and 2 seconds.
func simulateEndpointCall(event DummyBreakData) {
	// Record the time the endpoint is "called" (when it's fired)
	firedAt := time.Now()

	// Print the start of the endpoint call for the event
	fmt.Printf("Event: %s, Start Time: %s, Trigger Time: %s, Fired At: %s\n",
		event.name,
		event.startTime.Format(time.RFC3339),
		event.triggerTime.Format(time.RFC3339),
		firedAt.Format(time.RFC3339),
	)

	// Simulate random response time between 0 and 2 seconds
	randomDelay := time.Duration(rand.Intn(2000)) * time.Millisecond
	time.Sleep(randomDelay)

	// Print that the response is a 200 success
	fmt.Printf("Endpoint responded with 200 OK after %v\n", event.name, randomDelay.Seconds())
}

// scheduleEvent waits for the event's trigger time to fire
func scheduleEvent(data DummyBreakData, now time.Time, done chan<- bool) { // done channel
	// calculate how long we need to wait
	timeUntilTrigger := data.triggerTime.Sub(now)
	if timeUntilTrigger > 0 {
		time.Sleep(timeUntilTrigger) // wait until it's time to trigger the event
	}

	// simulate calling the endpoint for the break
	simulateEndpointCall(data)

	done <- true // signal that this goroutine is done
}

// processEvents handles each event concurrently, scheduling them by trigger time
func processEvents(data []DummyBreakData) {
	now := time.Now()
	done := make(chan bool, len(data)) // create a channel to track when all events are fired
	// each go routine will signal it is finished by sending true to the channel

	// for each event, start a goroutine
	for _, event := range data {
		go scheduleEvent(event, now, done)
	}

	// wait for all goroutines to complete
	for i := 0; i < len(data); i++ {
		//result := <-done // result unused
		<-done // done is our channel, <- is the receive operator. the value is received but discarded, we just want it to block execution
		// todo: look into non blocking alternative
	}
}

func main() {
	fmt.Println("begin")
	// advert break start time: now + 5 minutes
	adBreakStart := time.Now().Add(5 * time.Minute)

	// generate mock data
	mockData := generateMockData(adBreakStart)
	fmt.Println("Event List:")
	for _, event := range mockData {
		fmt.Printf("Name: %s, Start Time: %s\n", event.name, event.startTime.Format(time.RFC3339))
	}

	// sort mock data
	sortedData := sortByTriggerTime(mockData)

	// Print the sorted data
	fmt.Println("Sorted by trigger time:")
	for _, event := range sortedData {
		fmt.Printf("Event: %s, Start Time: %s, Trigger Time: %s\n",
			event.name,
			event.startTime.Format(time.RFC3339),
			event.triggerTime.Format(time.RFC3339),
		)
	}

	// Process events concurrently
	fmt.Println("Processing events...")
	processEvents(sortedData)
}
