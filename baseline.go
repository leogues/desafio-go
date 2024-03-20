package desafiogo

/*
	AS DESCRIBED IN THE RULES: DO NOT EDIT THIS FILE!!!
*/

func BaselineRun(data []float64) *Result {
	ave := calculateAverage(data)
	med := calculateMedian(data)
	p90 := calculatePercentile(data, 90)
	p99 := calculatePercentile(data, 99)

	mod, count := float64(0), 1
	for _, d := range data {
		c := 0
		for j := 0; j < len(data); j++ {
			if d == data[j] {
				c++
			}
		}
		if c > count {
			mod, count = d, c
		}
	}

	hasNumber47 := false
	for _, d := range data {
		if d == 47 {
			hasNumber47 = true
		}
	}

	return &Result{
		Average:      ave,
		Median:       med,
		Percentile90: p90,
		Percentile99: p99,
		Mode:         mod,
		HasNumber47:  hasNumber47,
	}
}
