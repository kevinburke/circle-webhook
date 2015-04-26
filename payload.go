package circle

import "time"

type Response struct {
	Payload Payload `json:"payload"`
}

type CommitSha string

type Payload struct {
	AllCommitDetails []CommitDetails `json:"all_commit_details"`
	AuthorDate       time.Time       `json:"author_date"`
	AuthorEmail      string          `json:"author_email"`
	AuthorName       string          `json:"author_name"`
	Body             string          `json:"body"`
	Branch           string          `json:"branch"`
	BuildNum         int             `json:"build_num"`
	BuildTimeMillis  int             `json:"build_time_millis"`
	BuildUrl         string          `json:"build_url"`
	Canceled         bool            `json:"canceled"`
	CircleYml        CircleYml       `json:"circle_yml"`
	CommitterDate    time.Time       `json:"committer_date"`
	CommitterEmail   string          `json:"committer_email"`
	CommitterName    string          `json:"committer_name"`
	Compare          string          `json:"compare"`
	// DontBuild null
	// Failed null
	// FeatureFlags null
	InfrastructureFail bool   `json:"infrastructure_fail"`
	IsFirstGreenBuild  bool   `json:"is_first_green_build"`
	Lifecycle          string `json:"lifecycle"`
	// Messages []
	// Node
	OSS                     bool      `json:"oss"`
	Outcome                 string    `json:"outcome"`
	Owners                  []string  `json:"owners"`
	Parallel                int16     `json:"parallel"`
	Previous                Build     `json:"previous"`
	PreviousSuccessfulBuild Build     `json:"previous_successful_build"`
	QueuedAt                time.Time `json:"queued_at"`
	Reponame                string    `json:"reponame"`
	// Retries null
	// RetryOf null
	// SSHEnabled null
	StartTime     time.Time   `json:"start_time"`
	Status        string      `json:"status"`
	Steps         []BuildStep `json:"steps"`
	StopTime      time.Time   `json:"stop_time"`
	Subject       string      `json:"subject"`
	Timedout      bool        `json:"timedout"`
	UsageQueuedAt time.Time   `json:"usage_queued_at"`
	User          User        `json:"user"`
	Username      string      `json:"username"`
	VcsRevision   CommitSha   `json:"vcs_revision"`
	VcsUrl        string      `json:"vcs_url"`
	Why           string      `json:"why"`
}

type CommitDetails struct {
	AuthorDate     time.Time `json:"author_date"`
	AuthorEmail    string    `json:"author_email"`
	AuthorLogin    string    `json:"author_login"`
	Body           string    `json:"body"`
	Branch         string    `json:"branch"`
	Commit         CommitSha `json:"commit"`
	CommitUrl      string    `json:"commit_url"`
	CommitterDate  time.Time `json:"committer_date"`
	CommitterEmail string    `json:"committer_email"`
	CommitterLogin string    `json:"committer_login"`
	CommitterName  string    `json:"committer_name"`
	Subject        string    `json:"subject"`
}

type CircleYml struct {
	String string `json:"string"`
}

type BuildStep struct {
	Actions []Action `json:"actions"`
	Name    string   `json:"name"`
}

type Action struct {
	// BashCommand null
	// Canceled null
	Command string `json:"command"`
	// Continue null
	EndTime  time.Time `json:"end_time"`
	ExitCode uint8
	// Failed null
	HasOutput bool `json:"has_output"`
	Index     int8 `json:"index"`
	// InfrastructureFail null
	// Messages []
	Name          string    `json:"name"`
	OutputUrl     string    `json:"output_url"`
	Parallel      bool      `json:"parallel"`
	RunTimeMillis int       `json:"run_time_millis"`
	Source        string    `json:"source"`
	StartTime     time.Time `json:"start_time"`
	Status        string    `json:"status"`
	Step          uint8     `json:"step"`
	// Timedout null
	Truncated bool   `json:"truncated"`
	Type      string `json:"type"`
}

type Build struct {
	BuildNum        int    `json:"build_num"`
	BuildTimeMillis int    `json:"build_time_millis"`
	Status          string `json:"status"`
}

type User struct {
	Email  string `json:"email"`
	IsUser bool   `json:"is_user"`
	Login  string `json:"login"`
	Name   string `json:"name"`
}
