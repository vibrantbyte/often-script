/*
@Time : 2019-04-11 14:56 
@Author : xiaoyueya
@File : weighted_round_robin
@Software: GoLand
*/
package main

import (
	"fmt"
	"time"
)


//RoundRobinPeer
type RoundRobinPeer struct {
	// 编号
	Serial int64
	// 名称
	Name string
	// 权重
	Weight int32
	// 挂掉
	Down bool

	//--- private ---
	// 当前权重
	CurrentWeight int32
	// CurrentWeight + Weight
	EffectiveWeight int32
	// 失败次数
	Fails uint
	//--- private ---

	//(毫秒)
	Accessed int64
	// 上次检查时间(毫秒)
	Checked int64

	// 最大失败次数
	MaxFails uint
	// 过期时间(毫秒)
	FailTimeout int64
}

//RoundRobinPeerData
type RoundRobinPeerData struct {
	// 有序轮询集合
	Peers []*RoundRobinPeer
	// 有序集合大小
	Len uint
	// 选中的peer
	Current *RoundRobinPeer
}

func main()  {

	fmt.Println("初始化数据")
	data := new(RoundRobinPeerData)
	data.Peers = make([]*RoundRobinPeer,0)
	data.Len = 0

	peer1 := new(RoundRobinPeer)
	peer1.Name = "a"
	peer1.Serial = 1
	peer1.Weight = 5
	peer1.Down = false
	data.Peers = append(data.Peers,peer1)
	data.Len = 1


	peer2 := new(RoundRobinPeer)
	peer2.Name = "b"
	peer2.Serial = 2
	peer2.Weight = 3
	peer2.Down = false
	data.Peers = append(data.Peers,peer2)
	data.Len = 2


	peer3 := new(RoundRobinPeer)
	peer3.Name = "c"
	peer3.Serial = 3
	peer3.Weight = 1
	peer3.Down = false
	data.Peers = append(data.Peers,peer3)
	data.Len = 3

	for i := 0; i< 9 ;i++  {
		peer := GetPeer(data)
		if peer != nil {
			println(peer.Name)
		}
	}




}




func GetPeer(data *RoundRobinPeerData) *RoundRobinPeer{

	var best *RoundRobinPeer
	//当前时间纳秒
	now := GetMillisecond()
	//权重总值
	var total int32 = 0

	var i uint


	//遍历peer列表
	for i = 0 ;i < data.Len ;i++  {

		// 获取当前peer
		peer := data.Peers[i]

		// 检查当前后端服务器的 down 标志位，若为 true 表示不参与策略选择，则 continue 检查下一个后端服务器
		if peer.Down {
			continue
		}

		// 当前后端服务器的 down 标志位为 false,接着检查当前后端服务器连接失败的次数是否已经达到 max_fails；
		// 且睡眠的时间还没到 fail_timeout，则当前后端服务器不被选择，continue 检查下一个后端服务器；
		if peer.MaxFails > 0 && peer.Fails >= peer.MaxFails && now - peer.Checked <= peer.FailTimeout {
			continue
		}

		// 若当前后端服务器可能被选中，则计算其权重

		/*
		 * 在上面初始化过程中 current_weight = 0，effective_weight = weight；
		 * 此时，设置当前后端服务器的权重 current_weight 的值为原始值加上 effective_weight；
		 * 设置总的权重为原始值加上 effective_weight；
		 */
		peer.CurrentWeight += peer.EffectiveWeight
		total += peer.EffectiveWeight

		// 服务器正常，调整 effective_weight 的值 - 恢复服务
		if peer.EffectiveWeight < peer.Weight {
			peer.EffectiveWeight++
		}

		// 若当前后端服务器的权重 current_weight 大于目前 best 服务器的权重，则当前后端服务器被选中
		if best == nil || peer.CurrentWeight > best.CurrentWeight {
			best = peer
		}
	}

	if best == nil {
		return nil
	}

	data.Current = best
	best.CurrentWeight -= total

	if (now - best.Checked) > best.FailTimeout {
		best.Checked = now
	}

	return best
}

// 获取当前毫秒数
func GetMillisecond() int64 {
	return time.Now().UnixNano() / 1e6
}