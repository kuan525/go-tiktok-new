package db

func GetFriendIdList(userId int64) ([]int64, error) {
	if !rdFollows.CheckFollow(userId) {
		following, err := getFollowIdList(userId)
		if err != nil {
			return *new([]int64), err
		}
		addFollowRelationToRedis(userId, following)
	}
	if !rdFollows.CheckFollower(userId) {
		followers, err := getFollowerIdList(userId)
		if err != nil {
			return *new([]int64), err
		}
		addFollowerRelationToRedis(userId, followers)
	}
	return rdFollows.GetFriend(userId), nil
}
