package main

// batch

var sharedCli = &esClientWithBuffer{}

type esClientWithBuffer struct {
	batchBuffer [][]interface{}
	messageCh   chan interface{} //本身就是个channel
	shortBuffer []interface{}
}

func (sli *esClientWithBuffer) pushBatch() {
	// todo 队列操作
}

func (cli *esClientWithBuffer) prepareBatch() {
	for msg := range cli.messageCh {
		if len(cli.shortBuffer) == batchSize {
			cli.batchBuffer = append(cli.batchBuffer, cli.shortBuffer) //从shortBuffer转移到batchBuffer
			cli.shortBuffer = []interface{}{}                          //清空了
		}
		cli.shortBuffer = append(cli.shortBuffer, msg)
	}
}

func pushToElasticSearchService(data interface{}) {
	sharedCli.messageCh <- data
}

var batchSize int = 20
