package nodecom

import (
	"github.com/heegspace/heegrpc"
	"github.com/heegspace/heegrpc/registry"
	"github.com/heegspace/heegrpc/rpc"
)

// 获取有数据节点客户端
//
// @param regi
// @param node
//
func DatanodeClient(regi *registry.Registry, s2sname string) *rpc.HeegClient {
	datanode_s2s, err := regi.Selector(s2sname)
	if nil != err {
		panic(err)
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: datanode_s2s.Host,
		Port: int(datanode_s2s.Port),
	})

	return client
}
