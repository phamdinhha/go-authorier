package models

type Policy struct {
	Subject string `json:"subject"`
	Object  string `json:"object"`
	Action  string `json:"action"`
}

type UpdatePolicyReq struct {
	OldPolicy Policy `json:"old_policy"`
	NewPolicy Policy `json:"new_policy"`
}
