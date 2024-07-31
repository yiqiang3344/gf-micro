package balancer

import "github.com/gogf/gf/v2/net/gsel"

type builderFlowColor struct{}

func NewBuilderFlowColor() gsel.Builder {
	return &builderFlowColor{}
}

func (*builderFlowColor) Name() string {
	return "BalancerFlowColor"
}

func (*builderFlowColor) Build() gsel.Selector {
	return NewSelectorFlowColor()
}
