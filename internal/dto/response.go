package dto

type DataReply struct {
	Data interface{} `json:"data"`
} // @name Response

type ErrorMessage struct {
	Message string `json:"message"`
} // @name ErrorMessage
type ErrorReply struct {
	Error ErrorMessage `json:"error"`
} // @name ErrorReply

func ReplyError(message string) ErrorReply {
	return ErrorReply{
		Error: ErrorMessage{
			Message: message,
		},
	}
}
