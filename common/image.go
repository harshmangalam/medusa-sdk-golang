package common

type Image struct {
	Url       string         `json:"url"`
	Id        string         `json:"id,omitempty"`
	CreatedAt string         `json:"created_at,omitempty"`
	UpdatedAt string         `json:"updated_at,omitempty"`
	DeletedAt string         `json:"deleted_at,omitempty"`
	Metadata  map[string]any `json:"metadata,omitempty"`
}
