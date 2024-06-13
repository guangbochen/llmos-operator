package chat

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rancher/apiserver/pkg/store/empty"
	"github.com/rancher/apiserver/pkg/types"
	"github.com/rancher/wrangler/v2/pkg/schemas/validation"
	"github.com/sirupsen/logrus"

	"github.com/llmos-ai/llmos-controller/pkg/constant"
	entv1 "github.com/llmos-ai/llmos-controller/pkg/generated/ent"
	"github.com/llmos-ai/llmos-controller/pkg/utils"
)

type Store struct {
	empty.Store
	handler Handler
}

func toAPIObject(c *entv1.Chat) types.APIObject {
	return types.APIObject{
		Type:   "chat",
		ID:     c.ID.String(),
		Object: c,
	}
}

func (s *Store) ByID(apiOp *types.APIRequest, schema *types.APISchema, id string) (types.APIObject, error) {
	userInfo, ok := apiOp.GetUserInfo()
	if !ok {
		return types.APIObject{}, fmt.Errorf("failed to get user info")
	}

	chat, err := s.handler.FindByID(id, userInfo.GetUID())
	if err != nil {
		return types.APIObject{}, err
	}

	return types.APIObject{
		Type:   "chat",
		ID:     chat.ID.String(),
		Object: chat,
	}, nil
}

func (s *Store) List(apiOp *types.APIRequest, schema *types.APISchema) (types.APIObjectList, error) {
	userInfo, ok := apiOp.GetUserInfo()
	if !ok {
		return types.APIObjectList{}, fmt.Errorf("failed to get user info")
	}

	// check if user is admin and listing all chats
	if utils.ArrayStringContains(userInfo.GetGroups(), constant.AdminRole) {
		return s.list("")
	}

	return s.list(userInfo.GetUID())
}

func (s *Store) list(uid string) (types.APIObjectList, error) {
	var chats = make([]*entv1.Chat, 0)
	var err error
	if uid != "" {
		chats, err = s.handler.ListByUser(uid)
	} else {
		chats, err = s.handler.ListAll()
		if err != nil {
			return types.APIObjectList{}, err
		}
	}

	objs := make([]types.APIObject, 0)

	for _, chat := range chats {
		objs = append(objs, toAPIObject(chat))
	}

	return types.APIObjectList{
		Objects: objs,
	}, nil
}

func (s *Store) Create(apiOp *types.APIRequest, schema *types.APISchema, data types.APIObject) (types.APIObject, error) {
	userInfo, ok := apiOp.GetUserInfo()
	if !ok {
		return types.APIObject{}, fmt.Errorf("failed to get user info")
	}

	logrus.Debugf("create new chat: %+v", data)
	//newChat := data.Object.(entv1.Chat)
	// Encode the map to JSON
	jsonData, err := json.Marshal(data.Object)
	if err != nil {
		return types.APIObject{}, fmt.Errorf("failed to encode data to json: %v", err)
	}

	// Decode the JSON into the struct
	var req NewChatRequest
	err = json.Unmarshal(jsonData, &req)
	if err != nil {
		return types.APIObject{}, fmt.Errorf("failed to decode data from json: %v", err)
	}
	logrus.Debugf("create new chat: %+v", req)

	chat, err := s.handler.Create(userInfo.GetUID(), req)
	if err != nil {
		return types.APIObject{}, fmt.Errorf("failed to create chat: %v", err)
	}
	return types.APIObject{
		Type:   "chat",
		ID:     chat.ID.String(),
		Object: chat,
	}, nil
}

func getRawData(request *http.Request) ([]byte, error) {
	return io.ReadAll(request.Body)
}

func (s *Store) Update(apiOp *types.APIRequest, schema *types.APISchema, data types.APIObject, id string) (types.APIObject, error) {
	return types.APIObject{}, validation.NotFound
}

func (s *Store) Delete(apiOp *types.APIRequest, schema *types.APISchema, id string) (types.APIObject, error) {
	err := s.handler.Delete(id)
	if err != nil {
		return types.APIObject{}, err
	}
	return types.APIObject{}, nil
}
