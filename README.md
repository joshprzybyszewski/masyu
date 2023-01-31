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

_Solve timeout: 1m30s_

|Puzzle|Min|p25|Median|p75|p95|max|sample size|
|-|-|-|-|-|-|-|-:|
|6x6 easy|231.39µs|444.33µs|482.68µs|530.3µs|625.59µs|1.53ms|383|
|8x8 easy|274.01µs|552µs|605.12µs|667.2µs|1.59ms|1.84ms|348|
|8x8 medium|304.81µs|560.37µs|823.12µs|1.51ms|1.81ms|2.34ms|325|
|8x8 hard|507.58µs|1.47ms|1.63ms|1.81ms|2.11ms|3.28ms|308|
|10x10 easy|319.74µs|675.88µs|800.63µs|1.02ms|1.9ms|2.73ms|300|
|10x10 medium|320.38µs|1.01ms|1.58ms|1.88ms|2.44ms|6.47ms|280|
|10x10 hard|992.85µs|1.77ms|2.03ms|2.39ms|3.29ms|6.66ms|261|
|15x15 easy|711.4µs|1.44ms|2.08ms|2.62ms|3.16ms|5.69ms|241|
|15x15 medium|1.3ms|2.68ms|3.18ms|3.94ms|7.95ms|17.21ms|191|
|15x15 hard|2.04ms|4.34ms|6.96ms|14.04ms|70.01ms|276.31ms|147|
|20x20 easy|1.81ms|3.19ms|3.78ms|4.27ms|5.29ms|8.7ms|183|
|20x20 medium|2.02ms|4.73ms|7ms|13.17ms|80.64ms|14.04s|152|
|20x20 hard|3.87ms|27.67ms|114.93ms|329.73ms|11.31s|1m30.01s|128|
|30x30 hard|30.08ms|30.08ms|30.08ms|30.08ms|30.08ms|30.08ms|1|
|35x35 hard|122.07ms|122.07ms|122.07ms|122.07ms|122.07ms|122.07ms|1|
|40x40 hard|774.48ms|774.48ms|774.48ms|774.48ms|774.48ms|774.48ms|1|
|25x25 easy|2.65ms|4.47ms|5.18ms|5.8ms|8.76ms|106.99ms|131|
|25x25 medium|4.9ms|11.93ms|28.28ms|109.8ms|852.88ms|10.06s|109|
|25x25 hard|7.1ms|188.9ms|1.9s|43.47s|1m30.01s|1m30.01s|103|

_Last Updated: 31 Jan 23 14:18 CST_
</resultsMarker>
