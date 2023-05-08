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
|6x6 easy|225.6µs|436.97µs|480.25µs|515.19µs|601.52µs|1.44ms|384|
|8x8 easy|323.96µs|566.18µs|627.29µs|690.44µs|1.55ms|1.95ms|349|
|8x8 medium|304.29µs|556.2µs|847.8µs|1.44ms|1.72ms|2.27ms|333|
|8x8 hard|580.07µs|1.35ms|1.48ms|1.68ms|1.93ms|3.52ms|309|
|10x10 easy|352.07µs|759.32µs|838.51µs|1.25ms|1.9ms|2.75ms|300|
|10x10 medium|360.38µs|1.24ms|1.6ms|1.8ms|2.35ms|5.88ms|280|
|10x10 hard|931.45µs|1.77ms|2.01ms|2.41ms|3.47ms|6.59ms|261|
|15x15 easy|722.8µs|1.52ms|2.31ms|2.68ms|3.23ms|4.5ms|241|
|15x15 medium|1.23ms|2.71ms|3.13ms|3.94ms|7.7ms|16.67ms|191|
|15x15 hard|2.09ms|4.12ms|7.15ms|14.43ms|68.53ms|271.02ms|147|
|20x20 easy|942.44µs|3.37ms|3.85ms|4.3ms|5.07ms|8.1ms|183|
|20x20 medium|2.28ms|4.96ms|7.73ms|14.05ms|80.4ms|12.79s|152|
|20x20 hard|4.29ms|26.85ms|105.59ms|340ms|11.17s|1m30s|128|
|30x30 hard|29.01ms|29.01ms|29.01ms|29.01ms|29.01ms|29.01ms|1|
|35x35 hard|115.13ms|115.13ms|115.13ms|115.13ms|115.13ms|115.13ms|1|
|40x40 hard|777.54ms|777.54ms|777.54ms|777.54ms|777.54ms|777.54ms|1|
|25x25 easy|2.36ms|5.04ms|5.78ms|6.56ms|9.66ms|107.38ms|131|
|25x25 medium|5.75ms|11.42ms|26.98ms|110.05ms|842.88ms|9.68s|109|
|25x25 hard|6.85ms|196.09ms|1.87s|42.24s|1m30.01s|1m30.01s|103|

_Last Updated: 06 May 23 21:26 CDT_
</resultsMarker>
