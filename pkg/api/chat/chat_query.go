package chat

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	entv1 "github.com/llmos-ai/llmos-controller/pkg/generated/ent"
	"github.com/llmos-ai/llmos-controller/pkg/generated/ent/chat"
)

type Handler struct {
	client *entv1.Client
	ctx    context.Context
}

func NewHandler(c *entv1.Client, ctx context.Context) Handler {
	return Handler{
		client: c,
		ctx:    ctx,
	}
}

func (h *Handler) ListAll() (entv1.Chats, error) {
	chats, err := h.client.Chat.Query().All(h.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list chats: %v", err)
	}
	logrus.Debugf("Listing charts, total found %d", len(chats))
	return chats, nil
}

func (h *Handler) ListByUser(uid string) ([]*entv1.Chat, error) {
	uuid, err := uuid.Parse(uid)
	if err != nil {
		return nil, fmt.Errorf("failed parsing user uid: %v", err)
	}

	chats, err := h.client.Chat.
		Query().
		Where(chat.UserId(uuid)).
		All(h.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list all chats by user id %s : %w", uid, err)
	}
	logrus.Debugf("Found total %d chats for user: %s", len(chats), uid)
	return chats, nil
}

func (h *Handler) Create(uid string, req NewChatRequest) (*entv1.Chat, error) {
	userId, err := uuid.Parse(uid)
	if err != nil {
		return nil, fmt.Errorf("failed parsing user uid: %v", err)
	}
	return h.client.Chat.
		Create().
		SetTitle(req.Title).
		SetHistory(req.History).
		SetMessages(req.Messages).
		SetModels(req.Models).
		SetTags(req.Tags).
		SetUserId(userId).
		Save(h.ctx)
}

func (h *Handler) Update(id string, req UpdateChatRequest) (*entv1.Chat, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("failed parsing user uid: %v", err)
	}
	client := h.client.Chat.UpdateOneID(uid).
		SetNillableHistory(&req.History).
		SetNillableTitle(&req.Title)

	if req.Messages != nil || len(req.Messages) > 0 {
		client.SetMessages(req.Messages)
	}

	return client.Save(h.ctx)
}

func (h *Handler) FindByID(id, userId string) (*entv1.Chat, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("failed parsing user uid: %v", err)
	}

	uuid, err := uuid.Parse(userId)
	if err != nil {
		return nil, fmt.Errorf("failed parsing user uid: %v", err)
	}

	chat, err := h.client.Chat.
		Query().
		Where(chat.ID(uid)).
		Where(chat.UserId(uuid)).
		First(h.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying chat %s: %v", id, err)
	}
	return chat, nil
}

func (h *Handler) Delete(id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("failed parsing user uid: %v", err)
	}

	return h.client.Chat.
		DeleteOneID(uid).
		Exec(h.ctx)
}
