package containers

import "context"

type RedisContainer struct {
	*Container
}

func NewRedisContainer(ctx context.Context) (*RedisContainer, error) {
	cont, err := NewContainer(ctx, "redis:latest", "Ready to accept connections", []string{"6379/tcp"})
	return &RedisContainer{Container: cont}, err
}

func (r *RedisContainer) GetRedisPort(ctx context.Context) int {
	p, _ := r.GetPort(ctx, "6379/tcp")
	return p
}
