package master

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"crontab/common"
)

// mongodb日志管理
type LogMgr struct {
	client *mongo.Client
	logCollection *mongo.Collection
}

var (
	G_logMgr *LogMgr
)

func InitLogMgr() (err error) {
	var (
		client *mongo.Client
	)
	ctx , cancel :=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel() //养成良好的习惯，在调用WithTimeout之后defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(G_config.MongodbUri))
	if err != nil {
		log.Print(err)
	}
	G_logMgr = &LogMgr{
		client: client,
		logCollection: client.Database("cron").Collection("log"),
	}
	return
}



//func InitLogMgr() (err error) {
//	var (
//		client *mongo.Client
//	)
//
//	// 建立mongodb连接
//	if client, err = mongo.Connect(
//		context.TODO(),
//		G_config.MongodbUri,
//		clientopt.ConnectTimeout(time.Duration(G_config.MongodbConnectTimeout) * time.Millisecond)); err != nil {
//		return
//	}
//
//	G_logMgr = &LogMgr{
//		client: client,
//		logCollection: client.Database("cron").Collection("log"),
//	}
//	return
//}

// 查看任务日志
func (logMgr *LogMgr) ListLog(name string, skip int, limit int) (logArr []*common.JobLog, err error){
	var (
		filter *common.JobLogFilter
		logSort *common.SortLogByStartTime
		cursor *mongo.Cursor
		jobLog *common.JobLog
	)

	// len(logArr)
	logArr = make([]*common.JobLog, 0)

	// 过滤条件
	filter = &common.JobLogFilter{JobName: name}

	// 按照任务开始时间倒排
	logSort = &common.SortLogByStartTime{SortOrder: -1}

	// 查询
	option := options.Find()
	option.SetSort(logSort)
	option.SetSkip(int64(skip))
	option.SetLimit(int64(limit))
	if cursor, err = logMgr.logCollection.Find(context.TODO(), filter, option); err != nil {
		return
	}
	// 延迟释放游标
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		jobLog = &common.JobLog{}

		// 反序列化BSON
		if err = cursor.Decode(jobLog); err != nil {
			continue // 有日志不合法
		}

		logArr = append(logArr, jobLog)
	}
	return
}