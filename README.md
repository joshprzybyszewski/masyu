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
|6x6 easy|186.059µs|426.034µs|452.915µs|534.111µs|1.500418ms|365|
|8x8 easy|263.764µs|598.299µs|651.981µs|1.580922ms|1.921716ms|331|
|8x8 medium|344.948µs|614.762µs|1.452749ms|1.73976ms|2.330115ms|309|
|8x8 hard|510.664µs|1.552967ms|1.700363ms|2.154755ms|3.417565ms|292|
|10x10 easy|354.799µs|802.008µs|1.45783ms|1.86714ms|2.278794ms|284|
|10x10 medium|426.565µs|1.623665ms|1.822695ms|2.310947ms|6.929059ms|266|
|10x10 hard|1.416741ms|2.008575ms|2.453387ms|4.273096ms|9.407792ms|247|
|15x15 easy|831.891µs|2.296378ms|2.594161ms|3.012318ms|3.76717ms|227|
|15x15 medium|1.144746ms|3.119787ms|3.992227ms|7.961123ms|126.937488ms|177|
|15x15 hard|2.415529ms|9.542062ms|14.979651ms|62.876494ms|613.029515ms|133|
|20x20 easy|817.753µs|3.422922ms|3.930433ms|4.944004ms|7.03508ms|169|
|20x20 medium|2.43787ms|9.866112ms|29.621979ms|269.052608ms|16.880640782s|138|
|20x20 hard|5.515366ms|408.68245ms|8.062463465s|30.00340995s|30.020371265s|115|
|25x25 easy|2.780786ms|5.46537ms|5.955526ms|10.495071ms|35.723983ms|118|
|25x25 medium|5.812958ms|51.121839ms|491.779601ms|17.462106519s|30.016599967s|97|
|25x25 hard|18.925271ms|30.002011248s|30.005996777s|30.014449447s|30.018367564s|81|

_Last Updated: 28 Jan 23 19:45 CST_
</resultsMarker>
