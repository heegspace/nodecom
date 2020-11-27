package nodecom

import (
	"github.com/heegspace/heegproto/loginnode"
	"github.com/heegspace/heegproto/registernode"

	"github.com/heegspace/heegproto/searchnode"

	"github.com/heegspace/heegproto/common"

	"github.com/heegspace/heegproto/cloudnode"
	"github.com/heegspace/heegproto/codenode"
	"github.com/heegspace/heegproto/datanode"
	"github.com/heegspace/heegproto/questionnode"
	"github.com/heegspace/heegproto/s2sname"
	"github.com/heegspace/heegproto/usernode"
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
	if nil == client {
		panic("New Heegrpc client is nil")
	}

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
	if nil == client {
		panic("New Heegrpc client is nil")
	}

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
	if nil == client {
		panic("New Heegrpc client is nil")
	}

	questionNode := questionnode.NewQuestionnodeServiceClient(client.Client())

	return questionNode
}

// 获取search节点客户端
//
// @param s2sname
//
func Searchnode(s2sname string) *searchnode.SearchnodeServiceClient {
	searchnode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		panic(err)
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: searchnode_s2s.Host,
		Port: int(searchnode_s2s.Port),
	})
	if nil == client {
		panic("New Heegrpc client is nil")
	}

	searchNode := searchnode.NewSearchnodeServiceClient(client.Client())

	return searchNode
}

// 获取cloud节点客户端
//
// @param s2sname
//
func Cloudnode(s2sname string) *cloudnode.CloudnodeServiceClient {
	cloudnode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		panic(err)
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: cloudnode_s2s.Host,
		Port: int(cloudnode_s2s.Port),
	})
	if nil == client {
		panic("New Heegrpc client is nil")
	}

	cloudNode := cloudnode.NewCloudnodeServiceClient(client.Client())

	return cloudNode
}

// 获取register节点客户端
//
// @param s2sname
//
func Registernode(s2sname string) *registernode.RegisternodeServiceClient {
	registernode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		panic(err)
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: registernode_s2s.Host,
		Port: int(registernode_s2s.Port),
	})
	if nil == client {
		panic("New Heegrpc client is nil")
	}

	registerNode := registernode.NewRegisternodeServiceClient(client.Client())

	return registerNode
}

// 获取login节点客户端
//
// @param s2sname
//
func Loginnode(s2sname string) *loginnode.LoginnodeServiceClient {
	loginnode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		panic(err)
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: loginnode_s2s.Host,
		Port: int(loginnode_s2s.Port),
	})
	if nil == client {
		panic("New Heegrpc client is nil")
	}

	loginNode := loginnode.NewLoginnodeServiceClient(client.Client())

	return loginNode
}

// 获取login节点客户端
//
// @param s2sname
//
func Usernode(s2sname string) *usernode.UsernodeServiceClient {
	usernode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		panic(err)
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: usernode_s2s.Host,
		Port: int(usernode_s2s.Port),
	})
	if nil == client {
		panic("New Heegrpc client is nil")
	}

	userNode := usernode.NewUsernodeServiceClient(client.Client())

	return userNode
}

// 获取ls2s节点客户端
//
// @param s2sname
//
func S2sname(s2sname string) *s2sname.S2snameServiceClient {
	s2sname_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		panic(err)
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: s2sname_s2s.Host,
		Port: int(s2sname_s2s.Port),
	})
	if nil == client {
		panic("New Heegrpc client is nil")
	}

	s2s := s2sname.NewS2snameServiceClient(client.Client())

	return s2s
}
