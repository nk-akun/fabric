package blockcutter

const (
	// pid parameters
	kp = 0.16
	ki = 0.01
	kd = 0.05

	// 期望值
	expectationMessageCount = 500
)

var (
	currentMessageCount int
	lastMessageCount    int
	currentError        int
	lastError           int
	proportional        int // 比例项
	integral            int // 积分项
	derivative          int // 微分项
)

func init() {
	integral = 0
	currentError = 0
	currentMessageCount = 1
}

func getMaxMessageCount() uint32 {
	return uint32(currentMessageCount)
}

func updateMessageCount(realCount int) {
	lastMessageCount = realCount
	lastError = currentError

	currentError = expectationMessageCount - lastMessageCount
	proportional = currentError
	integral += currentError
	derivative = lastError - currentError

	currentMessageCount = lastMessageCount + int(kp*float64(proportional)+ki*float64(integral)+kd*float64(derivative))
}
