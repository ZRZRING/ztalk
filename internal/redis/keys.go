package redis

const (
	// KeyPostCreateTimeZSet
	// type: zset;
	// score: timestamp;
	// value: post id;
	KeyPostCreateTimeZSet = "ztalk:post:create_time"
	// KeyPostScoreZSet
	// type: zset;
	// score: vote score;
	// value: post id;
	KeyPostScoreZSet = "ztalk:post:score"
	// KeyVoteHashPrefix
	// type: hash;
	// binding: KeyVoteHash
	KeyVoteHashPrefix = "ztalk:vote:post"
	// KeyCommunitySet
	// type: set;
	// value: community id;
	KeyCommunitySet = "ztalk:community"
)

func KeyVoteHash(postID string) string {
	return KeyVoteHashPrefix + ":" + postID
}
