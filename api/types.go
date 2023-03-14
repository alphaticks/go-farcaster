package api

import "strings"

func IsNoAccountFoundErr(err error) bool {
	return strings.Contains(err.Error(), "No account found")
}

func IsUnableToFindErr(err error) bool {
	return strings.Contains(err.Error(), "Unable to find")
}

func IsRateLimitExceededErr(err error) bool {
	return strings.Contains(err.Error(), "Rate limit exceeded")
}

func IsNoFIDRegisteredErr(err error) bool {
	return strings.Contains(err.Error(), "No FID registered")
}

type Error struct {
	Message string `json:"message"`
}

type Cursor struct {
	Cursor string `json:"cursor"`
}

type AuthRequest struct {
	Method string `json:"method"`
	Params struct {
		ExpiresAt uint64 `json:"expiresAt,omitempty"`
		Timestamp uint64 `json:"timestamp"`
	} `json:"params"`
}

type AuthResponse struct {
	Result struct {
		Token Token `json:"token"`
	} `json:"result"`
	Errors []Error `json:"errors"`
}

type Token struct {
	Secret    string `json:"secret"`
	ExpiresAt uint64 `json:"expiresAt"`
}

type SuccessResponse struct {
	Errors []Error `json:"errors"`
}

type LikeCastResponse struct {
	Result struct {
		Reaction Reaction `json:"like"`
	} `json:"result"`
	Errors []Error `json:"errors"`
}

type LikeCastRequest struct {
	Type     string `json:"type"`
	CastFid  uint   `json:"castFid"`
	CastHash string `json:"castHash"`
}

type DeleteCastRequest struct {
	Type     string `json:"type"`
	CastFid  uint   `json:"castFid"`
	CastHash string `json:"castHash"`
}

type DeleteCastResponse struct {
	Result struct {
		Success bool `json:"success"`
	} `json:"result"`
	Errors []Error `json:"errors"`
}

type PostCastRequest struct {
	Text   string  `json:"text"`
	Parent *Parent `json:"parent,omitempty"`
}

type Parent struct {
	Hash string `json:"hash"`
	Fid  uint   `json:"fid"`
}

type PostCastResponse struct {
	Result struct {
		Cast Cast `json:"cast"`
	} `json:"result"`
	Errors []Error `json:"errors"`
}

type RecastRequest struct {
	CastHash string `json:"castHash"`
}

type RecastResponse struct {
	Result struct {
		CastHash string `json:"castHash"`
	} `json:"result"`
	Errors []Error `json:"errors"`
}

type FollowRequest struct {
	TargetFid uint `json:"targetFid"`
}

type GetFollowersResponse struct {
	Result struct {
		Users []User `json:"users"`
	} `json:"result"`
	Errors []Error `json:"errors"`
	Next   Cursor  `json:"next"`
}

type GetFollowingResponse struct {
	Result struct {
		Users []User `json:"users"`
	} `json:"result"`
	Errors []Error `json:"errors"`
	Next   Cursor  `json:"next"`
}

type GetUserResponse struct {
	Result struct {
		User User `json:"user"`
	}
	Errors []Error `json:"errors"`
}

type GetUsersResponse struct {
	Result struct {
		Users []User `json:"users"`
	} `json:"result"`
	Errors []Error `json:"errors"`
	Next   Cursor  `json:"next"`
}

type GetCastResponse struct {
	Result struct {
		Cast Cast `json:"cast"`
	}
	Errors []Error `json:"errors"`
}

type GetUserCollectionsResponse struct {
	Result struct {
		Collections []Collection `json:"collections"`
	}
	Errors []Error `json:"errors"`
	Next   Cursor  `json:"next"`
}

type GetVerificationsResponse struct {
	Result struct {
		Verifications []Verification `json:"verifications"`
	} `json:"result"`
	Errors []Error `json:"errors"`
}

type GetNotificationsResponse struct {
	Result struct {
		Notifications []Notification `json:"notifications"`
	} `json:"result"`
	Errors []Error `json:"errors"`
	Next   Cursor  `json:"next"`
}

type GetCastsResponse struct {
	Result struct {
		Casts []Cast `json:"casts"`
	} `json:"result"`
	Errors []Error `json:"errors"`
	Next   Cursor  `json:"next"`
}

type GetRecastersResponse struct {
	Result struct {
		Users []User `json:"users"`
	} `json:"result"`
	Errors []Error `json:"errors"`
	Next   Cursor  `json:"next"`
}

type GetReactionsResponse struct {
	Result struct {
		Reactions []Reaction `json:"reactions"`
	} `json:"result"`
	Errors []Error `json:"errors"`
	Next   Cursor  `json:"next"`
}

type GetLikesResponse struct {
	Result struct {
		Likes []Reaction `json:"likes"`
	} `json:"result"`
	Errors []Error `json:"errors"`
	Next   Cursor  `json:"next"`
}

type User struct {
	Fid         uint   `json:"fid"`
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	Pfp         struct {
		Url      string `json:"url"`
		Verified bool   `json:"verified"`
	} `json:"pfp"`
	Profile struct {
		Bio struct {
			Text     string        `json:"text"`
			Mentions []interface{} `json:"mentions"`
		} `json:"bio"`
	} `json:"profile"`
	FollowerCount  uint `json:"followerCount"`
	FollowingCount uint `json:"followingCount"`
	ViewerContext  struct {
		Following  bool `json:"following"`
		FollowedBy bool `json:"followedBy"`
	} `json:"viewerContext"`
}

type Reaction struct {
	Type    string `json:"type"`
	Hash    string `json:"hash"`
	Reactor struct {
		Fid         uint   `json:"fid"`
		Username    string `json:"username"`
		DisplayName string `json:"displayName"`
		Pfp         struct {
			Url      string `json:"url"`
			Verified bool   `json:"verified"`
		} `json:"pfp"`
		FollowerCount    uint   `json:"followerCount"`
		FollowingCount   uint   `json:"followingCount"`
		ReferrerUsername string `json:"referrerUsername"`
		ViewerContext    struct {
			Following  bool `json:"following"`
			FollowedBy bool `json:"followedBy"`
		} `json:"viewerContext"`
	} `json:"reactor"`
	Timestamp int64  `json:"timestamp"`
	CastHash  string `json:"castHash"`
}

type Cast struct {
	Hash         string `json:"hash"`
	ThreadHash   string `json:"threadHash"`
	ParentHash   string `json:"parentHash"`
	ParentAuthor struct {
		Fid         int    `json:"fid"`
		Username    string `json:"username"`
		DisplayName string `json:"displayName"`
		Pfp         struct {
			Url      string `json:"url"`
			Verified bool   `json:"verified"`
		} `json:"pfp"`
		Profile struct {
			Bio struct {
				Text     string        `json:"text"`
				Mentions []interface{} `json:"mentions"`
			} `json:"bio"`
		} `json:"profile"`
		FollowerCount  uint `json:"followerCount"`
		FollowingCount uint `json:"followingCount"`
	} `json:"parentAuthor"`
	Author struct {
		Fid         uint   `json:"fid"`
		Username    string `json:"username"`
		DisplayName string `json:"displayName"`
		Pfp         struct {
			Url      string `json:"url"`
			Verified bool   `json:"verified"`
		} `json:"pfp"`
		Profile struct {
			Bio struct {
				Text     string        `json:"text"`
				Mentions []interface{} `json:"mentions"`
			} `json:"bio"`
		} `json:"profile"`
		FollowerCount  uint `json:"followerCount"`
		FollowingCount uint `json:"followingCount"`
	} `json:"author"`
	Text      string `json:"text"`
	Timestamp int64  `json:"timestamp"`
	Replies   struct {
		Count uint `json:"count"`
	} `json:"replies"`
	Reactions struct {
		Count uint `json:"count"`
	} `json:"reactions"`
	Recasts struct {
		Count     uint          `json:"count"`
		Recasters []interface{} `json:"recasters"`
	} `json:"recasts"`
	Watches struct {
		Count uint `json:"count"`
	} `json:"watches"`
	ViewerContext struct {
		Reacted bool `json:"reacted"`
		Recast  bool `json:"recast"`
		Watched bool `json:"watched"`
	} `json:"viewerContext"`
}

type Verification struct {
	Fid       uint   `json:"fid"`
	Address   string `json:"address"`
	Timestamp int64  `json:"timestamp"`
}

type Notification struct {
	Type      string `json:"type"`
	Id        string `json:"id"`
	Timestamp int64  `json:"timestamp"`
	Actor     User   `json:"actor"`
	Content   struct {
		Cast Cast `json:"cast"`
	} `json:"content"`
}

type Collection struct {
	Id                  string `json:"id"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	ItemCount           uint   `json:"itemCount"`
	OwnerCount          uint   `json:"ownerCount"`
	FarcasterOwnerCount uint   `json:"farcasterOwnerCount"`
	ImageUrl            string `json:"imageUrl"`
	VolumeTraded        string `json:"volumeTraded"`
	ExternalUrl         string `json:"externalUrl"`
	OpenSeaUrl          string `json:"openSeaUrl"`
	TwitterUsername     string `json:"twitterUsername"`
	SchemaName          string `json:"schemaName"`
}
