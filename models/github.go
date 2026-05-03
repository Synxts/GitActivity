package models

type Event struct {
	Type    string  `json:"type"`
	Payload Payload `json:"payload"`
	Repo    Repo    `json:"repo"`
}

type Repo struct {
	Name string `json:"name"`
}

type Payload struct {
	Action  string   `json:"action"`
	RefType string   `json:"ref_type"`
	Commits []Commit `json:"commits"`
}

type Commit struct {
	SHA string `json:"sha"`
}
