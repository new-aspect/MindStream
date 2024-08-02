```
type ServerError struct {
Code    int
Message string
}

type ErrorResponse struct {
StatusCode    int         `json:"statusCode"`
StatusMessage string      `json:"statusMessage"`
Succeed       bool        `json:"succeed"`
Data          interface{} `json:"data"`
}

func getServerError(err string) ServerError {
code, exist := Codes[err]

	if !exist {
		err = "Bad Request"
		code = 40000
	}

	return ServerError{
		Code:    code,
		Message: err,
	}
}

func ErrorHandler(w http.ResponseWriter, err string) {
serverError := getServerError(err)

	res := ErrorResponse{
		StatusCode:    serverError.Code,
		StatusMessage: serverError.Message,
		Succeed:       false,
		Data:          nil,
	}

	statusCode := int(serverError.Code / 100)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}
code.go
var Codes = map[string]int{
"NOT_AUTH": 20001,

	"REQUEST_BODY_ERROR": 40001,

	"DATABASE_ERROR": 50001,
}
```
设计的好处
1 细化错误的分类
通过定义具体的错误编码（如`20001`表示"NOT_AUTH"）,开发者可以清晰地标识和处理不同类型的错误，而不仅仅是400或500系列的错误代码

2 更清晰的错误信息
使用明确的错误信息Message字段，客户可以获得更加人性化和易于理解的错误说明。不仅提高了用户体验，也有助于开发者快速定位

3 统一的的错误响应结构
ErrorResponse结构体提供了一个统一的错误响应格式，确保所有错误的响应一致性。这种一致性有助于前端与后端的解耦，前端可以预期响应的格式，从而简化错误处理逻辑

4 防止信息泄露
通过控制返回的Message信息，可以避免在响应中直接暴露内部错误信息或铭感信息，从而提高系统的安全性

5 扩展性
这种设计可以使得扩展错误处理变得简单，只需要向Code字典里面添加新的错误码和描述，就可以支持新的错误类型，而不需要修改核心的错误逻辑