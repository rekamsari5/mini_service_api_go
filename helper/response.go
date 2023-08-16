package helper

type Response struct {
	ResponseCode   string      `json:"responseCode"`
	ResponseDesc   string      `json:"responseDesc"`
	ResponseData   interface{} `json:"responseData"`
	ResponseErrors interface{} `json:"responseErrors"`
}

func APIResponseSuccess(responseDesc string, responseData interface{}) Response {
	jsonResponse := Response{
		ResponseCode: "00",
		ResponseDesc: responseDesc,
		ResponseData: responseData,
	}

	return jsonResponse
}

func APIResponseError(responseDesc string, responseCode string, responseErrors interface{}) Response {
	jsonResponse := Response{
		ResponseCode:   responseCode,
		ResponseDesc:   responseDesc,
		ResponseErrors: responseErrors,
	}

	return jsonResponse
}
