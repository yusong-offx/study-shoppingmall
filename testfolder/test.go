package testfolder

import (
	"bufio"
	"log"
	"os"
)

func P() []string {
	file, err := os.Open("./MOCKDATA.sql")
	r := bufio.NewReader(file)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	ret := []string{}
	for {
		a, b := r.ReadString('\n')
		if b != nil {
			break
		}
		ret = append(ret, a[:len(a)-1])
	}
	return ret
}
