package usecase

import (
	"context"
	"time"

	"github.com/amitshekhariitbhu/go-backend-clean-architecture/domain"
)

type taskUsecase struct {
	taskRepository domain.PacketRepository
	contextTimeout time.Duration
}

func NewTaskUsecase(taskRepository domain.PacketRepository, timeout time.Duration) domain.PacketUsecase {
	return &taskUsecase{
		taskRepository: taskRepository,
		contextTimeout: timeout,
	}
}

func (tu *taskUsecase) Create(c context.Context, task *domain.Packet) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.Create(ctx, task)
}

func (tu *taskUsecase) FetchAll(c context.Context) ([]domain.Packet, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.FetchAll(ctx)
}
