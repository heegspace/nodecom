package nodecom

import (
	"fmt"

	"github.com/heegspace/heegproto/friendnode"
	"github.com/heegspace/heegproto/notenode"

	"github.com/heegspace/heegproto/loginnode"
	"github.com/heegspace/heegproto/registernode"

	"github.com/heegspace/heegproto/searchnode"

	"github.com/heegspace/heegproto/common"

	"github.com/heegspace/heegproto/authnode"
	"github.com/heegspace/heegproto/certnode"
	"github.com/heegspace/heegproto/cloudnode"
	"github.com/heegspace/heegproto/codenode"
	"github.com/heegspace/heegproto/cronnode"
	"github.com/heegspace/heegproto/dartynode"
	"github.com/heegspace/heegproto/datanode"
	"github.com/heegspace/heegproto/limitnode"
	"github.com/heegspace/heegproto/lognode"
	"github.com/heegspace/heegproto/macipnode"
	"github.com/heegspace/heegproto/questionnode"
	"github.com/heegspace/heegproto/s2sname"
	"github.com/heegspace/heegproto/sensinode"
	"github.com/heegspace/heegproto/teachnode"
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
func Datanode(s2sname string) (*datanode.DatanodeServiceClient, *thrift.TBufferedTransport) {
	datanode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		fmt.Println(s2sname, " node fail, 2s retry. ", err)

		return nil, nil
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: datanode_s2s.Host,
		Port: int(datanode_s2s.Port),
	})
	if nil == client {
		fmt.Println(s2sname, " client fail, 2s retry. ")

		return nil, nil
	}

	client1, trans1 := client.Client()
	if nil == client1 || nil == trans1 {
		return nil, nil
	}

	dataNode := datanode.NewDatanodeServiceClient(client1)
	return dataNode, trans1
}

// 获取有数据节点客户端
//
// @param s2sname
//
func Codenode(s2sname string) (*codenode.CodenodeServiceClient, *thrift.TBufferedTransport) {
	datanode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		fmt.Println(s2sname, " node fail, 2s retry. ", err)

		return nil, nil
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: datanode_s2s.Host,
		Port: int(datanode_s2s.Port),
	})
	if nil == client {
		fmt.Println(s2sname, " client fail, 2s retry. ")

		return nil, nil
	}

	client1, trans1 := client.Client()
	if nil == client1 || nil == trans1 {
		return nil, nil
	}

	codeNode := codenode.NewCodenodeServiceClient(client1)
	return codeNode, trans1
}

// 获取question节点客户端
//
// @param s2sname
//
func Questionnode(s2sname string) (*questionnode.QuestionnodeServiceClient, *thrift.TBufferedTransport) {
	questionnode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		fmt.Println(s2sname, " node fail, 2s retry. ", err)

		return nil, nil
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: questionnode_s2s.Host,
		Port: int(questionnode_s2s.Port),
	})
	if nil == client {
		fmt.Println(s2sname, " client fail, 2s retry. ")

		return nil, nil
	}

	client1, trans1 := client.Client()
	if nil == client1 || nil == trans1 {
		return nil, nil
	}

	questionNode := questionnode.NewQuestionnodeServiceClient(client1)
	return questionNode, trans1
}

// 获取search节点客户端
//
// @param s2sname
//
func Searchnode(s2sname string) (*searchnode.SearchnodeServiceClient, *thrift.TBufferedTransport) {
	searchnode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		fmt.Println(s2sname, " node fail, 2s retry. ", err)

		return nil, nil
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: searchnode_s2s.Host,
		Port: int(searchnode_s2s.Port),
	})
	if nil == client {
		fmt.Println(s2sname, " client fail, 2s retry. ")

		return nil, nil
	}

	client1, trans1 := client.Client()
	if nil == client1 || nil == trans1 {
		return nil, nil
	}

	searchNode := searchnode.NewSearchnodeServiceClient(client1)
	return searchNode, trans1
}

// 获取cloud节点客户端
//
// @param s2sname
//
func Cloudnode(s2sname string) (*cloudnode.CloudnodeServiceClient, *thrift.TBufferedTransport) {
	cloudnode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		fmt.Println(s2sname, " node fail, 2s retry. ", err)

		return nil, nil
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: cloudnode_s2s.Host,
		Port: int(cloudnode_s2s.Port),
	})
	if nil == client {
		fmt.Println(s2sname, " client fail, 2s retry. ")

		return nil, nil
	}

	client1, trans1 := client.Client()
	if nil == client1 || nil == trans1 {
		return nil, nil
	}

	cloudNode := cloudnode.NewCloudnodeServiceClient(client1)
	return cloudNode, trans1
}

// 获取register节点客户端
//
// @param s2sname
//
func Registernode(s2sname string) (*registernode.RegisternodeServiceClient, *thrift.TBufferedTransport) {
	registernode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		fmt.Println(s2sname, " node fail, 2s retry. ", err)

		return nil, nil
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: registernode_s2s.Host,
		Port: int(registernode_s2s.Port),
	})
	if nil == client {
		fmt.Println(s2sname, " client fail, 2s retry. ")

		return nil, nil
	}

	client1, trans1 := client.Client()
	if nil == client1 || nil == trans1 {
		return nil, nil
	}

	registerNode := registernode.NewRegisternodeServiceClient(client1)
	return registerNode, trans1
}

// 获取login节点客户端
//
// @param s2sname
//
func Loginnode(s2sname string) (*loginnode.LoginnodeServiceClient, *thrift.TBufferedTransport) {
	loginnode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		fmt.Println(s2sname, " node fail, 2s retry. ", err)

		return nil, nil
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: loginnode_s2s.Host,
		Port: int(loginnode_s2s.Port),
	})
	if nil == client {
		fmt.Println(s2sname, " client fail, 2s retry. ")

		return nil, nil
	}

	client1, trans1 := client.Client()
	if nil == client1 || nil == trans1 {
		return nil, nil
	}

	loginNode := loginnode.NewLoginnodeServiceClient(client1)
	return loginNode, trans1
}

// 获取login节点客户端
//
// @param s2sname
//
func Usernode(s2sname string) (*usernode.UsernodeServiceClient, *thrift.TBufferedTransport) {
	usernode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		fmt.Println(s2sname, " node fail, 2s retry. ", err)

		return nil, nil
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: usernode_s2s.Host,
		Port: int(usernode_s2s.Port),
	})
	if nil == client {
		fmt.Println(s2sname, " client fail, 2s retry. ")

		return nil, nil
	}

	client1, trans1 := client.Client()
	if nil == client1 || nil == trans1 {
		return nil, nil
	}

	userNode := usernode.NewUsernodeServiceClient(client1)
	return userNode, trans1
}

// 获取login节点客户端
//
// @param s2sname
//
func Authnode(s2sname string) (*authnode.AuthnodeServiceClient, *thrift.TBufferedTransport) {
	authnode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		fmt.Println(s2sname, " node fail, 2s retry. ")

		return nil, nil
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: authnode_s2s.Host,
		Port: int(authnode_s2s.Port),
	})
	if nil == client {
		fmt.Println(s2sname, " client fail, 2s retry. ")

		return nil, nil
	}

	client1, trans1 := client.Client()
	if nil == client1 || nil == trans1 {
		return nil, nil
	}

	authNode := authnode.NewAuthnodeServiceClient(client1)
	return authNode, trans1
}

// 获取note节点客户端
//
// @param s2sname
//
func Notenode(s2sname string) (*notenode.NotenodeServiceClient, *thrift.TBufferedTransport) {
	notenode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		fmt.Println(s2sname, " node fail, 2s retry. ")

		return nil, nil
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: notenode_s2s.Host,
		Port: int(notenode_s2s.Port),
	})
	if nil == client {
		fmt.Println(s2sname, " client fail, 2s retry. ")

		return nil, nil
	}

	client1, trans1 := client.Client()
	if nil == client1 || nil == trans1 {
		return nil, nil
	}

	noteNode := notenode.NewNotenodeServiceClient(client1)
	return noteNode, trans1
}

// 获取friend节点客户端
//
// @param s2sname
//
func Friendnode(s2sname string) (*friendnode.FriendnodeServiceClient, *thrift.TBufferedTransport) {
	friendnode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		fmt.Println(s2sname, " node fail, 2s retry. ")

		return nil, nil
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: friendnode_s2s.Host,
		Port: int(friendnode_s2s.Port),
	})
	if nil == client {
		fmt.Println(s2sname, " client fail, 2s retry. ")

		return nil, nil
	}

	client1, trans1 := client.Client()
	if nil == client1 || nil == trans1 {
		return nil, nil
	}

	friendNode := friendnode.NewFriendnodeServiceClient(client1)
	return friendNode, trans1
}

// 获取darty节点客户端
//
// @param s2sname
//
func Dartynode(s2sname string) (*dartynode.DartynodeServiceClient, *thrift.TBufferedTransport) {
	dartynode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		fmt.Println(s2sname, " node fail, 2s retry. ")

		return nil, nil
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: dartynode_s2s.Host,
		Port: int(dartynode_s2s.Port),
	})
	if nil == client {
		fmt.Println(s2sname, " client fail, 2s retry. ")

		return nil, nil
	}

	client1, trans1 := client.Client()
	if nil == client1 || nil == trans1 {
		return nil, nil
	}

	dartyNode := dartynode.NewDartynodeServiceClient(client1)
	return dartyNode, trans1
}

// 获取sensi节点客户端
//
// @param s2sname
//
func Sensinode(s2sname string) (*sensinode.SensinodeServiceClient, *thrift.TBufferedTransport) {
	sensinode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		fmt.Println(s2sname, " node fail, 2s retry. ")

		return nil, nil
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: sensinode_s2s.Host,
		Port: int(sensinode_s2s.Port),
	})
	if nil == client {
		fmt.Println(s2sname, " client fail, 2s retry. ")

		return nil, nil
	}

	client1, trans1 := client.Client()
	if nil == client1 || nil == trans1 {
		return nil, nil
	}

	sensiNode := sensinode.NewSensinodeServiceClient(client1)
	return sensiNode, trans1
}

// 获取limit节点客户端
//
// @param s2sname
//
func Limitnode(s2sname string) (*limitnode.LimitnodeServiceClient, *thrift.TBufferedTransport) {
	limitnode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		fmt.Println(s2sname, " node fail, 2s retry. ")

		return nil, nil
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: limitnode_s2s.Host,
		Port: int(limitnode_s2s.Port),
	})
	if nil == client {
		fmt.Println(s2sname, " client fail, 2s retry. ")

		return nil, nil
	}

	client1, trans1 := client.Client()
	if nil == client1 || nil == trans1 {
		return nil, nil
	}

	limitNode := limitnode.NewLimitnodeServiceClient(client1)
	return limitNode, trans1
}

// 获取teachnode节点
//
// @param s2sname
//
func Teachnode(s2sname string) (*teachnode.TeachnodeServiceClient, *thrift.TBufferedTransport) {
	teachnode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		fmt.Println(s2sname, " node fail, 2s retry.")

		return nil, nil
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: teachnode_s2s.Host,
		Port: int(teachnode_s2s.Port),
	})
	if nil == client {
		fmt.Println(s2sname, " client fail, 2s retry.")

		return nil, nil
	}

	client1, tran1 := client.Client()
	if nil == client1 || nil == trans1 {
		return nil, nil
	}

	teachNode := teachnode.NewTeachnodeServiceClient(client1)
	return teachNode, tran1
}

// 创建cert客户端
//
func Certnode(s2sname string) (*certnode.CertnodeServiceClient, *thrift.TBufferedTransport) {
	certnode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		fmt.Println(s2sname, " node fail, 2s retry.")

		return nil, nil
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: certnode_s2s.Host,
		Port: int(certnode_s2s.Port),
	})
	if nil == client {
		fmt.Println(s2sname, " client fail, 2s retry.")

		return nil, nil
	}

	client1, tran1 := client.Client()
	if nil == client1 || nil == trans1 {
		return nil, nil
	}

	certNode := certnode.NewCertnodeServiceClient(client1)
	return certNode, tran1
}

// 创建macipnode客户端
//
func Macipnode(s2sname string) (*macipnode.MacipnodeServiceClient, *thrift.TBufferedTransport) {
	macipnode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		fmt.Println(s2sname, " node fail, 2s retry. ")

		return nil, nil
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: macipnode_s2s.Host,
		Port: int(macipnode_s2s.Port),
	})
	if nil == client {
		fmt.Println(s2sname, " client fail, 2s retry. ")

		return nil, nil
	}

	client1, trans1 := client.Client()
	if nil == client1 || nil == trans1 {
		return nil, nil
	}

	macipNode := macipnode.NewMacipnodeServiceClient(client1)
	return macipNode, trans1
}

// 获取Lognode客户端
//
func Lognode(s2sname string) (*lognode.LognodeServiceClient, *thrift.TBufferedTransport) {
	lognode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		fmt.Println(s2sname, " node fail, 2s retry. ")

		return nil, nil
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: lognode_s2s.Host,
		Port: int(lognode_s2s.Port),
	})
	if nil == client {
		fmt.Println(s2sname, " client fail, 2s retry. ")

		return nil, nil
	}

	client1, trans1 := client.Client()
	if nil == client1 || nil == trans1 {
		return nil, nil
	}

	macipNode := lognode.NewLognodeServiceClient(client1)
	return macipNode, trans1
}

// 获取Cronnode客户端
//
func Cronnode(s2sname string) (*cronnode.CronnodeServiceClient, *thrift.TBufferedTransport) {
	cronnode_s2s, err := registry.NewRegistry().Selector(s2sname)
	if nil != err {
		fmt.Println(s2sname, " node fail, 2s retry. ")

		return nil, nil
	}

	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: cronnode_s2s.Host,
		Port: int(cronnode_s2s.Port),
	})
	if nil == client {
		fmt.Println(s2sname, " client fail, 2s retry. ")

		return nil, nil
	}

	client1, trans1 := client.Client()
	if nil == client1 || nil == trans1 {
		return nil, nil
	}

	cronNode := cronnode.NewCronnodeServiceClient(client1)
	return cronNode, trans1
}

// 获取ls2s节点客户端
//
// @param s2sname
//
func S2sname(host string, port int) (*s2sname.S2snameServiceClient, *thrift.TBufferedTransport) {
	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: host,
		Port: port,
	})
	if nil == client {
		panic("New NewHeegRpcClient  is nil")
	}

	client1, trans1 := client.Client()
	if nil == client1 || nil == trans1 {
		panic("New NewHeegRpcClient client is nil")
	}

	s2s := s2sname.NewS2snameServiceClient(client1)
	return s2s, trans1
}
