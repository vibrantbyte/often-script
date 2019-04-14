/*
@Time : 2019-04-12 12:02 
@Author : xiaoyueya
@File : hight_quality
@Software: GoLand
*/
package main

/**
根据现有的结构组合成需要递归的结构，封装成处理策略
 */

type Strategy interface {

	Add(strategy Strategy)
	//获取下一个策略
	Next() Strategy
	//获取大小
	Size() uint

	//Set
	Set(first Strategy)
}

type DefaultStrategy struct {
	//策略长度
	Len uint
	//是否是头指针
	IsHead bool

	//策略
	next Strategy
}

func NewDefaultStrategy() Strategy {
	var strategy Strategy
	def := new(DefaultStrategy)
	def.Len += 1
	def.IsHead = true
	def.Set(nil)
	strategy = def
	return strategy
}

func (def *DefaultStrategy) Next() Strategy{
	return def.next
}

func (def *DefaultStrategy) Add(strategy Strategy){
	strategy.Add(def.next)
	strategy.Set(def)
	def.next = strategy
	def.Len += 1
}

func (def *DefaultStrategy) Size() uint {
	return def.Len - 1
}

func (def *DefaultStrategy) Set(first Strategy)  {

}


type RightsStrategy struct {
	//策略
	first Strategy
	next Strategy
}

//Next
func (s *RightsStrategy) Next() Strategy{
	return s.next
}

//add
func (s *RightsStrategy) Add(strategy Strategy){
	s.next = strategy
}

//Size
func (s *RightsStrategy) Size() uint{
	return s.first.Size()
}

//Set
func (s *RightsStrategy) Set(first Strategy)  {
	s.first = first
}



//StrategyChain
type StrategyChain struct {

}

func main()  {

	var list = NewDefaultStrategy()

	partner := new(RightsStrategy)
	list.Add(partner)

	inspector := new(RightsStrategy)
	list.Add(inspector)

	inspector1 := new(RightsStrategy)
	list.Add(inspector1)

	inspector2 := new(RightsStrategy)
	list.Add(inspector2)

	for {
		if node := list.Next();node != nil{
			if item,ok := node.(*RightsStrategy);ok{
				println(item.Size())
			}

			list = node
		}else{
			break
		}
	}





}

func PanicTest(){
	panic("2222")
}


//PartnerInitial
//func PartnerInitial(node ){

	//data := new(RoundRobinPeerData)
	//
	//peer1 := new(RoundRobinPeer)
	//peer1.Name = "a"
	//peer1.Serial = 1
	//peer1.Weight = 5
	//peer1.EffectiveWeight = 5
	//peer1.Down = false
	//data.Append(peer1)
	//
	//
	//peer2 := new(RoundRobinPeer)
	//peer2.Name = "b"
	//peer2.Serial = 2
	//peer2.Weight = 3
	//peer2.EffectiveWeight = 3
	//peer2.Down = false
	//data.Append(peer2)
	//
	//
	//peer3 := new(RoundRobinPeer)
	//peer3.Name = "c"
	//peer3.Serial = 3
	//peer3.Weight = 1
	//peer3.EffectiveWeight = 1
	//peer3.Down = false
	//data.Append(peer3)
	//
	//peer4 := new(RoundRobinPeer)
	//peer4.Name = "d"
	//peer4.Serial = 4
	//peer4.Weight = 0
	//peer4.EffectiveWeight = 0
	//peer4.Down = false
	//data.Append(peer4)

//}

