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

|Puzzle|Min|p25|Median|p75|p95|max|sample size|
|-|-|-|-|-|-|-|-:|
|6x6 easy|176.643µs|486.669µs|450.409µs|486.669µs|588.989µs|1.358759ms|369|
|8x8 easy|258.459µs|631.157µs|584.068µs|631.157µs|1.551635ms|1.819653ms|335|
|8x8 medium|283.853µs|1.457753ms|594.06µs|1.457753ms|1.718335ms|2.373745ms|312|
|8x8 hard|478.662µs|1.621606ms|1.486937ms|1.621606ms|1.87342ms|4.111766ms|295|
|10x10 easy|307.996µs|999.317µs|767.347µs|999.317µs|1.855772ms|2.593089ms|287|
|10x10 medium|467.243µs|1.789111ms|1.568844ms|1.789111ms|2.31779ms|5.600314ms|269|
|10x10 hard|981.792µs|2.188232ms|1.926653ms|2.188232ms|2.979408ms|6.511793ms|250|
|15x15 easy|704.737µs|2.413922ms|1.923119ms|2.413922ms|2.855099ms|3.739854ms|230|
|15x15 medium|1.211633ms|3.844764ms|2.941965ms|3.844764ms|7.013648ms|137.764883ms|180|
|15x15 hard|2.70504ms|15.888966ms|8.249435ms|15.888966ms|74.041383ms|774.376549ms|136|
|20x20 easy|1.447868ms|4.05855ms|3.646946ms|4.05855ms|4.78296ms|5.980194ms|172|
|20x20 medium|2.252575ms|19.070908ms|9.209974ms|19.070908ms|250.775865ms|48.397190858s|141|
|20x20 hard|5.891944ms|2.773560162s|288.263742ms|2.773560162s|2m0.002272441s|2m0.016217852s|118|
|25x25 easy|2.804969ms|6.096055ms|5.252833ms|6.096055ms|8.748763ms|29.7489ms|121|
|25x25 medium|4.711035ms|342.221808ms|57.922618ms|342.221808ms|11.240667275s|2m0.012901129s|100|
|25x25 hard|9.253519ms|2m0.006197348s|2m0.001696382s|2m0.006197348s|2m0.018019636s|2m0.020902662s|84|

_Last Updated: 29 Jan 23 15:57 CST_
</resultsMarker>
