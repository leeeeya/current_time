package current_time

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
)

func CurrentTime() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(time)
}
