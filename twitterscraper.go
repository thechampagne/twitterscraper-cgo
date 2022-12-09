package main

/*
typedef struct {
   char* type;
   double*** coordinates;
} twitterscraper_bounding_box_t;

typedef struct {
   char* id;
   char* place_type;
   char* name;
   char* full_name;
   char* country_code;
   char* country;
   twitterscraper_bounding_box_t bounding_box;
} twitterscraper_place_t;

typedef struct {
   char*  avatar;
   char*  banner;
   char*  biography;
   char*  birthday;
   int    followers_count;
   int    following_count;
   int    friends_count;
   int    is_private; // bool
   int    is_verified; // bool
   char*  joined;
   int    likes_count;
   int    listed_count;
   char*  location;
   char*  name;
   char** pinned_tweet_ids;
   int    tweets_count;
   char*  url;
   char*  user_id;
   char*  username;
   char*  website;
} twitterscraper_profile_t;

typedef struct {
   char* id;
   char* preview;
   char* url;
} twitterscraper_video_t;
   
typedef struct twitterscraper_tweet_t twitterscraper_tweet_t;
   
struct twitterscraper_tweet_t {
   char** hashtags;
   char* html;
   char* id;
   twitterscraper_tweet_t* in_reply_to_status;
   is_quoted;  // bool
   is_pin;     // bool
   is_reply;   // bool
   is_retweet; // bool
   int likes;
   char* permanent_url;
   char** photos;
   twitterscraper_place_t* place;
   twitterscraper_tweet_t* quoted_status;
   int replies;
   int retweets;
   twitterscraper_tweet_t* retweeted_status;
   char* text;
   char* time_parsed;   // time.Time
   int64_t timestamp;
   char** urls;
   char* user_id;
   char* username;
   twitterscraper_video_t* video;
   int sensitive_content; // bool
};

typedef enum {
   TWITTERSCRAPER_SEARCH_TOP,
   TWITTERSCRAPER_SEARCH_LATEST,
   TWITTERSCRAPER_SEARCH_PHOTOS,
   TWITTERSCRAPER_SEARCH_VIDEOS,
   TWITTERSCRAPER_SEARCH_USERS
} twitterscraper_search_mode_t;

typedef struct {
   twitterscraper_profile_t profile;
   int is_err;
   char* err_msg;
} twitterscraper_profile_result_t;

typedef struct {
   twitterscraper_tweet_t tweet;
   int is_err;
   char* err_msg;
} twitterscraper_tweet_result_t;
*/
import "C"

func main() {}
