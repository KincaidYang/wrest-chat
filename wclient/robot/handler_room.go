package robot

import (
	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/wcferry"
)

func roomHandler() {

	if len(args.Bot.HostedRooms) == 0 {
		return
	}

	for _, v := range args.Bot.HostedRooms {
		v := v // copy it
		cmdkey := "/room:" + v.Mask
		handlers[cmdkey] = &Handler{
			Level:    0,
			ChatAble: true,
			RoomAble: false,
			Describe: "加入群聊 " + v.Name,
			Callback: func(msg *wcferry.WxMsg) string {
				resp := wc.CmdClient.InviteChatroomMembers(v.RoomId, msg.Sender)
				if resp == 1 {
					return "已发送群邀请，稍后请点击进入"
				} else {
					return "发送群邀请失败"
				}
			},
		}
	}

}