package main

import pb "github.com/ziyixi/monorepo/self_host/packages/todofy/proto"

const (
	sender       = "ziyixi@mailjet.ziyixi.science"
	senderName   = "Todofy"
	receiverName = "dida365"
)

var (
	allowedPopullateTodoMethod = map[pb.TodoApp][]pb.PopullateTodoMethod{
		pb.TodoApp_TODO_APP_DIDA365: {pb.PopullateTodoMethod_POPULLATE_TODO_METHOD_MAILJET},
	}
)
