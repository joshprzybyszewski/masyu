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

_Solve timeout: 30s_

|Puzzle|Min|Median|p75|p95|max|sample size|
|-|-|-|-|-|-|-:|
|6x6 easy|223.312µs|477.064µs|515.671µs|1.394432ms|1.726631ms|355|
|8x8 easy|220.534µs|583.425µs|1.414368ms|1.71112ms|2.148769ms|322|
|8x8 medium|390.568µs|1.501674ms|1.712752ms|2.072066ms|3.978002ms|300|
|8x8 hard|497.29µs|1.703205ms|1.897357ms|2.903503ms|11.019765ms|283|
|10x10 easy|332.985µs|850.083µs|1.766ms|2.113741ms|2.403579ms|275|
|10x10 medium|521.827µs|1.906567ms|2.234859ms|3.477579ms|12.650202ms|258|
|10x10 hard|1.222395ms|2.232876ms|2.875546ms|7.040536ms|12.751256ms|239|
|15x15 easy|1.08661ms|2.558725ms|2.873616ms|3.467896ms|6.908544ms|219|
|15x15 medium|1.14189ms|3.851121ms|6.033323ms|22.907881ms|1.53948882s|169|
|15x15 hard|2.847773ms|16.629952ms|52.792007ms|869.726161ms|30.00128953s|126|
|20x20 easy|1.786229ms|4.041495ms|4.436741ms|5.740916ms|33.136953ms|127|
|20x20 medium|3.578868ms|21.61837ms|116.089002ms|3.824698323s|30.002106959s|96|
|20x20 hard|4.593844ms|1.704187546s|30.001557118s|30.002361214s|30.011390325s|74|
|25x25 easy|4.187052ms|6.091488ms|6.990712ms|12.918043ms|75.527183ms|77|
|25x25 medium|9.786067ms|284.987059ms|3.500762527s|30.002764234s|30.003224743s|60|
|25x25 hard|284.973043ms|30.001709971s|30.00284483s|30.006617124s|30.020998395s|45|

_Last Updated: 28 Jan 23 10:39 CST_
</resultsMarker>
