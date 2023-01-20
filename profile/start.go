package profile

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

const (
	cpuProfileFilename    = `cpu.pprof`
	memoryProfileFilename = `mem.pprof`
)

func Start() func() {
	f, err := os.Create(cpuProfileFilename)
	if err != nil {
		log.Printf("got error attempting to create cpu file: %q\n", err)
		return func() {}
	}

	err = pprof.StartCPUProfile(f)
	if err != nil {
		log.Printf("got error attempting to start a profile: %q\n", err)
		return func() {}
	}

	return func() {
		pprof.StopCPUProfile()
		err = f.Close()
		if err != nil {
			log.Printf("got error attempting to close profile file: %q\n", err)
		}

		f, err := os.Create(memoryProfileFilename)
		if err != nil {
			log.Printf("got error attempting to create memory file: %q\n", err)
			return
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Printf("got error attempting to write memory file: %q\n", err)
		}
	}
}
