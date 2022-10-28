package api

type V1Response struct {
	Result interface{} `json:"result"`
	Error  string      `json:"error"`
}

type GetUsersResponse struct {
}

type GetCastsResponse struct {
	Result struct {
		Casts []Cast `json:"casts"`
	} `json:"result"`
}

type GetProfileResponse struct {
	Result struct {
		Profile Profile `json:"user"`
	} `json:"result"`
}

type Profile struct {
	Address     string `json:"address"`
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	Avatar      struct {
		Url        string `json:"url"`
		IsVerified bool   `json:"isVerified"`
	} `json:"avatar"`
	FollowerCount     int  `json:"followerCount"`
	FollowingCount    int  `json:"followingCount"`
	IsViewerFollowing bool `json:"isViewerFollowing"`
	IsFollowingViewer bool `json:"isFollowingViewer"`
	Profile           struct {
		Bio struct {
			Text     string        `json:"text"`
			Mentions []interface{} `json:"mentions"`
		} `json:"bio"`
		DirectMessageTargets struct {
			Telegram string `json:"telegram"`
		} `json:"directMessageTargets"`
	} `json:"profile"`
	ViewerCanSendDirectCasts bool `json:"viewerCanSendDirectCasts"`
}

type Cast struct {
	Body struct {
		Type        string `json:"type"`
		PublishedAt int64  `json:"publishedAt"`
		Sequence    int    `json:"sequence"`
		Address     string `json:"address"`
		Username    string `json:"username"`
		Data        struct {
			Text                  string      `json:"text"`
			ReplyParentMerkleRoot interface{} `json:"replyParentMerkleRoot"`
		} `json:"data"`
		PrevMerkleRoot string `json:"prevMerkleRoot"`
	} `json:"body"`
	Attachments struct {
		OpenGraph []interface{} `json:"openGraph"`
	} `json:"attachments"`
	Signature        string `json:"signature"`
	MerkleRoot       string `json:"merkleRoot"`
	ThreadMerkleRoot string `json:"threadMerkleRoot"`
	Meta             struct {
		DisplayName      string      `json:"displayName"`
		Avatar           string      `json:"avatar"`
		IsVerifiedAvatar bool        `json:"isVerifiedAvatar"`
		Mentions         interface{} `json:"mentions"`
		NumReplyChildren int         `json:"numReplyChildren"`
		Reactions        struct {
			Count int    `json:"count"`
			Type  string `json:"type"`
			Self  bool   `json:"self"`
		} `json:"reactions"`
		Recasters []interface{} `json:"recasters"`
		Recasts   struct {
			Count int  `json:"count"`
			Self  bool `json:"self"`
		} `json:"recasts"`
		Watches struct {
			Count int  `json:"count"`
			Self  bool `json:"self"`
		} `json:"watches"`
	} `json:"meta"`
}
