package msg

// 消息状态
const (
	StatusUnread   = 0 // 消息未读
	StatusHaveRead = 1 // 消息已读
)

type Type int

// 消息类型
const (
	TypeTopicComment   Type = 0 // 收到话题评论
	TypeCommentReply   Type = 1 // 收到他人回复
	TypeTopicLike      Type = 2 // 收到点赞
	TypeTopicFavorite  Type = 3 // 话题被收藏
	TypeTopicRecommend Type = 4 // 话题被设为推荐
	TypeTopicDelete    Type = 5 // 话题被删除
	TypeArticleComment Type = 6 // 收到文章评论
)

type TopicLikeExtraData struct {
	TopicId    int64 `json:"topicId"`
	LikeUserId int64 `json:"likeUserId"`
}

type TopicFavoriteExtraData struct {
	TopicId        int64 `json:"topicId"`
	FavoriteUserId int64 `json:"favoriteUserId"`
}

type TopicRecommendExtraData struct {
	TopicId int64 `json:"topicId"`
}

type TopicDeleteExtraData struct {
	TopicId      int64 `json:"topicId"`
	DeleteUserId int64 `json:"deleteUserId"`
}

type CommentExtraData struct {
	EntityType string `json:"entityType"`
	EntityId   int64  `json:"entityId"`
	CommentId  int64  `json:"commentId"`
	QuoteId    int64  `json:"quoteId"`
}
