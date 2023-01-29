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
|6x6 easy|208.966µs|432.568µs|470.886µs|556.964µs|1.493753ms|365|
|8x8 easy|228.734µs|590.019µs|644.909µs|1.584834ms|1.904925ms|331|
|8x8 medium|236.631µs|650.218µs|1.482872ms|1.739392ms|2.721269ms|309|
|8x8 hard|461.205µs|1.519375ms|1.723342ms|2.046272ms|3.503219ms|292|
|10x10 easy|344.917µs|836.232µs|1.41482ms|1.951329ms|2.773776ms|284|
|10x10 medium|490.628µs|1.631702ms|1.856467ms|2.343091ms|5.577548ms|266|
|10x10 hard|1.17942ms|2.02851ms|2.410772ms|4.010887ms|9.408063ms|247|
|15x15 easy|700.93µs|2.304653ms|2.58058ms|2.945211ms|3.746441ms|227|
|15x15 medium|1.287461ms|3.115656ms|3.940769ms|8.709362ms|134.771467ms|177|
|15x15 hard|2.09826ms|8.593841ms|16.025195ms|64.269371ms|602.952267ms|133|
|20x20 easy|1.551753ms|3.684739ms|4.151422ms|4.96738ms|8.609057ms|169|
|20x20 medium|2.855121ms|10.169641ms|29.350544ms|267.403868ms|16.999737209s|138|
|20x20 hard|3.890778ms|400.517937ms|8.294356705s|30.008654672s|30.017169339s|115|
|25x25 easy|2.896007ms|5.275033ms|6.041628ms|10.503508ms|32.706615ms|118|
|25x25 medium|5.729474ms|51.090466ms|512.155436ms|17.6863446s|30.014319499s|97|
|25x25 hard|18.879361ms|30.002181333s|30.004907308s|30.01885186s|30.020191011s|81|

_Last Updated: 28 Jan 23 20:53 CST_
</resultsMarker>
