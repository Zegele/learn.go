package main

// c.biancheng.net/view/4359.html
import (
	"fmt"
	"sync"
	"time"
)

// 模拟接力赛
// 4个跑步者围绕赛道轮流跑。第二个，第三个和第四个跑步者要接到前一位跑步者的接力棒后才能起跑。
// 比赛中最重要的部分时要传递接力棒，要求同步传递。
// 在同步接力棒的时候，参与接力的两个跑步者必须在同一时刻准备好交接。

// 这个示例展示如何用无缓冲的通道来模拟，4个goroutine间的接力

// wg 用来等待程序结束
var wg sync.WaitGroup

func main() {
	// 创建一个无缓冲的通道
	baton := make(chan int)

	// 为最后一位跑步者将计数加1
	wg.Add(1)

	// 第一位跑步者持有接力棒
	go Runner(baton)

	// 比赛开始
	baton <- 1
	// 等待比赛结束
	wg.Wait()
}

// Runner 模拟接力比赛中的一位跑步者
func Runner(baton chan int) {
	var newRunner int

	// 等待接力棒
	runner := <-baton
	// 开始绕着跑道跑步
	fmt.Printf("Runner %d Running With Baton\n", runner)

	//创建下一位跑步者
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		go Runner(baton)
	}

	// 围绕跑道跑
	time.Sleep(100 * time.Millisecond)

	// 比赛结束了吗？
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		wg.Done()
		return
	}

	// 将接力棒交给下一位跑步者
	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)

	baton <- newRunner
}

/*模拟打网球
// 在网球比赛中，两位选手会把球在两个人之间来回传递。
//选手总是处在以下两种状态之一，要么在等待接球，要么将球打向对方。

// wg 用来等待程序结束
var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	//创建一个无缓冲的通道
	court := make(chan int)

	// 计数加2， 表示要等待2个goroutine
	wg.Add(2)

	// 启动两个选手
	go player("Nadal", court)
	go player("DDDD", court)

	// 发球
	court <- 1 // 上面的两个goroutine都在等接收数据

	// 等待游戏结束
	wg.Wait()
}

//player 模拟一个选手在打网球
func player(name string, court chan int) {
	// 在函数退出时调用Done 来通知main函数工作已经完成
	defer wg.Done()

	for {
		// 等待球被打击过来
		ball, ok := <-court
		if !ok {
			// 如果通道被关闭，我们就赢了
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// 选随机数，然后用这个数来判断我们是否丢球
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			// 关闭通道， 表示我们输了
			close(court)
			return
		}

		// 显示击球数，并将击球数加1
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// 将球打向对手
		court <- ball
	}
}


*/
