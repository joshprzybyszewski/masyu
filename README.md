# masyu
Masyu Solver - golang

This is the self-proclaimed World's fastest solver for [the masyu puzzle](www.puzzle-masyu.com). It is similar to [my Shingoki Solver](https://github.com/joshprzybyszewski/shingokisolver).

To run, execute `make compete`.

## Results

Check the [Hall of Fame](www.puzzle-masyu.com/hall.php?hallsize=18) for the results recorded by the puzzle server (which include network duration of submission). Below are the results of the solver as recorded on my machine.

_NOTE: Update this table with `make results`._

<resultsMarker>

_GOOS: darwin_

_GOARCH: arm64_

_Solve timeout: 1m30s_

|Puzzle|Min|p25|Median|p75|p95|max|sample size|
|-|-|-|-|-|-|-|-:|
|5x5 easy|99.79µs|282.04µs|404.33µs|522.08µs|1.1ms|1.4ms|129|
|5x5 medium|85.45µs|248.45µs|445.25µs|637.25µs|1.2ms|1.37ms|124|
|7x7 easy|182.54µs|469.29µs|768.25µs|1.14ms|1.73ms|1.92ms|120|
|7x7 medium|186.79µs|468.91µs|694.91µs|1.06ms|1.67ms|2.31ms|116|
|7x7 hard|227.79µs|681.75µs|950.58µs|1.28ms|1.98ms|2.41ms|113|
|10x10 easy|445.29µs|1.05ms|1.59ms|2.04ms|2.67ms|3.49ms|108|
|10x10 medium|497.29µs|943.29µs|1.22ms|1.83ms|2.9ms|3.78ms|106|
|10x10 hard|661.54µs|1.2ms|1.82ms|2.4ms|3.85ms|7.95ms|106|
|15x15 easy|1.07ms|2.18ms|2.86ms|3.85ms|4.88ms|6.53ms|104|
|15x15 medium|869.16µs|2.17ms|2.72ms|3.81ms|4.71ms|8.64ms|103|
|15x15 hard|1.51ms|2.2ms|2.83ms|3.76ms|7.59ms|34.84ms|102|
|20x20 easy|2ms|3ms|4.03ms|4.9ms|6.41ms|8.08ms|98|
|20x20 medium|1.86ms|3.59ms|4.24ms|5.2ms|6.87ms|8.2ms|98|
|20x20 hard|2.69ms|3.9ms|4.93ms|8.07ms|147.06ms|498.85ms|97|
|25x25 easy|3.27ms|5.24ms|6.38ms|7.14ms|8.38ms|11.66ms|94|
|25x25 medium|3.33ms|5.35ms|6.39ms|7.37ms|8.57ms|9.43ms|97|
|25x25 hard|3.99ms|14.73ms|85.79ms|412.33ms|2.67s|22.93s|95|
|30x30 hard|5.67ms|5.67ms|5.67ms|5.67ms|5.67ms|5.67ms|1|

_Last Updated: 22 Oct 23 13:56 CDT_
</resultsMarker>
