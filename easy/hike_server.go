package easy

import(
	"net"
	"fmt"
	"strings"
	"time"
	"sync"
	"sync/atomic"
	"strconv"
)

const(
	ServerAddr = ":9191"
	IPScheme = "tcp4"
	Timeout = 4
	RotatingFactor = 10000
)
var store = make(map[int64]*pNode, 5000)
var connLookup = make(map[uint64]*pNode, 5000)
var mutexOld = &sync.Mutex{}
var mutex = &sync.Mutex{}
var ConnId uint64 = 0
var Beginning = time.Now().Unix()

type packet struct{
	id uint64
	created time.Time
}

type pNode struct{
	pac packet
	conn net.Conn
	next *pNode
	prev *pNode
}

func Listen(){
	tcpAddr,err := net.ResolveTCPAddr(IPScheme,ServerAddr)
	checkError(err)
	listener, err := net.ListenTCP(IPScheme, tcpAddr)
	checkError(err)
	i := 0
	go consumer()
	for{
		conn, err := listener.Accept()
		if err!=nil{
			fmt.Printf("Failed to accept incoming connection%v\n",err)
			continue
		}
		i++
		go handleClient(conn, i)
	}
}

func consumer(){
	{
		var pnode *pNode
		var ok bool
		var key int64 
		for{
				key=getKey(time.Now().Unix())
					if pnode,ok = removeFromStore(key);ok{
						fmt.Printf("Consuming packet with key %d at %v\n", key, time.Now().Unix())
						go process(pnode)
					}
			
		}
	}
}

func process(p *pNode){
	for ;p!=nil;p=p.next{
		p.conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\n Connection#%d freed", p.pac.id)))
		p.conn.Close()
		removeFromLookUp(p.pac.id)
	}
}

func handleClient(conn net.Conn, i int){
	request := make([]byte,128)//limit request length
	//defer conn.Close()

	var connId uint64 = 0
	readLength,err:= conn.Read(request)
	checkError(err)
	if readLength == 0{
		return //connection closed
	}else{
		requestData := string(request[:readLength])
		requestTokens := strings.Split(requestData," ")

		//method:=requestTokens[0]
		endpoint:= requestTokens[1]

		uri:=strings.Split(endpoint,"?")
		path:= strings.TrimSpace(uri[0])
		params:= ""
		if len(uri)>1{
			params = uri[1]
		}
		paramMap := getParameterMap(params)
		switch path {
		case "/debug":conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\nHey %v",getParameterMap(params))))
		case "/sleep":
			timeVal,err := strconv.Atoi(paramMap["timeout"]);
			if err!=nil{ 
				checkError(err)
				timeVal = Timeout	
			}
			connId = atomic.AddUint64(&ConnId,1)
			key:= getKey(time.Now().Unix()+int64(timeVal))
			now := time.Now().Unix()
				fmt.Printf("New packet recieved with key %d at %v\n", key, now)
				addToStore(key, connId, time.Now(), conn) //now + timeout
		case "/status":
			fmt.Printf("Number of active connections: %d\n",len(store))
			fmt.Printf("Number of active connections: %v\n",connLookup)
			conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\n\r\n\rNumber of active connections: %v",store)))
			conn.Close()
		case "/kill":
			connId,_ := strconv.Atoi(paramMap["connId"]);
			go process(connLookup[uint64(connId)])
			default:conn.Write([]byte("HTTP/1.1 200 OK"))
			conn.Close()
		}
		
	}
}

func addToStore(key int64,connId uint64, t time.Time, conn net.Conn){
	mutex.Lock()
	defer mutex.Unlock()
	val,ok := store[key]
	newNode := &pNode{packet{connId,t},conn,val, nil} 
	if ok{
		val.prev = newNode
	}
	go addToLookUp(connId,newNode) 
	store[key] = newNode
}

func removeFromStore(key int64) (*pNode,bool){
	mutex.Lock()
	defer mutex.Unlock()
		val,ok := store[key]
		if ok{
			delete(store,key)
			return val,true
		}
		return val,false
}


func addToLookUp(connId uint64, pn *pNode){
	mutexOld.Lock()
	defer mutexOld.Unlock()
	connLookup[connId] = pn
}

func removeFromLookUp(connId uint64) {
	mutexOld.Lock()
	defer mutexOld.Unlock()
		val,ok := connLookup[connId]
		if ok{
			delete(connLookup,connId)
			go process(val)
		}
}


func getKey(ts int64) int64 {
		return ts%RotatingFactor
}

func getParameterMap(paramString string) map[string]string{
	kvs:=	strings.Split(paramString,"&")
	ret := make(map[string]string, len(kvs))
	for _,kv:= range kvs{
		kvArr:= strings.Split(kv,"=")
		if len(kvArr)>1{
			ret[strings.TrimSpace(kvArr[0])] = strings.TrimSpace(kvArr[1])
		}
	}
	return ret
}

func checkError(err error){
	if err != nil{
		fmt.Printf("Error occured: %v \n", err)
	}
}