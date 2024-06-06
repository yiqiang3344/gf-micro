package rocketmq_client

import (
	"context"
	"errors"
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
	v2 "github.com/apache/rocketmq-clients/golang/v5/protocol/v2"
	"os"
	"strings"
	"time"
)

type MqClient struct {
	endpoint string
	producer rmq_client.Producer
}

// Client 获取mq客户端，endpoint 为mq服务器地址加端口，示例：127.0.0.1:9876
func Client(endpoint string) *MqClient {
	// todo 参数配置化，提升扩展性
	return &MqClient{
		endpoint: endpoint,
	}
}

// TopicType 主题类型
type TopicType string

const (
	TopicNormal      TopicType = "NORMAL"
	TopicFIFO        TopicType = "FIFO"
	TopicDelay       TopicType = "DELAY"
	TopicTransaction TopicType = "TRANSACTION"
)

// SendOptions 发送消息时的可选参数，todo 拆分生产者和发送的参数，生产者参数包含是否发送完及销毁配置
type SendOptions struct {
	NameSpace          string                          //命名空间，可选
	ConsumerGroup      string                          //消费者组，可选
	Credentials        *credentials.SessionCredentials //鉴权，可选
	Tag                string                          //标签，可选
	MessageGroup       string                          //消息组，FIFO类型主题用，可选
	Keys               []string                        //key列表，可选
	Properties         map[string]string               //消息属性，可选
	DeliveryTimestamp  *time.Time                      //延时时间，Delay类型主题用
	MaxAttempts        int32                           //重试次数，可选
	Async              bool                            //异步发送
	successFunc        SendSuccessFunc                 //消息成功处理方法，可选
	failedFunc         SendFailedFunc                  //消息失败处理方法，可选
	transactionChecker SendTransactionCheckerFunc      //事务检查器，事务消息必填
	sendEndFunc        SendTransactionEndFunc          //事务消息发送结束处理方法，可选
}

func (s *SendOptions) callSuccessFunc(message string, res []*rmq_client.SendReceipt) {
	if s.successFunc != nil {
		s.successFunc(message, res, *s)
	}
}

func (s *SendOptions) callFailedFunc(message string, err error) {
	if s.failedFunc != nil {
		s.failedFunc(message, err, *s)
	}
}

func (s *SendOptions) callSendEndFunc(message string, res []*rmq_client.SendReceipt) bool {
	if s.sendEndFunc != nil {
		return s.sendEndFunc(message, res, *s)
	}
	return true
}

// SendFailedFunc 发送失败的处理方法
type SendFailedFunc func(message string, err error, options SendOptions)

// SendSuccessFunc 发送成功的处理方法
type SendSuccessFunc func(message string, res []*rmq_client.SendReceipt, options SendOptions)

// SendOptionFunc 发送可选参数配置方法
type SendOptionFunc func(options *SendOptions)

// SendTransactionCheckerFunc 发送事务消息时的检查器
type SendTransactionCheckerFunc func(msg *rmq_client.MessageView) rmq_client.TransactionResolution

// SendTransactionEndFunc 事务消息发送完毕后的事务逻辑处理方法，返回 true 则消息事务commit, 否则 rollback
type SendTransactionEndFunc func(message string, res []*rmq_client.SendReceipt, options SendOptions) bool

func WithSendOptionNameSpace(namespace string) SendOptionFunc {
	return func(o *SendOptions) {
		o.NameSpace = namespace
	}
}
func WithSendOptionConsumerGroup(consumerGroup string) SendOptionFunc {
	return func(o *SendOptions) {
		o.ConsumerGroup = consumerGroup
	}
}
func WithSendOptionCredentials(credentials credentials.SessionCredentials) SendOptionFunc {
	return func(o *SendOptions) {
		o.Credentials = &credentials
	}
}
func WithSendOptionTag(tag string) SendOptionFunc {
	return func(o *SendOptions) {
		o.Tag = tag
	}
}
func WithSendOptionMessageGroup(messageGroup string) SendOptionFunc {
	return func(o *SendOptions) {
		o.MessageGroup = messageGroup
	}
}
func WithSendOptionKeys(keys []string) SendOptionFunc {
	return func(o *SendOptions) {
		o.Keys = keys
	}
}
func WithSendOptionProperties(properties map[string]string) SendOptionFunc {
	return func(o *SendOptions) {
		o.Properties = properties
	}
}
func WithSendOptionDeliveryTimestamp(deliveryTimestamp *time.Time) SendOptionFunc {
	return func(o *SendOptions) {
		o.DeliveryTimestamp = deliveryTimestamp
	}
}
func WithSendOptionMaxAttempts(maxAttempts int32) SendOptionFunc {
	return func(o *SendOptions) {
		o.MaxAttempts = maxAttempts
	}
}
func WithSendOptionAsync(async bool) SendOptionFunc {
	return func(o *SendOptions) {
		o.Async = async
	}
}
func WithSendOptionSuccessFunc(successFunc SendSuccessFunc) SendOptionFunc {
	return func(o *SendOptions) {
		o.successFunc = successFunc
	}
}
func WithSendOptionFailedFunc(failedFunc SendFailedFunc) SendOptionFunc {
	return func(o *SendOptions) {
		o.failedFunc = failedFunc
	}
}
func WithSendOptionSendEndFunc(sendEndFunc SendTransactionEndFunc) SendOptionFunc {
	return func(o *SendOptions) {
		o.sendEndFunc = sendEndFunc
	}
}
func WithSendOptionTransactionChecker(transactionChecker SendTransactionCheckerFunc) SendOptionFunc {
	return func(o *SendOptions) {
		o.transactionChecker = transactionChecker
	}
}

var defaultSendOption = SendOptions{
	NameSpace:     "",
	ConsumerGroup: "",
	Tag:           "",
	MessageGroup:  "",
	Keys:          []string{},
	Properties:    map[string]string{},
	MaxAttempts:   3,
	Async:         false,
}

func (s *MqClient) initMsg(topic, message string, options *SendOptions) *rmq_client.Message {
	//初始化消息体
	msg := &rmq_client.Message{
		Topic: topic,
		Body:  []byte(message),
	}
	//设置消息tag
	if options.Tag != "" {
		msg.SetTag(options.Tag)
	}
	//设置消息组（只有FIFO类型topic可用）
	if options.MessageGroup != "" {
		msg.SetMessageGroup(options.MessageGroup)
	}
	//设置消息key
	msg.SetKeys(options.Keys...)
	//设置消息属性
	if len(options.Properties) > 0 {
		for k, v := range options.Properties {
			msg.AddProperty(k, v)
		}
	}
	//设置延迟时间（只有Delay类型topic可用）
	if options.DeliveryTimestamp != nil {
		msg.SetDelayTimestamp(*options.DeliveryTimestamp)
	}
	return msg
}

// todo 初始化生产者时参数化，可传入topic列表
func (s *MqClient) initProducer(ctx context.Context, topicType TopicType, topic string, options *SendOptions) (producer rmq_client.Producer, err error) {
	switch topicType {
	case TopicNormal, TopicFIFO, TopicDelay:
		producer, err = rmq_client.NewProducer(
			&rmq_client.Config{
				Endpoint:      s.endpoint,
				NameSpace:     options.NameSpace,
				ConsumerGroup: options.ConsumerGroup,
				Credentials:   options.Credentials,
			},
			rmq_client.WithTopics(topic),
			rmq_client.WithMaxAttempts(options.MaxAttempts),
		)
	case TopicTransaction:
		producer, err = rmq_client.NewProducer(
			&rmq_client.Config{
				Endpoint:      s.endpoint,
				NameSpace:     options.NameSpace,
				ConsumerGroup: options.ConsumerGroup,
				Credentials:   options.Credentials,
			},
			rmq_client.WithTopics(topic),
			rmq_client.WithMaxAttempts(options.MaxAttempts),
			rmq_client.WithTransactionChecker(&rmq_client.TransactionChecker{
				Check: options.transactionChecker,
			}),
		)
	}
	return
}

//todo 提供手动销毁生产者的方法

func (s *MqClient) send(ctx context.Context, topicType TopicType, topic string, options *SendOptions, producer rmq_client.Producer, message string) {
	msg := s.initMsg(topic, message, options)

	switch topicType {
	case TopicNormal, TopicFIFO, TopicDelay:
		//发送消息并根据结果回调对应方法
		resp, err := producer.Send(ctx, msg)
		if err != nil {
			options.callFailedFunc(message, err)
		} else {
			options.callSuccessFunc(message, resp)
		}
	case TopicTransaction:
		transaction := producer.BeginTransaction()
		resp, err := producer.SendWithTransaction(ctx, msg, transaction)
		if err != nil {
			options.callFailedFunc(message, err)
			return
		}
		//调用消息发送完毕的事务处理方法，成功（未定义处理方法也算成功）则提交事务，否则回滚事务
		if options.callSendEndFunc(message, resp) {
			err = transaction.Commit()
		} else {
			err = transaction.RollBack()
		}
		if err != nil {
			options.callFailedFunc(message, err)
		} else {
			options.callSuccessFunc(message, resp)
		}
	}

	return
}

// SendMsgs 发送消息，包括Normal、FIFO、Delay、Transaction。 todo 支持发送不同消息类型的消息，提升生产者的使用范围
func (s *MqClient) SendMsgs(ctx context.Context, topicType TopicType, topic string, messages []string, sendOptionFunc ...SendOptionFunc) (err error) {
	if len(messages) == 0 {
		return
	}
	if strings.Trim(topic, "") == "" {
		err = errors.New("topic不能为空")
		return
	}

	//初始化可选项参数
	o := defaultSendOption
	options := &o
	if len(sendOptionFunc) > 0 {
		for _, f := range sendOptionFunc {
			f(options)
		}
	}

	switch topicType {
	case TopicNormal:
	case TopicFIFO:
		if strings.Trim(options.MessageGroup, "") == "" {
			err = errors.New("FIFO消息主题MessageGroup选项不能为空")
			return
		}
	case TopicDelay:
		if options.DeliveryTimestamp == nil {
			err = errors.New("DELAY消息主题DeliveryTimestamp选项不能为空")
			return
		}
	case TopicTransaction:
		if options.transactionChecker == nil {
			err = errors.New("事务消息主题transactionChecker选项不能为空")
			return
		}
	}

	//终端打印日志
	os.Setenv("mq.consoleAppender.enabled", "true")
	rmq_client.InitLogger()

	//初始化生产者
	producer, err := s.initProducer(ctx, topicType, topic, options)
	if err != nil {
		return
	}
	err = producer.Start()
	if err != nil {
		return
	}

	// 优雅的关闭生产者
	defer producer.GracefulStop()

	for _, message := range messages {
		s.send(ctx, topicType, topic, options, producer, message)
	}
	return
}

func IsTooManyRequest(err error) bool {
	//如果是重试失败，则判断是否设置了补偿机制，有则调用
	if e, ok := err.(*rmq_client.ErrRpcStatus); ok && e.GetCode() == int32(v2.Code_TOO_MANY_REQUESTS) {
		return true
	}
	return false
}
