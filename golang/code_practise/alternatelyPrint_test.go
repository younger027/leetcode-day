package code_practise

import (
	"encoding/csv"
	"fmt"
	simpleJson "github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestAlternatelyPrintNumAndLetter(t *testing.T) {
	//AlternatelyPrintNumAndLetter()

	//MutifyWriteReadChanOP()

	m := map[string]bool{"a": true, "b": true}

	if m["a"] {
		fmt.Println("aa")
	} else if m["b"] {
		fmt.Println("b")
	}
}

func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	keyLength := len(key)
	for i := 0; i < keyLength; i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}

func TestSomething(t *testing.T) {
	fmt.Println(fnv32("younger"))
}

var SHARD_COUNT = 32

// A "thread" safe map of type string:Anything.
// To avoid lock bottlenecks this map is dived to several (SHARD_COUNT) map shards.
type ConcurrentMap[V any] []*ConcurrentMapShared[V]

// A "thread" safe string to anything map.
type ConcurrentMapShared[V any] struct {
	items        map[string]V
	sync.RWMutex // Read Write mutex, guards access to internal map.
}

// Creates a new concurrent map.
func New[V any]() ConcurrentMap[V] {
	m := make(ConcurrentMap[V], SHARD_COUNT)
	for i := 0; i < SHARD_COUNT; i++ {
		m[i] = &ConcurrentMapShared[V]{items: make(map[string]V)}
	}
	return m
}

func (m ConcurrentMap[V]) GetShard(key string) *ConcurrentMapShared[V] {
	return m[uint(fnv32(key))%uint(SHARD_COUNT)]
}

func (c *ConcurrentMap[V]) Store(key string, value V) {

}

// MapKeys returns a slice of all the keys in m.
// The keys are not returned in any particular order.
func MapKeys[Key comparable, Val any](m map[Key]Val) []Key {
	s := make([]Key, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	return s
}

func TestMerFile(t *testing.T) {
	name1 := "/Users/rockey-lyy/GoWork/leetcode-day/result.txt"
	name2 := "/Users/rockey-lyy/GoWork/leetcode-day/hot_list.csv"

	MergeFile(name1, name2)
}

func MergeFile(f1, f2 string) []string {
	data, err := ioutil.ReadFile(f1)
	if err != nil {
		fmt.Printf("文件打开失败=%v\n", err)
	}

	file, err := os.Create(f2)
	if err != nil {
		fmt.Println("open file is failed, err: ", err)
	}

	defer file.Close()

	w := csv.NewWriter(file)

	filter := make(map[string]string)
	r := strings.Split(string(data), ",")
	for i := 0; i < len(r); i++ {
		if len(r[i]) == 0 {
			continue
		}

		res := NetGetData(r[i])
		filter[r[i]] = res

		hot := strings.ReplaceAll(r[i], "\"", "")
		fmt.Println("process---", hot, res)

		record := []string{hot, res}
		w.Write(record)
		w.Flush()
	}

	return nil
}

func NetGetData(k string) string {
	data := fmt.Sprintf(`{"key_word":%v}`, k)

	req, err := http.NewRequest("POST", "http://154.82.111.42:8888/ytsearch", strings.NewReader(data))
	if err != nil {
		fmt.Println("newrequest err:", err)
		return ""
	}

	client := http.Client{Timeout: 3 * time.Second}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client.Do(req) err:", err)
		return ""
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ReadAll err:", err)
		return ""
	}

	js, err := simpleJson.NewJson(b)
	if err != nil {
		fmt.Println("ReadAll err:", err)
		return ""
	}

	result := js.Get("result").MustArray()
	if len(result) == 0 {
		return ""
	}

	title := js.Get("result").GetIndex(0).Get("title").MustString()

	return title
}
