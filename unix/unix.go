package unix

import (
	"fmt"
	"time"
)

func Unix() {
	fmt.Println(time.Now().Unix())
}