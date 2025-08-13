package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"time"

	rng "github.com/leesper/go_rng"
)

func main() {
	k := flag.Int("k", 1000, "number of samplings")
	n := flag.Int("n", 100000, "size of uniform distribution population")
	maxx := flag.Int("max", 1000, "max value in population")
	m := flag.Int("m", 100, "number of samples per sampling")
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
	fmt.Fprintf(os.Stderr, "# Sample std dev %f\n", stddev/math.Sqrt(float64(*m)))

	// store for samples, re-used *k times
	sample := make([]float64, *m, *m)

	sampleMeans := make([]float64, *k, *k)

	// Loop over the number of samplings
	fmt.Printf("sum\tmean\tstddev\n")
	for i := 0; i < *k; i++ {
		// sample the population
		for j := 0; j < *m; j++ {
			idx := urng.Int32n(int32(*n))
			sample[j] = population[idx]
		}
		// find the sample's statistics
		sampleSum, sampleMean, sampleStdDev := StdDev(sample)
		sampleMeans[i] = sampleMean
		fmt.Printf("%.01f\t%.01f\t%.01f\n",
			sampleSum, sampleMean, sampleStdDev)
	}

	sumMeans := 0.0
	for i := 0; i < *k; i++ {
		sumMeans += sampleMeans[i]
	}
	meanSampleMeans := sumMeans / float64(*k)
	kurt := MomentCoefficientofKurtosis(sampleMeans, meanSampleMeans)
	fmt.Fprintf(os.Stderr, "# kurtosis of sample means: %.02f\n", kurt)
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

func MomentCoefficientofKurtosis(data []float64, mean float64) float64 {
	var sumOfDeviances2 float64
	var sumOfDeviances4 float64

	n := len(data)

	for i := 0; i < n; i++ {
		diff := data[i] - mean
		diff2 := diff * diff
		sumOfDeviances2 += diff2
		sumOfDeviances4 += diff2 * diff2
	}

	return float64(n) * sumOfDeviances4 / (sumOfDeviances2 * sumOfDeviances2)
}
