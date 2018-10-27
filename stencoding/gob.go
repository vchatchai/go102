package stencoding

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type User struct {
	FirstName string
	LastName  string
	Age       int
	Active    bool
}

func (u User) String() string {
	return fmt.Sprintf(`{"FirstName":%s,"LastName":%s,"Age":%d, "Active":%v`, u.FirstName, u.LastName, u.Age, u.Active)
}

type SimmpleUser struct {
	FirstName string
	LastName  string
}

func (u SimmpleUser) String() string {
	return fmt.Sprintf(`{"FirstName":%s, "LastName":%s}`, u.FirstName, u.LastName)
}

func Gob() {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)

	user := User{
		FirstName: "Radomir",
		LastName:  "Sohlich",
		Age:       30,
		Active:    true,
	}

	enc.Encode(user)
	fmt.Printf("%X\n", buff.Bytes())

	out := User{}
	dec := gob.NewDecoder(&buff)

	dec.Decode(&out)
	fmt.Println(out.String())

	enc.Encode(user)
	out2 := SimmpleUser{}
	dec.Decode(&out2)
	fmt.Println(out2.String())

}
