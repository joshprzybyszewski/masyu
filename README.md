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
|6x6 easy|215.511µs|481.049µs|523.104µs|1.428994ms|1.63479ms|355|
|8x8 easy|340.802µs|601.429µs|1.436114ms|1.694541ms|2.131447ms|322|
|8x8 medium|251.94µs|1.491884ms|1.715375ms|2.085179ms|3.927292ms|300|
|8x8 hard|505.285µs|1.665651ms|1.903927ms|2.94031ms|10.41404ms|283|
|10x10 easy|344.302µs|874.24µs|1.762519ms|2.088877ms|2.650511ms|275|
|10x10 medium|551.386µs|1.841553ms|2.221349ms|3.483415ms|13.587724ms|258|
|10x10 hard|1.066205ms|2.284202ms|2.931428ms|7.203698ms|12.523765ms|239|
|15x15 easy|1.016128ms|2.630274ms|2.960161ms|3.505787ms|6.560219ms|219|
|15x15 medium|1.28782ms|3.863986ms|6.427106ms|24.410226ms|1.542041319s|169|
|15x15 hard|2.797193ms|16.301714ms|55.989082ms|862.023198ms|30.001414226s|126|
|20x20 easy|1.888567ms|4.092278ms|4.535851ms|5.376431ms|36.241887ms|162|
|20x20 medium|3.857195ms|20.483888ms|121.020955ms|4.268777249s|30.002001878s|131|
|20x20 hard|4.740794ms|2.285029688s|30.001135411s|30.00255294s|30.014026783s|108|
|25x25 easy|2.799659ms|6.003886ms|6.99413ms|14.150342ms|76.263486ms|111|
|25x25 medium|2.762974ms|362.873294ms|3.138360338s|30.002799024s|30.003536814s|90|
|25x25 hard|27.007198ms|30.001836836s|30.002798418s|30.003411454s|30.013898779s|77|

_Last Updated: 28 Jan 23 17:40 CST_
</resultsMarker>
