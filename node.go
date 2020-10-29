package nodecom

import (
	"github.com/heegspace/heegproto/codenode"
	"github.com/heegspace/heegproto/datanode"
	"github.com/heegspace/heegrpc"
	"github.com/heegspace/heegrpc/registry"
	"github.com/heegspace/heegrpc/rpc"
)

// 获取有数据节点客户端
//
// @param regi
// @param node
//
func Datanode(s2sname string) *datanode.DatanodeServiceClient {
	datanode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		panic(err)
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: datanode_s2s.Host,
		Port: int(datanode_s2s.Port),
	})

	dataNode := datanode.NewDatanodeServiceClient(client.Client())

	return dataNode
}

// 获取有数据节点客户端
//
// @param regi
// @param node
//
func Codenode(s2sname string) *codenode.CodenodeServiceClient {
	datanode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		panic(err)
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: datanode_s2s.Host,
		Port: int(datanode_s2s.Port),
	})

	codeNode := codenode.NewCodenodeServiceClient(client.Client())

	return codeNode
}
