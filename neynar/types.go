package neynar

import "time"

type Cast struct {
	Object string `json:"object"`
	Hash   string `json:"hash"`
	Author struct {
		Object         string `json:"object"`
		Fid            int    `json:"fid"`
		Username       string `json:"username"`
		DisplayName    string `json:"display_name"`
		PfpUrl         string `json:"pfp_url"`
		CustodyAddress string `json:"custody_address"`
		Profile        struct {
			Bio struct {
				Text string `json:"text"`
			} `json:"bio"`
		} `json:"profile"`
		FollowerCount     int      `json:"follower_count"`
		FollowingCount    int      `json:"following_count"`
		Verifications     []string `json:"verifications"`
		VerifiedAddresses struct {
			EthAddresses []string      `json:"eth_addresses"`
			SolAddresses []interface{} `json:"sol_addresses"`
		} `json:"verified_addresses"`
		VerifiedAccounts []struct {
			Platform string `json:"platform"`
			Username string `json:"username"`
		} `json:"verified_accounts"`
		PowerBadge    bool `json:"power_badge"`
		ViewerContext struct {
			Following  bool `json:"following"`
			FollowedBy bool `json:"followed_by"`
			Blocking   bool `json:"blocking"`
			BlockedBy  bool `json:"blocked_by"`
		} `json:"viewer_context"`
	} `json:"author"`
	ThreadHash    string      `json:"thread_hash"`
	ParentHash    interface{} `json:"parent_hash"`
	ParentUrl     *string     `json:"parent_url"`
	RootParentUrl *string     `json:"root_parent_url"`
	ParentAuthor  struct {
		Fid interface{} `json:"fid"`
	} `json:"parent_author"`
	Text      string        `json:"text"`
	Timestamp time.Time     `json:"timestamp"`
	Embeds    []interface{} `json:"embeds"`
	Channel   *struct {
		Object        string `json:"object"`
		Id            string `json:"id"`
		Name          string `json:"name"`
		ImageUrl      string `json:"image_url"`
		ViewerContext struct {
			Following bool `json:"following"`
		} `json:"viewer_context"`
	} `json:"channel"`
	Reactions struct {
		LikesCount   int `json:"likes_count"`
		RecastsCount int `json:"recasts_count"`
		Likes        []struct {
			Fid   int    `json:"fid"`
			Fname string `json:"fname"`
		} `json:"likes"`
		Recasts []interface{} `json:"recasts"`
	} `json:"reactions"`
	Replies struct {
		Count int `json:"count"`
	} `json:"replies"`
	MentionedProfiles []struct {
		Object         string `json:"object"`
		Fid            int    `json:"fid"`
		CustodyAddress string `json:"custody_address"`
		Username       string `json:"username"`
		DisplayName    string `json:"display_name"`
		PfpUrl         string `json:"pfp_url"`
		Profile        struct {
			Bio struct {
				Text              string        `json:"text"`
				MentionedProfiles []interface{} `json:"mentioned_profiles"`
			} `json:"bio"`
		} `json:"profile"`
		FollowerCount     int           `json:"follower_count"`
		FollowingCount    int           `json:"following_count"`
		Verifications     []interface{} `json:"verifications"`
		VerifiedAddresses struct {
			EthAddresses []interface{} `json:"eth_addresses"`
			SolAddresses []interface{} `json:"sol_addresses"`
		} `json:"verified_addresses"`
		PowerBadge bool `json:"power_badge"`
	} `json:"mentioned_profiles"`
	ViewerContext struct {
		Liked    bool `json:"liked"`
		Recasted bool `json:"recasted"`
	} `json:"viewer_context"`
	AuthorChannelContext struct {
		Role      string `json:"role"`
		Following bool   `json:"following"`
	} `json:"author_channel_context,omitempty"`
}

type Notification struct {
	Object              string    `json:"object"`
	MostRecentTimestamp time.Time `json:"most_recent_timestamp"`
	Type                string    `json:"type"`
	Seen                bool      `json:"seen"`
	Cast                *Cast     `json:"cast,omitempty"`
	Reactions           []struct {
		Object string `json:"object"`
		Cast   struct {
			Object string `json:"object"`
			Hash   string `json:"hash"`
		} `json:"cast"`
		User struct {
			Object         string `json:"object"`
			Fid            int    `json:"fid"`
			Username       string `json:"username"`
			DisplayName    string `json:"display_name"`
			PfpUrl         string `json:"pfp_url"`
			CustodyAddress string `json:"custody_address"`
			Profile        struct {
				Bio struct {
					Text string `json:"text"`
				} `json:"bio"`
			} `json:"profile"`
			FollowerCount     int      `json:"follower_count"`
			FollowingCount    int      `json:"following_count"`
			Verifications     []string `json:"verifications"`
			VerifiedAddresses struct {
				EthAddresses []string `json:"eth_addresses"`
				SolAddresses []string `json:"sol_addresses"`
			} `json:"verified_addresses"`
			VerifiedAccounts interface{} `json:"verified_accounts"`
			PowerBadge       bool        `json:"power_badge"`
			ViewerContext    struct {
				Following  bool `json:"following"`
				FollowedBy bool `json:"followed_by"`
				Blocking   bool `json:"blocking"`
				BlockedBy  bool `json:"blocked_by"`
			} `json:"viewer_context"`
		} `json:"user"`
	} `json:"reactions,omitempty"`
	Count   int `json:"count,omitempty"`
	Follows []struct {
		Object string `json:"object"`
		User   struct {
			Object         string `json:"object"`
			Fid            int    `json:"fid"`
			Username       string `json:"username"`
			DisplayName    string `json:"display_name"`
			PfpUrl         string `json:"pfp_url"`
			CustodyAddress string `json:"custody_address"`
			Profile        struct {
				Bio struct {
					Text string `json:"text"`
				} `json:"bio"`
				Location struct {
					Latitude  float64 `json:"latitude"`
					Longitude float64 `json:"longitude"`
					Address   struct {
						City        string `json:"city"`
						State       string `json:"state"`
						StateCode   string `json:"state_code,omitempty"`
						Country     string `json:"country"`
						CountryCode string `json:"country_code"`
					} `json:"address"`
				} `json:"location,omitempty"`
			} `json:"profile"`
			FollowerCount     int      `json:"follower_count"`
			FollowingCount    int      `json:"following_count"`
			Verifications     []string `json:"verifications"`
			VerifiedAddresses struct {
				EthAddresses []string `json:"eth_addresses"`
				SolAddresses []string `json:"sol_addresses"`
			} `json:"verified_addresses"`
			VerifiedAccounts []struct {
				Platform string `json:"platform"`
				Username string `json:"username"`
			} `json:"verified_accounts"`
			PowerBadge    bool `json:"power_badge"`
			ViewerContext struct {
				Following  bool `json:"following"`
				FollowedBy bool `json:"followed_by"`
				Blocking   bool `json:"blocking"`
				BlockedBy  bool `json:"blocked_by"`
			} `json:"viewer_context"`
		} `json:"user"`
	} `json:"follows,omitempty"`
}

type GetNotificationsResponse struct {
	UnseenNotificationsCount int            `json:"unseen_notifications_count"`
	Notifications            []Notification `json:"notifications"`
	Next                     struct {
		Cursor interface{} `json:"cursor"`
	} `json:"next"`
}
