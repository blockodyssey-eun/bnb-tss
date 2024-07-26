package response

import "net/http"

// Error codes
const (
	ErrInvalidRequest = "ErrInvalidRequest"
	ErrInvalidNType   = "ErrInvalidNType"
	ErrInvalidMType   = "ErrInvalidMType"
	ErrNNotPositive   = "ErrNNotPositive"
	ErrMNotPositive   = "ErrMNotPositive"
	ErrMGreaterThanN  = "ErrMGreaterThanN"
	ErrKeyGeneration  = "ErrKeyGeneration"
	ErrKeyMarshalling = "ErrKeyMarshalling"
	ErrPodCreation    = "ErrPodCreation"
)

// Error code to HTTP status code mapping
var ErrorCodeToStatusCode = map[string]int{
	ErrInvalidRequest: http.StatusBadRequest,
	ErrNNotPositive:   http.StatusBadRequest,
	ErrMNotPositive:   http.StatusBadRequest,
	ErrMGreaterThanN:  http.StatusBadRequest,
	ErrKeyGeneration:  http.StatusInternalServerError,
	ErrKeyMarshalling: http.StatusInternalServerError,
	ErrPodCreation:    http.StatusInternalServerError,
}

// Error code to message mapping
var ErrorCodeToMessage = map[string]string{
	ErrInvalidRequest: "잘못된 요청 데이터입니다.",
	ErrNNotPositive:   "n은 양의 정수여야 합니다",
	ErrMNotPositive:   "m은 양의 정수여야 합니다",
	ErrMGreaterThanN:  "m은 n보다 클 수 없습니다",
	ErrKeyGeneration:  "키 생성 중 오류가 발생했습니다",
	ErrKeyMarshalling: "공개 키 변환 중 오류가 발생했습니다",
	ErrPodCreation:    "Pod 생성 중 오류가 발생했습니다",
}

// const (
// 	ErrMsgInvalidRequestBody           = "잘못된 요청 본문입니다"
// 	ErrMsgUnsupportedNetwork           = "지원되지 않는 네트워크입니다"
// 	ErrMsgDuplicateRequestID           = "중복된 요청 ID입니다"
// 	ErrMsgFailedRetrieveClientSecurity = "클라이언트 보안 정보를 가져오는데 실패했습니다"
// 	ErrMsgFailedSetupStreams           = "스트림 설정에 실패했습니다"
// 	ErrMsgFailedStartKeyGeneration     = "키 생성 시작에 실패했습니다"
// 	ErrMsgFailedDuringKeyGeneration    = "키 생성 중 실패했습니다"
// 	ErrMsgInvalidRequestID             = "유효하지 않은 요청 ID입니다"
// 	ErrMsgFailedConnectGRPC            = "gRPC 연결에 실패했습니다"
// 	ErrMsgInvalidSignRequest           = "서명 요청이 유효하지 않습니다"
// 	ErrMsgFailedStartSigning           = "서명 프로세스 시작에 실패했습니다"
// 	ErrMsgFailedDuringSigning          = "서명 프로세스 중 실패했습니다"
// )
