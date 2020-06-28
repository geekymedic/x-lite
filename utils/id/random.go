package id

import (
	"crypto/md5"
	"encoding/hex"
	"hash"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/lithammer/shortuuid/v3"
	"github.com/zentures/cityhash"
)

var timeRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func RangeBitsInt(low, hi int) int {
	if low > hi {
		panic("low must be less or equal hi")
	}
	return low + timeRand.Intn(hi-low)
}

var syncPool sync.Pool

func init() {
	syncPool.New = func() interface{} {
		return cityhash.New64()
	}
}

func MD5UUID4() string {
	hasher := md5.New()
	txt := uuid.New()
	hasher.Write(txt[:])
	return hex.EncodeToString(hasher.Sum(nil))
}

func ShortUUID() string {
	return shortuuid.New()
}

func RandomUint64(input ...[]byte) uint64 {
	input = append(input, []byte(MD5UUID4()))
	hash64 := syncPool.Get().(hash.Hash64)
	for _, b := range input {
		hash64.Write(b)
	}
	id := hash64.Sum64()
	hash64.Reset()
	return id
}

func ConvertBytesToUint64(input ...[]byte) uint64 {
	input = append(input, []byte(MD5UUID4()))
	hash64 := syncPool.Get().(hash.Hash64)
	id := hash64.Sum64()
	hash64.Reset()
	return id
}
