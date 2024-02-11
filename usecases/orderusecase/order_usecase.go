package orderusecase

import (
	"context"
	"fmt"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/repositories/masterrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/orderrepositoryinterface"
	"healthcare-capt-america/interfaces/usecases/orderusecaseinterface"
	"healthcare-capt-america/services"
)

type orderUsecase struct {
	orderrepo   orderrepositoryinterface.OrderRepository
	paymentrepo masterrepositoryinterface.PaymentRepository
}

// UpdateOrderStatus implements orderusecaseinterface.OrderUsecase.
func (usecase *orderUsecase) UpdateOrderStatus(ctx context.Context, id uint64) (*models.Order, error) {
	order, err := usecase.orderrepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	switch order.OrderStatus {
	case enums.Processed:
		order.OrderStatus = enums.Sent
	case enums.Sent:
		order.OrderStatus = enums.OrderConfirmed
	}
	err = usecase.orderrepo.Update(ctx, order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (usecase *orderUsecase) GetOrderById(ctx context.Context, id uint64) (*models.Order, error) {
	order, err := usecase.orderrepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, apperror.NewClientError(fmt.Errorf("can't find order with id %d", id))
	}
	uid := ctx.Value(enums.UserIdKey).(uint64)
	isPharmacyAdmin, err := services.Authority.CheckUserRole(uid, enums.PharmacyAdmin.Slug)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	if isPharmacyAdmin && order.Pharmacy.AdminPharmacies.UserId != uid {
		return nil, apperror.NewClientError(fmt.Errorf("this order isn't in your pharmacy"))
	}
	return order, nil
}

func (usecase *orderUsecase) DeleteOrderById(ctx context.Context, id uint64) error {
	order, err := usecase.orderrepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if order == nil {
		return apperror.NewClientError(fmt.Errorf("can't find order with id %d", id))
	}
	uid := ctx.Value(enums.UserIdKey).(uint64)
	isPharmacyAdmin, err := services.Authority.CheckUserRole(uid, enums.PharmacyAdmin.Slug)
	if err != nil {
		return apperror.NewServerError(err)
	}
	if isPharmacyAdmin && order.Pharmacy.AdminPharmacies.UserId != uid {
		return apperror.NewClientError(fmt.Errorf("this order isn't in your pharmacy"))
	}
	return usecase.orderrepo.Delete(ctx, id)
}

func (usecase *orderUsecase) GetAllOrders(ctx context.Context, qry *requests.GlobalQuery) ([]*models.Order, *responses.Pagination, error) {
	uid := ctx.Value(enums.UserIdKey).(uint64)
	isPharmacyAdmin, err := services.Authority.CheckUserRole(uid, enums.PharmacyAdmin.Slug)
	if err != nil {
		return nil, nil, apperror.NewServerError(err)
	}
	if isPharmacyAdmin {
		qry.AddCondition(`"Pharmacy__AdminPharmacies".user_id`, requests.Equal, uid)
	}
	return usecase.orderrepo.FindAll(ctx, qry)
}

func (usecase *orderUsecase) CreateOrder(ctx context.Context, order *models.Order) (*models.Order, error) {
	order, err := usecase.orderrepo.CreateOrder(ctx, order)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	order, err = usecase.orderrepo.FindByID(ctx, order.ID)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (usecase *orderUsecase) GetUserOrder(ctx context.Context, status string, qry *requests.GlobalQuery) ([]*models.Order, *responses.Pagination, error) {
	uid := ctx.Value(enums.UserIdKey).(uint64)
	if len(status) == 0 {
		status = "%%"
	}
	qry.AddCondition(`"orders"."user_id"`, requests.Equal, uid)
	qry.AddCondition("order_status", requests.Ilike, status)

	o, p, err := usecase.orderrepo.FindAll(ctx, qry)
	if err != nil {
		return nil, nil, err
	}
	return o, p, nil
}

func (usecase *orderUsecase) GetUserDetailOrder(ctx context.Context, id uint64) (*models.Order, error) {
	uid := ctx.Value(enums.UserIdKey).(uint64)
	order, err := usecase.orderrepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, apperror.NewClientError(fmt.Errorf("can't find order with id %d", id))
	}
	if order.UserId != uid {
		return nil, apperror.NewClientError(fmt.Errorf("this isn't your order"))
	}
	return order, nil
}

func (usecase *orderUsecase) UpdateUserOrder(ctx context.Context, order_id uint64, status string) error {
	order, err := usecase.orderrepo.FindByID(ctx, order_id)
	if err != nil {
		return err
	}
	if order == nil {
		return apperror.NewClientError(fmt.Errorf("can't find order with id %d", order_id))
	}
	uid := ctx.Value(enums.UserIdKey).(uint64)
	if order.UserId != uid {
		return apperror.NewClientError(fmt.Errorf("this isnt your order"))
	}
	if status == enums.OrderCanceled {
		if order.OrderStatus != enums.WaitingForPayment {
			return apperror.NewClientError(fmt.Errorf("you can't cancel this order"))
		}
	} else if status == enums.OrderConfirmed {
		if order.OrderStatus != enums.Sent {
			return apperror.NewClientError(fmt.Errorf("you can't confirm this order"))
		}
	} else {
		return apperror.NewClientError(fmt.Errorf("invalid status"))
	}
	order.OrderStatus = status
	err = usecase.orderrepo.Update(ctx, order)
	if err != nil {
		return err
	}
	if order.OrderStatus == enums.Processed {
		payment, err := usecase.paymentrepo.FindByID(ctx, order.PaymentId)
		if err != nil {
			return apperror.NewServerError(err)
		}
		if payment == nil {
			return apperror.NewClientError(fmt.Errorf(`payment not found`))
		}
		paymentStatus := enums.PaymentApproved
		payment.Status = paymentStatus
		_, err = usecase.paymentrepo.Update(ctx, payment)
		if err != nil {
			return apperror.NewServerError(err)
		}
	}
	return nil
}

func (usecase *orderUsecase) AdminUpdateOrder(ctx context.Context, order_id uint64, status string) error {
	order, err := usecase.orderrepo.FindByID(ctx, order_id)
	if err != nil {
		return err
	}
	if order == nil {
		return apperror.NewClientError(fmt.Errorf("can't find order with id %d", order_id))
	}
	uid := ctx.Value(enums.UserIdKey).(uint64)
	if order.Pharmacy.AdminPharmacies.UserId != uid {
		return apperror.NewClientError(fmt.Errorf("this order isn't in your pharmacy"))
	}
	order.OrderStatus = status
	err = usecase.orderrepo.Update(ctx, order)
	if err != nil {
		return err
	}
	status = order.OrderStatus
	paymentStatus := ""
	switch status {
	case enums.WaitingForPayment:
		paymentStatus = enums.PaymentRejected
	case enums.Processed:
		paymentStatus = enums.PaymentApproved
	}
	if len(paymentStatus) != 0 {
		payment, err := usecase.paymentrepo.FindByID(ctx, order.PaymentId)
		if err != nil {
			return apperror.NewServerError(err)
		}
		if payment == nil {
			return apperror.NewClientError(fmt.Errorf(`payment not found`))
		}
		payment.Status = paymentStatus
		_, err = usecase.paymentrepo.Update(ctx, payment)
		if err != nil {
			return apperror.NewServerError(err)
		}
	}
	return nil
}

func NewOrderUsecase(orderrepo orderrepositoryinterface.OrderRepository, paymentrepo masterrepositoryinterface.PaymentRepository) *orderUsecase {
	return &orderUsecase{
		orderrepo:   orderrepo,
		paymentrepo: paymentrepo,
	}
}

var _ orderusecaseinterface.OrderUsecase = &orderUsecase{}
