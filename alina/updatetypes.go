package alina

type UpdateType string

const (
	MessageNew           UpdateType = "message_new"
	MessageReply         UpdateType = "message_reply"
	MessageEdit          UpdateType = "message_edit"
	MessageDeny          UpdateType = "message_deny"
	PhotoNew             UpdateType = "photo_new"
	PhotoCommentEdit     UpdateType = "photo_comment_edit"
	PhotoCommentRestore  UpdateType = "photo_comment_restore"
	PhotoCommentDelete   UpdateType = "photo_comment_delete"
	AudioNew             UpdateType = "audio_new"
	VideoNew             UpdateType = "video_new"
	VideoCommentNew      UpdateType = "video_comment_new"
	VideoCommentEdit     UpdateType = "video_comment_edit"
	VideoCommentRestore  UpdateType = "video_comment_restore"
	VideoCommentDelete   UpdateType = "video_comment_delete"
	WallPostNew          UpdateType = "wall_post_new"
	WallRepost           UpdateType = "wall_repost"
	WallReplyNew         UpdateType = "wall_reply_new"
	WallReplyEdit        UpdateType = "wall_reply_edit"
	WallReplyRestore     UpdateType = "wall_reply_restore"
	WallReptyDelete      UpdateType = "wall_reply_delete"
	BoardPostNew         UpdateType = "board_post_new"
	BoardPostEdit        UpdateType = "board_post_edit"
	BoardPostRestore     UpdateType = "board_post_restore"
	BoardPostDelete      UpdateType = "board_post_delete"
	MarketCommentNew     UpdateType = "market_comment_new"
	MarketCommentEdit    UpdateType = "market_comment_edit"
	MarketCommentRestore UpdateType = "market_comment_restore"
	MarketCommentDelete  UpdateType = "market_comment_delete"
	GroupLeave           UpdateType = "group_leave"
	GroupJoin            UpdateType = "group_join"
	UserBlock            UpdateType = "user_block"
	UserUnblock          UpdateType = "user_unblock"
	PollVoteNew          UpdateType = "poll_vote_new"
	GroupOfficersEdit    UpdateType = "group_officers_edit"
	GroupChangeSettings  UpdateType = "group_change_settings"
	GroupChangePhoto     UpdateType = "group_change_photo"
)
