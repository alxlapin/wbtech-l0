package domain

import (
	"encoding/json"
	"time"
)

type Order struct {
	ID              string   `json:"order_id"`
	TrackNumber     string   `json:"track_number"`
	Entry           string   `json:"entry"`
	Locale          string   `json:"locale"`
	CustomerID      string   `json:"customer_id"`
	DeliveryService string   `json:"delivery_service"`
	DateCreated     string   `json:"date_created"`
	Delivery        Delivery `json:"delivery"`
	Payment         Payment  `json:"payment"`
	Items           []Item   `json:"items"`
}

type Delivery struct {
	ID      string `json:"-"`
	OrderID string `json:"-"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type Payment struct {
	ID           string   `json:"-"`
	OrderID      string   `json:"-"`
	Transaction  string   `json:"transaction"`
	RequestID    string   `json:"request_id"`
	Currency     string   `json:"currency"`
	Provider     string   `json:"provider"`
	Amount       int      `json:"amount"`
	PaymentDt    UnixTime `json:"payment_dt"`
	Bank         string   `json:"bank"`
	DeliveryCost int      `json:"delivery_cost"`
	GoodsTotal   int      `json:"goods_total"`
	CustomFee    int      `json:"custom_fee"`
}

type Item struct {
	ID         string `json:"-"`
	OrderID    string `json:"-"`
	ChrtID     int    `json:"chrt_id"`
	Price      int    `json:"price"`
	Rid        string `json:"rid"`
	Name       string `json:"name"`
	Sale       int    `json:"sale"`
	Size       string `json:"size"`
	TotalPrice int    `json:"total_price"`
	NmID       int    `json:"nm_id"`
	Brand      string `json:"brand"`
	Status     int    `json:"status"`
}

type UnixTime time.Time

func (ut *UnixTime) UnmarshalJSON(b []byte) error {
	var timestamp int64
	if err := json.Unmarshal(b, &timestamp); err != nil {
		return err
	}
	*ut = UnixTime(time.Unix(timestamp, 0))
	return nil
}

// TODO: разобраться, как работать с кастомным типом [UnixTime] при записи/чтении данных в/из БД, т.к нативно магии не случится.
// TODO: Возможно, есть более простые решения [работы с датой].
