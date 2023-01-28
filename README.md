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

_Solve timeout: 30s_

|Puzzle|Min|Median|p75|p95|max|sample size|
|-|-|-|-|-|-|-:|
|6x6 easy|221.855µs|461.913µs|510.799µs|1.376001ms|1.936908ms|355|
|8x8 easy|295.244µs|629.816µs|1.454135ms|1.745324ms|2.267046ms|322|
|8x8 medium|241.86µs|1.491302ms|1.6937ms|2.109128ms|4.231265ms|300|
|8x8 hard|540.798µs|1.649484ms|1.852005ms|2.639142ms|11.379495ms|283|
|10x10 easy|484.62µs|876.926µs|1.76859ms|2.145214ms|2.677424ms|275|
|10x10 medium|504.655µs|1.861432ms|2.163105ms|3.398385ms|12.478061ms|258|
|10x10 hard|1.104349ms|2.219863ms|2.979811ms|7.253454ms|11.672296ms|239|
|15x15 easy|673.922µs|2.583094ms|2.897993ms|3.489032ms|6.861222ms|219|
|15x15 medium|1.276295ms|3.90583ms|6.255544ms|24.684653ms|1.546771678s|169|
|15x15 hard|2.976081ms|17.043915ms|58.374303ms|846.610948ms|30.002085342s|126|
|20x20 easy|1.769259ms|4.09817ms|4.506196ms|5.809724ms|33.543131ms|127|
|20x20 medium|2.988211ms|21.236494ms|114.787223ms|3.84316997s|30.002309464s|96|
|20x20 hard|4.962328ms|1.720355918s|30.001343183s|30.002637315s|30.012308685s|74|
|25x25 easy|4.203745ms|6.10316ms|7.041999ms|12.249805ms|76.23554ms|77|
|25x25 medium|11.522673ms|286.606005ms|3.416517108s|30.002615458s|30.003354894s|60|
|25x25 hard|284.316252ms|30.001857836s|30.00277704s|30.006295943s|30.015547422s|45|

_Last Updated: 28 Jan 23 08:28 CST_
</resultsMarker>