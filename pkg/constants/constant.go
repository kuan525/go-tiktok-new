package constants

// 连接配置
const (
	MySQLDefaultDSN = "douyin:douyin123@tcp(127.0.0.1:18000)/douyin?charset=utf8&parseTime=True&loc=Local"

	MinioEndPoint        = "localhost:18001"
	MinioAccessID        = "douyin"
	MinioSecretAccessKey = "douyin123"
	MiniUseSSL           = false

	RedisAddr     = "localhost:18003"
	RedisPassword = "123456"

	MinioReverseProxyHost = "http://localhost:18001"
	HertzMainHost         = "0.0.0.0:18005"
)

// 项目常量
const (
	UserTableName      = "users"
	FollowsTableName   = "follows"
	VideosTableName    = "videos"
	MessageTableName   = "messages"
	FavoritesTableName = "likes"
	CommentTableName   = "comments"

	VideoFeedCount       = 30
	FavoriteActionType   = 1
	UnFavoriteActionType = 2

	MinioVideoBucketName = "videobucket"
	MinioImgBucketName   = "imagebucket"

	TestSign       = "测试账号！tiktok"
	TestAva        = "avatar/test1.jpg"
	TestBackground = "background/test1.png"
)
