package funcs

import (
	"fmt"

	"dhui.com/configs"
	"github.com/go-redis/redis"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisLikes() *RedisClient {
	return &RedisClient{
		client: redis.NewClient(&redis.Options{
			Addr:     configs.REDIS_IP + ":" + configs.REDIS_PORT, // 指定
			Password: configs.REDIS_PASSWORD,
			DB:       configs.REDIS_DB,
		}),
	}
}

// 用户点赞
func (r *RedisClient) Like(userId string, objectId int) error {
	userLikesKey := fmt.Sprintf("user:%s:likes", userId)
	objectLikesKey := fmt.Sprintf("object:%d:likes", objectId)
	// 添加用户点赞记录
	if err := r.client.SAdd(userLikesKey, objectId).Err(); err != nil {
		return err
	}
	// 添加对象被点赞记录
	if err := r.client.SAdd(objectLikesKey, userId).Err(); err != nil {
		// 如果添加失败，需要对用户点赞记录的添加进行回滚
		r.client.SRem(userLikesKey, objectId)
		return err
	}
	return nil
}
func (r *RedisClient) Unlike(userId string, objectId int) error {
	userLikesKey := fmt.Sprintf("user:%s:likes", userId)
	objectLikesKey := fmt.Sprintf("object:%d:likes", objectId)
	// 删除对象被点赞记录
	if err := r.client.SRem(objectLikesKey, userId).Err(); err != nil {
		return err
	}
	// 删除用户点赞记录
	if err := r.client.SRem(userLikesKey, objectId).Err(); err != nil {
		// 如果删除失败，需要回滚对象被点赞记录的删除
		r.client.SAdd(objectLikesKey, userId)
		return err
	}
	return nil
}
func (r *RedisClient) GetLikedObjects(userId string) ([]string, error) {
	userLikesKey := fmt.Sprintf("user:%s:likes", userId)
	// 获取用户点赞记录中所有对象的ID
	return r.client.SMembers(userLikesKey).Result()
}
func (r *RedisClient) GetObjectLikes(objectId int) ([]string, error) {
	objectLikerKey := fmt.Sprintf("object:%d:likers", objectId)
	return r.client.SMembers(objectLikerKey).Result()
}

// 用户收藏
func (r *RedisClient) Collection(userId string, objectId int) error {
	userCollectionsKey := fmt.Sprintf("user:%s:collections", userId)
	objectCollectionsKey := fmt.Sprintf("object:%d:collections", objectId)
	// 添加用户点赞记录
	if err := r.client.SAdd(userCollectionsKey, objectId).Err(); err != nil {
		return err
	}
	// 添加对象被点赞记录
	if err := r.client.SAdd(objectCollectionsKey, userId).Err(); err != nil {
		// 如果添加失败，需要对用户点赞记录的添加进行回滚
		r.client.SRem(userCollectionsKey, objectId)
		return err
	}
	return nil
}
func (r *RedisClient) Uncollection(userId string, objectId int) error {
	userCollectionsKey := fmt.Sprintf("user:%s:collections", userId)
	objectCollectionsKey := fmt.Sprintf("object:%d:collections", objectId)
	// 删除对象被点赞记录
	if err := r.client.SRem(objectCollectionsKey, userId).Err(); err != nil {
		return err
	}
	// 删除用户点赞记录
	if err := r.client.SRem(userCollectionsKey, objectId).Err(); err != nil {
		// 如果删除失败，需要回滚对象被点赞记录的删除
		r.client.SAdd(objectCollectionsKey, userId)
		return err
	}
	return nil
}
func (r *RedisClient) GetCollectionObjects(userId string) ([]string, error) {
	userCollectionsKey := fmt.Sprintf("user:%s:collections", userId)
	// 获取用户点赞记录中所有对象的ID
	return r.client.SMembers(userCollectionsKey).Result()
}
func (r *RedisClient) GetObjectCollections(objectId int) ([]string, error) {
	objectCollectionKey := fmt.Sprintf("object:%d:likers", objectId)
	return r.client.SMembers(objectCollectionKey).Result()
}
