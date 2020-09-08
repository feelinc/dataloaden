//go:generate go run github.com/feelinc/dataloader UserSliceLoader int []github.com/feelinc/dataloader/example.User

package slice

import (
	"strconv"
	"time"

	"github.com/feelinc/dataloader/example"
)

func NewLoader() *UserSliceLoader {
	return &UserSliceLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []int, params ...[]interface{}) ([][]example.User, []error) {
			users := make([][]example.User, len(keys))
			errors := make([]error, len(keys))

			for i, key := range keys {
				users[i] = []example.User{{ID: strconv.Itoa(key), Name: "user " + strconv.Itoa(key)}}
			}
			return users, errors
		},
	}
}
