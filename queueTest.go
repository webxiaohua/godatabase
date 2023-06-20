package main

import (
	"container/list"
	"fmt"
)

func main()  {
	queue := list.New()
	queue.PushBack(1)
	queue.PushBack(2)
	queue.PushBack(3)
	queue.PushBack(4)
	queue.PushBack(5)
	queue2 := make([]int,0)
	for {
		if queue.Len() == 0 {
			break
		}
		tmpItem := queue.Front()
		queue2 = append(queue2,tmpItem.Value.(int))
		queue.Remove(tmpItem)
	}
	fmt.Printf("%+v \n",queue2)


	guildMap := make(map[int64]*list.List)
	guildMap[1] = list.New()
	guildMap[2] = list.New()
	guildMap[1].PushBack(1101)
	guildMap[1].PushBack(1102)
	guildMap[2].PushBack(2101)
	guildMap[2].PushBack(2102)
	//guildMapStr, _ := json.Marshal(guildMap)


}