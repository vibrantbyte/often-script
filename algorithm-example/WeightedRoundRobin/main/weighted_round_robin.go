/*
@Time : 2019-04-11 14:56 
@Author : xiaoyueya
@File : weighted_round_robin
@Software: GoLand
*/
package main

import (
	"fmt"
)



var firstOpenApplication = true

func main()  {
	/**
	1、首次启动检测redis中数据是否存在，使用的是redis的hash数据结构来存储。
	2、如果不为空的话，需要和内存副本中获取的数据进行排序比较，然后重新初始化。
	3、开始进行权益公平算法调度。
	5、每次调度结果持久化到redis中，防止服务器宕机，找不到上一次的顺序。（如果不需要多实例下的强一致性，可以将缓存落盘到磁盘）
	6、实时动态的操作更新当前运行的算法。
	*/

	if firstOpenApplication {
		//1、首次！！！读取redis缓存中的上一次派单记录
		firstOpenApplication = false
	}

	data := new(RoundRobinPeerData)

	peer1 := new(RoundRobinPeer)
	peer1.Name = "a"
	peer1.Serial = 1
	peer1.Weight = 5
	peer1.EffectiveWeight = 5
	peer1.Down = false
	data.Append(peer1)


	peer2 := new(RoundRobinPeer)
	peer2.Name = "b"
	peer2.Serial = 2
	peer2.Weight = 3
	peer2.EffectiveWeight = 3
	peer2.Down = false
	data.Append(peer2)


	peer3 := new(RoundRobinPeer)
	peer3.Name = "c"
	peer3.Serial = 3
	peer3.Weight = 1
	peer3.EffectiveWeight = 1
	peer3.Down = false
	data.Append(peer3)

	peer4 := new(RoundRobinPeer)
	peer4.Name = "d"
	peer4.Serial = 4
	peer4.Weight = 0
	peer4.EffectiveWeight = 0
	peer4.Down = false
	data.Append(peer4)

	//2、比较后，生成对应的缓存信息 - 其他线程更新的时候操作的是本地缓存。

	//3、开始进行平衡权益算法调度
	for i := 0; i< 6 ;i++  {
		peer := GetPeer(data)
		if peer != nil {
			fmt.Println(peer.Name)
		}
		//3.1 递归

	}

	//4、将执行结果持久化到redis或其他存储介质
	var i uint
	for i=0;i<data.GetLen() ;i++  {
		peer := data.GetIndex(i)
		if !peer.Down {
			fmt.Printf("partnerId:%d,currentWeight:%d \n",peer.Serial,peer.CurrentWeight)
		}
	}



}

