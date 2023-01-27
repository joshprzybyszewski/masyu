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
|6x6 easy|198.718µs|461.63µs|509.021µs|1.42242ms|348|
|8x8 easy|249.111µs|621.716µs|1.45217ms|1.718441ms|315|
|8x8 medium|233.801µs|1.510652ms|1.690438ms|2.079029ms|293|
|8x8 hard|459.657µs|1.675624ms|1.857375ms|2.961232ms|276|
|10x10 easy|393.215µs|852.306µs|1.785071ms|2.073919ms|269|
|10x10 medium|410.148µs|1.850899ms|2.147861ms|3.389086ms|252|
|10x10 hard|932.628µs|2.210345ms|2.938531ms|7.143405ms|233|
|15x15 easy|749.378µs|2.617104ms|2.826925ms|3.472324ms|213|
|15x15 medium|1.215016ms|3.809013ms|6.057636ms|20.287609ms|163|
|15x15 hard|2.601484ms|17.732978ms|60.50996ms|875.572368ms|119|
|20x20 easy|2.005964ms|4.069719ms|4.32859ms|5.391534ms|122|
|20x20 medium|2.941205ms|22.393288ms|98.340796ms|1.006102027s|88|
|20x20 hard|46.359554ms|394.767468ms|2.313791525s|7.876374691s|9|
|25x25 easy|3.426586ms|6.073167ms|7.107932ms|10.982163ms|72|
|25x25 medium|13.419225ms|217.412824ms|1.269891495s|5.114112803s|47|
|25x25 hard|280.510155ms|4.090251659s|8.771279853s|8.771279853s|3|

_Last Updated: 27 Jan 23 12:55 CST_
</resultsMarker>