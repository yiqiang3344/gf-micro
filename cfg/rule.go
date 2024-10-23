package cfg

import (
	"github.com/gogf/gf/v2/container/gmap"
	"reflect"
)

var (
	commonRules = []checkOpt{
		{Pattern: APPNAME, Level: MustInput, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: JSONFORMATLOGS, Level: MustInput, Kind: reflect.String, Extra: map[ExtraKey]interface{}{CO: "access"}, Env: []ENV{PROD, DEV}},
		{Pattern: JSONFORMATLOGS, Level: MustInput, Kind: reflect.String, Extra: map[ExtraKey]interface{}{CO: "webclient"}, Env: []ENV{PROD, DEV}},
		{Pattern: ACCESSLOGLENGTHLIMIT, Level: OptionalInput, Kind: reflect.Int64, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER, Level: MustInput, Kind: reflect.Map, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".path", Level: MustInputNotZero, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".file", Level: MustInputNotZero, Kind: reflect.String, Extra: map[ExtraKey]interface{}{EQ: "app.{Ymd}.last.log"}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".prefix", Level: OptionalInput, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".level", Level: MustInputNotZero, Kind: reflect.String, Extra: map[ExtraKey]interface{}{EQ: "all"}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".timeFormat", Level: MustInputNotZero, Kind: reflect.String, Extra: map[ExtraKey]interface{}{EQ: "2006-01-02 15:04:05.000"}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".stdout", Level: MustInputZero, Kind: reflect.Bool, Env: []ENV{PROD}},
		{Pattern: LOGGER + ".stdout", Level: OptionalInput, Kind: reflect.Bool, Env: []ENV{DEV}},
		{Pattern: LOGGER + ".stdoutColorDisabled", Level: OptionalInput, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".rotateSize", Level: MustInputNotZero, Kind: reflect.Int64, Extra: map[ExtraKey]interface{}{LE: 100000000}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".rotateExpire", Level: MustInputNotZero, Kind: reflect.String, Extra: map[ExtraKey]interface{}{EQ: "24h"}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".rotateBackupLimit", Level: MustInputNotZero, Kind: reflect.Int64, Extra: map[ExtraKey]interface{}{EQ: 30}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".rotateBackupExpire", Level: MustInputNotZero, Kind: reflect.String, Extra: map[ExtraKey]interface{}{EQ: "720h"}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".rotateBackupCompress", Level: OptionalInput, Kind: reflect.Int64, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".rotateCheckInterval", Level: MustInputNotZero, Kind: reflect.String, Extra: map[ExtraKey]interface{}{EQ: "1h"}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".access.path", Level: MustInputNotZero, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".access.file", Level: MustInputNotZero, Kind: reflect.String, Extra: map[ExtraKey]interface{}{EQ: "access.{Ymd}.last.log"}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".access.prefix", Level: OptionalInput, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".access.level", Level: MustInputNotZero, Kind: reflect.String, Extra: map[ExtraKey]interface{}{EQ: "all"}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".access.timeFormat", Level: MustInputNotZero, Kind: reflect.String, Extra: map[ExtraKey]interface{}{EQ: "2006-01-02 15:04:05.000"}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".access.stdout", Level: MustInputZero, Kind: reflect.Bool, Env: []ENV{PROD}},
		{Pattern: LOGGER + ".access.stdout", Level: OptionalInput, Kind: reflect.Bool, Env: []ENV{DEV}},
		{Pattern: LOGGER + ".access.stdoutColorDisabled", Level: OptionalInput, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".access.rotateSize", Level: MustInputNotZero, Kind: reflect.Int64, Extra: map[ExtraKey]interface{}{LE: 100000000}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".access.rotateExpire", Level: MustInputNotZero, Kind: reflect.String, Extra: map[ExtraKey]interface{}{EQ: "24h"}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".access.rotateBackupLimit", Level: MustInputNotZero, Kind: reflect.Int64, Extra: map[ExtraKey]interface{}{EQ: 30}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".access.rotateBackupExpire", Level: MustInputNotZero, Kind: reflect.String, Extra: map[ExtraKey]interface{}{EQ: "720h"}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".access.rotateBackupCompress", Level: OptionalInput, Kind: reflect.Int64, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".access.rotateCheckInterval", Level: MustInputNotZero, Kind: reflect.String, Extra: map[ExtraKey]interface{}{EQ: "1h"}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".error.path", Level: MustInputNotZero, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".error.file", Level: MustInputNotZero, Kind: reflect.String, Extra: map[ExtraKey]interface{}{EQ: "error.{Ymd}.last.log"}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".error.prefix", Level: OptionalInput, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".error.level", Level: MustInputNotZero, Kind: reflect.String, Extra: map[ExtraKey]interface{}{EQ: "all"}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".error.timeFormat", Level: MustInputNotZero, Kind: reflect.String, Extra: map[ExtraKey]interface{}{EQ: "2006-01-02 15:04:05.000"}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".error.stdout", Level: MustInputZero, Kind: reflect.Bool, Env: []ENV{PROD}},
		{Pattern: LOGGER + ".error.stdout", Level: OptionalInput, Kind: reflect.Bool, Env: []ENV{DEV}},
		{Pattern: LOGGER + ".error.stdoutColorDisabled", Level: OptionalInput, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".error.rotateSize", Level: MustInputNotZero, Kind: reflect.Int64, Extra: map[ExtraKey]interface{}{LE: 100000000}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".error.rotateExpire", Level: MustInputNotZero, Kind: reflect.String, Extra: map[ExtraKey]interface{}{EQ: "24h"}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".error.rotateBackupLimit", Level: MustInputNotZero, Kind: reflect.Int64, Extra: map[ExtraKey]interface{}{EQ: 30}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".error.rotateBackupExpire", Level: MustInputNotZero, Kind: reflect.String, Extra: map[ExtraKey]interface{}{EQ: "720h"}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".error.rotateBackupCompress", Level: OptionalInput, Kind: reflect.Int64, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".error.rotateCheckInterval", Level: MustInputNotZero, Kind: reflect.String, Extra: map[ExtraKey]interface{}{EQ: "1h"}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".webclient.path", Level: MustInputNotZero, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".webclient.file", Level: MustInputNotZero, Kind: reflect.String, Extra: map[ExtraKey]interface{}{EQ: "webclient.{Ymd}.last.log"}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".webclient.prefix", Level: OptionalInput, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".webclient.level", Level: MustInputNotZero, Kind: reflect.String, Extra: map[ExtraKey]interface{}{EQ: "all"}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".webclient.timeFormat", Level: MustInputNotZero, Kind: reflect.String, Extra: map[ExtraKey]interface{}{EQ: "2006-01-02 15:04:05.000"}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".webclient.stdout", Level: MustInputZero, Kind: reflect.Bool, Env: []ENV{PROD}},
		{Pattern: LOGGER + ".webclient.stdout", Level: OptionalInput, Kind: reflect.Bool, Env: []ENV{DEV}},
		{Pattern: LOGGER + ".webclient.stdoutColorDisabled", Level: OptionalInput, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".webclient.rotateSize", Level: MustInputNotZero, Kind: reflect.Int64, Extra: map[ExtraKey]interface{}{LE: 100000000}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".webclient.rotateExpire", Level: MustInputNotZero, Kind: reflect.String, Extra: map[ExtraKey]interface{}{EQ: "24h"}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".webclient.rotateBackupLimit", Level: MustInputNotZero, Kind: reflect.Int64, Extra: map[ExtraKey]interface{}{EQ: 30}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".webclient.rotateBackupExpire", Level: MustInputNotZero, Kind: reflect.String, Extra: map[ExtraKey]interface{}{EQ: "720h"}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".webclient.rotateBackupCompress", Level: OptionalInput, Kind: reflect.Int64, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".webclient.rotateCheckInterval", Level: MustInputNotZero, Kind: reflect.String, Extra: map[ExtraKey]interface{}{EQ: "1h"}, Env: []ENV{PROD, DEV}},
	}
	apolloRules = []checkOpt{
		{Pattern: APOLLO, Level: MustInput, Kind: reflect.Map, Env: []ENV{PROD, DEV}},
		{Pattern: APOLLO + ".AppID", Level: MustInput, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: APOLLO + ".Cluster", Level: MustInput, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: APOLLO + ".IP", Level: MustInput, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: APOLLO + ".NamespaceName", Level: MustInput, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: APOLLO + ".IsBackupConfig", Level: OptionalInput, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
		{Pattern: APOLLO + ".BackupConfigPath", Level: OptionalInput, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: APOLLO + ".Secret", Level: ProposalNotZero, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: APOLLO + ".SyncServerTimeout", Level: OptionalInput, Kind: reflect.Int64, Env: []ENV{PROD, DEV}},
		{Pattern: APOLLO + ".MustStart", Level: ProposalNotZero, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
		{Pattern: APOLLO + ".Watch", Level: ProposalNotZero, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
	}
	grpcRules = []checkOpt{
		{Pattern: GRPC, Level: MustInput, Kind: reflect.Map, Env: []ENV{PROD, DEV}},
		{Pattern: GRPC + ".address", Level: MustInput, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: GRPC + ".name", Level: MustInput, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: GRPC + ".logStdout", Level: MustInputZero, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
		{Pattern: GRPC + ".errorStack", Level: MustInputZero, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
		{Pattern: GRPC + ".errorLogEnabled", Level: MustInputZero, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
		{Pattern: GRPC + ".accessLogEnabled", Level: MustInputZero, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
		{Pattern: GRPC + ".graceful", Level: MustInputNotZero, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
		{Pattern: GRPC + ".gracefulTimeout", Level: MustInputNotZero, Kind: reflect.Int64, Extra: map[ExtraKey]interface{}{GE: 2}, Env: []ENV{PROD, DEV}},
	}
	serverRules = []checkOpt{
		{Pattern: SERVER, Level: MustInput, Kind: reflect.Map, Env: []ENV{PROD, DEV}},
		{Pattern: SERVER + ".address", Level: MustInput, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: SERVER + ".openapiPath", Level: MustInput, Kind: reflect.String, Env: []ENV{DEV}},
		{Pattern: SERVER + ".openapiPath", Level: MustInputZero, Kind: reflect.String, Env: []ENV{PROD}},
		{Pattern: SERVER + ".swaggerPath", Level: MustInput, Kind: reflect.String, Env: []ENV{DEV}},
		{Pattern: SERVER + ".swaggerPath", Level: MustInputZero, Kind: reflect.String, Env: []ENV{PROD}},
		{Pattern: SERVER + ".logStdout", Level: MustInputZero, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
		{Pattern: SERVER + ".errorStack", Level: MustInputZero, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
		{Pattern: SERVER + ".errorLogEnabled", Level: MustInputZero, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
		{Pattern: SERVER + ".accessLogEnabled", Level: MustInputZero, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
		{Pattern: SERVER + ".graceful", Level: MustInputNotZero, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
		{Pattern: SERVER + ".gracefulTimeout", Level: MustInputNotZero, Kind: reflect.Int64, Extra: map[ExtraKey]interface{}{GE: 2}, Env: []ENV{PROD, DEV}},
	}
	rocketmqRules = []checkOpt{
		{Pattern: ROCKETMQ, Level: MustInput, Kind: reflect.Map, Env: []ENV{PROD, DEV}},
		{Pattern: ROCKETMQ + ".endpoint", Level: MustInputNotZero, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: ROCKETMQ + ".namespace", Level: MustInputNotZero, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: ROCKETMQ + ".accessKey", Level: ProposalNotZero, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: ROCKETMQ + ".accessSecret", Level: ProposalNotZero, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: ROCKETMQ + ".logPath", Level: MustInputNotZero, Kind: reflect.String, Env: []ENV{PROD}},
		{Pattern: ROCKETMQ + ".logPath", Level: OptionalInput, Kind: reflect.String, Env: []ENV{DEV}},
		{Pattern: ROCKETMQ + ".logStdout", Level: MustInputZero, Kind: reflect.Bool, Env: []ENV{PROD}},
		{Pattern: ROCKETMQ + ".logStdout", Level: OptionalInput, Kind: reflect.Bool, Env: []ENV{DEV}},
		{Pattern: ROCKETMQ + ".debug", Level: OptionalInput, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
	}
	xxljobRules = []checkOpt{
		{Pattern: XXLJOB, Level: MustInput, Kind: reflect.Map, Env: []ENV{PROD, DEV}},
		{Pattern: XXLJOB + ".serverAddr", Level: MustInputNotZero, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: XXLJOB + ".token", Level: MustInputNotZero, Kind: reflect.String, Env: []ENV{PROD}},
		{Pattern: XXLJOB + ".token", Level: OptionalInput, Kind: reflect.String, Env: []ENV{DEV}},
	}
	otlpRules = []checkOpt{
		{Pattern: OTLP, Level: MustInput, Kind: reflect.Map, Env: []ENV{PROD, DEV}},
		{Pattern: OTLP + ".endpoint", Level: MustInputNotZero, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: OTLP + ".traceToken", Level: MustInputNotZero, Kind: reflect.String, Env: []ENV{PROD, DEV}},
	}
	registryRules = []checkOpt{
		{Pattern: REGISTRY, Level: MustInput, Kind: reflect.Map, Env: []ENV{PROD, DEV}},
		{Pattern: REGISTRY_GRPC, Level: MustInputNotZero, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: REGISTRY_HTTP, Level: MustInputNotZero, Kind: reflect.String, Env: []ENV{PROD, DEV}},
	}
	redisRules = []checkOpt{
		{Pattern: REDIS, Level: MustInput, Kind: reflect.Map, Env: []ENV{PROD, DEV}},
		{Pattern: REDIS + ".default.address", Level: MustInputNotZero, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: REDIS + ".default.db", Level: MustInputNotZero, Kind: reflect.Int64, Env: []ENV{PROD, DEV}},
		{Pattern: REDIS + ".default.pass", Level: MustInputNotZero, Kind: reflect.String, Env: []ENV{PROD}},
		{Pattern: REDIS + ".default.pass", Level: OptionalInput, Kind: reflect.String, Env: []ENV{DEV}},
	}
	databaseRules = []checkOpt{
		{Pattern: DATABASE, Level: MustInput, Kind: reflect.Map, Env: []ENV{PROD, DEV}},
		{Pattern: DATABASE + ".default.link", Level: MustInputNotZero, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: DATABASE + ".default.debug", Level: ProposalZero, Kind: reflect.Bool, Env: []ENV{PROD}},
		{Pattern: DATABASE + ".default.debug", Level: OptionalInput, Kind: reflect.Bool, Env: []ENV{DEV}},
	}
)

func getRuleListMap() *gmap.ListMap {
	m := gmap.NewListMap()
	m.Set(GRPC, grpcRules)
	m.Set(SERVER, serverRules)
	m.Set(ROCKETMQ, rocketmqRules)
	m.Set(XXLJOB, xxljobRules)
	m.Set(OTLP, otlpRules)
	m.Set(REGISTRY, registryRules)
	m.Set(REDIS, redisRules)
	m.Set(DATABASE, databaseRules)
	return m
}
