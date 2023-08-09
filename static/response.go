package static

func Response(status int, result string) map[string]interface{} {
	var message string

	switch status {
	case 200:
		message = "OK"
	case 201:
		message = "Created"
	case 400:
		message = "Bad Request"
	case 401:
		message = "Unauthorized"
	case 500:
		message = "Internal Server Error"
	case 501:
		message = "Bad Gateway"
	case 304:
		message = "Not Modified"
	default:
		message = ""
	}

	type Map map[string]interface{}
	res := Map{
		"status":      status,
		"message":     message,
		"description": result,
	}
	return res
}
