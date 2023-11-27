package resources

import (
	_ "embed"
	"encoding/json"
)

//go:embed reddit.json
var ResourceRedditJSON []byte

type ResourceReddit struct {
	Kind string `json:"kind"`
	Data struct {
		After     string      `json:"after"`
		Dist      int         `json:"dist"`
		ModHash   string      `json:"modhash"`
		GeoFilter interface{} `json:"geo_filter"`
		Children  []struct {
			Kind string `json:"kind"`
			Data struct {
				ApprovedAtUtc     interface{} `json:"approved_at_utc"`
				Subreddit         string      `json:"subreddit"`
				Selftext          string      `json:"selftext"`
				AuthorFullname    string      `json:"author_fullname"`
				Saved             bool        `json:"saved"`
				ModReasonTitle    interface{} `json:"mod_reason_title"`
				Gilded            int         `json:"gilded"`
				Clicked           bool        `json:"clicked"`
				Title             string      `json:"title"`
				LinkFlairRichtext []struct {
					A string `json:"a,omitempty"`
					E string `json:"e"`
					U string `json:"u,omitempty"`
					T string `json:"t,omitempty"`
				} `json:"link_flair_richtext"`
				SubredditNamePrefixed      string      `json:"subreddit_name_prefixed"`
				Hidden                     bool        `json:"hidden"`
				Pwls                       *int        `json:"pwls"`
				LinkFlairCssClass          *string     `json:"link_flair_css_class"`
				Downs                      int         `json:"downs"`
				ThumbnailHeight            *int        `json:"thumbnail_height"`
				TopAwardedType             interface{} `json:"top_awarded_type"`
				HideScore                  bool        `json:"hide_score"`
				Name                       string      `json:"name"`
				Quarantine                 bool        `json:"quarantine"`
				LinkFlairTextColor         string      `json:"link_flair_text_color"`
				UpvoteRatio                float64     `json:"upvote_ratio"`
				AuthorFlairBackgroundColor *string     `json:"author_flair_background_color"`
				SubredditType              string      `json:"subreddit_type"`
				Ups                        int         `json:"ups"`
				TotalAwardsReceived        int         `json:"total_awards_received"`
				MediaEmbed                 struct {
				} `json:"media_embed"`
				ThumbnailWidth        *int          `json:"thumbnail_width"`
				AuthorFlairTemplateId *string       `json:"author_flair_template_id"`
				IsOriginalContent     bool          `json:"is_original_content"`
				UserReports           []interface{} `json:"user_reports"`
				SecureMedia           *struct {
					RedditVideo struct {
						BitrateKbps       int    `json:"bitrate_kbps"`
						FallbackUrl       string `json:"fallback_url"`
						HasAudio          bool   `json:"has_audio"`
						Height            int    `json:"height"`
						Width             int    `json:"width"`
						ScrubberMediaUrl  string `json:"scrubber_media_url"`
						DashUrl           string `json:"dash_url"`
						Duration          int    `json:"duration"`
						HlsUrl            string `json:"hls_url"`
						IsGif             bool   `json:"is_gif"`
						TranscodingStatus string `json:"transcoding_status"`
					} `json:"reddit_video"`
				} `json:"secure_media"`
				IsRedditMediaDomain bool        `json:"is_reddit_media_domain"`
				IsMeta              bool        `json:"is_meta"`
				Category            interface{} `json:"category"`
				SecureMediaEmbed    struct {
				} `json:"secure_media_embed"`
				LinkFlairText       *string     `json:"link_flair_text"`
				CanModPost          bool        `json:"can_mod_post"`
				Score               int         `json:"score"`
				ApprovedBy          interface{} `json:"approved_by"`
				IsCreatedFromAdsUi  bool        `json:"is_created_from_ads_ui"`
				AuthorPremium       bool        `json:"author_premium"`
				Thumbnail           string      `json:"thumbnail"`
				Edited              interface{} `json:"edited"`
				AuthorFlairCssClass *string     `json:"author_flair_css_class"`
				AuthorFlairRichtext []struct {
					A string `json:"a,omitempty"`
					E string `json:"e"`
					U string `json:"u,omitempty"`
					T string `json:"t,omitempty"`
				} `json:"author_flair_richtext"`
				Gildings struct {
				} `json:"gildings"`
				PostHint            string      `json:"post_hint,omitempty"`
				ContentCategories   []string    `json:"content_categories"`
				IsSelf              bool        `json:"is_self"`
				ModNote             interface{} `json:"mod_note"`
				Created             float64     `json:"created"`
				LinkFlairType       string      `json:"link_flair_type"`
				Wls                 *int        `json:"wls"`
				RemovedByCategory   interface{} `json:"removed_by_category"`
				BannedBy            interface{} `json:"banned_by"`
				AuthorFlairType     string      `json:"author_flair_type"`
				Domain              string      `json:"domain"`
				AllowLiveComments   bool        `json:"allow_live_comments"`
				SelftextHtml        *string     `json:"selftext_html"`
				Likes               interface{} `json:"likes"`
				SuggestedSort       *string     `json:"suggested_sort"`
				BannedAtUtc         interface{} `json:"banned_at_utc"`
				UrlOverriddenByDest string      `json:"url_overridden_by_dest,omitempty"`
				ViewCount           interface{} `json:"view_count"`
				Archived            bool        `json:"archived"`
				NoFollow            bool        `json:"no_follow"`
				IsCrosspostable     bool        `json:"is_crosspostable"`
				Pinned              bool        `json:"pinned"`
				Over18              bool        `json:"over_18"`
				Preview             struct {
					Images []struct {
						Source struct {
							Url    string `json:"url"`
							Width  int    `json:"width"`
							Height int    `json:"height"`
						} `json:"source"`
						Resolutions []struct {
							Url    string `json:"url"`
							Width  int    `json:"width"`
							Height int    `json:"height"`
						} `json:"resolutions"`
						Variants struct {
						} `json:"variants"`
						Id string `json:"id"`
					} `json:"images"`
					Enabled bool `json:"enabled"`
				} `json:"preview,omitempty"`
				AllAwardings             []interface{} `json:"all_awardings"`
				Awarders                 []interface{} `json:"awarders"`
				MediaOnly                bool          `json:"media_only"`
				CanGild                  bool          `json:"can_gild"`
				Spoiler                  bool          `json:"spoiler"`
				Locked                   bool          `json:"locked"`
				AuthorFlairText          *string       `json:"author_flair_text"`
				TreatmentTags            []interface{} `json:"treatment_tags"`
				Visited                  bool          `json:"visited"`
				RemovedBy                interface{}   `json:"removed_by"`
				NumReports               interface{}   `json:"num_reports"`
				Distinguished            interface{}   `json:"distinguished"`
				SubredditId              string        `json:"subreddit_id"`
				AuthorIsBlocked          bool          `json:"author_is_blocked"`
				ModReasonBy              interface{}   `json:"mod_reason_by"`
				RemovalReason            interface{}   `json:"removal_reason"`
				LinkFlairBackgroundColor string        `json:"link_flair_background_color"`
				Id                       string        `json:"id"`
				IsRobotIndexable         bool          `json:"is_robot_indexable"`
				ReportReasons            interface{}   `json:"report_reasons"`
				Author                   string        `json:"author"`
				DiscussionType           interface{}   `json:"discussion_type"`
				NumComments              int           `json:"num_comments"`
				SendReplies              bool          `json:"send_replies"`
				WhitelistStatus          *string       `json:"whitelist_status"`
				ContestMode              bool          `json:"contest_mode"`
				ModReports               []interface{} `json:"mod_reports"`
				AuthorPatreonFlair       bool          `json:"author_patreon_flair"`
				AuthorFlairTextColor     *string       `json:"author_flair_text_color"`
				Permalink                string        `json:"permalink"`
				ParentWhitelistStatus    *string       `json:"parent_whitelist_status"`
				Stickied                 bool          `json:"stickied"`
				Url                      string        `json:"url"`
				SubredditSubscribers     int           `json:"subreddit_subscribers"`
				CreatedUtc               float64       `json:"created_utc"`
				NumCrossposts            int           `json:"num_crossposts"`
				Media                    *struct {
					RedditVideo struct {
						BitrateKbps       int    `json:"bitrate_kbps"`
						FallbackUrl       string `json:"fallback_url"`
						HasAudio          bool   `json:"has_audio"`
						Height            int    `json:"height"`
						Width             int    `json:"width"`
						ScrubberMediaUrl  string `json:"scrubber_media_url"`
						DashUrl           string `json:"dash_url"`
						Duration          int    `json:"duration"`
						HlsUrl            string `json:"hls_url"`
						IsGif             bool   `json:"is_gif"`
						TranscodingStatus string `json:"transcoding_status"`
					} `json:"reddit_video"`
				} `json:"media"`
				IsVideo             bool   `json:"is_video"`
				LinkFlairTemplateId string `json:"link_flair_template_id,omitempty"`
				IsGallery           bool   `json:"is_gallery,omitempty"`
				GalleryData         struct {
					Items []struct {
						MediaId string `json:"media_id"`
						Id      int    `json:"id"`
					} `json:"items"`
				} `json:"gallery_data,omitempty"`
			} `json:"data"`
		} `json:"children"`
		Before interface{} `json:"before"`
	} `json:"data"`
}

func LoadResourceReddit() *ResourceReddit {
	var resourceReddit ResourceReddit
	err := json.Unmarshal(ResourceRedditJSON, &resourceReddit)
	if err != nil {
		panic(err)
	}
	return &resourceReddit
}
