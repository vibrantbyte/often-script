/*
@Time : 2019-04-11 14:56 
@Author : xiaoyueya
@File : weighted_round_robin
@Software: GoLand
*/
package main

import (
	"fmt"
	"sync"
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
	// 读写锁保证data内部对象操作的线程安全
	sync.RWMutex
	// 有序轮询集合 （线程共享）
	peers []*RoundRobinPeer
	// 有序集合大小 	(线程共享)
	len uint
}

//getLen （safe）
func (data *RoundRobinPeerData) GetLen() uint{
	data.RLock()
	defer data.RUnlock()
	return data.len
}

//Append (safe)
func (data *RoundRobinPeerData) Append(peer *RoundRobinPeer){
	data.Lock()
	defer data.Unlock()
	if data.peers == nil {
		data.peers = make([]*RoundRobinPeer,0)
	}
	data.peers = append(data.peers,peer)
	data.len += 1
}

//GetIndexPeer (safe)
func (data *RoundRobinPeerData) GetIndex(index uint) *RoundRobinPeer {
	data.RLock()
	defer data.RUnlock()
	return data.peers[index]
}

//UpdateIndex (safe)
func (data *RoundRobinPeerData) UpdateIndex(index uint,peer *RoundRobinPeer){
	data.Lock()
	defer data.Unlock()
	// 修改指针地址
	data.peers[index] = peer
}

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

//GetPeer
func GetPeer(data *RoundRobinPeerData) *RoundRobinPeer{
	var best *RoundRobinPeer
	//当前时间毫秒
	now := GetMillisecond()
	//权重总值
	var total int32 = 0

	var i uint


	//遍历peer列表
	for i = 0 ;i < data.GetLen() ;i++  {

		// 获取当前peer
		peer := data.GetIndex(i)

		//检查当前后端服务器的 down 标志位，若为 true 表示不参与策略选择，则 continue 检查下一个后端服务器
		if peer.Down {
			continue
		}

		// 当前后端服务器的 down 标志位为 false,接着检查当前后端服务器连接失败的次数是否已经达到 max_fails；
		// 且睡眠的时间还没到 fail_timeout，则当前后端服务器不被选择，continue 检查下一个后端服务器；
		//if peer.MaxFails > 0 && peer.Fails >= peer.MaxFails && now - peer.Checked <= peer.FailTimeout {
		//	continue
		//}

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