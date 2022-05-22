package judger

type Result int

const (
	ResultSuccess               Result = 0 - iota
	ResultWrongAnswer                  = -1
	ResultCpuTimeLimitExceeded         = 1
	ResultRealTimeLimitExceeded        = 2
	ResultMemoryLimitExceeded          = 3
	ResultRuntimeError                 = 4
	ResultSystemError                  = 5
)

type Error int

const (
	ErrorSuccess           Error = 0 - iota
	ErrorInvalidConfig           = -1
	ErrorForkFailed              = -2
	ErrorPthreadFailed           = -3
	ErrorWaitFailed              = -4
	ErrorRootRequired            = -5
	ErrorLoadSeccompFailed       = -6
	ErrorSetrlimitFailed         = -7
	ErrorDup2Failed              = -8
	ErrorSetuidFailed            = -9
	ErrorExecveFailed            = -10
	ErrorSpjError                = -11
)

const (
	IOModeStandard = "Standard IO"
	IOModeFile     = "File IO"
)
