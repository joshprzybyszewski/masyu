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

_Solve timeout: 10s_

|Puzzle|Min|Median|p75|p95|max|sample size|
|-|-|-|-|-|-|-:|
|6x6 easy|205.632µs|484.179µs|523.997µs|1.411517ms|1.720024ms|348|
|8x8 easy|249.927µs|631.171µs|1.477795ms|1.762709ms|1.980796ms|315|
|8x8 medium|216.857µs|1.496594ms|1.673794ms|1.996238ms|4.190782ms|293|
|8x8 hard|505.646µs|1.670832ms|1.859173ms|2.859876ms|8.498662ms|276|
|10x10 easy|394.07µs|867.72µs|1.769882ms|2.006424ms|2.414768ms|269|
|10x10 medium|562.129µs|1.837661ms|2.184093ms|3.35828ms|13.865531ms|252|
|10x10 hard|1.007764ms|2.228925ms|2.998667ms|7.015487ms|11.871118ms|233|
|15x15 easy|917.006µs|2.571906ms|2.833349ms|3.461336ms|6.773051ms|213|
|15x15 medium|1.169901ms|3.759425ms|6.444728ms|19.814753ms|1.525126097s|163|
|15x15 hard|1.954382ms|18.624851ms|62.087419ms|885.134137ms|10.001129563s|120|
|20x20 easy|1.794577ms|3.920609ms|4.263592ms|5.734413ms|33.890149ms|122|
|20x20 medium|2.610292ms|22.123558ms|116.76344ms|3.875769727s|10.002332887s|91|
|20x20 hard|3.313349ms|1.409946627s|10.001649492s|10.00272185s|10.010403898s|69|
|25x25 easy|2.365839ms|5.951253ms|7.172489ms|12.228838ms|73.652757ms|72|
|25x25 medium|14.196813ms|285.493249ms|4.475674254s|10.003095671s|10.003286352s|56|
|25x25 hard|281.371658ms|10.001878851s|10.003014889s|10.006842747s|10.014052793s|43|

_Last Updated: 27 Jan 23 13:31 CST_
</resultsMarker>