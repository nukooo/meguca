// Package types contains common shared types used throughout the project.
package types

// Board stores board metadata and the OPs of all threads
type Board struct {
	Ctr     int64    `json:"ctr"`
	Threads []Thread `json:"threads"`
}

// Thread is a transport/export wrapper that stores both the thread metada, its
// opening post data and its contained posts. The composite type itself is not
// stored in the database.
type Thread struct {
	Locked   bool  `json:"locked,omitempty" gorethink:"locked"`
	Archived bool  `json:"archived,omitempty" gorethink:"archived"`
	Sticky   bool  `json:"sticky,omitempty" gorethink:"sticky"`
	PostCtr  int16 `json:"postCtr" gorethink:"postCtr"`
	ImageCtr int16 `json:"imageCtr" gorethink:"imageCtr"`
	Post
	LogCtr    int64           `json:"logCtr" gorethink:"logCtr"`
	BumpTime  int64           `json:"bumpTime" gorethink:"bumpTime"`
	ReplyTime int64           `json:"replyTime" gorethink:"replyTime"`
	Subject   string          `json:"subject,omitempty" gorethink:"subject"`
	Posts     map[string]Post `json:"posts,omitempty" gorethink:"posts"`
}

// DatabaseThread is a template for wririting new threads to the database
type DatabaseThread struct {
	PostCtr   int16           `gorethink:"postCtr"`
	ImageCtr  int16           `gorethink:"imageCtr"`
	ID        int64           `gorethink:"id"`
	BumpTime  int64           `gorethink:"bumpTime"`
	ReplyTime int64           `gorethink:"replyTime"`
	Subject   string          `gorethink:"subject,omitempty"`
	Board     string          `gorethink:"board"`
	Posts     map[string]Post `gorethink:"posts"`
	Log       [][]byte        `gorethink:"log"`
}

// Post is a generic post. Either OP or reply.
type Post struct {
	Editing   bool    `json:"editing" gorethink:"editing"`
	Image     *Image  `json:"image,omitempty" gorethink:"image,omitempty"`
	OP        int64   `json:"op,omitempty" gorethink:"op"`
	ID        int64   `json:"id" gorethink:"id"`
	Time      int64   `json:"time" gorethink:"time"`
	Board     string  `json:"board" gorethink:"board"`
	IP        string  `json:"-" gorethink:"ip"`
	Nonce     string  `json:"-" gorethink:"nonce"`
	Body      string  `json:"body" gorethink:"body"`
	Name      string  `json:"name,omitempty" gorethink:"name,omitempty"`
	Trip      string  `json:"trip,omitempty" gorethink:"trip,omitempty"`
	Auth      string  `json:"auth,omitempty" gorethink:"auth,omitempty"`
	Email     string  `json:"email,omitempty" gorethink:"email,omitempty"`
	Backlinks LinkMap `json:"backlinks,omitempty" gorethink:"backlinks,omitempty"`
	Links     LinkMap `json:"links,omitempty" gorethink:"links,omitempty"`
}

// LinkMap contains a map of post numbers, this tread is linking, to
// corresponding Link structs
type LinkMap map[string]Link

// Link stores the target post's parent board and parent thread
type Link struct {
	Board string `json:"board" gorethink:"board"`
	OP    int    `json:"op" gorethink:"op"`
}
