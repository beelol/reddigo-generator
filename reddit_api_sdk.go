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



type PatchApiV1MePrefsAccept_pmsEnum string

const (
	PatchApiV1MePrefsAccept_pmsEnumEveryone PatchApiV1MePrefsAccept_pmsEnum = "everyone"
	PatchApiV1MePrefsAccept_pmsEnumWhitelisted PatchApiV1MePrefsAccept_pmsEnum = "whitelisted"
)

type PatchApiV1MePrefsBad_comment_autocollapseEnum string

const (
	PatchApiV1MePrefsBad_comment_autocollapseEnumOff PatchApiV1MePrefsBad_comment_autocollapseEnum = "off"
	PatchApiV1MePrefsBad_comment_autocollapseEnumLow PatchApiV1MePrefsBad_comment_autocollapseEnum = "low"
	PatchApiV1MePrefsBad_comment_autocollapseEnumMedium PatchApiV1MePrefsBad_comment_autocollapseEnum = "medium"
	PatchApiV1MePrefsBad_comment_autocollapseEnumHigh PatchApiV1MePrefsBad_comment_autocollapseEnum = "high"
)

type PatchApiV1MePrefsCountry_codeEnum string

const (
	PatchApiV1MePrefsCountry_codeEnumWF PatchApiV1MePrefsCountry_codeEnum = "WF"
	PatchApiV1MePrefsCountry_codeEnumJP PatchApiV1MePrefsCountry_codeEnum = "JP"
	PatchApiV1MePrefsCountry_codeEnumJM PatchApiV1MePrefsCountry_codeEnum = "JM"
	PatchApiV1MePrefsCountry_codeEnumJO PatchApiV1MePrefsCountry_codeEnum = "JO"
	PatchApiV1MePrefsCountry_codeEnumWS PatchApiV1MePrefsCountry_codeEnum = "WS"
	PatchApiV1MePrefsCountry_codeEnumJE PatchApiV1MePrefsCountry_codeEnum = "JE"
	PatchApiV1MePrefsCountry_codeEnumGW PatchApiV1MePrefsCountry_codeEnum = "GW"
	PatchApiV1MePrefsCountry_codeEnumGU PatchApiV1MePrefsCountry_codeEnum = "GU"
	PatchApiV1MePrefsCountry_codeEnumGT PatchApiV1MePrefsCountry_codeEnum = "GT"
	PatchApiV1MePrefsCountry_codeEnumGS PatchApiV1MePrefsCountry_codeEnum = "GS"
	PatchApiV1MePrefsCountry_codeEnumGR PatchApiV1MePrefsCountry_codeEnum = "GR"
	PatchApiV1MePrefsCountry_codeEnumGQ PatchApiV1MePrefsCountry_codeEnum = "GQ"
	PatchApiV1MePrefsCountry_codeEnumGP PatchApiV1MePrefsCountry_codeEnum = "GP"
	PatchApiV1MePrefsCountry_codeEnumGY PatchApiV1MePrefsCountry_codeEnum = "GY"
	PatchApiV1MePrefsCountry_codeEnumGG PatchApiV1MePrefsCountry_codeEnum = "GG"
	PatchApiV1MePrefsCountry_codeEnumGF PatchApiV1MePrefsCountry_codeEnum = "GF"
	PatchApiV1MePrefsCountry_codeEnumGE PatchApiV1MePrefsCountry_codeEnum = "GE"
	PatchApiV1MePrefsCountry_codeEnumGD PatchApiV1MePrefsCountry_codeEnum = "GD"
	PatchApiV1MePrefsCountry_codeEnumGB PatchApiV1MePrefsCountry_codeEnum = "GB"
	PatchApiV1MePrefsCountry_codeEnumGA PatchApiV1MePrefsCountry_codeEnum = "GA"
	PatchApiV1MePrefsCountry_codeEnumGN PatchApiV1MePrefsCountry_codeEnum = "GN"
	PatchApiV1MePrefsCountry_codeEnumGM PatchApiV1MePrefsCountry_codeEnum = "GM"
	PatchApiV1MePrefsCountry_codeEnumGL PatchApiV1MePrefsCountry_codeEnum = "GL"
	PatchApiV1MePrefsCountry_codeEnumGI PatchApiV1MePrefsCountry_codeEnum = "GI"
	PatchApiV1MePrefsCountry_codeEnumGH PatchApiV1MePrefsCountry_codeEnum = "GH"
	PatchApiV1MePrefsCountry_codeEnumPR PatchApiV1MePrefsCountry_codeEnum = "PR"
	PatchApiV1MePrefsCountry_codeEnumPS PatchApiV1MePrefsCountry_codeEnum = "PS"
	PatchApiV1MePrefsCountry_codeEnumPW PatchApiV1MePrefsCountry_codeEnum = "PW"
	PatchApiV1MePrefsCountry_codeEnumPT PatchApiV1MePrefsCountry_codeEnum = "PT"
	PatchApiV1MePrefsCountry_codeEnumPY PatchApiV1MePrefsCountry_codeEnum = "PY"
	PatchApiV1MePrefsCountry_codeEnumPA PatchApiV1MePrefsCountry_codeEnum = "PA"
	PatchApiV1MePrefsCountry_codeEnumPF PatchApiV1MePrefsCountry_codeEnum = "PF"
	PatchApiV1MePrefsCountry_codeEnumPG PatchApiV1MePrefsCountry_codeEnum = "PG"
	PatchApiV1MePrefsCountry_codeEnumPE PatchApiV1MePrefsCountry_codeEnum = "PE"
	PatchApiV1MePrefsCountry_codeEnumPK PatchApiV1MePrefsCountry_codeEnum = "PK"
	PatchApiV1MePrefsCountry_codeEnumPH PatchApiV1MePrefsCountry_codeEnum = "PH"
	PatchApiV1MePrefsCountry_codeEnumPN PatchApiV1MePrefsCountry_codeEnum = "PN"
	PatchApiV1MePrefsCountry_codeEnumPL PatchApiV1MePrefsCountry_codeEnum = "PL"
	PatchApiV1MePrefsCountry_codeEnumPM PatchApiV1MePrefsCountry_codeEnum = "PM"
	PatchApiV1MePrefsCountry_codeEnumZM PatchApiV1MePrefsCountry_codeEnum = "ZM"
	PatchApiV1MePrefsCountry_codeEnumZA PatchApiV1MePrefsCountry_codeEnum = "ZA"
	PatchApiV1MePrefsCountry_codeEnumZZ PatchApiV1MePrefsCountry_codeEnum = "ZZ"
	PatchApiV1MePrefsCountry_codeEnumZW PatchApiV1MePrefsCountry_codeEnum = "ZW"
	PatchApiV1MePrefsCountry_codeEnumME PatchApiV1MePrefsCountry_codeEnum = "ME"
	PatchApiV1MePrefsCountry_codeEnumMD PatchApiV1MePrefsCountry_codeEnum = "MD"
	PatchApiV1MePrefsCountry_codeEnumMG PatchApiV1MePrefsCountry_codeEnum = "MG"
	PatchApiV1MePrefsCountry_codeEnumMF PatchApiV1MePrefsCountry_codeEnum = "MF"
	PatchApiV1MePrefsCountry_codeEnumMA PatchApiV1MePrefsCountry_codeEnum = "MA"
	PatchApiV1MePrefsCountry_codeEnumMC PatchApiV1MePrefsCountry_codeEnum = "MC"
	PatchApiV1MePrefsCountry_codeEnumMM PatchApiV1MePrefsCountry_codeEnum = "MM"
	PatchApiV1MePrefsCountry_codeEnumML PatchApiV1MePrefsCountry_codeEnum = "ML"
	PatchApiV1MePrefsCountry_codeEnumMO PatchApiV1MePrefsCountry_codeEnum = "MO"
	PatchApiV1MePrefsCountry_codeEnumMN PatchApiV1MePrefsCountry_codeEnum = "MN"
	PatchApiV1MePrefsCountry_codeEnumMH PatchApiV1MePrefsCountry_codeEnum = "MH"
	PatchApiV1MePrefsCountry_codeEnumMK PatchApiV1MePrefsCountry_codeEnum = "MK"
	PatchApiV1MePrefsCountry_codeEnumMU PatchApiV1MePrefsCountry_codeEnum = "MU"
	PatchApiV1MePrefsCountry_codeEnumMT PatchApiV1MePrefsCountry_codeEnum = "MT"
	PatchApiV1MePrefsCountry_codeEnumMW PatchApiV1MePrefsCountry_codeEnum = "MW"
	PatchApiV1MePrefsCountry_codeEnumMV PatchApiV1MePrefsCountry_codeEnum = "MV"
	PatchApiV1MePrefsCountry_codeEnumMQ PatchApiV1MePrefsCountry_codeEnum = "MQ"
	PatchApiV1MePrefsCountry_codeEnumMP PatchApiV1MePrefsCountry_codeEnum = "MP"
	PatchApiV1MePrefsCountry_codeEnumMS PatchApiV1MePrefsCountry_codeEnum = "MS"
	PatchApiV1MePrefsCountry_codeEnumMR PatchApiV1MePrefsCountry_codeEnum = "MR"
	PatchApiV1MePrefsCountry_codeEnumMY PatchApiV1MePrefsCountry_codeEnum = "MY"
	PatchApiV1MePrefsCountry_codeEnumMX PatchApiV1MePrefsCountry_codeEnum = "MX"
	PatchApiV1MePrefsCountry_codeEnumMZ PatchApiV1MePrefsCountry_codeEnum = "MZ"
	PatchApiV1MePrefsCountry_codeEnumFR PatchApiV1MePrefsCountry_codeEnum = "FR"
	PatchApiV1MePrefsCountry_codeEnumFI PatchApiV1MePrefsCountry_codeEnum = "FI"
	PatchApiV1MePrefsCountry_codeEnumFJ PatchApiV1MePrefsCountry_codeEnum = "FJ"
	PatchApiV1MePrefsCountry_codeEnumFK PatchApiV1MePrefsCountry_codeEnum = "FK"
	PatchApiV1MePrefsCountry_codeEnumFM PatchApiV1MePrefsCountry_codeEnum = "FM"
	PatchApiV1MePrefsCountry_codeEnumFO PatchApiV1MePrefsCountry_codeEnum = "FO"
	PatchApiV1MePrefsCountry_codeEnumCK PatchApiV1MePrefsCountry_codeEnum = "CK"
	PatchApiV1MePrefsCountry_codeEnumCI PatchApiV1MePrefsCountry_codeEnum = "CI"
	PatchApiV1MePrefsCountry_codeEnumCH PatchApiV1MePrefsCountry_codeEnum = "CH"
	PatchApiV1MePrefsCountry_codeEnumCO PatchApiV1MePrefsCountry_codeEnum = "CO"
	PatchApiV1MePrefsCountry_codeEnumCN PatchApiV1MePrefsCountry_codeEnum = "CN"
	PatchApiV1MePrefsCountry_codeEnumCM PatchApiV1MePrefsCountry_codeEnum = "CM"
	PatchApiV1MePrefsCountry_codeEnumCL PatchApiV1MePrefsCountry_codeEnum = "CL"
	PatchApiV1MePrefsCountry_codeEnumCC PatchApiV1MePrefsCountry_codeEnum = "CC"
	PatchApiV1MePrefsCountry_codeEnumCA PatchApiV1MePrefsCountry_codeEnum = "CA"
	PatchApiV1MePrefsCountry_codeEnumCG PatchApiV1MePrefsCountry_codeEnum = "CG"
	PatchApiV1MePrefsCountry_codeEnumCF PatchApiV1MePrefsCountry_codeEnum = "CF"
	PatchApiV1MePrefsCountry_codeEnumCD PatchApiV1MePrefsCountry_codeEnum = "CD"
	PatchApiV1MePrefsCountry_codeEnumCZ PatchApiV1MePrefsCountry_codeEnum = "CZ"
	PatchApiV1MePrefsCountry_codeEnumCY PatchApiV1MePrefsCountry_codeEnum = "CY"
	PatchApiV1MePrefsCountry_codeEnumCX PatchApiV1MePrefsCountry_codeEnum = "CX"
	PatchApiV1MePrefsCountry_codeEnumCR PatchApiV1MePrefsCountry_codeEnum = "CR"
	PatchApiV1MePrefsCountry_codeEnumCW PatchApiV1MePrefsCountry_codeEnum = "CW"
	PatchApiV1MePrefsCountry_codeEnumCV PatchApiV1MePrefsCountry_codeEnum = "CV"
	PatchApiV1MePrefsCountry_codeEnumCU PatchApiV1MePrefsCountry_codeEnum = "CU"
	PatchApiV1MePrefsCountry_codeEnumSZ PatchApiV1MePrefsCountry_codeEnum = "SZ"
	PatchApiV1MePrefsCountry_codeEnumSY PatchApiV1MePrefsCountry_codeEnum = "SY"
	PatchApiV1MePrefsCountry_codeEnumSX PatchApiV1MePrefsCountry_codeEnum = "SX"
	PatchApiV1MePrefsCountry_codeEnumSS PatchApiV1MePrefsCountry_codeEnum = "SS"
	PatchApiV1MePrefsCountry_codeEnumSR PatchApiV1MePrefsCountry_codeEnum = "SR"
	PatchApiV1MePrefsCountry_codeEnumSV PatchApiV1MePrefsCountry_codeEnum = "SV"
	PatchApiV1MePrefsCountry_codeEnumST PatchApiV1MePrefsCountry_codeEnum = "ST"
	PatchApiV1MePrefsCountry_codeEnumSK PatchApiV1MePrefsCountry_codeEnum = "SK"
	PatchApiV1MePrefsCountry_codeEnumSJ PatchApiV1MePrefsCountry_codeEnum = "SJ"
	PatchApiV1MePrefsCountry_codeEnumSI PatchApiV1MePrefsCountry_codeEnum = "SI"
	PatchApiV1MePrefsCountry_codeEnumSH PatchApiV1MePrefsCountry_codeEnum = "SH"
	PatchApiV1MePrefsCountry_codeEnumSO PatchApiV1MePrefsCountry_codeEnum = "SO"
	PatchApiV1MePrefsCountry_codeEnumSN PatchApiV1MePrefsCountry_codeEnum = "SN"
	PatchApiV1MePrefsCountry_codeEnumSM PatchApiV1MePrefsCountry_codeEnum = "SM"
	PatchApiV1MePrefsCountry_codeEnumSL PatchApiV1MePrefsCountry_codeEnum = "SL"
	PatchApiV1MePrefsCountry_codeEnumSC PatchApiV1MePrefsCountry_codeEnum = "SC"
	PatchApiV1MePrefsCountry_codeEnumSB PatchApiV1MePrefsCountry_codeEnum = "SB"
	PatchApiV1MePrefsCountry_codeEnumSA PatchApiV1MePrefsCountry_codeEnum = "SA"
	PatchApiV1MePrefsCountry_codeEnumSG PatchApiV1MePrefsCountry_codeEnum = "SG"
	PatchApiV1MePrefsCountry_codeEnumSE PatchApiV1MePrefsCountry_codeEnum = "SE"
	PatchApiV1MePrefsCountry_codeEnumSD PatchApiV1MePrefsCountry_codeEnum = "SD"
	PatchApiV1MePrefsCountry_codeEnumYE PatchApiV1MePrefsCountry_codeEnum = "YE"
	PatchApiV1MePrefsCountry_codeEnumYT PatchApiV1MePrefsCountry_codeEnum = "YT"
	PatchApiV1MePrefsCountry_codeEnumLB PatchApiV1MePrefsCountry_codeEnum = "LB"
	PatchApiV1MePrefsCountry_codeEnumLC PatchApiV1MePrefsCountry_codeEnum = "LC"
	PatchApiV1MePrefsCountry_codeEnumLA PatchApiV1MePrefsCountry_codeEnum = "LA"
	PatchApiV1MePrefsCountry_codeEnumLK PatchApiV1MePrefsCountry_codeEnum = "LK"
	PatchApiV1MePrefsCountry_codeEnumLI PatchApiV1MePrefsCountry_codeEnum = "LI"
	PatchApiV1MePrefsCountry_codeEnumLV PatchApiV1MePrefsCountry_codeEnum = "LV"
	PatchApiV1MePrefsCountry_codeEnumLT PatchApiV1MePrefsCountry_codeEnum = "LT"
	PatchApiV1MePrefsCountry_codeEnumLU PatchApiV1MePrefsCountry_codeEnum = "LU"
	PatchApiV1MePrefsCountry_codeEnumLR PatchApiV1MePrefsCountry_codeEnum = "LR"
	PatchApiV1MePrefsCountry_codeEnumLS PatchApiV1MePrefsCountry_codeEnum = "LS"
	PatchApiV1MePrefsCountry_codeEnumLY PatchApiV1MePrefsCountry_codeEnum = "LY"
	PatchApiV1MePrefsCountry_codeEnumVA PatchApiV1MePrefsCountry_codeEnum = "VA"
	PatchApiV1MePrefsCountry_codeEnumVC PatchApiV1MePrefsCountry_codeEnum = "VC"
	PatchApiV1MePrefsCountry_codeEnumVE PatchApiV1MePrefsCountry_codeEnum = "VE"
	PatchApiV1MePrefsCountry_codeEnumVG PatchApiV1MePrefsCountry_codeEnum = "VG"
	PatchApiV1MePrefsCountry_codeEnumIQ PatchApiV1MePrefsCountry_codeEnum = "IQ"
	PatchApiV1MePrefsCountry_codeEnumVI PatchApiV1MePrefsCountry_codeEnum = "VI"
	PatchApiV1MePrefsCountry_codeEnumIS PatchApiV1MePrefsCountry_codeEnum = "IS"
	PatchApiV1MePrefsCountry_codeEnumIR PatchApiV1MePrefsCountry_codeEnum = "IR"
	PatchApiV1MePrefsCountry_codeEnumIT PatchApiV1MePrefsCountry_codeEnum = "IT"
	PatchApiV1MePrefsCountry_codeEnumVN PatchApiV1MePrefsCountry_codeEnum = "VN"
	PatchApiV1MePrefsCountry_codeEnumIM PatchApiV1MePrefsCountry_codeEnum = "IM"
	PatchApiV1MePrefsCountry_codeEnumIL PatchApiV1MePrefsCountry_codeEnum = "IL"
	PatchApiV1MePrefsCountry_codeEnumIO PatchApiV1MePrefsCountry_codeEnum = "IO"
	PatchApiV1MePrefsCountry_codeEnumIN PatchApiV1MePrefsCountry_codeEnum = "IN"
	PatchApiV1MePrefsCountry_codeEnumIE PatchApiV1MePrefsCountry_codeEnum = "IE"
	PatchApiV1MePrefsCountry_codeEnumID PatchApiV1MePrefsCountry_codeEnum = "ID"
	PatchApiV1MePrefsCountry_codeEnumBD PatchApiV1MePrefsCountry_codeEnum = "BD"
	PatchApiV1MePrefsCountry_codeEnumBE PatchApiV1MePrefsCountry_codeEnum = "BE"
	PatchApiV1MePrefsCountry_codeEnumBF PatchApiV1MePrefsCountry_codeEnum = "BF"
	PatchApiV1MePrefsCountry_codeEnumBG PatchApiV1MePrefsCountry_codeEnum = "BG"
	PatchApiV1MePrefsCountry_codeEnumBA PatchApiV1MePrefsCountry_codeEnum = "BA"
	PatchApiV1MePrefsCountry_codeEnumBB PatchApiV1MePrefsCountry_codeEnum = "BB"
	PatchApiV1MePrefsCountry_codeEnumBL PatchApiV1MePrefsCountry_codeEnum = "BL"
	PatchApiV1MePrefsCountry_codeEnumBM PatchApiV1MePrefsCountry_codeEnum = "BM"
	PatchApiV1MePrefsCountry_codeEnumBN PatchApiV1MePrefsCountry_codeEnum = "BN"
	PatchApiV1MePrefsCountry_codeEnumBO PatchApiV1MePrefsCountry_codeEnum = "BO"
	PatchApiV1MePrefsCountry_codeEnumBH PatchApiV1MePrefsCountry_codeEnum = "BH"
	PatchApiV1MePrefsCountry_codeEnumBI PatchApiV1MePrefsCountry_codeEnum = "BI"
	PatchApiV1MePrefsCountry_codeEnumBJ PatchApiV1MePrefsCountry_codeEnum = "BJ"
	PatchApiV1MePrefsCountry_codeEnumBT PatchApiV1MePrefsCountry_codeEnum = "BT"
	PatchApiV1MePrefsCountry_codeEnumBV PatchApiV1MePrefsCountry_codeEnum = "BV"
	PatchApiV1MePrefsCountry_codeEnumBW PatchApiV1MePrefsCountry_codeEnum = "BW"
	PatchApiV1MePrefsCountry_codeEnumBQ PatchApiV1MePrefsCountry_codeEnum = "BQ"
	PatchApiV1MePrefsCountry_codeEnumBR PatchApiV1MePrefsCountry_codeEnum = "BR"
	PatchApiV1MePrefsCountry_codeEnumBS PatchApiV1MePrefsCountry_codeEnum = "BS"
	PatchApiV1MePrefsCountry_codeEnumBY PatchApiV1MePrefsCountry_codeEnum = "BY"
	PatchApiV1MePrefsCountry_codeEnumBZ PatchApiV1MePrefsCountry_codeEnum = "BZ"
	PatchApiV1MePrefsCountry_codeEnumRU PatchApiV1MePrefsCountry_codeEnum = "RU"
	PatchApiV1MePrefsCountry_codeEnumRW PatchApiV1MePrefsCountry_codeEnum = "RW"
	PatchApiV1MePrefsCountry_codeEnumRS PatchApiV1MePrefsCountry_codeEnum = "RS"
	PatchApiV1MePrefsCountry_codeEnumRE PatchApiV1MePrefsCountry_codeEnum = "RE"
	PatchApiV1MePrefsCountry_codeEnumRO PatchApiV1MePrefsCountry_codeEnum = "RO"
	PatchApiV1MePrefsCountry_codeEnumOM PatchApiV1MePrefsCountry_codeEnum = "OM"
	PatchApiV1MePrefsCountry_codeEnumHR PatchApiV1MePrefsCountry_codeEnum = "HR"
	PatchApiV1MePrefsCountry_codeEnumHT PatchApiV1MePrefsCountry_codeEnum = "HT"
	PatchApiV1MePrefsCountry_codeEnumHU PatchApiV1MePrefsCountry_codeEnum = "HU"
	PatchApiV1MePrefsCountry_codeEnumHK PatchApiV1MePrefsCountry_codeEnum = "HK"
	PatchApiV1MePrefsCountry_codeEnumHN PatchApiV1MePrefsCountry_codeEnum = "HN"
	PatchApiV1MePrefsCountry_codeEnumHM PatchApiV1MePrefsCountry_codeEnum = "HM"
	PatchApiV1MePrefsCountry_codeEnumEH PatchApiV1MePrefsCountry_codeEnum = "EH"
	PatchApiV1MePrefsCountry_codeEnumEE PatchApiV1MePrefsCountry_codeEnum = "EE"
	PatchApiV1MePrefsCountry_codeEnumEG PatchApiV1MePrefsCountry_codeEnum = "EG"
	PatchApiV1MePrefsCountry_codeEnumEC PatchApiV1MePrefsCountry_codeEnum = "EC"
	PatchApiV1MePrefsCountry_codeEnumET PatchApiV1MePrefsCountry_codeEnum = "ET"
	PatchApiV1MePrefsCountry_codeEnumES PatchApiV1MePrefsCountry_codeEnum = "ES"
	PatchApiV1MePrefsCountry_codeEnumER PatchApiV1MePrefsCountry_codeEnum = "ER"
	PatchApiV1MePrefsCountry_codeEnumUY PatchApiV1MePrefsCountry_codeEnum = "UY"
	PatchApiV1MePrefsCountry_codeEnumUZ PatchApiV1MePrefsCountry_codeEnum = "UZ"
	PatchApiV1MePrefsCountry_codeEnumUS PatchApiV1MePrefsCountry_codeEnum = "US"
	PatchApiV1MePrefsCountry_codeEnumUM PatchApiV1MePrefsCountry_codeEnum = "UM"
	PatchApiV1MePrefsCountry_codeEnumUG PatchApiV1MePrefsCountry_codeEnum = "UG"
	PatchApiV1MePrefsCountry_codeEnumUA PatchApiV1MePrefsCountry_codeEnum = "UA"
	PatchApiV1MePrefsCountry_codeEnumVU PatchApiV1MePrefsCountry_codeEnum = "VU"
	PatchApiV1MePrefsCountry_codeEnumNI PatchApiV1MePrefsCountry_codeEnum = "NI"
	PatchApiV1MePrefsCountry_codeEnumNL PatchApiV1MePrefsCountry_codeEnum = "NL"
	PatchApiV1MePrefsCountry_codeEnumNO PatchApiV1MePrefsCountry_codeEnum = "NO"
	PatchApiV1MePrefsCountry_codeEnumNA PatchApiV1MePrefsCountry_codeEnum = "NA"
	PatchApiV1MePrefsCountry_codeEnumNC PatchApiV1MePrefsCountry_codeEnum = "NC"
	PatchApiV1MePrefsCountry_codeEnumNE PatchApiV1MePrefsCountry_codeEnum = "NE"
	PatchApiV1MePrefsCountry_codeEnumNF PatchApiV1MePrefsCountry_codeEnum = "NF"
	PatchApiV1MePrefsCountry_codeEnumNG PatchApiV1MePrefsCountry_codeEnum = "NG"
	PatchApiV1MePrefsCountry_codeEnumNZ PatchApiV1MePrefsCountry_codeEnum = "NZ"
	PatchApiV1MePrefsCountry_codeEnumNP PatchApiV1MePrefsCountry_codeEnum = "NP"
	PatchApiV1MePrefsCountry_codeEnumNR PatchApiV1MePrefsCountry_codeEnum = "NR"
	PatchApiV1MePrefsCountry_codeEnumNU PatchApiV1MePrefsCountry_codeEnum = "NU"
	PatchApiV1MePrefsCountry_codeEnumXK PatchApiV1MePrefsCountry_codeEnum = "XK"
	PatchApiV1MePrefsCountry_codeEnumXZ PatchApiV1MePrefsCountry_codeEnum = "XZ"
	PatchApiV1MePrefsCountry_codeEnumXX PatchApiV1MePrefsCountry_codeEnum = "XX"
	PatchApiV1MePrefsCountry_codeEnumKG PatchApiV1MePrefsCountry_codeEnum = "KG"
	PatchApiV1MePrefsCountry_codeEnumKE PatchApiV1MePrefsCountry_codeEnum = "KE"
	PatchApiV1MePrefsCountry_codeEnumKI PatchApiV1MePrefsCountry_codeEnum = "KI"
	PatchApiV1MePrefsCountry_codeEnumKH PatchApiV1MePrefsCountry_codeEnum = "KH"
	PatchApiV1MePrefsCountry_codeEnumKN PatchApiV1MePrefsCountry_codeEnum = "KN"
	PatchApiV1MePrefsCountry_codeEnumKM PatchApiV1MePrefsCountry_codeEnum = "KM"
	PatchApiV1MePrefsCountry_codeEnumKR PatchApiV1MePrefsCountry_codeEnum = "KR"
	PatchApiV1MePrefsCountry_codeEnumKP PatchApiV1MePrefsCountry_codeEnum = "KP"
	PatchApiV1MePrefsCountry_codeEnumKW PatchApiV1MePrefsCountry_codeEnum = "KW"
	PatchApiV1MePrefsCountry_codeEnumKZ PatchApiV1MePrefsCountry_codeEnum = "KZ"
	PatchApiV1MePrefsCountry_codeEnumKY PatchApiV1MePrefsCountry_codeEnum = "KY"
	PatchApiV1MePrefsCountry_codeEnumDO PatchApiV1MePrefsCountry_codeEnum = "DO"
	PatchApiV1MePrefsCountry_codeEnumDM PatchApiV1MePrefsCountry_codeEnum = "DM"
	PatchApiV1MePrefsCountry_codeEnumDJ PatchApiV1MePrefsCountry_codeEnum = "DJ"
	PatchApiV1MePrefsCountry_codeEnumDK PatchApiV1MePrefsCountry_codeEnum = "DK"
	PatchApiV1MePrefsCountry_codeEnumDE PatchApiV1MePrefsCountry_codeEnum = "DE"
	PatchApiV1MePrefsCountry_codeEnumDZ PatchApiV1MePrefsCountry_codeEnum = "DZ"
	PatchApiV1MePrefsCountry_codeEnumTZ PatchApiV1MePrefsCountry_codeEnum = "TZ"
	PatchApiV1MePrefsCountry_codeEnumTV PatchApiV1MePrefsCountry_codeEnum = "TV"
	PatchApiV1MePrefsCountry_codeEnumTW PatchApiV1MePrefsCountry_codeEnum = "TW"
	PatchApiV1MePrefsCountry_codeEnumTT PatchApiV1MePrefsCountry_codeEnum = "TT"
	PatchApiV1MePrefsCountry_codeEnumTR PatchApiV1MePrefsCountry_codeEnum = "TR"
	PatchApiV1MePrefsCountry_codeEnumTN PatchApiV1MePrefsCountry_codeEnum = "TN"
	PatchApiV1MePrefsCountry_codeEnumTO PatchApiV1MePrefsCountry_codeEnum = "TO"
	PatchApiV1MePrefsCountry_codeEnumTL PatchApiV1MePrefsCountry_codeEnum = "TL"
	PatchApiV1MePrefsCountry_codeEnumTM PatchApiV1MePrefsCountry_codeEnum = "TM"
	PatchApiV1MePrefsCountry_codeEnumTJ PatchApiV1MePrefsCountry_codeEnum = "TJ"
	PatchApiV1MePrefsCountry_codeEnumTK PatchApiV1MePrefsCountry_codeEnum = "TK"
	PatchApiV1MePrefsCountry_codeEnumTH PatchApiV1MePrefsCountry_codeEnum = "TH"
	PatchApiV1MePrefsCountry_codeEnumTF PatchApiV1MePrefsCountry_codeEnum = "TF"
	PatchApiV1MePrefsCountry_codeEnumTG PatchApiV1MePrefsCountry_codeEnum = "TG"
	PatchApiV1MePrefsCountry_codeEnumTD PatchApiV1MePrefsCountry_codeEnum = "TD"
	PatchApiV1MePrefsCountry_codeEnumTC PatchApiV1MePrefsCountry_codeEnum = "TC"
	PatchApiV1MePrefsCountry_codeEnumAE PatchApiV1MePrefsCountry_codeEnum = "AE"
	PatchApiV1MePrefsCountry_codeEnumAD PatchApiV1MePrefsCountry_codeEnum = "AD"
	PatchApiV1MePrefsCountry_codeEnumAG PatchApiV1MePrefsCountry_codeEnum = "AG"
	PatchApiV1MePrefsCountry_codeEnumAF PatchApiV1MePrefsCountry_codeEnum = "AF"
	PatchApiV1MePrefsCountry_codeEnumAI PatchApiV1MePrefsCountry_codeEnum = "AI"
	PatchApiV1MePrefsCountry_codeEnumAM PatchApiV1MePrefsCountry_codeEnum = "AM"
	PatchApiV1MePrefsCountry_codeEnumAL PatchApiV1MePrefsCountry_codeEnum = "AL"
	PatchApiV1MePrefsCountry_codeEnumAO PatchApiV1MePrefsCountry_codeEnum = "AO"
	PatchApiV1MePrefsCountry_codeEnumAN PatchApiV1MePrefsCountry_codeEnum = "AN"
	PatchApiV1MePrefsCountry_codeEnumAQ PatchApiV1MePrefsCountry_codeEnum = "AQ"
	PatchApiV1MePrefsCountry_codeEnumAS PatchApiV1MePrefsCountry_codeEnum = "AS"
	PatchApiV1MePrefsCountry_codeEnumAR PatchApiV1MePrefsCountry_codeEnum = "AR"
	PatchApiV1MePrefsCountry_codeEnumAU PatchApiV1MePrefsCountry_codeEnum = "AU"
	PatchApiV1MePrefsCountry_codeEnumAT PatchApiV1MePrefsCountry_codeEnum = "AT"
	PatchApiV1MePrefsCountry_codeEnumAW PatchApiV1MePrefsCountry_codeEnum = "AW"
	PatchApiV1MePrefsCountry_codeEnumAX PatchApiV1MePrefsCountry_codeEnum = "AX"
	PatchApiV1MePrefsCountry_codeEnumAZ PatchApiV1MePrefsCountry_codeEnum = "AZ"
	PatchApiV1MePrefsCountry_codeEnumQA PatchApiV1MePrefsCountry_codeEnum = "QA"
)

type PatchApiV1MePrefsDefault_comment_sortEnum string

const (
	PatchApiV1MePrefsDefault_comment_sortEnumConfidence PatchApiV1MePrefsDefault_comment_sortEnum = "confidence"
	PatchApiV1MePrefsDefault_comment_sortEnumTop PatchApiV1MePrefsDefault_comment_sortEnum = "top"
	PatchApiV1MePrefsDefault_comment_sortEnumNew PatchApiV1MePrefsDefault_comment_sortEnum = "new"
	PatchApiV1MePrefsDefault_comment_sortEnumControversial PatchApiV1MePrefsDefault_comment_sortEnum = "controversial"
	PatchApiV1MePrefsDefault_comment_sortEnumOld PatchApiV1MePrefsDefault_comment_sortEnum = "old"
	PatchApiV1MePrefsDefault_comment_sortEnumRandom PatchApiV1MePrefsDefault_comment_sortEnum = "random"
	PatchApiV1MePrefsDefault_comment_sortEnumQa PatchApiV1MePrefsDefault_comment_sortEnum = "qa"
	PatchApiV1MePrefsDefault_comment_sortEnumLive PatchApiV1MePrefsDefault_comment_sortEnum = "live"
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

type PatchApiV1MePrefsMedia_previewEnum string

const (
	PatchApiV1MePrefsMedia_previewEnumOn PatchApiV1MePrefsMedia_previewEnum = "on"
	PatchApiV1MePrefsMedia_previewEnumOff PatchApiV1MePrefsMedia_previewEnum = "off"
	PatchApiV1MePrefsMedia_previewEnumSubreddit PatchApiV1MePrefsMedia_previewEnum = "subreddit"
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



type PostApiV1CollectionsCreateCollectionDisplay_layoutEnum string

const (
	PostApiV1CollectionsCreateCollectionDisplay_layoutEnumTIMELINE PostApiV1CollectionsCreateCollectionDisplay_layoutEnum = "TIMELINE"
	PostApiV1CollectionsCreateCollectionDisplay_layoutEnumGALLERY PostApiV1CollectionsCreateCollectionDisplay_layoutEnum = "GALLERY"
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



