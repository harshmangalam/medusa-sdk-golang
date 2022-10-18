package regions

import "github.com/harshmngalam/medusa-sdk-golang/common"

type RegionQuery struct {
	Offset    int                    `json:"offset,omitempty" url:"offset,omitempty"`
	Limit     int                    `json:"limit,omitempty" url:"limit,omitempty"`
	CreatedAt *common.DateComparison `json:"created_at,omitempty" url:"created_at,omitempty"`
	UpdatedAt *common.DateComparison `json:"updated_at,omitempty" url:"updated_at,omitempty"`
}
