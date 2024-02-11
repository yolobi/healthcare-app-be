package paymentusecase

import (
	"context"
	"errors"
	"fmt"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/repositories/masterrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/orderrepositoryinterface"
	"healthcare-capt-america/interfaces/usecases/paymentusecaseinterface"
	"healthcare-capt-america/services"
	"healthcare-capt-america/utils"
	"mime/multipart"
	"strings"
)

type paymentUsecase struct {
	paymentrepo masterrepositoryinterface.PaymentRepository
	orderrepo   orderrepositoryinterface.OrderRepository
}

func (usecase *paymentUsecase) UpdatePaymentStatus(ctx context.Context, id uint64, status string) (*models.Payment, error) {
	var payment *models.Payment
	atomic := func(ctx context.Context, id uint64, status string) error {
		payment, err := usecase.paymentrepo.FindByID(ctx, id)
		if err != nil {
			return err
		}

		if payment.Status == status {
			s := fmt.Sprintf(`this payment status already %s`, status)
			return apperror.NewClientError(errors.New(s))
		}

		if err != nil {
			return apperror.NewServerError(err)
		}

		payment.Status = status
		order, err := usecase.orderrepo.FindByPaymentID(ctx, payment.ID)
		if err != nil {
			return err
		}
		if payment.Status == enums.PaymentApproved {
			order.OrderStatus = enums.Processed
		}

		err = usecase.orderrepo.Update(ctx, order)
		if err != nil {
			return err
		}
		return nil
	}
	err := atomic(ctx, id, status)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	return usecase.paymentrepo.Update(ctx, payment)
}

func (usecase *paymentUsecase) UploadPaymentFile(ctx context.Context, payment *models.Payment) (*models.Payment, error) {
	payment, err := usecase.paymentrepo.FindByID(ctx, payment.ID)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	userId := ctx.Value(enums.UserIdKey.Key).(uint64)
	paymentId := payment.ID
	isUsers, err := usecase.paymentrepo.ValidateUserPayment(ctx, paymentId, userId)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	if !isUsers {
		return nil, apperror.NewClientError(errors.New(`you cant update other payment`))
	}
	a := ctx.Value(enums.PaymentFileKey.Key).(*multipart.FileHeader)
	if a == nil {
		return nil, apperror.NewServerError(fmt.Errorf(`Error not found`))
	}
	extName, err := utils.GetFileType(a)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	fileName := strings.ReplaceAll(fmt.Sprintf("%d", paymentId), " ", "_") + extName
	path, err := services.UploadPaymentFile(ctx, a, fileName)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	payment.File = *path
	payment.Status = enums.WaitingApproval
	payment, err = usecase.paymentrepo.Update(ctx, payment)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	order, err := usecase.orderrepo.FindByPaymentID(ctx, paymentId)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	if order == nil {
		return nil, apperror.NewClientError(errors.New(`order not found`))
	}
	order.OrderStatus = enums.WaitingForConfirmation
	_ = usecase.orderrepo.Update(ctx, order)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	return payment, nil
}

func NewPaymentUsecase(paymentrepo masterrepositoryinterface.PaymentRepository, orderrepo orderrepositoryinterface.OrderRepository) *paymentUsecase {
	return &paymentUsecase{paymentrepo: paymentrepo, orderrepo: orderrepo}
}

var _ paymentusecaseinterface.PaymentUsecase = &paymentUsecase{}
