package judger

type ResultCode int

const (
	ResultSuccess               ResultCode = 0 - iota
	ResultWrongAnswer                      = -1
	ResultCpuTimeLimitExceeded             = 1
	ResultRealTimeLimitExceeded            = 2
	ResultMemoryLimitExceeded              = 3
	ResultRuntimeError                     = 4
	ResultSystemError                      = 5
)

type ErrorCode int

const (
	ErrorSuccess           ErrorCode = 0 - iota
	ErrorInvalidConfig               = -1
	ErrorForkFailed                  = -2
	ErrorPthreadFailed               = -3
	ErrorWaitFailed                  = -4
	ErrorRootRequired                = -5
	ErrorLoadSeccompFailed           = -6
	ErrorSetrlimitFailed             = -7
	ErrorDup2Failed                  = -8
	ErrorSetuidFailed                = -9
	ErrorExecveFailed                = -10
	ErrorSpjError                    = -11
)

const (
	IOModeStandard = "Standard IO"
	IOModeFile     = "File IO"
)

type Config struct {
	MaxCpuTime           int
	MaxRealTime          int
	MaxMemory            int
	MaxStack             int
	MaxProcessNumber     int
	MaxOutPutSize        int
	MemoryLimitCheckOnly int
	ExePath              string
	InputPath            string
	OutputPath           string
	ErrorPath            string
	LogPath              string
	SecCompRuleName      *string
	Uid                  int
	Gid                  int
	Args                 []string
	Env                  []string
}

type Result struct {
	CpuTime  int
	RealTime int
	Memory   int
	Signal   int
	ExitCode int
	Error    ErrorCode
	Result   ResultCode
}
