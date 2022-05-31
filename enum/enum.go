package enum

//处理进度
const (
	DOING   = "处理中"
	TOSERVE = "待人工咨询"
	SERVING = "人工咨询中"
	DONE    = "已完成"
)

//咨询交流记录类型
const (
	QUERY  = "query"
	ANSWER = "answer"
)

//回访记录状态
const (
	NORETURN      = "未回访"
	HAVEINTENTION = "有合作意向"
	NOINTENTION   = "无合作意向"
	COOPERATED    = "已合作"
)

type YesOrNo string

const (
	YES YesOrNo = "yes"
	NO  YesOrNo = "no"
)

type Cooperation string

const (
	MONITOR Cooperation = "monitor"
	PROTECT Cooperation = "protect"
)
