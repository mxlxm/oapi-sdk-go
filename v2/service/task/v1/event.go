// Package task code generated by lark suite oapi sdk gen
package task

import (
	"context"
	"github.com/larksuite/oapi-sdk-go/v2"
)

type taskUpdatedEventHandler struct {
	handler func(context.Context, *lark.RawRequest, *TaskUpdatedEvent) error
}

func (h *taskUpdatedEventHandler) Event() interface{} {
	return &TaskUpdatedEvent{}
}

func (h *taskUpdatedEventHandler) Handle(ctx context.Context, req *lark.RawRequest, event interface{}) error {
	return h.handler(ctx, req, event.(*TaskUpdatedEvent))
}

func (t *tasks) UpdatedEventHandler(handler func(ctx context.Context, req *lark.RawRequest, event *TaskUpdatedEvent) error) {
	t.app.Webhook.EventHandler("task.task.updated_v1", &taskUpdatedEventHandler{handler: handler})
}

type taskCommentUpdatedEventHandler struct {
	handler func(context.Context, *lark.RawRequest, *TaskCommentUpdatedEvent) error
}

func (h *taskCommentUpdatedEventHandler) Event() interface{} {
	return &TaskCommentUpdatedEvent{}
}

func (h *taskCommentUpdatedEventHandler) Handle(ctx context.Context, req *lark.RawRequest, event interface{}) error {
	return h.handler(ctx, req, event.(*TaskCommentUpdatedEvent))
}

func (c *comments) UpdatedEventHandler(handler func(ctx context.Context, req *lark.RawRequest, event *TaskCommentUpdatedEvent) error) {
	c.app.Webhook.EventHandler("task.task.comment.updated_v1", &taskCommentUpdatedEventHandler{handler: handler})
}
