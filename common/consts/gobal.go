package consts

const (
	// 当前支持平台
	PLATFORM_LOBBY = "lobby"
)

const (
	// 运行环境
	RUN_MODE_DEV  = "dev"
	RUN_MODE_PROD = "prod"
)

const (
	SQLCON    = "#"
	SQLSEP    = ","
	SEP       = "|"
	BLANK     = " "
	EMPTY_STR = ""
	ZERO      = 0
	ONE       = 1
)

const (
	DEFAULT_UID = 0
	ALL_UID     = 0
	ALL_GID     = 0
)

const (
	DEFAULT_NEWER = 0
	NEWER_NEW_ID  = 1
	NEWER_OLD_ID  = 2
)

const (
	DEFAULT_GENDER   = 0
	GENDER_MALE_ID   = 1
	GENDER_FEMALE_ID = 2
)

const (
	TIMEFORMAT = "2006-01-02"

	ONE_DAY   = 1
	ONE_WEEK  = 7
	ONE_MONTH = 30
)

const (
	MaxInt8   = 1<<7 - 1
	MinInt8   = -1 << 7
	MaxInt16  = 1<<15 - 1
	MinInt16  = -1 << 15
	MaxInt32  = 1<<31 - 1
	MinInt32  = -1 << 31
	MaxInt64  = 1<<63 - 1
	MinInt64  = -1 << 63
	MaxUint8  = 1<<8 - 1
	MaxUint16 = 1<<16 - 1
	MaxUint32 = 1<<32 - 1
	MaxUint64 = 1<<64 - 1
)
