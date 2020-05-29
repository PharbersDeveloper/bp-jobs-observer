package oss_task_start

import (
	"github.com/PharbersDeveloper/bp-jobs-observer/observers"
	ossTask "github.com/PharbersDeveloper/bp-jobs-observer/observers/oss_task_start/detail"
	"github.com/hashicorp/go-uuid"
	"os"
	"strconv"
)

const (
	EntryValue          = "oss-task-start"
)

func Run() {

	DbHost := os.Getenv(observers.DbHostKey)
	if DbHost == "" {
		println("Error! No DB_HOST env set.")
		return
	}
	DbPort := os.Getenv(observers.DbPortKey)
	if DbPort == "" {
		println("Error! No DB_PORT env set.")
		return
	}
	DbName := os.Getenv(observers.DbNameKey)
	if DbName == "" {
		println("Error! No DB_NAME env set.")
		return
	}
	DbColl := os.Getenv(observers.DbCollKey)
	if DbColl == "" {
		println("Error! No DB_COLL env set.")
		return
	}
	ParallelNumStr := os.Getenv(observers.ParallelNumKey)
	if ParallelNumStr == "" {
		println("Error! No PARALLEL_NUM env set.")
		return
	}
	ParallelNum, err := strconv.Atoi(ParallelNumStr)
	if err != nil {
		panic(err.Error())
	}
	ReqTopic := os.Getenv(observers.ReqTopicKey)
	if ReqTopic == "" {
		println("Error! No REQ_TOPIC env set.")
		return
	}

	newId, _ := uuid.GenerateUUID()
	//TODO: Conditions 配置抽离
	bpjo := ossTask.ObserverInfo{
		Id:         newId,
		DBHost:     DbHost,
		DBPort:     DbPort,
		Database:   DbName,
		Collection: DbColl,
		Conditions: map[string]interface{}{
			"$and": []map[string]interface{}{
				map[string]interface{}{"file": map[string]interface{}{"$exists": true, "$ne": ""}},
				map[string]interface{}{"isNewVersion": true},
				map[string]interface{}{"dfs": map[string]interface{}{"$exists": true, "$size": 0}},
			},
		},
		ParallelNumber: ParallelNum,
		RequestTopic:   ReqTopic,
	}
	bpjo.Open()
	bpjo.Exec()
	bpjo.Close()
}

//本地开发测试使用
func SetEntryEnv() {

	_ = os.Setenv(observers.EntryKey, EntryValue)
	_ = os.Setenv(observers.DbHostKey, "192.168.100.116")
	_ = os.Setenv(observers.DbPortKey, "27017")
	_ = os.Setenv(observers.DbNameKey, "pharbers-sandbox-merge")
	_ = os.Setenv(observers.DbCollKey, "assets")
	_ = os.Setenv(observers.ParallelNumKey, "1")
	_ = os.Setenv(observers.ReqTopicKey, "test528")

}