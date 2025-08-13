all: samples_dist.png

samplings1: samplings1.go
	go build samplings1.go

run.out: samplings1
	./samplings1 -m 35 -k 1000 > run.out 2> run.metadata

samples_dist.png: density_plot.r run.out
	R --no-save < density_plot.r
