package main

import (
	"fmt"
	"sync"
)

func main() {

	help := `
	gc#：GC 执行次数的编号，每次叠加。

	@#s：自程序启动后到当前的具体秒数。
	
	#%：自程序启动以来在GC中花费的时间百分比。
	
	#+...+#：GC 的标记工作共使用的 CPU 时间占总 CPU 时间的百分比。
	
	#->#-># MB：分别表示 GC 启动时, GC 结束时, GC 活动时的堆大小.
	
	#MB goal：下一次触发 GC 的内存占用阈值。
	
	#P：当前使用的处理器 P 的数量。
	`

	fmt.Println(help)

	wg := sync.WaitGroup{}

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(wg *sync.WaitGroup) {
			var counter int
			for i := 0; i < 1e10; i++ {
				counter++
			}
			wg.Done()
		}(&wg)
	}
	wg.Wait()

}
