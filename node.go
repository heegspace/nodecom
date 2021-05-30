package nodecom

import (
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
	"github.com/heegspace/heegproto/sensinode"
	"github.com/heegspace/heegproto/teachnode"
	"github.com/heegspace/heegproto/usernode"
	"github.com/asim/go-micro/v3/client"
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
func Datanode(s2sname string, cli client.Client) datanode.DatanodeService {
	reqcli := datanode.NewDatanodeService(s2sname, cli)

	return reqcli
}

// 获取有数据节点客户端
//
// @param s2sname
//
func Codenode(s2sname string, cli client.Client) codenode.CodenodeService {
	reqcli := codenode.NewCodenodeService(s2sname, cli)

	return reqcli
}

// 获取question节点客户端
//
// @param s2sname
//
func Questionnode(s2sname string, cli client.Client) questionnode.QuestionnodeService {
	reqcli := questionnode.NewQuestionnodeService(s2sname, cli)

	return reqcli
}

// 获取search节点客户端
//
// @param s2sname
//
func Searchnode(s2sname string, cli client.Client) searchnode.SearchnodeService {
	reqcli := searchnode.NewSearchnodeService(s2sname, cli)

	return reqcli
}

// 获取cloud节点客户端
//
// @param s2sname
//
func Cloudnode(s2sname string, cli client.Client) cloudnode.CloudnodeService {
	reqcli := cloudnode.NewCloudnodeService(s2sname, cli)

	return reqcli
}

// 获取register节点客户端
//
// @param s2sname
//
func Registernode(s2sname string, cli client.Client) registernode.RegisternodeService {
	reqcli := registernode.NewRegisternodeService(s2sname, cli)

	return reqcli
}

// 获取login节点客户端
//
// @param s2sname
//
func Loginnode(s2sname string, cli client.Client) loginnode.LoginnodeService {
	reqcli := loginnode.NewLoginnodeService(s2sname, cli)

	return reqcli
}

// 获取login节点客户端
//
// @param s2sname
//
func Usernode(s2sname string, cli client.Client) usernode.UsernodeService {
	reqcli := usernode.NewUsernodeService(s2sname, cli)

	return reqcli
}

// 获取login节点客户端
//
// @param s2sname
//
func Authnode(s2sname string, cli client.Client) authnode.AuthnodeService {
	reqcli := authnode.NewAuthnodeService(s2sname, cli)

	return reqcli
}

// 获取note节点客户端
//
// @param s2sname
//
func Notenode(s2sname string, cli client.Client) notenode.NotenodeService {
	reqcli := notenode.NewNotenodeService(s2sname, cli)

	return reqcli
}

// 获取friend节点客户端
//
// @param s2sname
//
func Friendnode(s2sname string, cli client.Client) friendnode.FriendnodeService {
	reqcli := friendnode.NewFriendnodeService(s2sname, cli)

	return reqcli
}

// 获取darty节点客户端
//
// @param s2sname
//
func Dartynode(s2sname string, cli client.Client) dartynode.DartynodeService {
	reqcli := dartynode.NewDartynodeService(s2sname, cli)

	return reqcli
}

// 获取sensi节点客户端
//
// @param s2sname
//
func Sensinode(s2sname string, cli client.Client) sensinode.SensinodeService {
	reqcli := sensinode.NewSensinodeService(s2sname, cli)

	return reqcli
}

// 获取limit节点客户端
//
// @param s2sname
//
func Limitnode(s2sname string, cli client.Client) limitnode.LimitnodeService {
	reqcli := limitnode.NewLimitnodeService(s2sname, cli)

	return reqcli
}

// 获取teachnode节点
//
// @param s2sname
//
func Teachnode(s2sname string, cli client.Client) teachnode.TeachnodeService {
	reqcli := teachnode.NewTeachnodeService(s2sname, cli)

	return reqcli
}

// 创建cert客户端
//
func Certnode(s2sname string, cli client.Client) certnode.CertnodeService {
	reqcli := certnode.NewCertnodeService(s2sname, cli)

	return reqcli
}

// 创建macipnode客户端
//
func Macipnode(s2sname string, cli client.Client) macipnode.MacipnodeService {
	reqcli := macipnode.NewMacipnodeService(s2sname, cli)

	return reqcli
}

// 获取Lognode客户端
//
func Lognode(s2sname string, cli client.Client) lognode.LognodeService {
	reqcli := lognode.NewLognodeService(s2sname, cli)

	return reqcli
}

// 获取Cronnode客户端
//
func Cronnode(s2sname string, cli client.Client) cronnode.CronnodeService {
	reqcli := cronnode.NewCronnodeService(s2sname, cli)

	return reqcli
}
