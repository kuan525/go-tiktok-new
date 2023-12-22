package relation

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"go-tiktok-new/biz/handler/relation"
)

func Register(r *server.Hertz) {
	root := r.Group("/", rootMw()...)
	{
		_douyin := root.Group("/douyin", _douyinMw()...)
		{
			_relation := _douyin.Group("/relation", _relationMw()...)
			{
				_action := _relation.Group("/action", _actionMw()...)
				_action.POST("/", append(_relationActionMw(), relation.RelationAction)...)
			}
			{
				_follow := _relation.Group("/follow", _followMw()...)
				{
					_list := _follow.Group("/list", _listMw()...)
					_list.GET("/", append(_relationFollowListMw(), relation.RelationFollowList)...)
				}
			}
			{
				_follower := _relation.Group("follower", _followerMw()...)
				{
					_list := _follower.Group("/list", _listMw()...)
					_list.GET("/", append(_relationFollowerListMw(), relation.RelationFollowerList)...)
				}
			}
			{
				_friend := _relation.Group("/firend", _friendMw()...)
				{
					_list := _friend.Group("/list", _listMw()...)
					_list.GET("/", append(_relationFriendListMw(), relation.RelationFriendList)...)
				}
			}
		}
	}
}
