package nodecom

import (
	"fmt"
	"time"

	"github.com/heegspace/heegproto/loginnode"
	"github.com/heegspace/heegproto/registernode"

	"github.com/heegspace/heegproto/searchnode"

	"github.com/heegspace/heegproto/common"

	"github.com/heegspace/heegproto/authnode"
	"github.com/heegspace/heegproto/cloudnode"
	"github.com/heegspace/heegproto/codenode"
	"github.com/heegspace/heegproto/datanode"
	"github.com/heegspace/heegproto/questionnode"
	"github.com/heegspace/heegproto/s2sname"
	"github.com/heegspace/heegproto/usernode"
	"github.com/heegspace/heegrpc"
	"github.com/heegspace/heegrpc/registry"
	"github.com/heegspace/heegrpc/rpc"
	"github.com/heegspace/thrift"
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
func Datanode(s2sname string) *datanode.DatanodeServiceClient,*thrift.TBufferedTransport {
retry:
	datanode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		fmt.Println(s2sname, " node fail, 2s retry. ", err)

		time.Sleep(2 * time.Second)
		goto retry
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: datanode_s2s.Host,
		Port: int(datanode_s2s.Port),
	})
	if nil == client {
		fmt.Println(s2sname, " client fail, 2s retry. ")

		time.Sleep(2 * time.Second)
		goto retry
	}

	client,trans := client.Client()
	dataNode := datanode.NewDatanodeServiceClient(client)

	return dataNode, trans
}

// 获取有数据节点客户端
//
// @param s2sname
//
func Codenode(s2sname string) *codenode.CodenodeServiceClient,*thrift.TBufferedTransport {
retry:
	datanode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		fmt.Println(s2sname, " node fail, 2s retry. ", err)

		time.Sleep(2 * time.Second)
		goto retry
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: datanode_s2s.Host,
		Port: int(datanode_s2s.Port),
	})
	if nil == client {
		fmt.Println(s2sname, " client fail, 2s retry. ")

		time.Sleep(2 * time.Second)
		goto retry
	}

	client,trans := client.Client()
	codeNode := codenode.NewCodenodeServiceClient(client)

	return codeNode,trans
}

// 获取question节点客户端
//
// @param s2sname
//
func Questionnode(s2sname string) *questionnode.QuestionnodeServiceClient,*thrift.TBufferedTransport {
retry:
	questionnode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		fmt.Println(s2sname, " node fail, 2s retry. ", err)

		time.Sleep(2 * time.Second)
		goto retry
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: questionnode_s2s.Host,
		Port: int(questionnode_s2s.Port),
	})
	if nil == client {
		fmt.Println(s2sname, " client fail, 2s retry. ")

		time.Sleep(2 * time.Second)
		goto retry
	}

	client,trans := client.Client()
	questionNode := questionnode.NewQuestionnodeServiceClient(client)

	return questionNode,trans
}

// 获取search节点客户端
//
// @param s2sname
//
func Searchnode(s2sname string) *searchnode.SearchnodeServiceClient,*thrift.TBufferedTransport {
retry:
	searchnode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		fmt.Println(s2sname, " node fail, 2s retry. ", err)

		time.Sleep(2 * time.Second)
		goto retry
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: searchnode_s2s.Host,
		Port: int(searchnode_s2s.Port),
	})
	if nil == client {
		fmt.Println(s2sname, " client fail, 2s retry. ")

		time.Sleep(2 * time.Second)
		goto retry
	}

	client,trans := client.Client()
	searchNode := searchnode.NewSearchnodeServiceClient(client)

	return searchNode,trans
}

// 获取cloud节点客户端
//
// @param s2sname
//
func Cloudnode(s2sname string) *cloudnode.CloudnodeServiceClient,*thrift.TBufferedTransport {
retry:
	cloudnode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		fmt.Println(s2sname, " node fail, 2s retry. ", err)

		time.Sleep(2 * time.Second)
		goto retry
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: cloudnode_s2s.Host,
		Port: int(cloudnode_s2s.Port),
	})
	if nil == client {
		fmt.Println(s2sname, " client fail, 2s retry. ")

		time.Sleep(2 * time.Second)
		goto retry
	}

	client,trans := client.Client()
	cloudNode := cloudnode.NewCloudnodeServiceClient(client)

	return cloudNode,trans
}

// 获取register节点客户端
//
// @param s2sname
//
func Registernode(s2sname string) *registernode.RegisternodeServiceClient,*thrift.TBufferedTransport {
retry:
	registernode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		fmt.Println(s2sname, " node fail, 2s retry. ", err)

		time.Sleep(2 * time.Second)
		goto retry
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: registernode_s2s.Host,
		Port: int(registernode_s2s.Port),
	})
	if nil == client {
		fmt.Println(s2sname, " client fail, 2s retry. ")

		time.Sleep(2 * time.Second)
		goto retry
	}

	client,trans := client.Client()
	registerNode := registernode.NewRegisternodeServiceClient(client)

	return registerNode,trans
}

// 获取login节点客户端
//
// @param s2sname
//
func Loginnode(s2sname string) *loginnode.LoginnodeServiceClient,*thrift.TBufferedTransport {
retry:
	loginnode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		fmt.Println(s2sname, " node fail, 2s retry. ", err)

		time.Sleep(2 * time.Second)
		goto retry
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: loginnode_s2s.Host,
		Port: int(loginnode_s2s.Port),
	})
	if nil == client {
		fmt.Println(s2sname, " client fail, 2s retry. ")

		time.Sleep(2 * time.Second)
		goto retry
	}

	client,trans := client.Client()
	loginNode := loginnode.NewLoginnodeServiceClient(client)

	return loginNode,trans
}

// 获取login节点客户端
//
// @param s2sname
//
func Usernode(s2sname string) *usernode.UsernodeServiceClient,*thrift.TBufferedTransport {
retry:
	usernode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		fmt.Println(s2sname, " node fail, 2s retry. ", err)

		time.Sleep(2 * time.Second)
		goto retry
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: usernode_s2s.Host,
		Port: int(usernode_s2s.Port),
	})
	if nil == client {
		fmt.Println(s2sname, " client fail, 2s retry. ")

		time.Sleep(2 * time.Second)
		goto retry
	}

	client,trans := client.Client()
	userNode := usernode.NewUsernodeServiceClient(client)

	return userNode,trans
}

// 获取login节点客户端
//
// @param s2sname
//
func Authnode(s2sname string) *authnode.AuthnodeServiceClient,*thrift.TBufferedTransport {
retry:
	authnode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		fmt.Println(s2sname, " node fail, 2s retry. ")

		time.Sleep(2 * time.Second)
		goto retry
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: authnode_s2s.Host,
		Port: int(authnode_s2s.Port),
	})
	if nil == client {
		fmt.Println(s2sname, " client fail, 2s retry. ")

		time.Sleep(2 * time.Second)
		goto retry
	}

	client,trans := client.Client()
	authNode := authnode.NewAuthnodeServiceClient(client)

	return authNode,trans
}

// 获取ls2s节点客户端
//
// @param s2sname
//
func S2sname(host string, port int) *s2sname.S2snameServiceClient,*thrift.TBufferedTransport {
	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: host,
		Port: port,
	})
	if nil == client {
		panic("New Heegrpc client is nil")
	}

	client,trans := client.Client()
	s2s := s2sname.NewS2snameServiceClient(client)

	return s2s,trans
}
