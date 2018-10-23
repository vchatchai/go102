package sttime

import (
	"encoding/json"
	"fmt"
	"time"
)

func Serialize() {
	eur, err := time.LoadLocation("Europe/Vienna")
	if err != nil {
		panic(err)
	}

	t := time.Date(2018, 10, 23, 11, 16, 11, 0, eur)

	b, err := t.MarshalJSON()
	if err != nil {
		panic(err)
	}

	fmt.Println("Serialized as RFC 3339:", string(b))

	t2 := time.Time{}
	t2.UnmarshalJSON(b)
	fmt.Println("Deserialize from RFC 3339:", t2)

	epoch := t.Unix()
	fmt.Println("Serialized as Epoch:", epoch)

	jsonStr := fmt.Sprintf(`{"created":%d}`, epoch)

	data := struct {
		Created int64 `json:"created"`
	}{}
	json.Unmarshal([]byte(jsonStr), &data)

	deserialize := time.Unix(data.Created, 0)
	// deserialize.In(eur)
	fmt.Println("Deserialize from Epoch:", deserialize)
}
