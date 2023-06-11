package handlers

import (
	dto "dumbmerch/dto/result"
	transactiondto "dumbmerch/dto/transaction"
	"dumbmerch/models"
	"dumbmerch/repository"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/yudapc/go-rupiah"
	"gopkg.in/gomail.v2"
)

type TransactionsHandler struct {
	Transaction repository.TransactionRepository
}

func HandlerTransactions(Transaction repository.TransactionRepository) *TransactionsHandler {
	return &TransactionsHandler{Transaction}
}

func (h *TransactionsHandler) GetAllTransactions(c echo.Context) error {
	dataResponse, err := h.Transaction.GetAllTransactions()
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	for i, p := range dataResponse {
		dataResponse[i].Attachment = path_file + p.Attachment
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: dataResponse,
	})
}

func (h *TransactionsHandler) GetTransactionById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	Trans, err := h.Transaction.GetTransactionById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	Trans.Attachment = path_file + Trans.Attachment

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: Trans,
	})
}

func (h *TransactionsHandler) CreateTransaction(c echo.Context) error {
	request := new(transactiondto.CreateTransactions)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	getUser, err := h.Transaction.GetUserByIdForTrans(int(userId))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	getTrip, err := h.Transaction.GetTripByIdForTrans(request.TripId)
	fmt.Println(getTrip, "ini trip id")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	// generate id random
	var transactionIsMatch = false
	var transactionId int

	for !transactionIsMatch {
		transactionId = int(time.Now().Unix())
		transactionData, _ := h.Transaction.GetTransactionById(transactionId)
		if transactionData.Id == 0 {
			transactionIsMatch = true
		}
	}

	transaction := models.Transaction{
		Id: transactionId,
		CounterQty: request.CounterQty,
		Total: request.Total,
		Status: request.Status,
		UserId: int(userId),
		User: models.UserResponse(getUser),
		TripId: request.TripId,
		Trip: models.TripResponse(getTrip),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	
	dataResponse, err := h.Transaction.CreateTransaction(transaction)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	// request token from midtrans
	var s = snap.Client{}
	s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: strconv.Itoa(dataResponse.Id),
			GrossAmt: int64(dataResponse.Total),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: dataResponse.User.FullName,
			Email: dataResponse.User.Email,
		},
	}

	snapResp, _ := s.CreateTransaction(req)

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: snapResp,
	})
}

func (h *TransactionsHandler) DeleteTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	
	Trans, err := h.Transaction.GetTransactionById(id)
	fmt.Println(Trans, "ini adalah trans")

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	dataResponse, err := h.Transaction.DeleteTransaction(Trans)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: dataResponse,
	})
}


// func (h *TransactionsHandler) EditTransaction(c echo.Context) error {
// 	// dataFile := c.Get("dataFile").(string)

// 	request := new(transactiondto.UpdateTransactions)
// 	if err := c.Bind(request); err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
// 			Code: http.StatusBadRequest,
// 			Message: err.Error(),
// 		})
// 	}

// 	id, _ := strconv.Atoi(c.Param("id"))
// 	transaction, err := h.Transaction.GetTransactionById(id)

// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
// 			Code: http.StatusBadRequest,
// 			Message: err.Error(),
// 		})
// 	}

// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
// 			Code: http.StatusBadRequest,
// 			Message: err.Error(),
// 		})
// 	}

// 	status := c.FormValue("status")

// 	if request.CounterQty != 0 {
// 		transaction.CounterQty = request.CounterQty
// 	}
// 	if request.Total != 0 {
// 		transaction.Total = request.Total
// 	}
// 	if status != "" {
// 		transaction.Status = status
// 	}

// 	if request.Status != "" {
// 		transaction.Status = request.Status
// 	}



// 	transaction.UpdatedAt = time.Now()

// 	dataResponse, err := h.Transaction.EditTransaction(transaction)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
// 			Code: http.StatusInternalServerError,
// 			Message: err.Error(),
// 		})
// 	}

// 	return c.JSON(http.StatusOK, dto.SuccessResult{
// 		Code: http.StatusOK,
// 		Data: dataResponse,
// 	})
// }

func sendEmail(status string, transaction models.Transaction) {
	if status != transaction.Status && status == "success" {
		var CONFIG_SMTP_HOST = "smtp.gmail.com"
		var CONFIG_SMTP_PORT = 587
		var CONFIG_SENDER_NAME = "DeweTour <ibnu.ibnualinsani23@gmail.com>"
		var CONFIG_AUTH_EMAIL = os.Getenv("EMAIL_SYSTEM")
		var CONFIG_AUTH_PASSWORD = os.Getenv("PASSWORD_SYSTEM")

		var tripName = transaction.Trip.Title
		var price = rupiah.FormatRupiah(float64(transaction.Trip.Price * transaction.CounterQty))
		fmt.Println(transaction.User.Email)

		mailer := gomail.NewMessage()
		mailer.SetHeader("From", CONFIG_SENDER_NAME)
		mailer.SetHeader("To", transaction.User.Email)
		mailer.SetHeader("Subject", "Transaction Status")
		mailer.SetBody("text/html", fmt.Sprintf(`
			<!DOCTYPE html>
				<html lang="en">
				<head>
				<meta charset="UTF-8" />
				<meta http-equiv="X-UA-Compatible" content="IE=edge" />
				<meta name="viewport" content="width=device-width, initial-scale=1.0" />
				<title>Document</title>
				<style>
					h1 {
					color: brown;
					}
				</style>
				</head>
				<body>
				<h2>Product payment :</h2>
				<ul style="list-style-type:none;">
					<li>Trip : %s</li>
					<li>Total payment: %s</li>
					<li>Status : <b>%s</b></li>
					<li>Thank you for making the order, please wait for the trip schedule, Enjoy Your Trip</li>
				</ul>
				</body>
			</html>
		`, tripName, price, status))

		dialer := gomail.NewDialer(
			CONFIG_SMTP_HOST,
			CONFIG_SMTP_PORT,
			CONFIG_AUTH_EMAIL,
			CONFIG_AUTH_PASSWORD,
		)

		err := dialer.DialAndSend(mailer)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Println("Mail sent! to " + transaction.User.Email)
	}
}

func (h *TransactionsHandler) Notification(c echo.Context) error {
	var notificationPayload map[string]interface{}

	if err := c.Bind(&notificationPayload); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	
	transactionStatus := notificationPayload["transaction_status"].(string)
	fraudStatus := notificationPayload["fraud_status"].(string)
	orderId := notificationPayload["order_id"].(string)

	order_id, _ := strconv.Atoi(orderId)

	transaction, err := h.Transaction.GetTransactionById(order_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadGateway,
			Message: err.Error(),
		})
	}

	fmt.Println("ini payload", notificationPayload)

	if transactionStatus == "capture" {
		if fraudStatus == "accept" {
			h.Transaction.EditTransaction("success", order_id, transaction.CounterQty)
			sendEmail("success", transaction)
			fmt.Println("SUKSES BANG")
		} else if fraudStatus == "deny" {
			h.Transaction.EditTransaction("failed", order_id, transaction.CounterQty)
			fmt.Println("GAGAL BANG")
		}
	} else if transactionStatus == "settlement" {
		h.Transaction.EditTransaction("success", order_id, transaction.CounterQty)
		sendEmail("success", transaction)
		fmt.Println("SUKSES SETTLE BANG")
	} else if transactionStatus == "deny" {
		h.Transaction.EditTransaction("failed", order_id, transaction.CounterQty)
		fmt.Println("FAILED DENY BANG")
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
		fmt.Println("CANCEL BANG")
		h.Transaction.EditTransaction("failed", order_id, transaction.CounterQty)
	} else if transactionStatus == "pending" {
		fmt.Println("PENDING BANG")
		h.Transaction.EditTransaction("pending", order_id, transaction.CounterQty)
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK, Data: notificationPayload,
	})
}