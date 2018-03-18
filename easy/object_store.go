package easy

import(
	"fmt"
	"time"
	"hash/fnv"
)

//id -> hash
var primary = make(map[string]uint64,1000)
//hash->text
var secondary = make(map[uint64][]byte,1000)

//gaurentee uniqueness 
func generateId() uint64{
	return uint64(time.Now().Unix())
}

func hash(text []byte)uint64{
	h:=fnv.New64a()
	h.Write(text)
	return h.Sum64()
}

func put(text []byte) uint64{
	id := generateId()
	key := hash(text) 
	primary[id]=key
	secondary[key]=text
}

func get()