package payment

import (
	"campaign/user"
	"os"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type Service interface {
	GetPaymentUrl(trx Transaction, user user.User) (string, error)
}

type service struct {
}

// GetToken implements Service.
func (s *service) GetPaymentUrl(trx Transaction, user user.User) (string, error) {
	snapClient := snap.Client{}
	snapClient.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(trx.ID),
			GrossAmt: int64(trx.Amount),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Name,
			Email: user.Email,
		},
	}

	snapResp, err := snapClient.CreateTransaction(req)
	if err != nil {
		return "", err
	}

	return snapResp.RedirectURL, nil

}

func NewService() Service {
	return &service{}
}
