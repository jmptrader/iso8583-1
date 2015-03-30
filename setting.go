package iso8583

const (
	MAX_ISO_SIZE   = 1024
	MIN_ISO_SIZE   = 20
	MAX_FIELD_SIZE = 1024
)

// 长度类型
const (
	CVAR   = iota // 定长
	LLVAR         // 变长，用2字节（BCD之后1字节）表示长度
	LLLVAR        // 变长，用3字节（BCD之后2字节）表示长度
)

// 数据类型
const (
	DT_BCD = iota // BCD
	DT_ASC        // ASCII
	DT_BIN        // 二进制
)

// 填充类型
const (
	FILL_ZERO  = iota // 填充'0'
	FILL_SPACE        // 填充空格
)

// 对齐类型
const (
	JUST_RIGHT = iota // 右对齐
	JUST_LEFT         // 左对齐
)

type FieldDef struct {
	Name       string
	LengthType int // 长度类型
	Length     int // 长度（对于定长）或最大长度（对于变长）
	DataType   int // 数据类型
	Just       int // 对齐方式
	Fill       int // 填充字符
}

// 银联的
var IsoSettingYL = []FieldDef{
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO}, // 第5域
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO}, // 10
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO}, // 15
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO}, // 20
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO}, // 25
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO}, // 30
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO}, // 35
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO}, // 40
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO}, // 45
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO}, // 50
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO}, // 55
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO}, // 60
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO},
	{"field 0", CVAR, 1, DT_BIN, JUST_RIGHT, FILL_ZERO}, // 64
}
