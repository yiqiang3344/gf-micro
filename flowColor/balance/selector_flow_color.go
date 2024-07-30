package balancer

import (
	"context"
	"github.com/gogf/gf/v2/net/gsel"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/yiqiang3344/gf-micro/flowColor"
	"sync"

	"github.com/gogf/gf/v2/util/grand"
)

type selectorFlowColor struct {
	mu    sync.RWMutex
	nodes map[string]gsel.Nodes
}

func NewSelectorFlowColor() gsel.Selector {
	return &selectorFlowColor{
		nodes: make(map[string]gsel.Nodes),
	}
}

func (s *selectorFlowColor) Update(ctx context.Context, nodes gsel.Nodes) error {
	var str string
	for _, node := range nodes {
		if str != "" {
			str += ","
		}
		str += node.Address() + "[" + node.Service().GetMetadata().Get(flowColor.FlowColor).String() + "]"
	}
	glog.Debugf(ctx, "Update nodes: %s\n", str)
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, v := range nodes {
		if f := v.Service().GetMetadata().Get(flowColor.FlowColor); f != nil {
			fs := f.String()
			if v1, ok := s.nodes[fs]; !ok {
				s.nodes[fs] = []gsel.Node{v}
			} else {
				s.nodes[fs] = append(v1, v)
			}
		}
	}
	return nil
}

func (s *selectorFlowColor) Pick(ctx context.Context) (node gsel.Node, done gsel.DoneFunc, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if len(s.nodes) == 0 {
		return nil, nil, nil
	}
	//先找指定的染色服务，流量染色标识为空的时候，忽略，往下找基准服务
	fc := *flowColor.GetCtxFlowColor(ctx)
	if v, ok := s.nodes[fc]; ok && fc != "" {
		node = v[grand.Intn(len(v))]
	}
	//再找基准服务
	if node == nil {
		if v, ok := s.nodes[flowColor.ColorBase]; ok {
			node = v[grand.Intn(len(v))]
		}
	}
	if node == nil {
		return nil, nil, nil
	}
	glog.Debugf(ctx, "flow color[%s] Picked node[%s,%s]", fc, node.Address(), node.Service().GetMetadata().Get(flowColor.FlowColor).String())
	return node, nil, nil
}
