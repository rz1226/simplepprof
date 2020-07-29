package simplepprof

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

var cpupprofFile *os.File

func init() {
	//锁
	runtime.SetMutexProfileFraction(1)

	//block
	runtime.SetBlockProfileRate(1)

	//cpu
	var err error
	cpupprofFile, err = os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
		return
	}
	pprof.StartCPUProfile(cpupprofFile)
}

//使用方法，只需要引入这个包，在main结束的时候运行此函数即可。
func DoPProfWhenExit() {
	//堆
	pprofMemFile, err := os.Create("mem.pprof")
	if err != nil {
		fmt.Println("can create mem.pprof")
		log.Fatal(err)
	}
	defer pprofMemFile.Close()
	pprof.WriteHeapProfile(pprofMemFile)

	//锁
	mutexpprofFile, err := os.Create("mutex.pprof")
	if err != nil {
		log.Fatal(err)
	}
	defer mutexpprofFile.Close()
	pprof.Lookup("mutex").WriteTo(mutexpprofFile, 1)

	//block
	blockpprofFile, err := os.Create("block.pprof")
	if err != nil {
		log.Fatal(err)
	}
	defer blockpprofFile.Close()
	pprof.Lookup("block").WriteTo(blockpprofFile, 1)

	//cpu
	pprof.StopCPUProfile()
	defer cpupprofFile.Close()

}
