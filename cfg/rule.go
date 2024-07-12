package cfg

import (
	"github.com/gogf/gf/v2/container/gmap"
	"reflect"
)

var (
	commonRules = []checkOpt{
		{Pattern: APPNAME, Level: MustInput, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: JSONFORMATLOGS, Level: MustInput, Kind: reflect.Slice, Extra: map[ExtraKey]interface{}{CO: "access"}, Env: []ENV{PROD, DEV}},
		{Pattern: JSONFORMATLOGS, Level: MustInput, Kind: reflect.Slice, Extra: map[ExtraKey]interface{}{CO: "webclient"}, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER, Level: MustInput, Kind: reflect.Map, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".path", Level: MustInputNotZero, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: LOGGER + ".file", Level: MustInputNotZero, Kind: reflect.String, Extra: map[ExtraKey]interface{}{EQ: "app.log"}, Env: []ENV{PROD, DEV}},
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
		{Pattern: LOGGER + ".access.file", Level: MustInputNotZero, Kind: reflect.String, Extra: map[ExtraKey]interface{}{EQ: "access.log"}, Env: []ENV{PROD, DEV}},
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
		{Pattern: LOGGER + ".error.file", Level: MustInputNotZero, Kind: reflect.String, Extra: map[ExtraKey]interface{}{EQ: "error.log"}, Env: []ENV{PROD, DEV}},
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
		{Pattern: LOGGER + ".webclient.file", Level: MustInputNotZero, Kind: reflect.String, Extra: map[ExtraKey]interface{}{EQ: "webclient.log"}, Env: []ENV{PROD, DEV}},
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
		{Pattern: "apollo", Level: MustInput, Kind: reflect.Map, Env: []ENV{PROD, DEV}},
		{Pattern: "apollo.AppID", Level: MustInput, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: "apollo.Cluster", Level: MustInput, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: "apollo.IP", Level: MustInput, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: "apollo.NamespaceName", Level: MustInput, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: "apollo.IsBackupConfig", Level: OptionalInput, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
		{Pattern: "apollo.BackupConfigPath", Level: OptionalInput, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: "apollo.Secret", Level: ProposalNotZero, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: "apollo.SyncServerTimeout", Level: OptionalInput, Kind: reflect.Int64, Env: []ENV{PROD, DEV}},
		{Pattern: "apollo.MustStart", Level: ProposalNotZero, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
		{Pattern: "apollo.Watch", Level: ProposalNotZero, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
	}
	grpcRules = []checkOpt{
		{Pattern: "grpc", Level: MustInput, Kind: reflect.Map, Env: []ENV{PROD, DEV}},
		{Pattern: "grpc.address", Level: MustInput, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: "grpc.name", Level: MustInput, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: "grpc.logStdout", Level: MustInputZero, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
		{Pattern: "grpc.errorStack", Level: MustInputZero, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
		{Pattern: "grpc.errorLogEnabled", Level: MustInputZero, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
		{Pattern: "grpc.accessLogEnabled", Level: MustInputZero, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
		{Pattern: "grpc.graceful", Level: MustInputNotZero, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
		{Pattern: "grpc.gracefulTimeout", Level: MustInputNotZero, Kind: reflect.Int64, Extra: map[ExtraKey]interface{}{GE: 2}, Env: []ENV{PROD, DEV}},
	}
	serverRules = []checkOpt{
		{Pattern: "server", Level: MustInput, Kind: reflect.Map, Env: []ENV{PROD, DEV}},
		{Pattern: "server.address", Level: MustInput, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: "server.openapiPath", Level: MustInput, Kind: reflect.String, Env: []ENV{DEV}},
		{Pattern: "server.openapiPath", Level: MustInputZero, Kind: reflect.String, Env: []ENV{PROD}},
		{Pattern: "server.swaggerPath", Level: MustInput, Kind: reflect.String, Env: []ENV{DEV}},
		{Pattern: "server.swaggerPath", Level: MustInputZero, Kind: reflect.String, Env: []ENV{PROD}},
		{Pattern: "server.logStdout", Level: MustInputZero, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
		{Pattern: "server.errorStack", Level: MustInputZero, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
		{Pattern: "server.errorLogEnabled", Level: MustInputZero, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
		{Pattern: "server.accessLogEnabled", Level: MustInputZero, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
		{Pattern: "server.graceful", Level: MustInputNotZero, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
		{Pattern: "server.gracefulTimeout", Level: MustInputNotZero, Kind: reflect.Int64, Extra: map[ExtraKey]interface{}{GE: 2}, Env: []ENV{PROD, DEV}},
	}
	rocketmqRules = []checkOpt{
		{Pattern: "rocketmq", Level: MustInput, Kind: reflect.Map, Env: []ENV{PROD, DEV}},
		{Pattern: "rocketmq.endpoint", Level: MustInputNotZero, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: "rocketmq.namespace", Level: MustInputNotZero, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: "rocketmq.accessKey", Level: ProposalNotZero, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: "rocketmq.accessSecret", Level: ProposalNotZero, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: "rocketmq.logPath", Level: MustInputNotZero, Kind: reflect.String, Env: []ENV{PROD}},
		{Pattern: "rocketmq.logPath", Level: OptionalInput, Kind: reflect.String, Env: []ENV{DEV}},
		{Pattern: "rocketmq.logStdout", Level: MustInputZero, Kind: reflect.Bool, Env: []ENV{PROD}},
		{Pattern: "rocketmq.logStdout", Level: OptionalInput, Kind: reflect.Bool, Env: []ENV{DEV}},
		{Pattern: "rocketmq.debug", Level: OptionalInput, Kind: reflect.Bool, Env: []ENV{PROD, DEV}},
	}
	xxljobRules = []checkOpt{
		{Pattern: "xxljob", Level: MustInput, Kind: reflect.Map, Env: []ENV{PROD, DEV}},
		{Pattern: "xxljob.serverAddr", Level: MustInputNotZero, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: "xxljob.token", Level: MustInputNotZero, Kind: reflect.String, Env: []ENV{PROD}},
		{Pattern: "xxljob.token", Level: OptionalInput, Kind: reflect.String, Env: []ENV{DEV}},
	}
	otlpRules = []checkOpt{
		{Pattern: "otlp", Level: MustInput, Kind: reflect.Map, Env: []ENV{PROD, DEV}},
		{Pattern: "otlp.endpoint", Level: MustInputNotZero, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: "otlp.traceToken", Level: MustInputNotZero, Kind: reflect.String, Env: []ENV{PROD, DEV}},
	}
	registryRules = []checkOpt{
		{Pattern: "registry", Level: MustInput, Kind: reflect.Map, Env: []ENV{PROD, DEV}},
		{Pattern: "registry.grpcEtcd", Level: MustInputNotZero, Kind: reflect.String, Env: []ENV{PROD, DEV}},
	}
	redisRules = []checkOpt{
		{Pattern: "redis", Level: MustInput, Kind: reflect.Map, Env: []ENV{PROD, DEV}},
		{Pattern: "redis.default.address", Level: MustInputNotZero, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: "redis.default.db", Level: MustInputNotZero, Kind: reflect.Int64, Env: []ENV{PROD, DEV}},
		{Pattern: "redis.default.pass", Level: MustInputNotZero, Kind: reflect.String, Env: []ENV{PROD}},
		{Pattern: "redis.default.pass", Level: OptionalInput, Kind: reflect.String, Env: []ENV{DEV}},
	}
	databaseRules = []checkOpt{
		{Pattern: "database", Level: MustInput, Kind: reflect.Map, Env: []ENV{PROD, DEV}},
		{Pattern: "database.default.link", Level: MustInputNotZero, Kind: reflect.String, Env: []ENV{PROD, DEV}},
		{Pattern: "database.default.debug", Level: ProposalZero, Kind: reflect.Bool, Env: []ENV{PROD}},
		{Pattern: "database.default.debug", Level: OptionalInput, Kind: reflect.Bool, Env: []ENV{DEV}},
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
