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

_Solve timeout: 45s_

|Puzzle|Min|p25|Median|p75|p95|max|sample size|
|-|-|-|-|-|-|-|-:|
|6x6 easy|232.571µs|434.764µs|467.619µs|494.67µs|561.95µs|1.598504ms|373|
|8x8 easy|270.914µs|527.948µs|570.036µs|631.452µs|1.818014ms|2.968374ms|339|
|8x8 medium|344.917µs|499.553µs|661.307µs|1.759394ms|2.738823ms|5.519513ms|316|
|8x8 hard|485.113µs|1.723131ms|2.288504ms|2.939977ms|4.680117ms|7.263835ms|299|
|10x10 easy|314.003µs|727.573µs|790.576µs|1.687516ms|2.726689ms|5.64842ms|291|
|10x10 medium|430.375µs|1.545694ms|2.278164ms|2.970803ms|5.299194ms|11.568975ms|273|
|10x10 hard|1.454834ms|3.14072ms|4.36678ms|5.578068ms|8.500345ms|12.735686ms|254|
|15x15 easy|563.136µs|1.505947ms|2.643781ms|3.809499ms|6.312599ms|11.367359ms|234|
|15x15 medium|1.019455ms|5.095626ms|8.384954ms|9.862035ms|16.546006ms|81.365449ms|184|
|15x15 hard|4.992662ms|13.020788ms|18.734639ms|46.34839ms|153.673564ms|308.960824ms|140|
|20x20 easy|1.617791ms|4.219382ms|6.273429ms|9.710331ms|12.397982ms|18.356131ms|176|
|20x20 medium|3.492789ms|14.038911ms|22.840999ms|45.331497ms|225.680915ms|24.005670528s|145|
|20x20 hard|17.544537ms|113.504671ms|384.454715ms|3.595957637s|45.010967212s|45.020173601s|121|
|25x25 easy|2.566265ms|6.926929ms|10.720036ms|12.806081ms|24.062753ms|46.512259ms|124|
|25x25 medium|14.943583ms|61.11585ms|188.289747ms|1.197868644s|45.0040565s|45.022052802s|102|
|25x25 hard|75.899197ms|4.540321776s|45.00235454s|45.010535458s|45.018551104s|45.023067752s|97|

_Last Updated: 30 Jan 23 14:39 CST_
</resultsMarker>
