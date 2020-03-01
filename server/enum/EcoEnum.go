package enum

const (
	VERSION     = "2.1"
	TITLE       = "禪念 Bot Go " + VERSION
	WORDS_LIMIT = 200
	QUICK_MENU  = "快速選單"
	IMG_URL_ZEN = "https://firebasestorage.googleapis.com/v0/b/walter-bot-a2142.appspot.com/o/line-bot%2Fimage%2Fother%2Fzen%2F3.png?alt=media&token=44770a07-a661-40dc-960e-45da1699e4f2"
	DEFAULT_IMG = "https://firebasestorage.googleapis.com/v0/b/walter-bot-a2142.appspot.com/o/line-bot%2Fimage%2Fother%2Fgirl_img%2F27367-5nYPUB.jpg?alt=media&token=9ec89929-5b2d-478c-b8da-c37f61f338a0"
)

type EcoEnum int

const (
	GOOD_WORDS    = iota // 好語
	WISDOM_ADAGE         // 自在語
	PHORISM              // 靜思語
	INSPIRATIONAL        // 勵志語
	INPUT_WORDS          // 請輸入
)

func (t EcoEnum) String() string {
	return [...]string{
		"好語",
		"自在語",
		"靜思語",
		"勵志語",
		"請輸入",
	}[t]
}
