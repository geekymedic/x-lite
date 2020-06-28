package types

const (
	StateName          = "neon.bff.state"
	ResponseStatusCode = "neon.bff.response.status_code"
	ResponseErr        = "neon.bff.response.msg"
	ResponseBody       = "neon.bff.response.body"
	NeonSession        = "neon.bff.session"
)

// resolve cycle import
const (
	CodeSuccess     = 0
	CodeServerError = 1006
)
