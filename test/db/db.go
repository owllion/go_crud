package mock_main

import "fmt"

//模擬DB行為
type DB interface {
	Get(key string) (int, error)
}

func GetFromDB(db DB, key string) int {
	if value, err := db.Get(key); err == nil {
		fmt.Println("print value----", value)
		return value
	}
	//有err就會回傳-1，所以要是外面mock expect有return err,那不管怎樣都只會回傳-1
	return -1
}
