package main

import "fmt"

// 结果需要弄出的6升水
const (
	rw   = 6
	capA = 7
	capB = 5
)

type cup struct {
	name   string // 名字
	cap    int    // 剩余容积 单位升
	water  int    // 当前水量
	maxCap int    // 最大容积
}

// 装满水
func (c *cup) full() {
	c.water = c.maxCap
	c.cap = 0
}

// 倒掉水
func (c *cup) clean() {
	c.water = 0
	c.cap = c.maxCap
}

// 检查是否装满
func (c *cup) checkFull() bool {
	return c.water >= c.maxCap
}

// 检查是否倒空了
func (c *cup) checkEmpty() bool {
	return c.water <= 0
}

// 重启新计算剩余容积
func (c *cup) rCap() {
	c.cap = c.maxCap - c.water
}

// 从当前杯子向另一个杯子倒水倒满为止
func (fm *cup) swap(to *cup) {
	if fm.water > to.cap {
		fm.water = fm.water - to.cap
		fm.rCap()
		to.full()
	} else {
		fm.clean()
		to.water = to.water + fm.water
		to.rCap()
	}
}

// 展示水量
func (c *cup) showWater() {
	fmt.Printf("Cup %s water is %d L\n", c.name, c.water)
}

// 设置最大容积
func newCup(name string, maxCap int) *cup {
	return &cup{
		name:   name,
		cap:    maxCap,
		maxCap: maxCap,
	}
}

// 第一步 将容量小的杯子装满水
func first(cs ...*cup) (fm, to *cup) {
	if cs[0].maxCap < cs[1].maxCap {
		cs[0].full()
		fm = cs[0]
		to = cs[1]
	} else {
		cs[1].full()
		fm = cs[1]
		to = cs[0]
	}
	return
}

// 初始两个杯子
func initCups() (*cup, *cup) {
	return newCup("A", capA), newCup("B", capB)
}

func main() {
	water()
}

/*
 * 现在有两个 5L 和 7L 装的杯子， 水随便用， 要求倒出 6L 水
 */
func water(cs ...*cup) {
	time.Sleep(500 * time.)
	// 初始
	if len(cs) == 0 {
		water(first(initCups()))
	}

	cs[0].swap(cs[1])
	cs[0].showWater()
	cs[1].showWater()
	if cs[0].water == rw || cs[1].water == rw {
		fmt.Println("结束了:")
		return
	}
	if cs[1].checkFull() {
		cs[1].clean()
		water(cs...)
	}
	if cs[0].checkEmpty() {
		cs[0].full()
		water(cs[1], cs[0])
	}

}
