package rocketmq_client

import (
	"context"
	"errors"
	"fmt"
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
	v2 "github.com/apache/rocketmq-clients/golang/v5/protocol/v2"
	"os"
	"strings"
	"time"
)

// TopicType 主题类型
type TopicType string

const (
	TopicNormal      TopicType = "Normal"
	TopicFIFO        TopicType = "FIFO"
	TopicDelay       TopicType = "Delay"
	TopicTransaction TopicType = "Transaction"
)

type Client interface {
	StartProducer(ctx context.Context, oFunc ...ProducerOptionFunc) error                                                //启动生产者
	StopProducer() error                                                                                                 //注销消费者
	Send(ctx context.Context, topicType TopicType, msg Message) (resp []*rmq_client.SendReceipt, err error)              //同步发送消息
	SendAsync(ctx context.Context, topicType TopicType, msg Message, dealFunc SendAsyncDealFunc) error                   //异步发送消息
	SendTransaction(ctx context.Context, message Message, confirmFunc ConfirmFunc) error                                 //发送事务消息
	SimpleConsume(ctx context.Context, consumeFuc ConsumeFunc, oFunc ...ConsumerOptionFunc) (stopFunc func(), err error) //简单模式消费消息
}

type ClientCfg struct {
	Endpoint         string           //必填
	NameSpace        string           //必填
	ConsumerGroup    string           //使用消费者时，必填
	AccessKey        string           //可选
	AccessSecret     string           //可选
	LogPath          string           //官方rocketmq日志文件路径
	LogStdout        bool             //是否在终端输出官方rocketmq日志，输出的话则不会记录日志文件
	Debug            bool             //是否在终端输出本客户端的debug信息
	DebugHandlerFunc debugHandlerFunc //本客户端的debug信息处理方法，不管debug开没开，有debug信息的时候都会调用
}

type defaultClient struct {
	Cfg          *rmq_client.Config
	debug        bool
	debugHandler debugHandlerFunc
	producer     rmq_client.Producer
}

type debugHandlerFunc func(msg string)

// GetClient 获取mq客户端
func GetClient(cfg *ClientCfg) (client Client, err error) {
	if cfg.LogStdout {
		os.Setenv("mq.consoleAppender.enabled", "true")
	} else {
		os.Setenv("mq.consoleAppender.enabled", "false")
	}
	os.Setenv("rocketmq.client.logRoot", cfg.LogPath)
	rmq_client.ResetLogger()

	if strings.Trim(cfg.Endpoint, "") == "" {
		err = errors.New("Endpoint不能为空")
		return
	}

	if strings.Trim(cfg.NameSpace, "") == "" {
		err = errors.New("NameSpace不能为空")
		return
	}

	client = &defaultClient{
		Cfg: &rmq_client.Config{
			Endpoint:      cfg.Endpoint,
			NameSpace:     cfg.NameSpace,
			ConsumerGroup: cfg.ConsumerGroup,
			Credentials: &credentials.SessionCredentials{
				AccessKey:     cfg.AccessKey,
				AccessSecret:  cfg.AccessSecret,
				SecurityToken: "",
			},
		},
		debug:        cfg.Debug,
		debugHandler: cfg.DebugHandlerFunc,
	}
	return
}

func (s *defaultClient) debugLog(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	if s.debugHandler != nil {
		s.debugHandler(msg)
	}
	if !s.debug {
		return
	}
	fmt.Printf("%s %10s %s\n", time.Now().Format("2006-01-02 15:04:05.000"), "DEBUG", msg)
	return
}

func WithProducerOptionTopics(Topics []string) ProducerOptionFunc {
	return func(o *ProducerOptions) {
		o.Topics = Topics
	}
}

type ProducerOptionFunc func(options *ProducerOptions)

func WithProducerOptionMaxAttempts(maxAttempts int32) ProducerOptionFunc {
	return func(o *ProducerOptions) {
		o.MaxAttempts = maxAttempts
	}
}

type SendTransactionCheckerFunc func(msg *rmq_client.MessageView) rmq_client.TransactionResolution

func WithProducerOptionTransactionChecker(transactionChecker SendTransactionCheckerFunc) ProducerOptionFunc {
	return func(o *ProducerOptions) {
		o.transactionChecker = transactionChecker
	}
}

type ProducerOptions struct {
	Topics             []string                   //支持的主题列表，可选
	MaxAttempts        int32                      //重试次数，可选
	transactionChecker SendTransactionCheckerFunc //事务检查器，事务消息必填
}

// StartProducer 启动生产者
func (s *defaultClient) StartProducer(ctx context.Context, oFunc ...ProducerOptionFunc) error {
	o := ProducerOptions{
		MaxAttempts: 3,
	}
	options := &o
	if len(oFunc) > 0 {
		for _, f := range oFunc {
			f(options)
		}
	}

	producer, err := rmq_client.NewProducer(
		s.Cfg,
		rmq_client.WithTopics(options.Topics...),
		rmq_client.WithMaxAttempts(options.MaxAttempts),
		rmq_client.WithTransactionChecker(&rmq_client.TransactionChecker{
			Check: options.transactionChecker,
		}),
	)
	if err != nil {
		s.debugLog("生产者初始化失败:%v", err)
		return err
	}
	err = producer.Start()
	if err != nil {
		s.debugLog("生产者启动失败:%v", err)
		return err
	}
	s.producer = producer
	return nil
}

// StopProducer 注销生产者
func (s *defaultClient) StopProducer() error {
	err := s.producer.GracefulStop()
	if err != nil {
		s.debugLog("生产者注销失败:%v", err)
		return nil
	}
	s.producer = nil
	s.debugLog("生产者注销成功")
	return err
}

type Message struct {
	Body              string            //消息内容，必填
	Topic             string            //主题，必填
	Tag               string            //标签，可选
	MessageGroup      string            //消息组，FIFO消息类型必填，其他可选
	Keys              []string          //索引列表，可选
	Properties        map[string]string //自定义属性，可选
	DeliveryTimestamp *time.Time        //延迟时间，Delay消息类型必填，其他可选
}

// initMsg 包装消息
func (s *defaultClient) initMsg(ctx context.Context, topicType TopicType, message Message) (msg *rmq_client.Message, err error) {
	//校验
	if strings.Trim(message.Topic, "") == "" {
		err = errors.New("topic必填")
		s.debugLog("消息初始化失败:%v", err)
		return
	}
	if strings.Trim(message.Body, "") == "" {
		err = errors.New("body必填")
		s.debugLog("消息初始化失败:%v", err)
		return
	}
	switch topicType {
	case TopicFIFO:
		if strings.Trim(message.MessageGroup, "") == "" {
			err = errors.New("FIFO消息类型messageGroup必填")
			s.debugLog("消息初始化失败:%v", err)
			return
		}
	case TopicDelay:
		if message.DeliveryTimestamp == nil {
			err = errors.New("Delay消息类型deliveryTimestamp必填")
			s.debugLog("消息初始化失败:%v", err)
			return
		}
	}

	//初始化消息体
	msg = &rmq_client.Message{
		Topic: message.Topic,
		Body:  []byte(message.Body),
	}
	//设置消息tag
	if message.Tag != "" {
		msg.SetTag(message.Tag)
	}
	//设置消息组（只有FIFO类型topic可用）
	if message.MessageGroup != "" {
		msg.SetMessageGroup(message.MessageGroup)
	}
	//设置消息key
	msg.SetKeys(message.Keys...)
	//设置消息属性
	// todo 添加trace_id
	if len(message.Properties) > 0 {
		for k, v := range message.Properties {
			msg.AddProperty(k, v)
		}
	}
	//设置延迟时间（只有Delay类型topic可用）
	if message.DeliveryTimestamp != nil {
		msg.SetDelayTimestamp(*message.DeliveryTimestamp)
	}
	return
}

// IsTooManyRequest 是否触发了流控
func IsTooManyRequest(err error) bool {
	//如果是重试失败，则判断是否设置了补偿机制，有则调用
	if e, ok := err.(*rmq_client.ErrRpcStatus); ok && e.GetCode() == int32(v2.Code_TOO_MANY_REQUESTS) {
		return true
	}
	return false
}

// IsNoNewMessage 是否没有新消息
func IsNoNewMessage(err error) bool {
	//如果是重试失败，则判断是否设置了补偿机制，有则调用
	if e, ok := err.(*rmq_client.ErrRpcStatus); ok && e.GetCode() == int32(v2.Code_MESSAGE_NOT_FOUND) {
		return true
	}
	return false
}

// Send 同步发送消息
// 可支持普通、延迟、顺序类型的消息，不支持事务消息
func (s *defaultClient) Send(ctx context.Context, topicType TopicType, msg Message) (resp []*rmq_client.SendReceipt, err error) {
	if s.producer == nil {
		err = errors.New("请先初始化生产者")
		s.debugLog("消息发送失败:%v", err)
		return
	}

	if topicType == TopicTransaction {
		err = errors.New("此方法不支持发送Transaction消息")
		s.debugLog("消息发送失败:%v", err)
		return
	}
	message, err := s.initMsg(ctx, topicType, msg)
	if err != nil {
		s.debugLog("消息发送失败:%v", err)
		return
	}

	resp, err = s.producer.Send(ctx, message)
	if err != nil {
		s.debugLog("消息发送失败:%v", err)
		return
	}
	return
}

type SendAsyncDealFunc func(ctx context.Context, msg Message, resp []*rmq_client.SendReceipt, err error)

// SendAsync 异步发送消息
// 可支持普通、延迟、顺序类型的消息，不支持事务消息
func (s *defaultClient) SendAsync(ctx context.Context, topicType TopicType, msg Message, dealFunc SendAsyncDealFunc) (err error) {
	if s.producer == nil {
		err = errors.New("请先初始化生产者")
		s.debugLog("消息发送失败:%v", err)
		return
	}

	if dealFunc == nil {
		err = errors.New("dealFunc必填")
		s.debugLog("消息发送失败:%v", err)
		return
	}

	if topicType == TopicTransaction {
		err = errors.New("此方法不支持发送Transaction消息")
		s.debugLog("消息发送失败:%v", err)
		return
	}

	message, err := s.initMsg(ctx, topicType, msg)
	if err != nil {
		return err
	}

	s.producer.SendAsync(ctx, message, func(ctx context.Context, receipts []*rmq_client.SendReceipt, err error) {
		dealFunc(ctx, msg, receipts, err)
	})
	return
}

// ConfirmFunc 二次确认方法
// 注意：不要异步处理，本地事务逻辑提交时返回true，否则返回false
type ConfirmFunc func(msg Message, resp []*rmq_client.SendReceipt) bool

// SendTransaction 发送事务消息
// 注意：事务消息的生产者不能和其他类型消息的生产者共用
func (s *defaultClient) SendTransaction(ctx context.Context, message Message, confirmFunc ConfirmFunc) (err error) {
	if s.producer == nil {
		err = errors.New("请先初始化生产者")
		s.debugLog("消息发送失败:%v", err)
		return
	}

	if confirmFunc == nil {
		err = errors.New("confirmFunc必填")
		s.debugLog("消息发送失败:%v", err)
		return
	}

	msg, err := s.initMsg(ctx, TopicTransaction, message)
	if err != nil {
		return
	}

	transaction := s.producer.BeginTransaction()
	resp, err := s.producer.SendWithTransaction(ctx, msg, transaction)
	if err != nil {
		s.debugLog("消息发送失败:%v", err)
		return
	}
	if confirmFunc(message, resp) {
		return transaction.Commit()
	}
	return transaction.RollBack()
}

type ConsumerOptionFunc func(options *ConsumerOptions)

func WithConsumerOptionAwaitDuration(AwaitDuration time.Duration) ConsumerOptionFunc {
	return func(o *ConsumerOptions) {
		o.AwaitDuration = AwaitDuration
	}
}

func WithConsumerOptionMaxMessageNum(MaxMessageNum int32) ConsumerOptionFunc {
	return func(o *ConsumerOptions) {
		o.MaxMessageNum = MaxMessageNum
	}
}

func WithConsumerOptionInvisibleDuration(InvisibleDuration time.Duration) ConsumerOptionFunc {
	return func(o *ConsumerOptions) {
		o.InvisibleDuration = InvisibleDuration
	}
}

func WithConsumerOptionSubExpressions(SubExpressions map[string]*rmq_client.FilterExpression) ConsumerOptionFunc {
	return func(o *ConsumerOptions) {
		o.SubExpressions = SubExpressions
	}
}

type ConsumerOptions struct {
	AwaitDuration     time.Duration                           //消息处理超时时间，超时会触发消费重试
	MaxMessageNum     int32                                   //每次接收的消息数量
	InvisibleDuration time.Duration                           //接收到的消息的不可见时间
	SubExpressions    map[string]*rmq_client.FilterExpression //订阅表达式，key为topic，简单消费类型只支持tag和sql匹配
}

// ConsumeFunc 消费方法
// 方法内消费成功时需要调用consumer.Ack()；
// 消费时间可能超过消费者MaxMessageNum设置的时间时，可调用consumer.ChangeInvisibleDuration()或consumer.ChangeInvisibleDurationAsync()方法调整消息消费超时时间；
type ConsumeFunc func(ctx context.Context, msg *rmq_client.MessageView, consumer rmq_client.SimpleConsumer)

// SimpleConsume 简单消费类型消费
func (s *defaultClient) SimpleConsume(ctx context.Context, consumeFunc ConsumeFunc, oFunc ...ConsumerOptionFunc) (stopFunc func(), err error) {
	o := ConsumerOptions{
		AwaitDuration:     time.Second * 5,
		MaxMessageNum:     10,
		InvisibleDuration: time.Second * 5,
	}
	options := &o
	if len(oFunc) > 0 {
		for _, f := range oFunc {
			f(options)
		}
	}

	if len(options.SubExpressions) == 0 {
		err = errors.New("SubExpressions不能为空")
		s.debugLog("消费者参数不合法:%v", err)
		return
	}

	if strings.Trim(s.Cfg.ConsumerGroup, "") == "" {
		err = errors.New("ConsumerGroup不能为空")
		s.debugLog("消费者参数不合法:%v", err)
		return
	}

	consumer, err := rmq_client.NewSimpleConsumer(
		s.Cfg,
		rmq_client.WithAwaitDuration(options.AwaitDuration),
		rmq_client.WithSubscriptionExpressions(options.SubExpressions),
	)
	if err != nil {
		s.debugLog("初始化消费者失败:%v", err)
		return nil, err
	}

	err = consumer.Start()
	if err != nil {
		s.debugLog("消费者启动失败:%v", err)
		return nil, err
	}

	stopFunc = func() {
		err1 := consumer.GracefulStop()
		if err1 != nil {
			s.debugLog("消费者注销失败:%v", err1)
			return
		}
		s.debugLog("消费者注销成功")
	}

	go func() {
		for {
			mvs, err1 := consumer.Receive(ctx, options.MaxMessageNum, options.InvisibleDuration)
			if err1 != nil {
				if IsNoNewMessage(err1) {
					time.Sleep(1 * time.Second)
					continue
				}
				s.debugLog("获取消息失败:%v", err1)
			}
			for _, mv := range mvs {
				consumeFunc(ctx, mv, consumer)
			}
		}
	}()
	return
}
