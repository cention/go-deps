package gohubspot

type IdentityProfile struct {
	Vid        int        `json:"vid"`
	Properties Properties `json:"properties"`
}
