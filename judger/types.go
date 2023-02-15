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

type RunConfig struct {
	MaxCpuTime           int
	MaxRealTime          int
	MaxMemory            int
	MaxStack             int
	MaxProcessNumber     int
	MaxOutputSize        int
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

type RunResult struct {
	CpuTime  int
	Error    ErrorCode
	ExitCode int
	Memory   int
	RealTime int
	Result   ResultCode
	Signal   int
}

type JudgeRunConfig struct {
	Command     string
	SeccompRule string
	Env         []string
}

type JudgeParams struct {
	RunConfig     JudgeRunConfig
	ExePath       string
	MaxCpuTime    int
	MaxMemory     int
	SubmissionDir string
	TestCaseId    string
}

type JudgeResult struct {
	CpuTime   int        `json:"cpu_time"`
	Error     ErrorCode  `json:"error"`
	ExitCode  int        `json:"exit_code"`
	Memory    int        `json:"memory"`
	OutputMd5 string     `json:"output_md5"`
	Output    string     `json:"output"`
	RealTime  int        `json:"real_time"`
	Result    ResultCode `json:"result"`
	Signal    int        `json:"signal"`
	TestCase  string     `json:"test_case"`
}

type TestCaseDetail struct {
	// InputName         string `json:"input_name"`
	StrippedOutputMd5 string `json:"stripped_output_md5"`
	OutputSize        *int   `json:"output_size"`
}

type TestCaseInfo struct {
	TestCases map[string]TestCaseDetail `json:"test_cases"`
}

type JudgeOneParams struct {
	Args           []string
	Env            []string
	ExePath        string
	MaxCpuTime     int
	MaxMemory      int
	SeccompRule    string
	SubmissionDir  string
	TestCaseId     string
	TestCaseDir    string
	TestCaseName   string
	TestCaseDetail TestCaseDetail
}
