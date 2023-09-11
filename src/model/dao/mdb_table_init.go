package dao

import (
	"sync"

	"github.com/assimon/luuu/model/mdb"
	"github.com/gookit/color"
)

var once sync.Once

func MdbTableInit() {
	once.Do(func() {
		if err := Mdb.AutoMigrate(&mdb.Orders{}); err != nil {
			color.Red.Printf("[store_db] AutoMigrate DB(Orders),err=%s\n", err)
			// panic(err)
			return
		}
		if err := Mdb.AutoMigrate(&mdb.WalletAddress{}); err != nil {
			color.Red.Printf("[store_db] AutoMigrate DB(Orders),err=%s\n", err)
			// panic(err)
			return
		}
	})
}
