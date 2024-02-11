package services

import (
	"context"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/interfaces/repositories/orderrepositoryinterface"

	"github.com/robfig/cron/v3"
)

type CronService struct {
	Scheduler *cron.Cron
	orderRepo orderrepositoryinterface.OrderRepository
}

func NewCronService(repo orderrepositoryinterface.OrderRepository) *CronService {
	s := cron.New()
	return &CronService{
		Scheduler: s,
		orderRepo: repo,
	}
}

func (cs *CronService) RunAllJob() error {
	err := cs.AutoConfirmOrder()
	if err != nil {
		return apperror.NewServerError(err)
	}

	err = cs.AutoCancelOrder()
	if err != nil {
		return apperror.NewServerError(err)
	}
	return nil
}

func (cs *CronService) AutoConfirmOrder() error {
	_, err := cs.Scheduler.AddFunc("@hourly", func() {
		cs.orderRepo.SetConfirmed(context.Background())
	})
	return err
}

func (cs *CronService) AutoCancelOrder() error {
	_, err := cs.Scheduler.AddFunc("@hourly", func() {
		cs.orderRepo.SetCanceled(context.Background())
	})
	return err
}
