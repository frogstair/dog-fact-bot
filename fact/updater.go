package fact

import (
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

const DELAY = 30

func Update() {
	fh, err := os.Open("doglist.txt")
	if err != nil {
		panic("Can't open doglist.txt")
	}
	defer fh.Close()
	for {
		text, err := ioutil.ReadAll(fh)
		fh.Seek(0, io.SeekStart)
		if err != nil {
			panic("Can't open doglist.txt")
		}

		lines := strings.Split(string(text), "\n")
		lines = removeDuplicateValues(lines)

		plen := len(List)
		List = lines
		shuffle(List)
		nlen := len(List)

		log.Printf("Read in %d total dog facts, %d new\n", nlen, nlen-plen)
		log.Printf("Waiting %d seconds...\n", DELAY)
		time.Sleep(DELAY * time.Second)
	}
}

func removeDuplicateValues(arr []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range arr {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func shuffle(arr []string) []string {
	rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr
}
