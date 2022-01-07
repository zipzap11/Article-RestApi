package response

type StandardResponse struct {
	Code   int
	Status string
	Data   interface{}
}
