package rocketmq_client

import (
	"context"
	"errors"
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	v2 "github.com/apache/rocketmq-clients/golang/v5/protocol/v2"
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
	StartProducer(ctx context.Context, oFunc ...ProducerOptionFunc) error                                   //启动生产者
	StopProducer() error                                                                                    //注销消费者
	Send(ctx context.Context, topicType TopicType, msg Message) (resp []*rmq_client.SendReceipt, err error) //同步发送消息
	SendAsync(ctx context.Context, topicType TopicType, msg Message, dealFunc SendAsyncDealFunc) error      //异步发送消息
	SendTransaction(ctx context.Context, message Message, confirmFunc ConfirmFunc) error                    //发送事务消息
	SimpleConsume(ctx context.Context, consumeFuc ConsumeFunc, oFunc ...ConsumerOptionFunc) error           //简单模式消费消息
}

// GetClient 获取mq客户端
func GetClient(cfg *rmq_client.Config) Client {
	return &defaultClient{
		Cfg: cfg,
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

type defaultClient struct {
	Cfg      *rmq_client.Config
	producer rmq_client.Producer
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
		return err
	}
	err = producer.Start()
	if err != nil {
		return err
	}
	s.producer = producer
	return nil
}

// StopProducer 注销生产者
func (s *defaultClient) StopProducer() error {
	err := s.producer.GracefulStop()
	s.producer = nil
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
func initMsg(topicType TopicType, message Message) (msg *rmq_client.Message, err error) {
	//校验
	if strings.Trim(message.Topic, "") == "" {
		err = errors.New("topic必填")
		return
	}
	if strings.Trim(message.Body, "") == "" {
		err = errors.New("body必填")
		return
	}
	switch topicType {
	case TopicFIFO:
		if strings.Trim(message.MessageGroup, "") == "" {
			err = errors.New("FIFO消息类型messageGroup必填")
			return
		}
	case TopicDelay:
		if message.DeliveryTimestamp == nil {
			err = errors.New("Delay消息类型deliveryTimestamp必填")
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

// Send 同步发送消息
// 可支持普通、延迟、顺序类型的消息，不支持事务消息
func (s *defaultClient) Send(ctx context.Context, topicType TopicType, msg Message) (resp []*rmq_client.SendReceipt, err error) {
	if s.producer == nil {
		return nil, errors.New("请先初始化生产者")
	}

	if topicType == TopicTransaction {
		return nil, errors.New("此方法不支持发送Transaction消息")
	}

	message, err := initMsg(topicType, msg)
	if err != nil {
		return
	}

	resp, err = s.producer.Send(ctx, message)
	return
}

type SendAsyncDealFunc func(ctx context.Context, msg Message, resp []*rmq_client.SendReceipt, err error)

// SendAsync 异步发送消息
// 可支持普通、延迟、顺序类型的消息，不支持事务消息
func (s *defaultClient) SendAsync(ctx context.Context, topicType TopicType, msg Message, dealFunc SendAsyncDealFunc) error {
	if s.producer == nil {
		return errors.New("请先初始化生产者")
	}

	if dealFunc == nil {
		return errors.New("dealFunc必填")
	}

	if topicType == TopicTransaction {
		return errors.New("此方法不支持发送Transaction消息")
	}

	message, err := initMsg(topicType, msg)
	if err != nil {
		return err
	}

	s.producer.SendAsync(ctx, message, func(ctx context.Context, receipts []*rmq_client.SendReceipt, err error) {
		dealFunc(ctx, msg, receipts, err)
	})
	return nil
}

// ConfirmFunc 二次确认方法
// 注意：不要异步处理，本地事务逻辑提交时返回true，否则返回false
type ConfirmFunc func(msg Message, resp []*rmq_client.SendReceipt) bool

// SendTransaction 发送事务消息
// 注意：事务消息的生产者不能和其他类型消息的生产者共用
func (s *defaultClient) SendTransaction(ctx context.Context, message Message, confirmFunc ConfirmFunc) error {
	if s.producer == nil {
		return errors.New("请先初始化生产者")
	}

	if confirmFunc == nil {
		return errors.New("confirmFunc必填")
	}

	msg, err := initMsg(TopicTransaction, message)
	if err != nil {
		return err
	}

	transaction := s.producer.BeginTransaction()
	resp, err := s.producer.SendWithTransaction(ctx, msg, transaction)
	if err != nil {
		return err
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

type ConsumerOptions struct {
	AwaitDuration     time.Duration                           //消息处理超时时间，超时会触发消费重试
	MaxMessageNum     int32                                   //每次接收的消息数量
	InvisibleDuration time.Duration                           //接收到的消息的不可见时间
	SubExpressions    map[string]*rmq_client.FilterExpression //订阅表达式，key为topic，简单消费类型只支持tag和sql匹配
}

// ConsumeFunc 消费方法
// 方法内消费成功时需要调用consumer.Ack()；
// 消费时间可能超过消费者MaxMessageNum设置的时间时，可调用consumer.ChangeInvisibleDuration()或consumer.ChangeInvisibleDurationAsync()方法调整消息消费超时时间；
type ConsumeFunc func(ctx context.Context, msg *rmq_client.MessageView, consumer rmq_client.SimpleConsumer) bool

// SimpleConsume 简单消费类型消费
func (s *defaultClient) SimpleConsume(ctx context.Context, consumeFunc ConsumeFunc, oFunc ...ConsumerOptionFunc) error {
	o := ConsumerOptions{
		AwaitDuration: time.Second * 5,
		MaxMessageNum: 10,
	}
	options := &o
	if len(oFunc) > 0 {
		for _, f := range oFunc {
			f(options)
		}
	}

	consumer, err := rmq_client.NewSimpleConsumer(
		s.Cfg,
		rmq_client.WithAwaitDuration(options.AwaitDuration),
		rmq_client.WithSubscriptionExpressions(options.SubExpressions),
	)
	if err != nil {
		return err
	}

	err = consumer.Start()
	if err != nil {
		return err
	}

	// 优雅的停止
	defer consumer.GracefulStop()

	for {
		mvs, err1 := consumer.Receive(ctx, options.MaxMessageNum, options.InvisibleDuration)
		if err1 != nil {
			return err1
		}
		for _, mv := range mvs {
			consumeFunc(ctx, mv, consumer)
		}
	}
}
