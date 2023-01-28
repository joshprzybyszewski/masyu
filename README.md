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
|6x6 easy|234.114µs|462.879µs|504.368µs|1.372282ms|1.774429ms|355|
|8x8 easy|251.038µs|624.911µs|1.404939ms|1.686326ms|2.151234ms|322|
|8x8 medium|280.08µs|1.448204ms|1.686698ms|2.069082ms|3.743023ms|300|
|8x8 hard|543.196µs|1.697013ms|1.896839ms|2.948658ms|10.041375ms|283|
|10x10 easy|367.243µs|876.811µs|1.753858ms|2.093682ms|2.405527ms|275|
|10x10 medium|388.021µs|1.856784ms|2.161487ms|3.472172ms|13.409987ms|258|
|10x10 hard|833.446µs|2.154286ms|2.918024ms|7.008692ms|10.769771ms|239|
|15x15 easy|1.01746ms|2.644183ms|2.898781ms|3.483214ms|6.742431ms|219|
|15x15 medium|1.371403ms|3.820233ms|5.919119ms|24.726262ms|1.523760932s|169|
|15x15 hard|2.874299ms|16.913347ms|58.816986ms|852.157553ms|30.001113073s|126|
|20x20 easy|1.771891ms|3.983184ms|4.386323ms|5.436219ms|33.827211ms|127|
|20x20 medium|3.770224ms|21.421598ms|115.03368ms|3.850878408s|30.002374296s|96|
|20x20 hard|4.466431ms|1.72084183s|30.001299515s|30.00247033s|30.006559739s|74|
|25x25 easy|4.347552ms|5.994877ms|7.396517ms|12.550681ms|76.271804ms|77|
|25x25 medium|11.867927ms|272.569492ms|3.409668468s|30.003192548s|30.003294718s|60|
|25x25 hard|277.455431ms|30.001879318s|30.00264347s|30.004829661s|30.017510464s|45|

_Last Updated: 28 Jan 23 07:46 CST_
</resultsMarker>