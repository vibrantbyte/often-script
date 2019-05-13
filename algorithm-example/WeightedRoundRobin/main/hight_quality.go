/*
@Time : 2019-04-12 12:02 
@Author : xiaoyueya
@File : hight_quality
@Software: GoLand
*/
package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
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





/**
根据现有的结构组合成需要递归的结构，封装成处理策略
 */

type Strategy interface {

	Add(strategy Strategy)
	//获取下一个策略
	Next() Strategy
	//获取大小
	Size() uint
}

//头结点
type DefaultStrategy struct {
	//策略长度
	Len uint
	//是否是头指针
	IsHead bool

	//策略
	next Strategy
	last Strategy
}

func NewDefaultStrategy() Strategy {
	var strategy Strategy
	def := new(DefaultStrategy)
	def.Len = 0
	def.IsHead = true
	strategy = def
	return strategy
}

func (def *DefaultStrategy) Next() Strategy{
	return def.next
}

func (def *DefaultStrategy) Add(strategy Strategy){
	if def.last == nil {
		def.next = strategy
	}else {
		def.last.Add(strategy)
	}
	def.last = strategy
	def.Len += 1
}

func (def *DefaultStrategy) Size() uint {
	return def.Len
}

//结点
type StrategyImplement struct {
	//previous
	previous Strategy
	//next node
	next Strategy
	//首节点
	first Strategy
}

//Next
func (s *StrategyImplement) Next() Strategy{
	return s.next
}

//Add
func (s *StrategyImplement) Add(strategy Strategy){
	s.next = strategy
}

//empty func(root effective) 请不要使用
func (s *StrategyImplement) Size() uint{
	return s.first.Size()
}

//权益策略
type RightsStrategy struct {
	StrategyImplement
	//Name
	Name string
	//RoundRobinPeerData
	Data *RoundRobinPeerData
}

func main()  {
	var list = NewDefaultStrategy()

	partner := new(RightsStrategy)
	partner.Name = "partner"
	PartnerInitial(partner)
	list.Add(partner)
	p := make(map[string]*RightsStrategy)

	for i := 0;i< 15 ;i++  {
		SerialFlow(p,list)
	}

}


//可忽视的执行
func (r *RightsStrategy) NeglectExecute(f func()){
	f()
}

//串行
func (r *RightsStrategy) SerialExecute(f func()){
	f()
}

//子执行
func (r *RightsStrategy) SubExecute(f func()) {
	f()
}


func SerialFlow(p map[string]*RightsStrategy,root Strategy){

	for {
		if node := root.Next();node != nil{
			if item,ok := node.(*RightsStrategy);ok{
				switch item.Name {
					case "partner":{
						item.SerialExecute(func() {
							//执行策略
							peer :=GetPeer(item.Data)
							fmt.Printf("合伙人：%s \n",peer.Name)
							//执行下一个函数，并且当前执行结果到下一个执行器上
							if p[peer.Name] == nil {
								inspector := new(RightsStrategy)
								inspector.Name = "inspector"
								InspectorInitial(inspector,peer.Name,int(peer.EffectiveWeight))
								p[peer.Name] = inspector
							}
							fmt.Println(GetPeer(p[peer.Name].Data).Name)
						})
					}
					case "inspector":{
						item.SubExecute(func() {

						})

					}
					default:
						break
				}
			}

			root = node
		}else{
			break
		}
	}
}

//PartnerInitial
func PartnerInitial(node *RightsStrategy){

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

	node.Data = data
}


func InspectorInitial(node *RightsStrategy,name string,count int){
	data := new(RoundRobinPeerData)

	for i:=0;i<count ;i++  {
		peer1 := new(RoundRobinPeer)
		peer1.Name = fmt.Sprintf("%s-%s-%d",node.Name,name,i)
		peer1.Serial = 1
		peer1.Weight = 1
		peer1.EffectiveWeight = 1
		peer1.Down = false
		data.Append(peer1)
	}

	node.Data = data
}















//生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}
