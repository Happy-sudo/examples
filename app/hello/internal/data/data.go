package data

import (
	"context"
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/cloudwego/kitex/pkg/klog"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"hello/internal/conf"
	"hello/internal/data/ent"
	"time"
)

var ProviderSet = wire.NewSet(NewData, NewHelloRepo, NewDBClient, NewRedisClient)

type Data struct {
	DB    *ent.Client
	Redis *redis.Client
}

// NewData mysql 连接池
func NewData(log klog.CtxLogger, entDbClient *ent.Client, redisClient *redis.Client) (*Data, func(), error) {

	d := &Data{
		DB:    entDbClient,
		Redis: redisClient,
	}
	return d, func() {
		ctx := context.Background()
		if err := d.Redis.Close(); err != nil {
			log.CtxErrorf(ctx, "redis close failed, err:", err.Error())
		}
		if err := d.DB.Close(); err != nil {
			log.CtxErrorf(context.Background(), "mysql close failed, err:", err)
		}
		log.CtxInfof(context.Background(), "closing the data resources")
	}, nil
}

// NewDBClient mysql 连接池
func NewDBClient(log klog.CtxLogger, config *conf.Config) *ent.Client {

	if config.MysqlOptions.Enable {
		//创建数据库实例
		db, err := sql.Open(config.MysqlOptions.Driver, config.MysqlOptions.Source)
		if err != nil {
			log.CtxFatalf(context.Background(), "failed opening connection to mysql: %v", err)
		}
		// 配置连接池
		db.SetMaxIdleConns(config.MysqlOptions.MaxIdleConns)
		db.SetConnMaxLifetime(time.Duration(config.MysqlOptions.ConnMaxLifetime) * time.Second)
		db.SetConnMaxIdleTime(time.Duration(config.MysqlOptions.ConnMaxIdleTime) * time.Second)
		db.SetMaxOpenConns(config.MysqlOptions.MaxOpenConns)

		// 测试数据库连接
		err = db.Ping()
		if err != nil {
			log.CtxFatalf(context.Background(), "failed connection to mysql: %v", err)
		}
		log.CtxInfof(context.Background(), "mysql connection successful")

		client := ent.NewClient(ent.Driver(entsql.OpenDB(config.MysqlOptions.Driver, db)))
		if err := client.Schema.Create(context.Background()); err != nil {
			log.CtxErrorf(context.Background(), "err:%s", err.Error())
		}
		return client
	}

	log.CtxInfof(context.Background(), "mysql 未配置")

	return nil
}

// NewRedisClient redis 连接池
func NewRedisClient(log klog.CtxLogger, config *conf.Config) *redis.Client {

	if config.RedisOptions.Enable {
		rdb := redis.NewClient(&redis.Options{
			Network:  config.RedisOptions.Network,
			Addr:     config.RedisOptions.Addr,
			Username: config.RedisOptions.Username,
			Password: config.RedisOptions.Password,
			DB:       int(config.RedisOptions.DB),

			//连接池
			PoolSize:     int(config.RedisOptions.PoolSize),
			MinIdleConns: int(config.RedisOptions.MinIdleConns),
			MaxIdleConns: config.RedisOptions.MaxIdleConns,
			//超时配置
			DialTimeout:  time.Duration(config.RedisOptions.DialTimeout),
			ReadTimeout:  time.Duration(config.RedisOptions.ReadTimeout),
			WriteTimeout: time.Duration(config.RedisOptions.WriteTimeout),
			PoolTimeout:  time.Duration(config.RedisOptions.PoolTimeout),

			ConnMaxIdleTime: time.Duration(config.RedisOptions.ConnMaxIdleTime),
			ConnMaxLifetime: time.Duration(config.RedisOptions.ConnMaxLifetime),

			//命令执行失败时的重试策略
			MaxRetries:      int(config.RedisOptions.MaxRetries),
			MinRetryBackoff: time.Duration(config.RedisOptions.MinRetryBackoff),
			MaxRetryBackoff: time.Duration(config.RedisOptions.MaxRetryBackoff),
		})

		if _, err := rdb.Ping(context.Background()).Result(); err != nil {
			log.CtxFatalf(context.Background(), "failed connection to redis: %v", err)
		}
		log.CtxInfof(context.Background(), "redis connection successful")
		return rdb
	}

	log.CtxInfof(context.Background(), "redis 未配置")

	return nil
}
