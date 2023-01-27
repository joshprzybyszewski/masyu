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
|6x6 easy|228.015µs|493.247µs|538.867µs|1.400521ms|1.682479ms|348|
|8x8 easy|249.416µs|579.134µs|1.172199ms|1.641871ms|2.743029ms|315|
|8x8 medium|260.502µs|1.425275ms|1.609727ms|1.961894ms|3.449436ms|293|
|8x8 hard|239.146µs|1.688017ms|1.888165ms|3.006003ms|6.683725ms|276|
|10x10 easy|399.363µs|848.355µs|1.753863ms|2.039347ms|2.433413ms|269|
|10x10 medium|429.677µs|1.788923ms|2.063201ms|3.302303ms|14.384818ms|252|
|10x10 hard|909.484µs|2.224298ms|2.885473ms|7.055078ms|13.039431ms|233|
|15x15 easy|673.52µs|2.563994ms|2.834168ms|3.314582ms|6.699005ms|213|
|15x15 medium|1.363569ms|3.797204ms|6.346514ms|20.414325ms|1.551340867s|163|
|15x15 hard|2.624312ms|18.314131ms|63.236645ms|904.023049ms|30.001396827s|120|
|20x20 easy|2.017367ms|4.011324ms|4.447589ms|5.38401ms|33.18855ms|122|
|20x20 medium|2.589611ms|21.896561ms|115.015041ms|3.866483538s|30.002247636s|91|
|20x20 hard|4.390476ms|1.41880537s|30.001402442s|30.002225353s|30.011686604s|69|
|25x25 easy|2.994106ms|5.939569ms|7.35207ms|12.203552ms|79.945288ms|72|
|25x25 medium|16.485253ms|281.9003ms|4.486196506s|30.002833407s|30.002968462s|56|
|25x25 hard|278.963733ms|30.001760384s|30.002736437s|30.003492975s|30.014600655s|43|

_Last Updated: 27 Jan 23 14:12 CST_
</resultsMarker>