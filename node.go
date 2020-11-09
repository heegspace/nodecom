package nodecom

import (
	"github.com/heegspace/heegproto/common"

	"github.com/heegspace/heegproto/codenode"
	"github.com/heegspace/heegproto/datanode"
	"github.com/heegspace/heegproto/questionnode"
	"github.com/heegspace/heegrpc"
	"github.com/heegspace/heegrpc/registry"
	"github.com/heegspace/heegrpc/rpc"
)

// 检查是否有权限
//
// @param auth 	授权参数
// @return bool
//
func Authorize(auth *common.Authorize) bool {
	if nil == auth {
		return false
	}

	return true
}

// 获取有数据节点客户端
//
// @param s2sname
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
// @param s2sname
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

// 获取question节点客户端
//
// @param s2sname
//
func Questionnode(s2sname string) *questionnode.QuestionnodeServiceClient {
	questionnode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		panic(err)
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: questionnode_s2s.Host,
		Port: int(questionnode_s2s.Port),
	})

	questionNode := questionnode.NewQuestionnodeServiceClient(client.Client())

	return questionNode
}
