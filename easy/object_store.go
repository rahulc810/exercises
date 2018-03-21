package easy

import (
	"fmt"
	"hash/fnv"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"bytes"

	"github.com/rahulc810/exercises/utils"
)

type data struct {
	text    []byte
	refs    uint32
	version int64
	next    *data
}

type event struct {
	eKey string
	op   uint8
}

//id -> k1
var primary = make(map[uint64]string, 1000)

//k1->*text
var secondary = make(map[string]*data, 1000)
var q = utils.NewObjQueue(10)

var hashMap = make(map[uint64]*data, 1000)

var pMutex = sync.RWMutex{}
var sMutex = sync.RWMutex{}
var iMutex = sync.Mutex{}
var count uint64 = 0

func Start() {
	go processQ()
	go func() {
		for {
			itrsec()
			fmt.Println()
			//time.Sleep(5 * time.Millisecond)
		}
	}()
}

//gaurentee uniqueness
func generateID() uint64 {
	iMutex.Lock()
	defer iMutex.Unlock()
	return uint64(atomic.AddUint64(&count, 1))
}

func getInternalID(pID uint64) string {
	return strconv.FormatUint(pID, 10) + "sec"
}

func hash(text []byte) uint64 {
	h := fnv.New64a()
	h.Write(text)
	return h.Sum64()
}

func Put(text []byte) uint64 {
	id := generateID()
	iid := getInternalID(id)

	pMutex.Lock()
	sMutex.Lock()
	defer pMutex.Unlock()
	defer sMutex.Unlock()

	primary[id] = iid
	secondary[iid] = &data{text, 1, time.Now().Unix(), nil}

	q.Enqueue(event{iid, 1})

	return id
}

func Get(id uint64) ([]byte, error) {
	pMutex.RLock()
	sMutex.RLock()
	defer pMutex.RUnlock()
	defer sMutex.RUnlock()

	iid, ok := primary[id]
	if !ok {
		return nil, fmt.Errorf("Key not found: %d", id)
	}
	ret, ok := secondary[iid]
	if !ok {
		return nil, fmt.Errorf("[Internal] Key not found: %s", iid)
	}
	return ret.text, nil
}

func Del(id uint64) error {
	pMutex.Lock()
	sMutex.Lock()
	defer pMutex.Unlock()
	defer sMutex.Unlock()

	iid, ok := primary[id]
	if !ok {
		return fmt.Errorf("Key not found: %d", id)
	}
	delete(primary, id)
	q.Enqueue(event{iid, 2})
	return nil
}

func processQ() {
	for {
		eve, err := q.Dequeue()
		if err == nil && eve != nil {
			e := eve.(event)
			switch e.op {
			case 1: //put
				possibleDup := secondary[e.eKey]
				h := hash(possibleDup.text)
				duplicate, commonDataValue := putInMap(h, possibleDup)
				if duplicate {
					secondary[e.eKey] = commonDataValue
				}
			case 2: //del
				dataValue := secondary[e.eKey]
				h := hash(dataValue.text)
				removeFromMap(e, h, dataValue.text)
			}
			fmt.Printf("HASHMAP => %v\n", hashMap)
		}
	}
}

func getFromMap(key uint64) (*data, bool) {
	ret, ok := hashMap[key]
	return ret, ok
}

func putInMap(key uint64, d *data) (bool, *data) {
	ret, ok := hashMap[key]
	if !ok {
		hashMap[key] = d
		return false, nil
	} else {
		//traverse the linked list and euqate with each element
		//if found increment the ref count else
		//append to the linked list
		node := ret
		found := false
		for ok := true; ok; ok = (node != nil) {
			if bytes.Equal(node.text, d.text) {
				node.refs++
				found = true
				break
			} else if node.next == nil {
				break
			}
			node = node.next
		}

		if !found {
			node.next = d
		}

		return true, ret
	}
}

func removeFromMap(e event, key uint64, text []byte) {
	ret, ok := hashMap[key]
	if !ok {
	} else {
		fmt.Printf("Removing from secondary => %v <%v>\n", ret, secondary[e.eKey])
		sMutex.Lock()
		delete(secondary, e.eKey)
		sMutex.Unlock()

		//traverse the linked list and euqate with each element
		//if found decrement the ref count if ref count > 1
		//else remove the node
		var prev *data
		for node := ret; node != nil; prev, node = node, node.next {
			if bytes.Equal(node.text, text) {
				if node.refs <= 1 {
					//delete node
					if prev == nil {
						//first node
						delete(hashMap, key)
						//there are other nodes in teh bucket
						if node.next != nil {
							hashMap[key] = node.next
						}
					} else {
						prev.next = node.next
					}
				} else {
					node.refs--
				}
			}
		}
	}
}

func itrsec() {
	sMutex.RLock()
	defer sMutex.RUnlock()

	var out string
	for k, v := range secondary {
		out += fmt.Sprintf("%v : %v |", k, v)
	}

	fmt.Printf("Secondary => %v\n", out)
}
