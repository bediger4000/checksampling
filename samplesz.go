package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"time"

	rng "github.com/leesper/go_rng"
)

func main() {
	k := flag.Int("k", 1000, "number of samplings")
	n := flag.Int("n", 100000, "size of uniform distribution population")
	maxx := flag.Int("max", 1000, "max value in population")
	flag.Parse()

	population := make([]float64, *n, *n)
	urng := rng.NewUniformGenerator(time.Now().UnixNano() + int64(os.Getpid()))
	max := int32(*maxx)

	// make up the population
	for i := 0; i < *n; i++ {
		population[i] = float64(urng.Int32n(max))
	}

	sum, mean, stddev := StdDev(population)

	fmt.Fprintf(os.Stderr, "# %d values in population\n", *n)
	fmt.Fprintf(os.Stderr, "# %d max value in population\n", max)
	fmt.Fprintf(os.Stderr, "# Population sum %.01f\n", sum)
	fmt.Fprintf(os.Stderr, "# Population mean %.01f\n", mean)
	fmt.Fprintf(os.Stderr, "# Population std dev %f\n", stddev)

	sampleSizes := []int{15, 30, 50, 100, 500, 1000, 10000}

	fmt.Printf("sample size\tmin\tmed\tmax\tmean\tsample sdev\tsdev\n")
	for _, m := range sampleSizes {
		sample := make([]float64, m, m) // reuse for all samples of a single sample size
		sampleMeans := make([]float64, *k, *k)
		for i := 0; i < *k; i++ {
			for j := 0; j < m; j++ {
				idx := urng.Int32n(int32(*n))
				sample[j] = population[idx]
			}

			sampleMeans[i] = calculateMean(sample)
		}
		characteristics(m, sampleMeans, stddev)
	}
}

func characteristics(sampleSize int, data []float64, populationSdev float64) {
	sort.Float64s(data)

	median := findMedian(data)
	_, mean, stddev := StdDev(data)

	fmt.Printf("%d\t\t%.01f\t%.01f\t%.01f\t%.01f\t%.01f\t%.01f\n",
		sampleSize, data[0], median, data[len(data)-1], mean, stddev,
		populationSdev/math.Sqrt(float64(sampleSize)),
	)
}

func calculateMean(data []float64) float64 {
	var sum float64
	max := len(data)
	for i := 0; i < max; i++ {
		sum += data[i]
	}
	mean := sum / float64(max)
	return mean
}

func StdDev(data []float64) (float64, float64, float64) {
	var sum, sumOfSquares float64

	max := len(data)

	for i := 0; i < max; i++ {
		x := data[i]
		sum += x
		sumOfSquares += x * x
	}

	mean := sum / float64(max)

	return sum, mean, math.Sqrt(sumOfSquares/float64(max) - mean*mean)
}

func findMedian(data []float64) float64 {
	l := len(data)
	q := l / 2          // rounds down
	a := 1 - (l & 0x01) // 1 for odd l, 0 for even l

	low := data[q-a]
	high := data[q]
	mid := low + ((high - low) / 2.)

	return mid
}
