# Test Central Limit Theorem

From [Statistics Libre Text](https://stats.libretexts.org/Bookshelves/Introductory_Statistics/Introductory_Statistics_(Shafer_and_Zhang)/06%3A_Sampling_Distributions/6.02%3A_The_Sampling_Distribution_of_the_Sample_Mean):

> For samples of size 30 or more, the sample mean is approximately
> normally distributed, with mean 
> &mu;<sub>X</sub> = &mu;
> and standard deviation &sigma;<<sub>X</sub> = &sigma;/&#8730;n
> where n is the sample size.
> The larger the sample size, the better the approximation.

The above is a pretty typical assertion in statistics texts.
The derivation must be a doozy,
because there's rarely a discussion of the Central Limit Theory.

## Building and running

I wrote and tested this on Arch Linux, but Golang is pretty darn portable.
You should be able to try this on almost any distro with a [Go](https://go.dev/) compiler
and the [R statistical language](https://www.r-project.org/) installed.
You really only need R to get fancy histogram images.

```
$ make all
```

That will yield an executable, `samplings1`, an image in a file named `samples_dist.png`
a record of the samples of the population in `run.out`,
metadata about the population and samples in `run.metadata`.
The R "summary" of the means of the samples will be in `r.summary`.

#### Defaults

Executing `samplings1` will generate a population of values of 100,000,
whose values are between 0  and 1000.
It will take 1000 samplings, each with 100 values, from the population,
and calculate mean and standard deviation for each sampling.
It outputs the sum, the arithmetic mean and the standard deviation
for each sampling on stdout.


## Results

![Histogram, density and normal distribution](example_samples_dist.png)

That's a density histogram of a run of `samplings`.
There were 100,000 values in the generated population.

The blue line is the normal curve for the mean and standard deviation of
the population.
The red line is R's `density` function for the means of the samplings.
I believe they're comparable as theoretic and experimental distributions.

The file `run.metadata` for that population gives it a mean of 501.4,
and a standard deviation of 288.19.
The theoretical sample standard deviation is 48.7
The Kurtosis of the  sample means is 2.96,
extremely close to the kurtosis of a normal distribution.
The file `r.summary` gives the sample means a mean of 500.1,
very close to the population mean.
Running the second column of the `run.out` file through GNU `datamash`:

```
$ awk 'NR>1{print $2}' run.out | datamash mean 1 sstdev 1
500.2961        50.275935241967
```

The samplings means have a mean of 500.3 and a standard deviation of 50.3.
That's really close to the theoretical values.

# Experiment with sample size

```
$ make samplesz
$ ./samplesz
# 100000 values in population
# 1000 max value in population
# Population sum 49994428.0
# Population mean 499.9
# Population std dev 288.624605
sample size     min     med     max     mean    sample sdev     sdev
15              265.6   500.3   751.0   501.4   75.7            74.5
30              331.0   499.1   646.6   499.0   52.8            52.7
50              363.7   500.5   600.8   498.8   40.2            40.8
100             404.5   498.8   588.5   498.1   28.3            28.9
500             466.1   500.8   540.7   500.0   12.6            12.9
1000            474.5   500.6   528.0   500.2   9.0             9.1
10000           489.1   500.2   509.6   500.0   2.9             2.9
```

Take 1000 samples of a given size, calculate some indicators about the
1000 samples.
The number of values in each sampe (15, 30, 50, 100...) is hardcoded.
I wanted to see if the assertion that a "large size sample" is more
than 30 actually worked any different.
