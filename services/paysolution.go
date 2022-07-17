package services

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"piennews/helper/config"
	"piennews/helper/database"
	"piennews/helper/logs"
	"piennews/helper/util"
	"piennews/models"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func (sv *service) NewInitPaysolution(payment_code string) (int64, error) {

	db, err := sql.Open("mysql", database.Connect().ConnectionString())
	if err != nil {
		return 0, err
	}
	defer db.Close()
	ref_no := util.GetUniqNumber()
	statement, err := db.Prepare(`INSERT INTO paysolution (payment_code, paysolution_ref_no) VALUES(?, ?);`)

	if err != nil {
		return 0, err
	}
	_, err = statement.Exec(payment_code, ref_no)
	if err != nil {
		return 0, err
	}

	return 0, nil

}

func (sv *service) InquiryPaysolution(ref_no string) (*models.InquiryModel, error) {
	lg := &logs.LogExternalParams{}

	defer func(begin time.Time) {
		lg.Begin = begin
		logs.ExternalLogs(lg).WriteExternalLogs()
	}(time.Now())
	// http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	merchantId := config.GetPaysolutionEnv().MERCHANTID
	merchant_secret_key := config.GetPaysolutionEnv().MERCHANT_SECRET_KEY
	merchant_api_key := config.GetPaysolutionEnv().MERCHANT_API_KEY

	values := map[string]string{"merchantId": merchantId, "refNo": ref_no}

	body, _ := json.Marshal(values)

	url := "https://apis.paysolutions.asia/order/orderdetailpost"
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("merchantId", merchantId)
	req.Header.Set("merchantSecretKey", merchant_secret_key)
	req.Header.Set("apikey", merchant_api_key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		lg.Error = err.Error()
		return nil, err
	}
	defer resp.Body.Close()
	resbody, _ := ioutil.ReadAll(resp.Body)
	resstr := string(resbody)

	lg.Url = url
	lg.Request = bytes.NewBuffer(body)

	transactions := models.PaysolutionInquiry{}
	err = json.Unmarshal([]byte(resstr), &transactions)
	if err != nil {
		lg.Error = err.Error()
		return nil, err
	}

	if len(transactions) == 0 {
		lg.Error = "Data Not Found"
		return nil, errors.New("Data Not Found")
	}
	lg.Response = transactions[0]
	return &transactions[0], nil

}

func (sv *service) GetOrderPrice(refno string) (float64, error) {
	lg := &logs.LogExternalParams{}

	defer func(begin time.Time) {
		lg.Begin = begin
		logs.ExternalLogs(lg).WriteExternalLogs()
	}(time.Now())

	db, err := sql.Open("mysql", database.Connect().ConnectionString())

	if err != nil {
		lg.Error = err.Error()
		return 0, err
	}

	defer db.Close()
	lg.Url = "db"
	lg.Request = refno

	var amount float64
	err = db.QueryRow(`SELECT total_baht  
	FROM orders o inner join paysolution p on o.payment_code = p.payment_code 
	where p.paysolution_ref_no = ?`, refno).Scan(
		&amount,
	)

	if err != nil {
		lg.Error = err.Error()
		return 0, err
	}
	lg.Response = amount
	return amount, nil
}

func (s *service) EnquipryNextStep(ref_no string, status string) error {
	db, err := sql.Open("mysql", database.Connect().ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()
	statement, err := db.Prepare(`UPDATE orders o 
	inner join paysolution p on  o.payment_code  = p.payment_code 
	set o.status=?,
	p.retry =p.retry + 1,
	p.last_update =now() 
	where p.paysolution_ref_no =?;`)

	if err != nil {
		return err
	}
	_, err = statement.Exec(status, status, ref_no)
	if err != nil {
		return err
	}

	return nil
}
