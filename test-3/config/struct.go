package config

// ConfigStruct : struct config
type ConfigStruct struct {
	Logger      Logger                 `mapstructure:"logger"`
	Databases   DatabasesConfig        `mapstructure:"databases"`
	Controllers ControllerConfig       `mapstructure:"controller"`
	Etc         map[string]interface{} `mapstructure:"etc"`
}

// Logger :
type Logger struct {
	Level       string `mapstructure:"level"`
	FileLog     bool   `mapstructure:"file_log"`
	ConsoleLog  bool   `mapstructure:"console_log"`
	LogAsJson   bool   `mapstructure:"log_as_json"`
	MaxSize     int    `mapstructure:"max_size"`
	MaxBackups  int    `mapstructure:"max_backups"`
	MaxAge      int    `mapstructure:"max_age"`
	DailyRotate bool   `mapstructure:"daily_rotate"`
	CompressLog bool   `mapstructure:"compress_log"`
}

// DatabasesConfig :
type DatabasesConfig struct {
	Redis   RedisConfig   `mapstructure:"redis"`
	Mongo   MongoConfig   `mapstructure:"mongo"`
	Postgre PostgreConfig `mapstructure:"postgre"`
	Oracle  OracleConfig  `mapstructure:"oracle"`
	Sqlite  SqliteConfig  `mapstructure:"sqlite"`
}

// ControllerConfig :
type ControllerConfig struct {
	CronJob    CronJobConfig    `mapstructure:"cronjob"`
	Daemon     DaemonConfig     `mapstructure:"daemon"`
	Restapi    RestapiConfig    `mapstructure:"restapi"`
	Grpc       GrpcConfig       `mapstructure:"grpc"`
	Kafka      KafkaConfig      `mapstructure:"kafka"`
	Prometheus PrometheusConfig `mapstructure:"prometheus"`
}

// RedisConfig :
type RedisConfig struct {
	Host      string `mapstructure:"host"`
	Auth      string `mapstructure:"auth"`
	DB        int    `mapstructure:"db"`
	MaxIdle   int    `mapstructure:"max_idle"`
	MaxActive int    `mapstructure:"max_active"`
}

type SqliteConfig struct {
	Path string `mapstructure:"path"`
	File string `mapstructure:"file"`
}

// MongoConfig : mongo config
type MongoConfig struct {
	Host    map[int]string `mapstructure:"host"`
	User    string         `mapstructure:"user"`
	Pass    string         `mapstructure:"pass"`
	DB      string         `mapstructure:"db"`
	Srv     bool           `mapstructure:"srv"`
	Cluster bool           `mapstructure:"cluster"`
	RsName  string         `mapstructure:"rs_name"`
	Timeout int            `mapstructure:"timeout"`
}

// PostgreConfig :
type PostgreConfig struct {
	Host    string `mapstructure:"host"`
	Port    int    `mapstructure:"port"`
	User    string `mapstructure:"user"`
	Pass    string `mapstructure:"pass"`
	DB      string `mapstructure:"db"`
	Schema  string `mapstructure:"schema"`
	MaxPool int    `mapstructure:"max_pool"`
}

// OracleConfig :
type OracleConfig struct {
	Host    map[int]string `mapstructure:"host"`
	User    string         `mapstructure:"user"`
	Pass    string         `mapstructure:"pass"`
	DB      string         `mapstructure:"db"`
	Timeout string         `mapstructure:"timeout"`
}

// CronJobConfig :
type CronJobConfig struct {
	Jobs map[string]CronJobJobsConfig `mapstructure:"jobs"`
}

// CronJobJobsConfig :
type CronJobJobsConfig struct {
	Every string   `mapstructure:"every"`
	Hours []string `mapstructure:"hours"`
}

// DaemonConfig :
type DaemonConfig struct {
	Sleep     int  `mapstructure:"sleep"`
	WaitGroup bool `mapstructure:"waitGroup"`
}

// RestapiConfig : REST API Config
type RestapiConfig struct {
	Port     string            `mapstructure:"port"`
	BasePath string            `mapstructure:"base_path"`
	Swagger  SwaggerConfig     `mapstructure:"swagger"`
	Cors     map[string]string `mapstructure:"cors"`
}

// GrpcConfig : Grpc Config
type GrpcConfig struct {
	Port string `mapstructure:"port"`
}

// SwaggerConfig : REST API Swagger Config
type SwaggerConfig struct {
	Title       string   `mapstructure:"title"`
	Description string   `mapstructure:"description"`
	Schemes     []string `mapstructure:"schemes"`
}

// KafkaConfig : kafka config
type KafkaConfig struct {
	Brokers           string               `mapstructure:"brokers"`
	Assignor          string               `mapstructure:"assignor"`
	Version           string               `mapstructure:"version"`
	Verbose           bool                 `mapstructure:"verbose"`
	DialTimeout       int                  `mapstructure:"dial_timeout"`
	TableReplication  int                  `mapstructure:"table_replication"`
	StreamReplication int                  `mapstructure:"stream_replication"`
	PartitionNumber   int                  `mapstructure:"partition_number"`
	ShutdownTimeout   int                  `mapstructure:"shutdown_timeout"`
	Consumer          KafkaConsumerConfig  `mapstructure:"consumer"`
	Publisher         KafkaPublisherConfig `mapstructure:"publisher"`
}

// KafkaConsumerConfig :
type KafkaConsumerConfig struct {
	Oldest bool `mapstructure:"oldest"`
}

// KafkaPublisherConfig :
type KafkaPublisherConfig struct {
	RetryMax   int  `mapstructure:"retrymax"`
	Timeout    int  `mapstructure:"timeout"`
	Idempotent bool `mapstructure:"idempotent"`
}

// PrometheusConfig :
type PrometheusConfig struct {
}
