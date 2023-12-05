package worker

import (
	"context"

	"github.com/hibiken/asynq"
	db "github.com/mitidevus/simplebank/db/sqlc"
)

const (
	QueueCritical = "critical"
	QueueDefault  = "default"
)

type TaskProcessor interface {
	Start() error // Register the task before starting the asynq server
	ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error
}

type RedisTaskProcessor struct {
	server *asynq.Server
	store  db.Store
}

func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt, store db.Store) TaskProcessor {
	server := asynq.NewServer(
		redisOpt,
		asynq.Config{
			Queues: map[string]int{
				QueueCritical: 10,
				QueueDefault:  5,
			},
		}, // This config object allow us to control many different parameters of the asynq Server
	)

	return &RedisTaskProcessor{
		server: server,
		store:  store,
	}
}

func (processor *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()

	mux.HandleFunc(TaskSendVerifyEmail, processor.ProcessTaskSendVerifyEmail)

	return processor.server.Start(mux)
}
