package redis

const (
	// KeyPostCreateTimeZSet
	// type: zset;
	// score: timestamp;
	// value: post id;
	KeyPostCreateTimeZSet = "post:create_time"

	// KeyPostScoreZSet
	// type: zset;
	// score: vote score;
	// value: post id;
	KeyPostScoreZSet = "post:score"

	// KeyVoteHash
	// type: hash;
	// param: postID;
	// field: userID;
	// value: vote direction;
	// oneOf: -1 0 1;
	KeyVoteHash = "vote:post-"

	// KeyCommunitySet
	// type: set;
	// value: community id;
	KeyCommunitySet = "community"
)
