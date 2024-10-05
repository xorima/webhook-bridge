package containers

import (
	"context"
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"slices"
	"strings"
	"time"
)

type Container struct {
	container testcontainers.Container
	ports     []string
}

func NewContainer(ctx context.Context, image, readyLog string, ports []string) (*Container, error) {
	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        image,
			ExposedPorts: ports,
			WaitingFor:   wait.ForLog(readyLog),
		},
	}
	cont, err := testcontainers.GenericContainer(ctx, req)
	if err != nil {
		return nil, err
	}
	return &Container{container: cont, ports: ports}, nil
}

func (c *Container) Start(ctx context.Context) error {
	return c.container.Start(ctx)

}
func (c *Container) Stop(ctx context.Context) error {
	t := time.Second
	return c.container.Stop(ctx, &t)

}
func (c *Container) GetPort(ctx context.Context, port string) (int, error) {
	if !slices.Contains(c.ports, port) {
		return 0, fmt.Errorf("port %s not found in defined ports at runtime, %v+", port, c.ports)
	}
	p := strings.Split(port, "/")
	// the error here is impossible as new container checks ports already
	prt, _ := nat.NewPort(p[1], p[0])
	ports, _ := c.container.MappedPort(ctx, prt)
	return ports.Int(), nil
}
