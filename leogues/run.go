package desafiogo

import (
	"sort"
	"sync"
)

func calculateMode(data []float64) float64 {
	dataCpy := append(make([]float64, 0, len(data)), data...)

	if !sort.IsSorted(sort.Float64Slice(dataCpy)) {
		sort.Float64s(dataCpy)
	}

	var mode float64
	var maxFrequency, frequencyCount int

	for i := 0; i < len(dataCpy)-1; i++ {
		if dataCpy[i] == dataCpy[i+1] {
			frequencyCount++
			if frequencyCount > maxFrequency {
				maxFrequency = frequencyCount
				mode = dataCpy[i]
			}
		} else {
			frequencyCount = 1
		}
	}

	return mode
}

func hasNumber(data []float64, number float64) bool {
	for _, d := range data {
		if d == number {
			return true
		}
	}
	return false

}

func Run(data []float64) *Result {
	sort.Float64s(data)

	var wg sync.WaitGroup
	wg.Add(6)

	var (
		ave         float64
		med         float64
		p90         float64
		p99         float64
		mod         float64
		hasNumber47 bool
	)

	go func() {
		ave = calculateAverage(data)
		wg.Done()
	}()

	go func() {
		med = calculateMedian(data)
		wg.Done()
	}()

	go func() {
		p90 = calculatePercentile(data, 90)
		wg.Done()
	}()

	go func() {
		p99 = calculatePercentile(data, 99)
		wg.Done()
	}()

	go func() {
		mod = calculateMode(data)
		wg.Done()
	}()

	go func() {
		hasNumber47 = hasNumber(data, 47)
		wg.Done()
	}()

	wg.Wait()

	return &Result{
		Average:      ave,
		Median:       med,
		Percentile90: p90,
		Percentile99: p99,
		Mode:         mod,
		HasNumber47:  hasNumber47,
	}

}
