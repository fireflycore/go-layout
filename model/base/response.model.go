package base

type ResponseEntity struct {
	Code    uint        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
