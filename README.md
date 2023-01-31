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
|6x6 easy|235.413µs|408.238µs|446.553µs|484.409µs|569.609µs|2.333501ms|382|
|8x8 easy|223.247µs|487.244µs|554.08µs|609.476µs|1.44906ms|1.981521ms|347|
|8x8 medium|215.558µs|437.827µs|662.497µs|1.167168ms|1.531533ms|2.137098ms|324|
|8x8 hard|468.81µs|1.316771ms|1.451456ms|1.627743ms|2.031612ms|7.67072ms|307|
|10x10 easy|306.422µs|682.153µs|762.814µs|1.475739ms|1.882465ms|2.488835ms|299|
|10x10 medium|431.235µs|1.142615ms|1.573071ms|1.78561ms|2.34598ms|6.169488ms|279|
|10x10 hard|853.03µs|1.737005ms|1.982781ms|2.331145ms|3.559717ms|5.235965ms|260|
|15x15 easy|884.36µs|1.457742ms|2.215338ms|2.578608ms|3.004402ms|4.156662ms|240|
|15x15 medium|1.224894ms|2.533978ms|3.051525ms|3.868775ms|8.173286ms|15.482397ms|190|
|15x15 hard|2.4402ms|4.188364ms|7.26325ms|15.376869ms|68.713605ms|267.085677ms|146|
|20x20 easy|1.519874ms|3.238392ms|3.710106ms|4.206411ms|4.917073ms|8.213881ms|182|
|20x20 medium|2.082205ms|4.712486ms|7.598068ms|13.135445ms|70.751397ms|13.632524907s|151|
|20x20 hard|4.166376ms|26.805718ms|106.297602ms|329.601355ms|11.597320543s|45.006771288s|127|
|25x25 easy|2.234247ms|4.117569ms|5.156644ms|5.931444ms|8.854141ms|107.754041ms|130|
|25x25 medium|5.602239ms|11.671643ms|27.0944ms|108.572935ms|836.012641ms|9.986674795s|108|
|25x25 hard|7.867334ms|178.467896ms|1.821871405s|42.489879791s|45.013387088s|45.0318067s|102|

_Last Updated: 31 Jan 23 12:07 CST_
</resultsMarker>
