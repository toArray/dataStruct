package dataStruct

//NodeDouble 双向循环链表节点
type NodeDouble struct {
	Data interface{} //数据域
	Next *NodeTwoWay //下个指针域
	Prev *NodeTwoWay //上个指针域
}

/*
CreateNodeDouble
@Desc	创建新节点
@Param	data interface{}	节点数据
*/
func CreateNodeDouble(data interface{}) *NodeDouble {
	return &NodeDouble{
		Data: data,
		Next: nil,
		Prev: nil,
	}
}
