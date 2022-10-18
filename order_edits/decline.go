package orderedits

type DeclineBody struct {
	DeclinedReason string `json:"declined_reason,omitempty"`
}

func NewDeclineBody() *DeclineBody {
	return new(DeclineBody)
}

func (d *DeclineBody) SetDeclineReason(reason string) *DeclineBody {
	d.DeclinedReason = reason
	return d
}
