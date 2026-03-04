package dto

import (
	"errors"
	"github.com/shopspring/decimal"
	"net/http"
	"sf-go/pkg/common"
)

type EmailRegisterReq struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	Captcha    string `json:"captcha"`
	InviteUuid string `json:"inviteUuid"`
}
type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateKaiReq struct {
	StrategyId string          `json:"strategyId"`
	SlPx       decimal.Decimal `json:"slPx"`
	TpPx       decimal.Decimal `json:"tpPx"`
}

type CreateOrderReq struct {
	User   string  `json:"user"`
	Option string  `json:"option"` //web业务
	Level  string  `json:"level"`
	Amount float64 `json:"amount"`
}

type BotCreateReq struct {
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	ApiId    int    `json:"apiId"`
	Password string `json:"password"`
}
type FollowReq struct {
	BotId  int `json:"botId"`
	Status int `json:"status"`
}
type StartReq struct {
	ApiId     int    `json:"apiId"`
	Amount    int    `json:"amount"`
	Parms     string `json:"parms"`
	BotCoinId int    `json:"botCoinId"`
	Status    int    `json:"status"`
}
type UpdatePasswordReq struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
type AdminPlaceOrderReq struct {
	SpotId     string  `json:"spotId"`
	ContractId string  `json:"contractId"`
	Symbol     string  `json:"symbol"`
	Exchange   string  `json:"exchange"`
	Level      string  `json:"level"`
	OrdType    string  `json:"ordType"`
	Sz         float64 `json:"sz"`
	TpOrdPx    string  `json:"tpOrdPx"`
	SlOrdPx    string  `json:"slOrdPx"`
}

type PlaceOrder struct {
	Exchange    string `json:"exchange"`
	ApiId       string `json:"apiId"`
	Symbol      string `json:"symbol"`
	Side        string `json:"side"`
	Amount      string `json:"amount"`
	OrdType     string `json:"ordType"`
	TpTriggerPx string `json:"tpTriggerPx"`
	SlTriggerPx string `json:"slTriggerPx"`
	Px          string `json:"px"`
	Lever       string `json:"lever"`
	IsClose     string `json:"isClose"`
}

type ClosePos struct {
	Exchange string `json:"exchange"`
	ApiId    string `json:"apiId"`
	OrdId    string `json:"ordId"`
	AlgoId   string `json:"algoId"`
	Symbol   string `json:"symbol"`
	IsAll    string `json:"isAll"`
}

type SetPositionMode struct {
	ApiId   string `json:"apiId"`
	PosMode string `json:"posMode"` //long_short_mode：开平仓模式 net_mode：买卖模式
}
type ClosePosReq struct {
	ApiId   string `json:"apiId"`
	Symbol  string `json:"symbol"`
	MgnMode string `json:"mgnMode"` //保证金模式  cross：全仓 ； isolated：逐仓
	Side    string `json:"side"`    //long：平多 ，short：平空
	PosSide string `json:"posSide"` //long：平多 ，short：平空
}
type SetLeverReq struct {
	ApiId   string `json:"apiId"`
	Symbol  string `json:"symbol"`
	Lever   string `json:"lever"`
	MgnMode string `json:"mgnMode"` //isolated：逐仓 cross：全仓
	PosSide string `json:"posSide"` //long：开平仓模式开多 ，short：开平仓模式开空
}

type UserPlaceOrderReq struct {
	ApiId        int64   `json:"apiId"`
	Amount       float64 `json:"amount"` //U
	RenkoPercent float64 `json:"renkoPercent"`
	Multiple     float64 `json:"multiple"`
	Price        float64 `json:"price"`

	Symbol string `json:"symbol"`

	Side    string `json:"side"`    //long：平多 ，short：平空
	PosSide string `json:"posSide"` //long：平多 ，short：平空
}

type CancelOrderReq struct {
	ApiId    string `json:"apiId"`
	Exchange string `json:"exchange"`
	Symbol   string `json:"symbol"`
	OrdId    string `json:"ordId"`
}
type MarkPriceCandlesReq struct {
	InstId string `json:"instId"`
	After  string `json:"after"`
	Before string `json:"before"`
	Bar    string `json:"bar"`
	Limit  string `json:"limit"`
}

// 	Symbol          string    `gorm:"column:symbol;size:32;not null"`     // 币种
// OpportunityType string    `gorm:"column:op_type;size:16"`             // 类型（做多/做空/套利/空投/其他）
// Side            string    `gorm:"column:side;size:8;not null"`        // long/short
// Tags            string    `gorm:"column:tags;type:text"`              // 标签，英文逗号分隔
// Reason          string    `gorm:"column:reason;type:text"`            // 主要理由
// Source          string    `gorm:"column:source;type:text"`            // 信息来源
// Start           time.Time `gorm:"column:start;type:datetime"`         // 机会开始时间
// Status          int8      `gorm:"column:status;default:0"`            // 0未下单,1已下单,2已完成,3忽略
// FollowUp        string    `gorm:"column:follow_up;type:text"`         // 跟踪动态/进展
// IsActive        bool      `gorm:"column:is_active;default:true"`      // 是否重点关注
// Note            string    `gorm:"column:note;type:text"`              // 备注
// CreatedAt       time.Time `gorm:"column:created_at;autoCreateTime"`
// UpdatedAt       time.Time `gorm:"column:updated_at;autoUpdateTime"`

func NewEmailRegisterReq() *EmailRegisterReq {
	return &EmailRegisterReq{}
}
func (e *EmailRegisterReq) Bind(g *Gin) error {
	if err := g.C.ShouldBindJSON(e); err != nil {
		g.DirectResponse(http.StatusBadRequest, err.Error(), nil)
		return err
	}
	if e.InviteUuid == "" {
		e.InviteUuid = "000000"
	}
	okEmail := common.VerifyEmailFormat(e.Email)
	if !okEmail {
		g.Response(http.StatusBadRequest, EMAIL_ERROR, nil)
		return errors.New("email error")
	}
	return nil
}

func NewLoginReq() *LoginReq {
	return &LoginReq{}
}
func (l *LoginReq) Bind(g *Gin) error {
	if err := g.C.ShouldBindJSON(l); err != nil {
		g.DirectResponse(http.StatusBadRequest, err.Error(), nil)
		return err
	}
	return nil
}
func NewUpdateKaiReq() *UpdateKaiReq {
	return &UpdateKaiReq{}
}
func (l *UpdateKaiReq) Bind(g *Gin) error {
	if err := g.C.ShouldBindJSON(l); err != nil {
		g.DirectResponse(http.StatusBadRequest, err.Error(), nil)
		return err
	}
	return nil
}

func NewPlaceOrder() *PlaceOrder {
	return &PlaceOrder{}
}
func (l *PlaceOrder) Bind(g *Gin) error {
	if err := g.C.ShouldBindJSON(l); err != nil {
		g.DirectResponse(http.StatusBadRequest, err.Error(), nil)
		return err
	}
	return nil
}

func NewClosePos() *ClosePos {
	return &ClosePos{}
}
func (l *ClosePos) Bind(g *Gin) error {
	if err := g.C.ShouldBindJSON(l); err != nil {
		g.DirectResponse(http.StatusBadRequest, err.Error(), nil)
		return err
	}
	return nil
}
func NewSetPositionMode() *SetPositionMode {
	return &SetPositionMode{}
}
func (l *SetPositionMode) Bind(g *Gin) error {
	if err := g.C.ShouldBindJSON(l); err != nil {
		g.DirectResponse(http.StatusBadRequest, err.Error(), nil)
		return err
	}
	return nil
}

func NewCreateOrderReq() *CreateOrderReq {
	return &CreateOrderReq{}
}

func (c *CreateOrderReq) Bind(g *Gin) error {
	if err := g.C.ShouldBindJSON(c); err != nil {
		g.DirectResponse(http.StatusBadRequest, err.Error(), nil)
		return err
	}
	return nil
}

func NewBotCreateReq() *BotCreateReq {
	return &BotCreateReq{}
}

func (b *BotCreateReq) Bind(g *Gin) error {
	if err := g.C.ShouldBindJSON(b); err != nil {
		g.DirectResponse(http.StatusBadRequest, err.Error(), nil)
		return err
	}
	return nil
}

func NewFollowReq() *FollowReq {
	return &FollowReq{}
}

func (b *FollowReq) Bind(g *Gin) error {
	if err := g.C.ShouldBindJSON(b); err != nil {
		g.DirectResponse(http.StatusBadRequest, err.Error(), nil)
		return err
	}
	return nil
}

func NewStartReq() *StartReq {
	return &StartReq{}
}

func (s *StartReq) Bind(g *Gin) error {
	if err := g.C.ShouldBindJSON(s); err != nil {
		g.DirectResponse(http.StatusBadRequest, err.Error(), nil)
		return err
	}
	return nil
}

func NewUpdatePasswordReq() *UpdatePasswordReq {
	return &UpdatePasswordReq{}
}

func (u *UpdatePasswordReq) Bind(g *Gin) error {
	if err := g.C.ShouldBindJSON(u); err != nil {
		g.DirectResponse(http.StatusBadRequest, err.Error(), nil)
		return err
	}
	return nil
}

func NewAdminPlaceOrderReq() *AdminPlaceOrderReq {
	return &AdminPlaceOrderReq{}
}

func (c *AdminPlaceOrderReq) Bind(g *Gin) error {
	if err := g.C.ShouldBindJSON(c); err != nil {
		g.DirectResponse(http.StatusBadRequest, err.Error(), nil)
		return err
	}
	return nil
}
func NewUserPlaceOrderReq() *UserPlaceOrderReq {
	return &UserPlaceOrderReq{}
}

func (c *UserPlaceOrderReq) Bind(g *Gin) error {
	if err := g.C.ShouldBindJSON(c); err != nil {
		g.DirectResponse(http.StatusBadRequest, err.Error(), nil)
		return err
	}
	return nil
}

func NewClosePosReq() *ClosePosReq {
	return &ClosePosReq{}
}

func (c *ClosePosReq) Bind(g *Gin) error {
	if err := g.C.ShouldBindJSON(c); err != nil {
		g.DirectResponse(http.StatusBadRequest, err.Error(), nil)
		return err
	}
	return nil
}
func NewSetLeverReq() *SetLeverReq {
	return &SetLeverReq{}
}

func (c *SetLeverReq) Bind(g *Gin) error {
	if err := g.C.ShouldBindJSON(c); err != nil {
		g.DirectResponse(http.StatusBadRequest, err.Error(), nil)
		return err
	}
	return nil
}

func NewCancelOrderReq() *CancelOrderReq {
	return &CancelOrderReq{}
}

func (c *CancelOrderReq) Bind(g *Gin) error {
	if err := g.C.ShouldBindJSON(c); err != nil {
		g.DirectResponse(http.StatusBadRequest, err.Error(), nil)
		return err
	}
	return nil
}
