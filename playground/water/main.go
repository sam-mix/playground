package main

import "fmt"

type cup struct {
	name   string // 名字
	cap    int    // 剩余容积 单位升
	water  int    // 当前水量
	maxCap int    // 最大容积
}

func main() {
	water()
}

// 结果需要弄出的6升水
const (
	rw   = 6
	capA = 7
	capB = 5
)

/*
 * 现在有两个 5L 和 7L 装的杯子， 水随便用， 要求倒出 6L 水
 */
func water(cs ...cup) {

	// 初始
	if len(cs) == 0 {
		ca := cup{
			name:   "A",
			cap:    0,
			water:  capA,
			maxCap: capA,
		}
		cb := cup{
			name:   "B",
			cap:    capB,
			water:  0,
			maxCap: capB,
		}
		water(ca, cb)

	}

	if cs[0].water == rw || cs[1].water == rw {
		fmt.Println("结束了")
		return
	}
	fm := cs[0]
	to := cs[1]

	fm.water = to.cap - fm.water

}
