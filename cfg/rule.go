package cfg

import "reflect"

var (
	commonRules = []checkOpt{
		{Pattern: "appName", Level: MustInput, Kind: reflect.String},
		{Pattern: "jsonFormatLogs", Level: MustInput, Kind: reflect.Slice},
		{Pattern: "logger", Level: MustInput, Kind: reflect.Map},
	}
	apolloRules = []checkOpt{
		{Pattern: "apollo", Level: MustInput, Kind: reflect.Map},
		{Pattern: "apollo.AppID", Level: MustInput, Kind: reflect.String},
		{Pattern: "apollo.Cluster", Level: MustInput, Kind: reflect.String},
		{Pattern: "apollo.IP", Level: MustInput, Kind: reflect.String},
		{Pattern: "apollo.NamespaceName", Level: MustInput, Kind: reflect.String},
		{Pattern: "apollo.IsBackupConfig", Level: OptionalInput, Kind: reflect.Bool},
		{Pattern: "apollo.BackupConfigPath", Level: OptionalInput, Kind: reflect.String},
		{Pattern: "apollo.Secret", Level: ProposalNotEmpty, Kind: reflect.String},
		{Pattern: "apollo.SyncServerTimeout", Level: OptionalInput, Kind: reflect.Int},
		{Pattern: "apollo.MustStart", Level: ProposalNotEmpty, Kind: reflect.Bool},
		{Pattern: "apollo.Watch", Level: ProposalNotEmpty, Kind: reflect.Bool},
	}
	grpcRules = []checkOpt{
		{Pattern: "grpc", Level: MustInput, Kind: reflect.Map},
	}
	serverRules = []checkOpt{
		{Pattern: "grpc", Level: MustInput, Kind: reflect.Map},
		{Pattern: "grpc.address", Level: MustInput, Kind: reflect.String},
		{Pattern: "grpc.name", Level: MustInput, Kind: reflect.String},
		{Pattern: "grpc.logStdout", Level: MustInputZero, Kind: reflect.Bool},
		{Pattern: "grpc.errorStack", Level: MustInputZero, Kind: reflect.Bool},
		{Pattern: "grpc.errorLogEnabled", Level: MustInputZero, Kind: reflect.Bool},
		{Pattern: "grpc.accessLogEnabled", Level: MustInputZero, Kind: reflect.Bool},
		{Pattern: "grpc.graceful", Level: MustInputNotZero, Kind: reflect.Bool},
		{Pattern: "grpc.gracefulTimeout", Level: MustInputNotZero, Kind: reflect.Int, Extra: map[ExtraKey]interface{}{GE: 2}},
	}
	rocketmqRules = []checkOpt{
		{Pattern: "grpc", Level: MustInput, Kind: reflect.Map},
	}
	xxljobRules = []checkOpt{
		{Pattern: "grpc", Level: MustInput, Kind: reflect.Map},
	}
	otlpRules = []checkOpt{
		{Pattern: "grpc", Level: MustInput, Kind: reflect.Map},
	}
	registryRules = []checkOpt{
		{Pattern: "grpc", Level: MustInput, Kind: reflect.Map},
	}
	redisRules = []checkOpt{
		{Pattern: "grpc", Level: MustInput, Kind: reflect.Map},
	}
	databaseRules = []checkOpt{
		{Pattern: "grpc", Level: MustInput, Kind: reflect.Map},
	}
)
