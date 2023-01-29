# masyu
Masyu Solver - golang

This is the self-proclaimed World's fastest solver for [the masyu puzzle](www.puzzle-masyu.com). It is similar to [my Shingoki Solver](https://github.com/joshprzybyszewski/shingokisolver).

To run, execute `make compete`.

## Results

Check the [Hall of Fame](www.puzzle-masyu.com/hall.php?hallsize=18) for the results recorded by the puzzle server (which include network duration of submission). Below are the results of the solver as recorded on my machine.

_NOTE: Update this table with `make results`._

<resultsMarker>

_GOOS: linux_

_GOARCH: amd64_

_cpu: Intel(R) Core(TM) i5-3570 CPU @ 3.40GHz_

_Solve timeout: 10m0s_

|Puzzle|Min|Median|p75|p95|max|sample size|
|-|-|-|-|-|-|-:|
|6x6 easy|217.489µs|455.798µs|488.951µs|574.881µs|1.571799ms|365|
|8x8 easy|265.099µs|582.14µs|635.166µs|1.555636ms|2.076139ms|331|
|8x8 medium|251.254µs|639.387µs|1.491168ms|1.728563ms|2.560641ms|309|
|8x8 hard|530.104µs|1.574084ms|1.713519ms|2.05674ms|2.717803ms|292|
|10x10 easy|374.214µs|804.183µs|1.500118ms|1.949799ms|2.882077ms|284|
|10x10 medium|348.863µs|1.645162ms|1.836567ms|2.297833ms|6.441404ms|266|
|10x10 hard|955.698µs|1.964213ms|2.395234ms|4.024258ms|9.653363ms|247|
|15x15 easy|684.76µs|2.335323ms|2.60949ms|2.994633ms|3.580615ms|227|
|15x15 medium|1.224524ms|2.952847ms|3.932663ms|8.656005ms|133.277348ms|177|
|15x15 hard|2.175221ms|8.279355ms|15.712966ms|62.26496ms|619.205159ms|133|
|20x20 easy|1.223357ms|3.790601ms|4.237061ms|5.116185ms|9.025728ms|169|
|20x20 medium|1.868819ms|9.909126ms|29.813732ms|270.064234ms|16.762765374s|138|
|20x20 hard|6.337712ms|393.318792ms|8.002627801s|3m38.960553367s|10m0.011565175s|115|
|25x25 easy|2.861715ms|5.498312ms|6.108223ms|9.781951ms|34.907577ms|118|
|25x25 medium|4.561067ms|49.716478ms|496.324749ms|17.484690696s|3m8.637007753s|97|
|25x25 hard|19.180944ms|2m41.319942548s|10m0.002971999s|10m0.011584809s|10m0.015606939s|81|

_Last Updated: 29 Jan 23 12:31 CST_
</resultsMarker>
