/**
 * @Copyright: 2025 Shopee
 * @Author: xiang.zhong
 * @Email: xiang.zhong@shopee.com
 * @Date: 2025/01/24 15:39
 */
package fsm

import (
	"context"

	"github.com/looplab/fsm"
	"go.uber.org/zap"
)

type Event struct {
	From []string `json:"From"` // 来源状态
	To   string   `json:"To"`   // 目标状态
}

type FsmConfig struct {
	InitialState PaymentPayoutFsmStatus `json:"initial_state"` // 初始状态
	Events       map[string]Event       `json:"events"`        // 状态迁移事件配置
	Callbacks    map[string]string      `json:"callbacks"`     // 回调函数配置
}

type PaymentFsm struct {
	FSM     *fsm.FSM        // 状态机实例
	Config  *FsmConfig      // 状态机配置
	Context context.Context // 上下文，用于回调函数
	Logger  *zap.Logger     // 日志，用于调试或记录
}

type PayoutModel struct {
	Id           int64
	PayoutStatus PaymentPayoutFsmStatus
	Operator     string
	CreateTime   int64
	UpdateTime   int64
}
