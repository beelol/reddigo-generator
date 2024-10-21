package reddigo

import (
	"fmt"
	"io"
	"net/http"
)

type ReddigoSDK struct {
	clientID     string
	clientSecret string
	accessToken  string
	refreshToken string
	userAgent    string
	httpClient   *http.Client
}

func NewReddigoSDK(config RedditConfig) *ReddigoSDK {
	return &ReddigoSDK{
		clientID:     config.ClientID,
		clientSecret: config.ClientSecret,
		accessToken:  config.AccessToken,
		refreshToken: config.RefreshToken,
		userAgent:    config.UserAgent,
		httpClient:   &http.Client{},
	}
}

func (sdk *ReddigoSDK) MakeRequest(method, endpoint string, body io.Reader) (*http.Response, error) {
	url := fmt.Sprintf("https://oauth.reddit.com%s", endpoint)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", sdk.accessToken))
	req.Header.Set("User-Agent", sdk.userAgent)

	resp, err := sdk.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// if resp.StatusCode == http.StatusUnauthorized {
	// Handle token refresh if needed
	// }

	return resp, nil
}


/*
GetApiV1Me makes a GET request to /api/v1/me
ID: GET /api/v1/me
Description: Returns the identity of the user.
*/
func (sdk *ReddigoSDK) GetApiV1Me() (GetApiV1MeResponse, error) {
	url := fmt.Sprintf("/api/v1/me")
	// Construct the request for GET method
	url := fmt.Sprintf("/api/v1/me")
	req, err := sdk.MakeRequest("GET", url, nil)
	if err != nil {
		return GetApiV1MeResponse{}, err
	}
	defer resp.Body.Close()
	var response GetApiV1MeResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return GetApiV1MeResponse{}, err
	}
	return response, nil
}



/*
GetApiV1MeKarma makes a GET request to /api/v1/me/karma
ID: GET /api/v1/me/karma
Description: Return a breakdown of subreddit karma.
*/
func (sdk *ReddigoSDK) GetApiV1MeKarma() (GetApiV1MeKarmaResponse, error) {
	url := fmt.Sprintf("/api/v1/me/karma")
	// Construct the request for GET method
	url := fmt.Sprintf("/api/v1/me/karma")
	req, err := sdk.MakeRequest("GET", url, nil)
	if err != nil {
		return GetApiV1MeKarmaResponse{}, err
	}
	defer resp.Body.Close()
	var response GetApiV1MeKarmaResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return GetApiV1MeKarmaResponse{}, err
	}
	return response, nil
}



// GetApiV1MePrefsResponse represents the response for GET /api/v1/me/prefs
type GetApiV1MePrefsResponse struct {
	Fields string `json:"fields"` /* A comma-separated list of items from this set:beta
	threaded_messages
	hide_downs
	private_feeds
	activity_relevant_ads
	enable_reddit_pro_analytics_emails
	profile_opt_out
	bad_comment_autocollapse
	third_party_site_data_personalized_content
	show_link_flair
	live_bar_recommendations_enabled
	show_trending
	top_karma_subreddits
	country_code
	theme_selector
	monitor_mentions
	email_comment_reply
	newwindow
	email_new_user_welcome
	research
	ignore_suggested_sort
	show_presence
	email_upvote_comment
	email_digests
	whatsapp_comment_reply
	num_comments
	feed_recommendations_enabled
	clickgadget
	use_global_defaults
	label_nsfw
	domain_details
	show_stylesheets
	live_orangereds
	highlight_controversial
	mark_messages_read
	no_profanity
	email_unsubscribe_all
	whatsapp_enabled
	lang
	in_redesign_beta
	email_messages
	third_party_data_personalized_ads
	email_chat_request
	allow_clicktracking
	hide_from_robots
	show_gold_expiration
	show_twitter
	compress
	store_visits
	video_autoplay
	email_upvote_post
	email_username_mention
	media_preview
	email_user_new_follower
	nightmode
	enable_default_themes
	geopopular
	third_party_site_data_personalized_ads
	survey_last_seen_time
	threaded_modmail
	enable_followers
	hide_ups
	min_comment_score
	public_votes
	show_location_based_recommendations
	email_post_reply
	collapse_read_messages
	show_flair
	send_crosspost_messages
	search_include_over_18
	hide_ads
	third_party_personalized_ads
	min_link_score
	over_18
	sms_notifications_enabled
	numsites
	media
	legacy_search
	email_private_message
	send_welcome_messages
	email_community_discovery
	highlight_new_comments
	default_comment_sort
	accept_pms */
}

/*
GetApiV1MePrefs makes a GET request to /api/v1/me/prefs
ID: GET /api/v1/me/prefs
Description: Return the preference settings of the logged in user
*/
func (sdk *ReddigoSDK) GetApiV1MePrefs() (GetApiV1MePrefsResponse, error) {
	url := fmt.Sprintf("/api/v1/me/prefs")
	// Construct the request for GET method
	url := fmt.Sprintf("/api/v1/me/prefs")
	req, err := sdk.MakeRequest("GET", url, nil)
	if err != nil {
		return GetApiV1MePrefsResponse{}, err
	}
	defer resp.Body.Close()
	var response GetApiV1MePrefsResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return GetApiV1MePrefsResponse{}, err
	}
	return response, nil
}



type PatchApiV1MePrefsAcceptPmsEnum string

const (
	PatchApiV1MePrefsAcceptPmsEnumEveryone PatchApiV1MePrefsAcceptPmsEnum = "everyone"
	PatchApiV1MePrefsAcceptPmsEnumWhitelisted PatchApiV1MePrefsAcceptPmsEnum = "whitelisted"
)

type PatchApiV1MePrefsBadCommentAutocollapseEnum string

const (
	PatchApiV1MePrefsBadCommentAutocollapseEnumOff PatchApiV1MePrefsBadCommentAutocollapseEnum = "off"
	PatchApiV1MePrefsBadCommentAutocollapseEnumLow PatchApiV1MePrefsBadCommentAutocollapseEnum = "low"
	PatchApiV1MePrefsBadCommentAutocollapseEnumMedium PatchApiV1MePrefsBadCommentAutocollapseEnum = "medium"
	PatchApiV1MePrefsBadCommentAutocollapseEnumHigh PatchApiV1MePrefsBadCommentAutocollapseEnum = "high"
)

type PatchApiV1MePrefsCountryCodeEnum string

const (
	PatchApiV1MePrefsCountryCodeEnumWF PatchApiV1MePrefsCountryCodeEnum = "WF"
	PatchApiV1MePrefsCountryCodeEnumJP PatchApiV1MePrefsCountryCodeEnum = "JP"
	PatchApiV1MePrefsCountryCodeEnumJM PatchApiV1MePrefsCountryCodeEnum = "JM"
	PatchApiV1MePrefsCountryCodeEnumJO PatchApiV1MePrefsCountryCodeEnum = "JO"
	PatchApiV1MePrefsCountryCodeEnumWS PatchApiV1MePrefsCountryCodeEnum = "WS"
	PatchApiV1MePrefsCountryCodeEnumJE PatchApiV1MePrefsCountryCodeEnum = "JE"
	PatchApiV1MePrefsCountryCodeEnumGW PatchApiV1MePrefsCountryCodeEnum = "GW"
	PatchApiV1MePrefsCountryCodeEnumGU PatchApiV1MePrefsCountryCodeEnum = "GU"
	PatchApiV1MePrefsCountryCodeEnumGT PatchApiV1MePrefsCountryCodeEnum = "GT"
	PatchApiV1MePrefsCountryCodeEnumGS PatchApiV1MePrefsCountryCodeEnum = "GS"
	PatchApiV1MePrefsCountryCodeEnumGR PatchApiV1MePrefsCountryCodeEnum = "GR"
	PatchApiV1MePrefsCountryCodeEnumGQ PatchApiV1MePrefsCountryCodeEnum = "GQ"
	PatchApiV1MePrefsCountryCodeEnumGP PatchApiV1MePrefsCountryCodeEnum = "GP"
	PatchApiV1MePrefsCountryCodeEnumGY PatchApiV1MePrefsCountryCodeEnum = "GY"
	PatchApiV1MePrefsCountryCodeEnumGG PatchApiV1MePrefsCountryCodeEnum = "GG"
	PatchApiV1MePrefsCountryCodeEnumGF PatchApiV1MePrefsCountryCodeEnum = "GF"
	PatchApiV1MePrefsCountryCodeEnumGE PatchApiV1MePrefsCountryCodeEnum = "GE"
	PatchApiV1MePrefsCountryCodeEnumGD PatchApiV1MePrefsCountryCodeEnum = "GD"
	PatchApiV1MePrefsCountryCodeEnumGB PatchApiV1MePrefsCountryCodeEnum = "GB"
	PatchApiV1MePrefsCountryCodeEnumGA PatchApiV1MePrefsCountryCodeEnum = "GA"
	PatchApiV1MePrefsCountryCodeEnumGN PatchApiV1MePrefsCountryCodeEnum = "GN"
	PatchApiV1MePrefsCountryCodeEnumGM PatchApiV1MePrefsCountryCodeEnum = "GM"
	PatchApiV1MePrefsCountryCodeEnumGL PatchApiV1MePrefsCountryCodeEnum = "GL"
	PatchApiV1MePrefsCountryCodeEnumGI PatchApiV1MePrefsCountryCodeEnum = "GI"
	PatchApiV1MePrefsCountryCodeEnumGH PatchApiV1MePrefsCountryCodeEnum = "GH"
	PatchApiV1MePrefsCountryCodeEnumPR PatchApiV1MePrefsCountryCodeEnum = "PR"
	PatchApiV1MePrefsCountryCodeEnumPS PatchApiV1MePrefsCountryCodeEnum = "PS"
	PatchApiV1MePrefsCountryCodeEnumPW PatchApiV1MePrefsCountryCodeEnum = "PW"
	PatchApiV1MePrefsCountryCodeEnumPT PatchApiV1MePrefsCountryCodeEnum = "PT"
	PatchApiV1MePrefsCountryCodeEnumPY PatchApiV1MePrefsCountryCodeEnum = "PY"
	PatchApiV1MePrefsCountryCodeEnumPA PatchApiV1MePrefsCountryCodeEnum = "PA"
	PatchApiV1MePrefsCountryCodeEnumPF PatchApiV1MePrefsCountryCodeEnum = "PF"
	PatchApiV1MePrefsCountryCodeEnumPG PatchApiV1MePrefsCountryCodeEnum = "PG"
	PatchApiV1MePrefsCountryCodeEnumPE PatchApiV1MePrefsCountryCodeEnum = "PE"
	PatchApiV1MePrefsCountryCodeEnumPK PatchApiV1MePrefsCountryCodeEnum = "PK"
	PatchApiV1MePrefsCountryCodeEnumPH PatchApiV1MePrefsCountryCodeEnum = "PH"
	PatchApiV1MePrefsCountryCodeEnumPN PatchApiV1MePrefsCountryCodeEnum = "PN"
	PatchApiV1MePrefsCountryCodeEnumPL PatchApiV1MePrefsCountryCodeEnum = "PL"
	PatchApiV1MePrefsCountryCodeEnumPM PatchApiV1MePrefsCountryCodeEnum = "PM"
	PatchApiV1MePrefsCountryCodeEnumZM PatchApiV1MePrefsCountryCodeEnum = "ZM"
	PatchApiV1MePrefsCountryCodeEnumZA PatchApiV1MePrefsCountryCodeEnum = "ZA"
	PatchApiV1MePrefsCountryCodeEnumZZ PatchApiV1MePrefsCountryCodeEnum = "ZZ"
	PatchApiV1MePrefsCountryCodeEnumZW PatchApiV1MePrefsCountryCodeEnum = "ZW"
	PatchApiV1MePrefsCountryCodeEnumME PatchApiV1MePrefsCountryCodeEnum = "ME"
	PatchApiV1MePrefsCountryCodeEnumMD PatchApiV1MePrefsCountryCodeEnum = "MD"
	PatchApiV1MePrefsCountryCodeEnumMG PatchApiV1MePrefsCountryCodeEnum = "MG"
	PatchApiV1MePrefsCountryCodeEnumMF PatchApiV1MePrefsCountryCodeEnum = "MF"
	PatchApiV1MePrefsCountryCodeEnumMA PatchApiV1MePrefsCountryCodeEnum = "MA"
	PatchApiV1MePrefsCountryCodeEnumMC PatchApiV1MePrefsCountryCodeEnum = "MC"
	PatchApiV1MePrefsCountryCodeEnumMM PatchApiV1MePrefsCountryCodeEnum = "MM"
	PatchApiV1MePrefsCountryCodeEnumML PatchApiV1MePrefsCountryCodeEnum = "ML"
	PatchApiV1MePrefsCountryCodeEnumMO PatchApiV1MePrefsCountryCodeEnum = "MO"
	PatchApiV1MePrefsCountryCodeEnumMN PatchApiV1MePrefsCountryCodeEnum = "MN"
	PatchApiV1MePrefsCountryCodeEnumMH PatchApiV1MePrefsCountryCodeEnum = "MH"
	PatchApiV1MePrefsCountryCodeEnumMK PatchApiV1MePrefsCountryCodeEnum = "MK"
	PatchApiV1MePrefsCountryCodeEnumMU PatchApiV1MePrefsCountryCodeEnum = "MU"
	PatchApiV1MePrefsCountryCodeEnumMT PatchApiV1MePrefsCountryCodeEnum = "MT"
	PatchApiV1MePrefsCountryCodeEnumMW PatchApiV1MePrefsCountryCodeEnum = "MW"
	PatchApiV1MePrefsCountryCodeEnumMV PatchApiV1MePrefsCountryCodeEnum = "MV"
	PatchApiV1MePrefsCountryCodeEnumMQ PatchApiV1MePrefsCountryCodeEnum = "MQ"
	PatchApiV1MePrefsCountryCodeEnumMP PatchApiV1MePrefsCountryCodeEnum = "MP"
	PatchApiV1MePrefsCountryCodeEnumMS PatchApiV1MePrefsCountryCodeEnum = "MS"
	PatchApiV1MePrefsCountryCodeEnumMR PatchApiV1MePrefsCountryCodeEnum = "MR"
	PatchApiV1MePrefsCountryCodeEnumMY PatchApiV1MePrefsCountryCodeEnum = "MY"
	PatchApiV1MePrefsCountryCodeEnumMX PatchApiV1MePrefsCountryCodeEnum = "MX"
	PatchApiV1MePrefsCountryCodeEnumMZ PatchApiV1MePrefsCountryCodeEnum = "MZ"
	PatchApiV1MePrefsCountryCodeEnumFR PatchApiV1MePrefsCountryCodeEnum = "FR"
	PatchApiV1MePrefsCountryCodeEnumFI PatchApiV1MePrefsCountryCodeEnum = "FI"
	PatchApiV1MePrefsCountryCodeEnumFJ PatchApiV1MePrefsCountryCodeEnum = "FJ"
	PatchApiV1MePrefsCountryCodeEnumFK PatchApiV1MePrefsCountryCodeEnum = "FK"
	PatchApiV1MePrefsCountryCodeEnumFM PatchApiV1MePrefsCountryCodeEnum = "FM"
	PatchApiV1MePrefsCountryCodeEnumFO PatchApiV1MePrefsCountryCodeEnum = "FO"
	PatchApiV1MePrefsCountryCodeEnumCK PatchApiV1MePrefsCountryCodeEnum = "CK"
	PatchApiV1MePrefsCountryCodeEnumCI PatchApiV1MePrefsCountryCodeEnum = "CI"
	PatchApiV1MePrefsCountryCodeEnumCH PatchApiV1MePrefsCountryCodeEnum = "CH"
	PatchApiV1MePrefsCountryCodeEnumCO PatchApiV1MePrefsCountryCodeEnum = "CO"
	PatchApiV1MePrefsCountryCodeEnumCN PatchApiV1MePrefsCountryCodeEnum = "CN"
	PatchApiV1MePrefsCountryCodeEnumCM PatchApiV1MePrefsCountryCodeEnum = "CM"
	PatchApiV1MePrefsCountryCodeEnumCL PatchApiV1MePrefsCountryCodeEnum = "CL"
	PatchApiV1MePrefsCountryCodeEnumCC PatchApiV1MePrefsCountryCodeEnum = "CC"
	PatchApiV1MePrefsCountryCodeEnumCA PatchApiV1MePrefsCountryCodeEnum = "CA"
	PatchApiV1MePrefsCountryCodeEnumCG PatchApiV1MePrefsCountryCodeEnum = "CG"
	PatchApiV1MePrefsCountryCodeEnumCF PatchApiV1MePrefsCountryCodeEnum = "CF"
	PatchApiV1MePrefsCountryCodeEnumCD PatchApiV1MePrefsCountryCodeEnum = "CD"
	PatchApiV1MePrefsCountryCodeEnumCZ PatchApiV1MePrefsCountryCodeEnum = "CZ"
	PatchApiV1MePrefsCountryCodeEnumCY PatchApiV1MePrefsCountryCodeEnum = "CY"
	PatchApiV1MePrefsCountryCodeEnumCX PatchApiV1MePrefsCountryCodeEnum = "CX"
	PatchApiV1MePrefsCountryCodeEnumCR PatchApiV1MePrefsCountryCodeEnum = "CR"
	PatchApiV1MePrefsCountryCodeEnumCW PatchApiV1MePrefsCountryCodeEnum = "CW"
	PatchApiV1MePrefsCountryCodeEnumCV PatchApiV1MePrefsCountryCodeEnum = "CV"
	PatchApiV1MePrefsCountryCodeEnumCU PatchApiV1MePrefsCountryCodeEnum = "CU"
	PatchApiV1MePrefsCountryCodeEnumSZ PatchApiV1MePrefsCountryCodeEnum = "SZ"
	PatchApiV1MePrefsCountryCodeEnumSY PatchApiV1MePrefsCountryCodeEnum = "SY"
	PatchApiV1MePrefsCountryCodeEnumSX PatchApiV1MePrefsCountryCodeEnum = "SX"
	PatchApiV1MePrefsCountryCodeEnumSS PatchApiV1MePrefsCountryCodeEnum = "SS"
	PatchApiV1MePrefsCountryCodeEnumSR PatchApiV1MePrefsCountryCodeEnum = "SR"
	PatchApiV1MePrefsCountryCodeEnumSV PatchApiV1MePrefsCountryCodeEnum = "SV"
	PatchApiV1MePrefsCountryCodeEnumST PatchApiV1MePrefsCountryCodeEnum = "ST"
	PatchApiV1MePrefsCountryCodeEnumSK PatchApiV1MePrefsCountryCodeEnum = "SK"
	PatchApiV1MePrefsCountryCodeEnumSJ PatchApiV1MePrefsCountryCodeEnum = "SJ"
	PatchApiV1MePrefsCountryCodeEnumSI PatchApiV1MePrefsCountryCodeEnum = "SI"
	PatchApiV1MePrefsCountryCodeEnumSH PatchApiV1MePrefsCountryCodeEnum = "SH"
	PatchApiV1MePrefsCountryCodeEnumSO PatchApiV1MePrefsCountryCodeEnum = "SO"
	PatchApiV1MePrefsCountryCodeEnumSN PatchApiV1MePrefsCountryCodeEnum = "SN"
	PatchApiV1MePrefsCountryCodeEnumSM PatchApiV1MePrefsCountryCodeEnum = "SM"
	PatchApiV1MePrefsCountryCodeEnumSL PatchApiV1MePrefsCountryCodeEnum = "SL"
	PatchApiV1MePrefsCountryCodeEnumSC PatchApiV1MePrefsCountryCodeEnum = "SC"
	PatchApiV1MePrefsCountryCodeEnumSB PatchApiV1MePrefsCountryCodeEnum = "SB"
	PatchApiV1MePrefsCountryCodeEnumSA PatchApiV1MePrefsCountryCodeEnum = "SA"
	PatchApiV1MePrefsCountryCodeEnumSG PatchApiV1MePrefsCountryCodeEnum = "SG"
	PatchApiV1MePrefsCountryCodeEnumSE PatchApiV1MePrefsCountryCodeEnum = "SE"
	PatchApiV1MePrefsCountryCodeEnumSD PatchApiV1MePrefsCountryCodeEnum = "SD"
	PatchApiV1MePrefsCountryCodeEnumYE PatchApiV1MePrefsCountryCodeEnum = "YE"
	PatchApiV1MePrefsCountryCodeEnumYT PatchApiV1MePrefsCountryCodeEnum = "YT"
	PatchApiV1MePrefsCountryCodeEnumLB PatchApiV1MePrefsCountryCodeEnum = "LB"
	PatchApiV1MePrefsCountryCodeEnumLC PatchApiV1MePrefsCountryCodeEnum = "LC"
	PatchApiV1MePrefsCountryCodeEnumLA PatchApiV1MePrefsCountryCodeEnum = "LA"
	PatchApiV1MePrefsCountryCodeEnumLK PatchApiV1MePrefsCountryCodeEnum = "LK"
	PatchApiV1MePrefsCountryCodeEnumLI PatchApiV1MePrefsCountryCodeEnum = "LI"
	PatchApiV1MePrefsCountryCodeEnumLV PatchApiV1MePrefsCountryCodeEnum = "LV"
	PatchApiV1MePrefsCountryCodeEnumLT PatchApiV1MePrefsCountryCodeEnum = "LT"
	PatchApiV1MePrefsCountryCodeEnumLU PatchApiV1MePrefsCountryCodeEnum = "LU"
	PatchApiV1MePrefsCountryCodeEnumLR PatchApiV1MePrefsCountryCodeEnum = "LR"
	PatchApiV1MePrefsCountryCodeEnumLS PatchApiV1MePrefsCountryCodeEnum = "LS"
	PatchApiV1MePrefsCountryCodeEnumLY PatchApiV1MePrefsCountryCodeEnum = "LY"
	PatchApiV1MePrefsCountryCodeEnumVA PatchApiV1MePrefsCountryCodeEnum = "VA"
	PatchApiV1MePrefsCountryCodeEnumVC PatchApiV1MePrefsCountryCodeEnum = "VC"
	PatchApiV1MePrefsCountryCodeEnumVE PatchApiV1MePrefsCountryCodeEnum = "VE"
	PatchApiV1MePrefsCountryCodeEnumVG PatchApiV1MePrefsCountryCodeEnum = "VG"
	PatchApiV1MePrefsCountryCodeEnumIQ PatchApiV1MePrefsCountryCodeEnum = "IQ"
	PatchApiV1MePrefsCountryCodeEnumVI PatchApiV1MePrefsCountryCodeEnum = "VI"
	PatchApiV1MePrefsCountryCodeEnumIS PatchApiV1MePrefsCountryCodeEnum = "IS"
	PatchApiV1MePrefsCountryCodeEnumIR PatchApiV1MePrefsCountryCodeEnum = "IR"
	PatchApiV1MePrefsCountryCodeEnumIT PatchApiV1MePrefsCountryCodeEnum = "IT"
	PatchApiV1MePrefsCountryCodeEnumVN PatchApiV1MePrefsCountryCodeEnum = "VN"
	PatchApiV1MePrefsCountryCodeEnumIM PatchApiV1MePrefsCountryCodeEnum = "IM"
	PatchApiV1MePrefsCountryCodeEnumIL PatchApiV1MePrefsCountryCodeEnum = "IL"
	PatchApiV1MePrefsCountryCodeEnumIO PatchApiV1MePrefsCountryCodeEnum = "IO"
	PatchApiV1MePrefsCountryCodeEnumIN PatchApiV1MePrefsCountryCodeEnum = "IN"
	PatchApiV1MePrefsCountryCodeEnumIE PatchApiV1MePrefsCountryCodeEnum = "IE"
	PatchApiV1MePrefsCountryCodeEnumID PatchApiV1MePrefsCountryCodeEnum = "ID"
	PatchApiV1MePrefsCountryCodeEnumBD PatchApiV1MePrefsCountryCodeEnum = "BD"
	PatchApiV1MePrefsCountryCodeEnumBE PatchApiV1MePrefsCountryCodeEnum = "BE"
	PatchApiV1MePrefsCountryCodeEnumBF PatchApiV1MePrefsCountryCodeEnum = "BF"
	PatchApiV1MePrefsCountryCodeEnumBG PatchApiV1MePrefsCountryCodeEnum = "BG"
	PatchApiV1MePrefsCountryCodeEnumBA PatchApiV1MePrefsCountryCodeEnum = "BA"
	PatchApiV1MePrefsCountryCodeEnumBB PatchApiV1MePrefsCountryCodeEnum = "BB"
	PatchApiV1MePrefsCountryCodeEnumBL PatchApiV1MePrefsCountryCodeEnum = "BL"
	PatchApiV1MePrefsCountryCodeEnumBM PatchApiV1MePrefsCountryCodeEnum = "BM"
	PatchApiV1MePrefsCountryCodeEnumBN PatchApiV1MePrefsCountryCodeEnum = "BN"
	PatchApiV1MePrefsCountryCodeEnumBO PatchApiV1MePrefsCountryCodeEnum = "BO"
	PatchApiV1MePrefsCountryCodeEnumBH PatchApiV1MePrefsCountryCodeEnum = "BH"
	PatchApiV1MePrefsCountryCodeEnumBI PatchApiV1MePrefsCountryCodeEnum = "BI"
	PatchApiV1MePrefsCountryCodeEnumBJ PatchApiV1MePrefsCountryCodeEnum = "BJ"
	PatchApiV1MePrefsCountryCodeEnumBT PatchApiV1MePrefsCountryCodeEnum = "BT"
	PatchApiV1MePrefsCountryCodeEnumBV PatchApiV1MePrefsCountryCodeEnum = "BV"
	PatchApiV1MePrefsCountryCodeEnumBW PatchApiV1MePrefsCountryCodeEnum = "BW"
	PatchApiV1MePrefsCountryCodeEnumBQ PatchApiV1MePrefsCountryCodeEnum = "BQ"
	PatchApiV1MePrefsCountryCodeEnumBR PatchApiV1MePrefsCountryCodeEnum = "BR"
	PatchApiV1MePrefsCountryCodeEnumBS PatchApiV1MePrefsCountryCodeEnum = "BS"
	PatchApiV1MePrefsCountryCodeEnumBY PatchApiV1MePrefsCountryCodeEnum = "BY"
	PatchApiV1MePrefsCountryCodeEnumBZ PatchApiV1MePrefsCountryCodeEnum = "BZ"
	PatchApiV1MePrefsCountryCodeEnumRU PatchApiV1MePrefsCountryCodeEnum = "RU"
	PatchApiV1MePrefsCountryCodeEnumRW PatchApiV1MePrefsCountryCodeEnum = "RW"
	PatchApiV1MePrefsCountryCodeEnumRS PatchApiV1MePrefsCountryCodeEnum = "RS"
	PatchApiV1MePrefsCountryCodeEnumRE PatchApiV1MePrefsCountryCodeEnum = "RE"
	PatchApiV1MePrefsCountryCodeEnumRO PatchApiV1MePrefsCountryCodeEnum = "RO"
	PatchApiV1MePrefsCountryCodeEnumOM PatchApiV1MePrefsCountryCodeEnum = "OM"
	PatchApiV1MePrefsCountryCodeEnumHR PatchApiV1MePrefsCountryCodeEnum = "HR"
	PatchApiV1MePrefsCountryCodeEnumHT PatchApiV1MePrefsCountryCodeEnum = "HT"
	PatchApiV1MePrefsCountryCodeEnumHU PatchApiV1MePrefsCountryCodeEnum = "HU"
	PatchApiV1MePrefsCountryCodeEnumHK PatchApiV1MePrefsCountryCodeEnum = "HK"
	PatchApiV1MePrefsCountryCodeEnumHN PatchApiV1MePrefsCountryCodeEnum = "HN"
	PatchApiV1MePrefsCountryCodeEnumHM PatchApiV1MePrefsCountryCodeEnum = "HM"
	PatchApiV1MePrefsCountryCodeEnumEH PatchApiV1MePrefsCountryCodeEnum = "EH"
	PatchApiV1MePrefsCountryCodeEnumEE PatchApiV1MePrefsCountryCodeEnum = "EE"
	PatchApiV1MePrefsCountryCodeEnumEG PatchApiV1MePrefsCountryCodeEnum = "EG"
	PatchApiV1MePrefsCountryCodeEnumEC PatchApiV1MePrefsCountryCodeEnum = "EC"
	PatchApiV1MePrefsCountryCodeEnumET PatchApiV1MePrefsCountryCodeEnum = "ET"
	PatchApiV1MePrefsCountryCodeEnumES PatchApiV1MePrefsCountryCodeEnum = "ES"
	PatchApiV1MePrefsCountryCodeEnumER PatchApiV1MePrefsCountryCodeEnum = "ER"
	PatchApiV1MePrefsCountryCodeEnumUY PatchApiV1MePrefsCountryCodeEnum = "UY"
	PatchApiV1MePrefsCountryCodeEnumUZ PatchApiV1MePrefsCountryCodeEnum = "UZ"
	PatchApiV1MePrefsCountryCodeEnumUS PatchApiV1MePrefsCountryCodeEnum = "US"
	PatchApiV1MePrefsCountryCodeEnumUM PatchApiV1MePrefsCountryCodeEnum = "UM"
	PatchApiV1MePrefsCountryCodeEnumUG PatchApiV1MePrefsCountryCodeEnum = "UG"
	PatchApiV1MePrefsCountryCodeEnumUA PatchApiV1MePrefsCountryCodeEnum = "UA"
	PatchApiV1MePrefsCountryCodeEnumVU PatchApiV1MePrefsCountryCodeEnum = "VU"
	PatchApiV1MePrefsCountryCodeEnumNI PatchApiV1MePrefsCountryCodeEnum = "NI"
	PatchApiV1MePrefsCountryCodeEnumNL PatchApiV1MePrefsCountryCodeEnum = "NL"
	PatchApiV1MePrefsCountryCodeEnumNO PatchApiV1MePrefsCountryCodeEnum = "NO"
	PatchApiV1MePrefsCountryCodeEnumNA PatchApiV1MePrefsCountryCodeEnum = "NA"
	PatchApiV1MePrefsCountryCodeEnumNC PatchApiV1MePrefsCountryCodeEnum = "NC"
	PatchApiV1MePrefsCountryCodeEnumNE PatchApiV1MePrefsCountryCodeEnum = "NE"
	PatchApiV1MePrefsCountryCodeEnumNF PatchApiV1MePrefsCountryCodeEnum = "NF"
	PatchApiV1MePrefsCountryCodeEnumNG PatchApiV1MePrefsCountryCodeEnum = "NG"
	PatchApiV1MePrefsCountryCodeEnumNZ PatchApiV1MePrefsCountryCodeEnum = "NZ"
	PatchApiV1MePrefsCountryCodeEnumNP PatchApiV1MePrefsCountryCodeEnum = "NP"
	PatchApiV1MePrefsCountryCodeEnumNR PatchApiV1MePrefsCountryCodeEnum = "NR"
	PatchApiV1MePrefsCountryCodeEnumNU PatchApiV1MePrefsCountryCodeEnum = "NU"
	PatchApiV1MePrefsCountryCodeEnumXK PatchApiV1MePrefsCountryCodeEnum = "XK"
	PatchApiV1MePrefsCountryCodeEnumXZ PatchApiV1MePrefsCountryCodeEnum = "XZ"
	PatchApiV1MePrefsCountryCodeEnumXX PatchApiV1MePrefsCountryCodeEnum = "XX"
	PatchApiV1MePrefsCountryCodeEnumKG PatchApiV1MePrefsCountryCodeEnum = "KG"
	PatchApiV1MePrefsCountryCodeEnumKE PatchApiV1MePrefsCountryCodeEnum = "KE"
	PatchApiV1MePrefsCountryCodeEnumKI PatchApiV1MePrefsCountryCodeEnum = "KI"
	PatchApiV1MePrefsCountryCodeEnumKH PatchApiV1MePrefsCountryCodeEnum = "KH"
	PatchApiV1MePrefsCountryCodeEnumKN PatchApiV1MePrefsCountryCodeEnum = "KN"
	PatchApiV1MePrefsCountryCodeEnumKM PatchApiV1MePrefsCountryCodeEnum = "KM"
	PatchApiV1MePrefsCountryCodeEnumKR PatchApiV1MePrefsCountryCodeEnum = "KR"
	PatchApiV1MePrefsCountryCodeEnumKP PatchApiV1MePrefsCountryCodeEnum = "KP"
	PatchApiV1MePrefsCountryCodeEnumKW PatchApiV1MePrefsCountryCodeEnum = "KW"
	PatchApiV1MePrefsCountryCodeEnumKZ PatchApiV1MePrefsCountryCodeEnum = "KZ"
	PatchApiV1MePrefsCountryCodeEnumKY PatchApiV1MePrefsCountryCodeEnum = "KY"
	PatchApiV1MePrefsCountryCodeEnumDO PatchApiV1MePrefsCountryCodeEnum = "DO"
	PatchApiV1MePrefsCountryCodeEnumDM PatchApiV1MePrefsCountryCodeEnum = "DM"
	PatchApiV1MePrefsCountryCodeEnumDJ PatchApiV1MePrefsCountryCodeEnum = "DJ"
	PatchApiV1MePrefsCountryCodeEnumDK PatchApiV1MePrefsCountryCodeEnum = "DK"
	PatchApiV1MePrefsCountryCodeEnumDE PatchApiV1MePrefsCountryCodeEnum = "DE"
	PatchApiV1MePrefsCountryCodeEnumDZ PatchApiV1MePrefsCountryCodeEnum = "DZ"
	PatchApiV1MePrefsCountryCodeEnumTZ PatchApiV1MePrefsCountryCodeEnum = "TZ"
	PatchApiV1MePrefsCountryCodeEnumTV PatchApiV1MePrefsCountryCodeEnum = "TV"
	PatchApiV1MePrefsCountryCodeEnumTW PatchApiV1MePrefsCountryCodeEnum = "TW"
	PatchApiV1MePrefsCountryCodeEnumTT PatchApiV1MePrefsCountryCodeEnum = "TT"
	PatchApiV1MePrefsCountryCodeEnumTR PatchApiV1MePrefsCountryCodeEnum = "TR"
	PatchApiV1MePrefsCountryCodeEnumTN PatchApiV1MePrefsCountryCodeEnum = "TN"
	PatchApiV1MePrefsCountryCodeEnumTO PatchApiV1MePrefsCountryCodeEnum = "TO"
	PatchApiV1MePrefsCountryCodeEnumTL PatchApiV1MePrefsCountryCodeEnum = "TL"
	PatchApiV1MePrefsCountryCodeEnumTM PatchApiV1MePrefsCountryCodeEnum = "TM"
	PatchApiV1MePrefsCountryCodeEnumTJ PatchApiV1MePrefsCountryCodeEnum = "TJ"
	PatchApiV1MePrefsCountryCodeEnumTK PatchApiV1MePrefsCountryCodeEnum = "TK"
	PatchApiV1MePrefsCountryCodeEnumTH PatchApiV1MePrefsCountryCodeEnum = "TH"
	PatchApiV1MePrefsCountryCodeEnumTF PatchApiV1MePrefsCountryCodeEnum = "TF"
	PatchApiV1MePrefsCountryCodeEnumTG PatchApiV1MePrefsCountryCodeEnum = "TG"
	PatchApiV1MePrefsCountryCodeEnumTD PatchApiV1MePrefsCountryCodeEnum = "TD"
	PatchApiV1MePrefsCountryCodeEnumTC PatchApiV1MePrefsCountryCodeEnum = "TC"
	PatchApiV1MePrefsCountryCodeEnumAE PatchApiV1MePrefsCountryCodeEnum = "AE"
	PatchApiV1MePrefsCountryCodeEnumAD PatchApiV1MePrefsCountryCodeEnum = "AD"
	PatchApiV1MePrefsCountryCodeEnumAG PatchApiV1MePrefsCountryCodeEnum = "AG"
	PatchApiV1MePrefsCountryCodeEnumAF PatchApiV1MePrefsCountryCodeEnum = "AF"
	PatchApiV1MePrefsCountryCodeEnumAI PatchApiV1MePrefsCountryCodeEnum = "AI"
	PatchApiV1MePrefsCountryCodeEnumAM PatchApiV1MePrefsCountryCodeEnum = "AM"
	PatchApiV1MePrefsCountryCodeEnumAL PatchApiV1MePrefsCountryCodeEnum = "AL"
	PatchApiV1MePrefsCountryCodeEnumAO PatchApiV1MePrefsCountryCodeEnum = "AO"
	PatchApiV1MePrefsCountryCodeEnumAN PatchApiV1MePrefsCountryCodeEnum = "AN"
	PatchApiV1MePrefsCountryCodeEnumAQ PatchApiV1MePrefsCountryCodeEnum = "AQ"
	PatchApiV1MePrefsCountryCodeEnumAS PatchApiV1MePrefsCountryCodeEnum = "AS"
	PatchApiV1MePrefsCountryCodeEnumAR PatchApiV1MePrefsCountryCodeEnum = "AR"
	PatchApiV1MePrefsCountryCodeEnumAU PatchApiV1MePrefsCountryCodeEnum = "AU"
	PatchApiV1MePrefsCountryCodeEnumAT PatchApiV1MePrefsCountryCodeEnum = "AT"
	PatchApiV1MePrefsCountryCodeEnumAW PatchApiV1MePrefsCountryCodeEnum = "AW"
	PatchApiV1MePrefsCountryCodeEnumAX PatchApiV1MePrefsCountryCodeEnum = "AX"
	PatchApiV1MePrefsCountryCodeEnumAZ PatchApiV1MePrefsCountryCodeEnum = "AZ"
	PatchApiV1MePrefsCountryCodeEnumQA PatchApiV1MePrefsCountryCodeEnum = "QA"
)

type PatchApiV1MePrefsDefaultCommentSortEnum string

const (
	PatchApiV1MePrefsDefaultCommentSortEnumConfidence PatchApiV1MePrefsDefaultCommentSortEnum = "confidence"
	PatchApiV1MePrefsDefaultCommentSortEnumTop PatchApiV1MePrefsDefaultCommentSortEnum = "top"
	PatchApiV1MePrefsDefaultCommentSortEnumNew PatchApiV1MePrefsDefaultCommentSortEnum = "new"
	PatchApiV1MePrefsDefaultCommentSortEnumControversial PatchApiV1MePrefsDefaultCommentSortEnum = "controversial"
	PatchApiV1MePrefsDefaultCommentSortEnumOld PatchApiV1MePrefsDefaultCommentSortEnum = "old"
	PatchApiV1MePrefsDefaultCommentSortEnumRandom PatchApiV1MePrefsDefaultCommentSortEnum = "random"
	PatchApiV1MePrefsDefaultCommentSortEnumQa PatchApiV1MePrefsDefaultCommentSortEnum = "qa"
	PatchApiV1MePrefsDefaultCommentSortEnumLive PatchApiV1MePrefsDefaultCommentSortEnum = "live"
)

type PatchApiV1MePrefsGEnum string

const (
	PatchApiV1MePrefsGEnumGLOBAL PatchApiV1MePrefsGEnum = "GLOBAL"
	PatchApiV1MePrefsGEnumUS PatchApiV1MePrefsGEnum = "US"
	PatchApiV1MePrefsGEnumAR PatchApiV1MePrefsGEnum = "AR"
	PatchApiV1MePrefsGEnumAU PatchApiV1MePrefsGEnum = "AU"
	PatchApiV1MePrefsGEnumBG PatchApiV1MePrefsGEnum = "BG"
	PatchApiV1MePrefsGEnumCA PatchApiV1MePrefsGEnum = "CA"
	PatchApiV1MePrefsGEnumCL PatchApiV1MePrefsGEnum = "CL"
	PatchApiV1MePrefsGEnumCO PatchApiV1MePrefsGEnum = "CO"
	PatchApiV1MePrefsGEnumHR PatchApiV1MePrefsGEnum = "HR"
	PatchApiV1MePrefsGEnumCZ PatchApiV1MePrefsGEnum = "CZ"
	PatchApiV1MePrefsGEnumFI PatchApiV1MePrefsGEnum = "FI"
	PatchApiV1MePrefsGEnumFR PatchApiV1MePrefsGEnum = "FR"
	PatchApiV1MePrefsGEnumDE PatchApiV1MePrefsGEnum = "DE"
	PatchApiV1MePrefsGEnumGR PatchApiV1MePrefsGEnum = "GR"
	PatchApiV1MePrefsGEnumHU PatchApiV1MePrefsGEnum = "HU"
	PatchApiV1MePrefsGEnumIS PatchApiV1MePrefsGEnum = "IS"
	PatchApiV1MePrefsGEnumIN PatchApiV1MePrefsGEnum = "IN"
	PatchApiV1MePrefsGEnumIE PatchApiV1MePrefsGEnum = "IE"
	PatchApiV1MePrefsGEnumIT PatchApiV1MePrefsGEnum = "IT"
	PatchApiV1MePrefsGEnumJP PatchApiV1MePrefsGEnum = "JP"
	PatchApiV1MePrefsGEnumMY PatchApiV1MePrefsGEnum = "MY"
	PatchApiV1MePrefsGEnumMX PatchApiV1MePrefsGEnum = "MX"
	PatchApiV1MePrefsGEnumNZ PatchApiV1MePrefsGEnum = "NZ"
	PatchApiV1MePrefsGEnumPH PatchApiV1MePrefsGEnum = "PH"
	PatchApiV1MePrefsGEnumPL PatchApiV1MePrefsGEnum = "PL"
	PatchApiV1MePrefsGEnumPT PatchApiV1MePrefsGEnum = "PT"
	PatchApiV1MePrefsGEnumPR PatchApiV1MePrefsGEnum = "PR"
	PatchApiV1MePrefsGEnumRO PatchApiV1MePrefsGEnum = "RO"
	PatchApiV1MePrefsGEnumRS PatchApiV1MePrefsGEnum = "RS"
	PatchApiV1MePrefsGEnumSG PatchApiV1MePrefsGEnum = "SG"
	PatchApiV1MePrefsGEnumES PatchApiV1MePrefsGEnum = "ES"
	PatchApiV1MePrefsGEnumSE PatchApiV1MePrefsGEnum = "SE"
	PatchApiV1MePrefsGEnumTW PatchApiV1MePrefsGEnum = "TW"
	PatchApiV1MePrefsGEnumTH PatchApiV1MePrefsGEnum = "TH"
	PatchApiV1MePrefsGEnumTR PatchApiV1MePrefsGEnum = "TR"
	PatchApiV1MePrefsGEnumGB PatchApiV1MePrefsGEnum = "GB"
	PatchApiV1MePrefsGEnumUS_WA PatchApiV1MePrefsGEnum = "US_WA"
	PatchApiV1MePrefsGEnumUS_DE PatchApiV1MePrefsGEnum = "US_DE"
	PatchApiV1MePrefsGEnumUS_DC PatchApiV1MePrefsGEnum = "US_DC"
	PatchApiV1MePrefsGEnumUS_WI PatchApiV1MePrefsGEnum = "US_WI"
	PatchApiV1MePrefsGEnumUS_WV PatchApiV1MePrefsGEnum = "US_WV"
	PatchApiV1MePrefsGEnumUS_HI PatchApiV1MePrefsGEnum = "US_HI"
	PatchApiV1MePrefsGEnumUS_FL PatchApiV1MePrefsGEnum = "US_FL"
	PatchApiV1MePrefsGEnumUS_WY PatchApiV1MePrefsGEnum = "US_WY"
	PatchApiV1MePrefsGEnumUS_NH PatchApiV1MePrefsGEnum = "US_NH"
	PatchApiV1MePrefsGEnumUS_NJ PatchApiV1MePrefsGEnum = "US_NJ"
	PatchApiV1MePrefsGEnumUS_NM PatchApiV1MePrefsGEnum = "US_NM"
	PatchApiV1MePrefsGEnumUS_TX PatchApiV1MePrefsGEnum = "US_TX"
	PatchApiV1MePrefsGEnumUS_LA PatchApiV1MePrefsGEnum = "US_LA"
	PatchApiV1MePrefsGEnumUS_NC PatchApiV1MePrefsGEnum = "US_NC"
	PatchApiV1MePrefsGEnumUS_ND PatchApiV1MePrefsGEnum = "US_ND"
	PatchApiV1MePrefsGEnumUS_NE PatchApiV1MePrefsGEnum = "US_NE"
	PatchApiV1MePrefsGEnumUS_TN PatchApiV1MePrefsGEnum = "US_TN"
	PatchApiV1MePrefsGEnumUS_NY PatchApiV1MePrefsGEnum = "US_NY"
	PatchApiV1MePrefsGEnumUS_PA PatchApiV1MePrefsGEnum = "US_PA"
	PatchApiV1MePrefsGEnumUS_CA PatchApiV1MePrefsGEnum = "US_CA"
	PatchApiV1MePrefsGEnumUS_NV PatchApiV1MePrefsGEnum = "US_NV"
	PatchApiV1MePrefsGEnumUS_VA PatchApiV1MePrefsGEnum = "US_VA"
	PatchApiV1MePrefsGEnumUS_CO PatchApiV1MePrefsGEnum = "US_CO"
	PatchApiV1MePrefsGEnumUS_AK PatchApiV1MePrefsGEnum = "US_AK"
	PatchApiV1MePrefsGEnumUS_AL PatchApiV1MePrefsGEnum = "US_AL"
	PatchApiV1MePrefsGEnumUS_AR PatchApiV1MePrefsGEnum = "US_AR"
	PatchApiV1MePrefsGEnumUS_VT PatchApiV1MePrefsGEnum = "US_VT"
	PatchApiV1MePrefsGEnumUS_IL PatchApiV1MePrefsGEnum = "US_IL"
	PatchApiV1MePrefsGEnumUS_GA PatchApiV1MePrefsGEnum = "US_GA"
	PatchApiV1MePrefsGEnumUS_IN PatchApiV1MePrefsGEnum = "US_IN"
	PatchApiV1MePrefsGEnumUS_IA PatchApiV1MePrefsGEnum = "US_IA"
	PatchApiV1MePrefsGEnumUS_OK PatchApiV1MePrefsGEnum = "US_OK"
	PatchApiV1MePrefsGEnumUS_AZ PatchApiV1MePrefsGEnum = "US_AZ"
	PatchApiV1MePrefsGEnumUS_ID PatchApiV1MePrefsGEnum = "US_ID"
	PatchApiV1MePrefsGEnumUS_CT PatchApiV1MePrefsGEnum = "US_CT"
	PatchApiV1MePrefsGEnumUS_ME PatchApiV1MePrefsGEnum = "US_ME"
	PatchApiV1MePrefsGEnumUS_MD PatchApiV1MePrefsGEnum = "US_MD"
	PatchApiV1MePrefsGEnumUS_MA PatchApiV1MePrefsGEnum = "US_MA"
	PatchApiV1MePrefsGEnumUS_OH PatchApiV1MePrefsGEnum = "US_OH"
	PatchApiV1MePrefsGEnumUS_UT PatchApiV1MePrefsGEnum = "US_UT"
	PatchApiV1MePrefsGEnumUS_MO PatchApiV1MePrefsGEnum = "US_MO"
	PatchApiV1MePrefsGEnumUS_MN PatchApiV1MePrefsGEnum = "US_MN"
	PatchApiV1MePrefsGEnumUS_MI PatchApiV1MePrefsGEnum = "US_MI"
	PatchApiV1MePrefsGEnumUS_RI PatchApiV1MePrefsGEnum = "US_RI"
	PatchApiV1MePrefsGEnumUS_KS PatchApiV1MePrefsGEnum = "US_KS"
	PatchApiV1MePrefsGEnumUS_MT PatchApiV1MePrefsGEnum = "US_MT"
	PatchApiV1MePrefsGEnumUS_MS PatchApiV1MePrefsGEnum = "US_MS"
	PatchApiV1MePrefsGEnumUS_SC PatchApiV1MePrefsGEnum = "US_SC"
	PatchApiV1MePrefsGEnumUS_KY PatchApiV1MePrefsGEnum = "US_KY"
	PatchApiV1MePrefsGEnumUS_OR PatchApiV1MePrefsGEnum = "US_OR"
	PatchApiV1MePrefsGEnumUS_SD PatchApiV1MePrefsGEnum = "US_SD"
)

type PatchApiV1MePrefsMediaEnum string

const (
	PatchApiV1MePrefsMediaEnumOn PatchApiV1MePrefsMediaEnum = "on"
	PatchApiV1MePrefsMediaEnumOff PatchApiV1MePrefsMediaEnum = "off"
	PatchApiV1MePrefsMediaEnumSubreddit PatchApiV1MePrefsMediaEnum = "subreddit"
)

type PatchApiV1MePrefsMediaPreviewEnum string

const (
	PatchApiV1MePrefsMediaPreviewEnumOn PatchApiV1MePrefsMediaPreviewEnum = "on"
	PatchApiV1MePrefsMediaPreviewEnumOff PatchApiV1MePrefsMediaPreviewEnum = "off"
	PatchApiV1MePrefsMediaPreviewEnumSubreddit PatchApiV1MePrefsMediaPreviewEnum = "subreddit"
)

/*
PatchApiV1MePrefs makes a PATCH request to /api/v1/me/prefs
ID: PATCH /api/v1/me/prefs
Description: No description available
*/
func (sdk *ReddigoSDK) PatchApiV1MePrefs(acceptPms string, activityRelevantAds bool, allowClicktracking bool, badCommentAutocollapse string, beta bool, clickgadget bool, collapseReadMessages bool, compress bool, countryCode string, defaultCommentSort string, domainDetails bool, emailChatRequest bool, emailCommentReply bool, emailCommunityDiscovery bool, emailDigests bool, emailMessages bool, emailNewUserWelcome bool, emailPostReply bool, emailPrivateMessage bool, emailUnsubscribeAll bool, emailUpvoteComment bool, emailUpvotePost bool, emailUserNewFollower bool, emailUsernameMention bool, enableDefaultThemes bool, enableFollowers bool, enableRedditProAnalyticsEmails bool, feedRecommendationsEnabled bool, g string, hideAds bool, hideDowns bool, hideFromRobots bool, hideUps bool, highlightControversial bool, highlightNewComments bool, ignoreSuggestedSort bool, inRedesignBeta bool, labelNsfw bool, lang interface{}, legacySearch bool, liveBarRecommendationsEnabled bool, liveOrangereds bool, markMessagesRead bool, media string, mediaPreview string, minCommentScore int, minLinkScore int, monitorMentions bool, newwindow bool, nightmode bool, noProfanity bool, numComments int, numsites int, over18 bool, privateFeeds bool, profileOptOut bool, publicVotes bool, research bool, searchIncludeOver18 bool, sendCrosspostMessages bool, sendWelcomeMessages bool, showFlair bool, showGoldExpiration bool, showLinkFlair bool, showLocationBasedRecommendations bool, showPresence bool, showStylesheets bool, showTrending bool, showTwitter bool, smsNotificationsEnabled bool, storeVisits bool, surveyLastSeenTime int, themeSelector interface{}, thirdPartyDataPersonalizedAds bool, thirdPartyPersonalizedAds bool, thirdPartySiteDataPersonalizedAds bool, thirdPartySiteDataPersonalizedContent bool, threadedMessages bool, threadedModmail bool, topKarmaSubreddits bool, useGlobalDefaults bool, videoAutoplay bool, whatsappCommentReply bool, whatsappEnabled bool) (PatchApiV1MePrefsResponse, error) {
	url := fmt.Sprintf("/api/v1/me/prefs")
	payload := map[string]interface{}{
		"accept_pms": acceptPms,
		"activity_relevant_ads": activityRelevantAds,
		"allow_clicktracking": allowClicktracking,
		"bad_comment_autocollapse": badCommentAutocollapse,
		"beta": beta,
		"clickgadget": clickgadget,
		"collapse_read_messages": collapseReadMessages,
		"compress": compress,
		"country_code": countryCode,
		"default_comment_sort": defaultCommentSort,
		"domain_details": domainDetails,
		"email_chat_request": emailChatRequest,
		"email_comment_reply": emailCommentReply,
		"email_community_discovery": emailCommunityDiscovery,
		"email_digests": emailDigests,
		"email_messages": emailMessages,
		"email_new_user_welcome": emailNewUserWelcome,
		"email_post_reply": emailPostReply,
		"email_private_message": emailPrivateMessage,
		"email_unsubscribe_all": emailUnsubscribeAll,
		"email_upvote_comment": emailUpvoteComment,
		"email_upvote_post": emailUpvotePost,
		"email_user_new_follower": emailUserNewFollower,
		"email_username_mention": emailUsernameMention,
		"enable_default_themes": enableDefaultThemes,
		"enable_followers": enableFollowers,
		"enable_reddit_pro_analytics_emails": enableRedditProAnalyticsEmails,
		"feed_recommendations_enabled": feedRecommendationsEnabled,
		"g": g,
		"hide_ads": hideAds,
		"hide_downs": hideDowns,
		"hide_from_robots": hideFromRobots,
		"hide_ups": hideUps,
		"highlight_controversial": highlightControversial,
		"highlight_new_comments": highlightNewComments,
		"ignore_suggested_sort": ignoreSuggestedSort,
		"in_redesign_beta": inRedesignBeta,
		"label_nsfw": labelNsfw,
		"lang": lang,
		"legacy_search": legacySearch,
		"live_bar_recommendations_enabled": liveBarRecommendationsEnabled,
		"live_orangereds": liveOrangereds,
		"mark_messages_read": markMessagesRead,
		"media": media,
		"media_preview": mediaPreview,
		"min_comment_score": minCommentScore,
		"min_link_score": minLinkScore,
		"monitor_mentions": monitorMentions,
		"newwindow": newwindow,
		"nightmode": nightmode,
		"no_profanity": noProfanity,
		"num_comments": numComments,
		"numsites": numsites,
		"over_18": over18,
		"private_feeds": privateFeeds,
		"profile_opt_out": profileOptOut,
		"public_votes": publicVotes,
		"research": research,
		"search_include_over_18": searchIncludeOver18,
		"send_crosspost_messages": sendCrosspostMessages,
		"send_welcome_messages": sendWelcomeMessages,
		"show_flair": showFlair,
		"show_gold_expiration": showGoldExpiration,
		"show_link_flair": showLinkFlair,
		"show_location_based_recommendations": showLocationBasedRecommendations,
		"show_presence": showPresence,
		"show_stylesheets": showStylesheets,
		"show_trending": showTrending,
		"show_twitter": showTwitter,
		"sms_notifications_enabled": smsNotificationsEnabled,
		"store_visits": storeVisits,
		"survey_last_seen_time": surveyLastSeenTime,
		"theme_selector": themeSelector,
		"third_party_data_personalized_ads": thirdPartyDataPersonalizedAds,
		"third_party_personalized_ads": thirdPartyPersonalizedAds,
		"third_party_site_data_personalized_ads": thirdPartySiteDataPersonalizedAds,
		"third_party_site_data_personalized_content": thirdPartySiteDataPersonalizedContent,
		"threaded_messages": threadedMessages,
		"threaded_modmail": threadedModmail,
		"top_karma_subreddits": topKarmaSubreddits,
		"use_global_defaults": useGlobalDefaults,
		"video_autoplay": videoAutoplay,
		"whatsapp_comment_reply": whatsappCommentReply,
		"whatsapp_enabled": whatsappEnabled,
	}
	// Construct the request for PATCH method
	url := fmt.Sprintf("/api/v1/me/prefs")
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return PatchApiV1MePrefsResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Body = io.NopCloser(bytes.NewBuffer(jsonPayload))
	req, err := sdk.MakeRequest("PATCH", url, nil)
	if err != nil {
		return PatchApiV1MePrefsResponse{}, err
	}
	defer resp.Body.Close()
	var response PatchApiV1MePrefsResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return PatchApiV1MePrefsResponse{}, err
	}
	return response, nil
}



/*
GetApiV1MeTrophies makes a GET request to /api/v1/me/trophies
ID: GET /api/v1/me/trophies
Description: Return a list of trophies for the current user.
*/
func (sdk *ReddigoSDK) GetApiV1MeTrophies() (GetApiV1MeTrophiesResponse, error) {
	url := fmt.Sprintf("/api/v1/me/trophies")
	// Construct the request for GET method
	url := fmt.Sprintf("/api/v1/me/trophies")
	req, err := sdk.MakeRequest("GET", url, nil)
	if err != nil {
		return GetApiV1MeTrophiesResponse{}, err
	}
	defer resp.Body.Close()
	var response GetApiV1MeTrophiesResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return GetApiV1MeTrophiesResponse{}, err
	}
	return response, nil
}



// GetPrefsWhereResponse represents the response for GET /prefs/{where}
type GetPrefsWhereResponse struct {
	After string `json:"after"` // fullname of a thing
	Before string `json:"before"` // fullname of a thing
	Count int `json:"count"` // a positive integer (default: 0)
	Limit interface{} `json:"limit"` // the maximum number of items desired (default: 25, maximum: 100)
	Show string `json:"show"` // (optional) the string all
	SrDetail bool `json:"sr_detail"` // (optional) expand subreddits
}

/*
GetPrefsWhere makes a GET request to /prefs/{where}
ID: GET /prefs/{where}
Description: This endpoint is a listing.
*/
func (sdk *ReddigoSDK) GetPrefsWhere(where string, where string, after string, before string, count string, limit string) (GetPrefsWhereResponse, error) {
	url := fmt.Sprintf("/prefs/%s", where)
	queryParams := url.Values{}
	queryParams.Add("after", after)
	queryParams.Add("before", before)
	queryParams.Add("count", count)
	queryParams.Add("limit", limit)
	url += "?" + queryParams.Encode()
	// Construct the request for GET method
	url := fmt.Sprintf("/prefs/%s", where)
	req, err := sdk.MakeRequest("GET", url, nil)
	if err != nil {
		return GetPrefsWhereResponse{}, err
	}
	defer resp.Body.Close()
	var response GetPrefsWhereResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return GetPrefsWhereResponse{}, err
	}
	return response, nil
}



/*
GetApiNeedsCaptcha makes a GET request to /api/needs_captcha
ID: GET /api/needs_captcha
Description: Check whether ReCAPTCHAs are needed for API methods
*/
func (sdk *ReddigoSDK) GetApiNeedsCaptcha() (GetApiNeedsCaptchaResponse, error) {
	url := fmt.Sprintf("/api/needs_captcha")
	// Construct the request for GET method
	url := fmt.Sprintf("/api/needs_captcha")
	req, err := sdk.MakeRequest("GET", url, nil)
	if err != nil {
		return GetApiNeedsCaptchaResponse{}, err
	}
	defer resp.Body.Close()
	var response GetApiNeedsCaptchaResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return GetApiNeedsCaptchaResponse{}, err
	}
	return response, nil
}



/*
PostApiV1CollectionsAddPostToCollection makes a POST request to /api/v1/collections/add_post_to_collection
ID: POST /api/v1/collections/add_post_to_collection
Description: Add a post to a collection
*/
func (sdk *ReddigoSDK) PostApiV1CollectionsAddPostToCollection(collectionId interface{}, linkFullname string) (PostApiV1CollectionsAddPostToCollectionResponse, error) {
	url := fmt.Sprintf("/api/v1/collections/add_post_to_collection")
	payload := map[string]interface{}{
		"collection_id": collectionId,
		"link_fullname": linkFullname,
	}
	// Construct the request for POST method
	url := fmt.Sprintf("/api/v1/collections/add_post_to_collection")
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return PostApiV1CollectionsAddPostToCollectionResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Body = io.NopCloser(bytes.NewBuffer(jsonPayload))
	req, err := sdk.MakeRequest("POST", url, nil)
	if err != nil {
		return PostApiV1CollectionsAddPostToCollectionResponse{}, err
	}
	defer resp.Body.Close()
	var response PostApiV1CollectionsAddPostToCollectionResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return PostApiV1CollectionsAddPostToCollectionResponse{}, err
	}
	return response, nil
}



// GetApiV1CollectionsCollectionResponse represents the response for GET /api/v1/collections/collection
type GetApiV1CollectionsCollectionResponse struct {
	CollectionId interface{} `json:"collection_id"` // the UUID of a collection
	IncludeLinks bool `json:"include_links"` // boolean value
}

/*
GetApiV1CollectionsCollection makes a GET request to /api/v1/collections/collection
ID: GET /api/v1/collections/collection
Description: Fetch a collection including all the links
*/
func (sdk *ReddigoSDK) GetApiV1CollectionsCollection() (GetApiV1CollectionsCollectionResponse, error) {
	url := fmt.Sprintf("/api/v1/collections/collection")
	// Construct the request for GET method
	url := fmt.Sprintf("/api/v1/collections/collection")
	req, err := sdk.MakeRequest("GET", url, nil)
	if err != nil {
		return GetApiV1CollectionsCollectionResponse{}, err
	}
	defer resp.Body.Close()
	var response GetApiV1CollectionsCollectionResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return GetApiV1CollectionsCollectionResponse{}, err
	}
	return response, nil
}



type PostApiV1CollectionsCreateCollectionDisplayLayoutEnum string

const (
	PostApiV1CollectionsCreateCollectionDisplayLayoutEnumTIMELINE PostApiV1CollectionsCreateCollectionDisplayLayoutEnum = "TIMELINE"
	PostApiV1CollectionsCreateCollectionDisplayLayoutEnumGALLERY PostApiV1CollectionsCreateCollectionDisplayLayoutEnum = "GALLERY"
)

/*
PostApiV1CollectionsCreateCollection makes a POST request to /api/v1/collections/create_collection
ID: POST /api/v1/collections/create_collection
Description: Create a collection
*/
func (sdk *ReddigoSDK) PostApiV1CollectionsCreateCollection(description string, displayLayout string, srFullname string, title string) (PostApiV1CollectionsCreateCollectionResponse, error) {
	url := fmt.Sprintf("/api/v1/collections/create_collection")
	payload := map[string]interface{}{
		"description": description,
		"display_layout": displayLayout,
		"sr_fullname": srFullname,
		"title": title,
	}
	// Construct the request for POST method
	url := fmt.Sprintf("/api/v1/collections/create_collection")
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return PostApiV1CollectionsCreateCollectionResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Body = io.NopCloser(bytes.NewBuffer(jsonPayload))
	req, err := sdk.MakeRequest("POST", url, nil)
	if err != nil {
		return PostApiV1CollectionsCreateCollectionResponse{}, err
	}
	defer resp.Body.Close()
	var response PostApiV1CollectionsCreateCollectionResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return PostApiV1CollectionsCreateCollectionResponse{}, err
	}
	return response, nil
}



/*
PostApiV1CollectionsDeleteCollection makes a POST request to /api/v1/collections/delete_collection
ID: POST /api/v1/collections/delete_collection
Description: Delete a collection
*/
func (sdk *ReddigoSDK) PostApiV1CollectionsDeleteCollection(collectionId interface{}) (PostApiV1CollectionsDeleteCollectionResponse, error) {
	url := fmt.Sprintf("/api/v1/collections/delete_collection")
	payload := map[string]interface{}{
		"collection_id": collectionId,
	}
	// Construct the request for POST method
	url := fmt.Sprintf("/api/v1/collections/delete_collection")
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return PostApiV1CollectionsDeleteCollectionResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Body = io.NopCloser(bytes.NewBuffer(jsonPayload))
	req, err := sdk.MakeRequest("POST", url, nil)
	if err != nil {
		return PostApiV1CollectionsDeleteCollectionResponse{}, err
	}
	defer resp.Body.Close()
	var response PostApiV1CollectionsDeleteCollectionResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return PostApiV1CollectionsDeleteCollectionResponse{}, err
	}
	return response, nil
}



