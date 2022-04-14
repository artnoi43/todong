package store

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"

	"github.com/artnoi43/todong/data/model"
	"github.com/artnoi43/todong/domain/usecase"
	"github.com/artnoi43/todong/lib/redishelper"
)

type redisDataGateway struct {
	client redisDB
}

var (
	usersKey = "users"
)

func NewRedisDataGateway(rdb redisDB) usecase.DataGateway {
	return &redisDataGateway{
		client: rdb,
	}
}

func handleRedisNil(err error) error {
	if err != nil && errors.Is(err, redis.Nil) {
		return ErrRecordNotFound
	}
	return err
}

func (r *redisDataGateway) DeleteUser(
	ctx context.Context,
	user *model.User,
) error {
	k, _ := redishelper.FromUser(user)
	key := k.String()
	username, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	// Delete user (hash store) from key "users"
	if _, err := r.client.HDel(ctx, usersKey, username).Result(); err != nil {
		return err
	}

	todoKey := redishelper.TodoKey{
		UserUUID: k.UUID,
	}

	// Delete user's todos and username from Redis
	if _, err := r.client.Del(ctx, todoKey.String(), k.String()).Result(); err != nil {
		return err
	}

	return handleRedisNil(err)
}

func (r *redisDataGateway) DeleteTodo(
	ctx context.Context,
	todo *model.Todo,
) error {
	k := redishelper.KeyFromTodo(todo)
	kStr := k.String()
	_, err := r.client.HDel(ctx, kStr, todo.UUID).Result()
	if err != nil {
		return handleRedisNil(err)
	}
	return handleRedisNil(err)
}

func (r *redisDataGateway) CreateUser(
	ctx context.Context,
	user *model.User,
) error {
	k, v := redishelper.FromUser(user)
	vStr := v.Marshal()
	// User info
	if _, err := r.client.HSet(ctx, usersKey, user.Username, vStr).Result(); err != nil {
		return handleRedisNil(err)
	}
	// Map username - UUID
	_, err := r.client.Set(ctx, k.String(), user.Username, 0).Result()
	if err != nil {
		return handleRedisNil(err)
	}
	return handleRedisNil(err)
}

func (r *redisDataGateway) CreateTodo(
	ctx context.Context,
	todo *model.Todo,
) error {
	k := redishelper.KeyFromTodo(todo)
	kStr := k.String()
	vStr := todo.Marshal()
	_, err := r.client.HSet(ctx, kStr, todo.UUID, vStr).Result()
	return handleRedisNil(err)
}

func (r *redisDataGateway) GetUserByUsername(
	ctx context.Context,
	username string,
	dst *model.User,
) error {
	v, err := r.client.HGet(ctx, usersKey, username).Result()
	if err != nil {
		return handleRedisNil(errors.Wrapf(err, "failed to get hash key: %s", username))
	}
	var user redishelper.UserVal
	if err := json.Unmarshal([]byte(v), &user); err != nil {
		return handleRedisNil(errors.Wrapf(err, "failed to unmarshal %s", v))
	}
	redishelper.CopyUser(user, dst)
	return handleRedisNil(err)
}

func (r *redisDataGateway) GetUserByUuid(
	ctx context.Context,
	uuid string,
	dst *model.User,
) error {
	if len(uuid) == 0 {
		return errors.New("empty uuid")
	}
	k := redishelper.UserKey{
		UUID: uuid,
	}
	key := k.String()
	username, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return handleRedisNil(errors.Wrapf(err, "failed to get username for uuid %s", uuid))
	}
	v, err := r.client.HGet(ctx, usersKey, username).Result()
	if err != nil {
		return handleRedisNil(errors.Wrapf(err, "failed to get hash key %s", username))
	}
	var user redishelper.UserVal
	err = json.Unmarshal([]byte(v), &user)
	if err != nil {
		return handleRedisNil(errors.Wrapf(err, "failed to unmarshal %s", v))
	}
	redishelper.CopyUser(user, dst)
	return handleRedisNil(err)
}

func (r *redisDataGateway) GetOneTodo(
	ctx context.Context,
	where *model.Todo,
	dst *model.Todo,
) error {
	k := redishelper.KeyFromTodo(where)
	kStr := k.String()
	v, err := r.client.HGet(ctx, kStr, where.UUID).Result()
	if err != nil {
		return handleRedisNil(errors.Wrapf(err, "failed to get key %s", kStr))
	}
	if err := json.Unmarshal([]byte(v), &dst); err != nil {
		return handleRedisNil(errors.Wrapf(err, "failed to unmarshal %s", v))
	}
	return nil
}

func (r *redisDataGateway) GetUserTodos(
	ctx context.Context,
	where *model.Todo,
	dst interface{},
) error {
	k := redishelper.KeyFromTodo(where)
	kStr := k.String()
	hgetAllMap, err := r.client.HGetAll(ctx, kStr).Result()
	if err != nil {
		return errors.Wrapf(err, "failed to get key %s", kStr)
	}
	var todos []model.Todo
	for _, todoVal := range hgetAllMap {
		var todo model.Todo
		if err := json.Unmarshal([]byte(todoVal), &todo); err != nil {
			return errors.Wrapf(err, "failed to unmarshal %s", todoVal)
		}
		todos = append(todos, todo)
	}
	switch dst := dst.(type) {
	case *[]model.Todo:
		*dst = todos
	default:
		return errors.New("interface{} conversion to *[]model.Todo failed: dst is of wrong type")
	}
	return handleRedisNil(err)
}

func (r *redisDataGateway) UpdateTodo(
	ctx context.Context,
	where *model.Todo,
	todo *model.Todo,
) error {
	return r.CreateTodo(ctx, todo)
}

func (r *redisDataGateway) ChangePassword(
	ctx context.Context,
	where *model.User,
	hashedPassword []byte,
) error {
	where.Password = hashedPassword
	return r.CreateUser(ctx, where)
}

func (r *redisDataGateway) Shutdown() {
	switch client := r.client.(type) {
	case *redis.Client:
		client.Close()
	}
}
