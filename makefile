all: samples_dist.png

samplings1: samplings1.go
	go build samplings1.go

run.out run.metadata: samplings1
	./samplings1 -m 35 -k 1000 > run.out 2> run.metadata

samples_dist.png r.summary: density_plot.r run.out
	R --no-save < density_plot.r

clean:
	rm -rf samplings1
	rm -rf run.out run.metadata
	rm -rf r.summary
	rm -rf samples_dist.png
