package clock

import (
	"time"

	"github.com/hilqiqi0/golang-utils/common/consts"
)

//获取某一天的0点时间
func GetZeroTimeDaysAgo(day int) time.Time {
	d := time.Now().AddDate(consts.ZERO, consts.ZERO, -day)
	return time.Date(d.Year(), d.Month(), d.Day(), consts.ZERO, consts.ZERO, consts.ZERO, consts.ZERO, d.Location())
}

func getDaysAgo(n int) string {
	return time.Now().AddDate(consts.ZERO, consts.ZERO, -n).Format(consts.TIMEFORMAT)
}

func GetTodayFormat() string {
	return getDaysAgo(consts.ONE_DAY - 1)
}

func GetYesterdayFormat() string {
	return getDaysAgo(consts.ONE_DAY)
}

func GetWeekAgoFormat() string {
	return getDaysAgo(consts.ONE_WEEK)
}

func GetMonthAgoFormat() string {
	// 约定一个月为30天前
	return getDaysAgo(consts.ONE_MONTH)
}
