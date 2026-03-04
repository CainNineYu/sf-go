package consts

import "time"

// TODO: 统一为指数增加重试
const (
	DefaultPingInterval          = 15 * time.Second   // ping的间隔秒
	DefaultReconnectInterval     = 2 * time.Second    // 重连间隔秒
	DefaultMaxRetry              = 10                 // 重连最大次数
	DefaultRetryInterval         = 2 * time.Second    // WS重试间隔秒
	DefaultRetrySubInterval      = 1 * time.Second    // WS订阅失败后重试间隔秒
	DefaultReceiptInterval       = 1200 * time.Second // 订阅更新时间间隔，超过时间无响应重新订阅
	DefaultReconnectLongInterval = 600 * time.Second  // WS断开后长时间重试间隔
)

// 对冲方式
type HedgeMethod string

const (
	HedgeMethodManual = HedgeMethod("manual") //手动对冲
	HedgeMethodAuto   = HedgeMethod("auto")   //自动对冲
	HedgeMethodMix    = HedgeMethod("mix")    //混合对冲
)

// 交易类型
type TradeCategory string

const (
	TradeCategorySpot      = TradeCategory("spot")
	TradeCategoryFutures   = TradeCategory("futures")
	TradeCategoryPerpetual = TradeCategory("perpetual")
)

// 交易平台
type Platform string

const (
	PlatformOKX     = Platform("OKX")
	PlatformBinance = Platform("BINANCE")
	PlatformDTX     = Platform("DTX")
)

// 对冲交易对配置状态
type HedgeSymbolConfigStatus int8

const (
	HedgeSymbolConfigStatusEnable  = HedgeSymbolConfigStatus(0) //启用
	HedgeSymbolConfigStatusDisable = HedgeSymbolConfigStatus(1) //停用
	HedgeSymbolConfigStatusAll     = HedgeSymbolConfigStatus(2) //用于校验全部查询
)

// 对冲方式
type StrategyType string

const (
	StrategyTypeMirror = StrategyType("mirror") //镜像
	StrategyTypeSwamp  = StrategyType("swap")
)

const (
	TOKEN_KEY = "Bearer "
)

// 刷量类型
type BrushType int64

const (
	BrushTypeCustom  = BrushType(1) //完全自定义刷量
	BrushTypeBinance = BrushType(2) //币安完全跟随刷量
	BrushTypeOkx     = BrushType(3) //欧意完全跟随刷量
)

const PxLast = "last"

// NATS
const (
	NATS_SUBJECT_SPOT_PRICE = "spot.price.%s.%s" // %s为exchange，%s为symbol
	NATS_SUBJECT_PERP_PRICE = "perp.price.%s.%s" // %s为exchange，%s为symbol
	NATS_SUBJECT_PERP_PX    = "px.perp.%s.%s.%s" // %s为exchange，%s为symbol，%s为last最新成交价格
)

// redis
const (
	PerpLastPxExpiration = 2 * time.Minute
	REDIS_KEY_SPOT_PRICE = "px:%s:%s:%s" // %s为exchange, %s为symbol, %s为market（last最新价）
	// Hset 存储 field：p（最新成交价），ts（更新时间戳）
	REDIS_KEY_PERP_PRICE = "px:%s:%s:%s" // %s为exchange, %s为symbol, %s为market（last最新价）
)

const HostIdStr = "quant-virtual-ord-%d.quant-virtual-ord%s" // 服务id,第二个为+ :端口，有冒号
