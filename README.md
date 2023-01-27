# masyu
Masyu Solver - golang

Here's a solver for [the masyu puzzle](www.puzzle-masyu.com). It is similar to [my Shingoki Solver](https://github.com/joshprzybyszewski/shingokisolver).

To run, execute `make compete`.

## Results

Check the Hall of Fame for the results recorded by the puzzle server (which include network duration of submission). Below are the results of the solver as recorded on my machine.

_NOTE: Update this table with `make results`._

<resultsMarker>
_GOOS: linux_

_GOARCH: amd64_

_cpu: Intel(R) Core(TM) i5-3570 CPU @ 3.40GHz_

|Puzzle|Min|Median|p75|p95|sample size|
|-|-|-|-|-|-:|
|6x6 easy|410.258µs|524.036µs|553.732µs|1.366089ms|10|
|8x8 easy|533.048µs|668.548µs|686.31µs|1.429323ms|10|
|8x8 medium|202.777µs|1.700456ms|1.852931ms|3.060469ms|10|
|8x8 hard|730.124µs|1.239221ms|2.109289ms|6.730459ms|10|
|10x10 easy|326.857µs|821.995µs|1.257429ms|1.409966ms|10|
|10x10 medium|1.219039ms|1.801411ms|1.845709ms|3.236479ms|10|
|10x10 hard|1.16202ms|1.802755ms|1.971914ms|18.134432ms|10|
|15x15 easy|778.375µs|2.100882ms|2.603295ms|3.179069ms|10|
|15x15 medium|1.200555ms|4.251559ms|18.752492ms|133.913446ms|10|
|15x15 hard|2.63341ms|27.479924ms|36.804052ms|288.407098ms|10|
|20x20 easy|1.297954ms|2.42244ms|3.153105ms|3.517365ms|10|
|20x20 medium|5.727719ms|45.536249ms|361.447198ms|497.900183ms|10|
|20x20 hard|56.356391ms|213.703091ms|576.951628ms|1.001666626s|7|
|25x25 easy|2.679198ms|5.290893ms|6.564947ms|8.822726ms|10|
|25x25 medium|26.103676ms|65.186098ms|115.679457ms|302.148579ms|5|
|25x25 hard|300.924925ms|359.978664ms|359.978664ms|359.978664ms|2|

_Last Updated: 27 Jan 23 11:45 CST_
</resultsMarker>