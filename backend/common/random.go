package common

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

type CryptoSource struct{}

type randomString struct{}

type RandomString interface {
	Get() string
}

func NewRandomString() RandomString {
	return &randomString{}
}

func (s *randomString) Get() string {
	rand.Seed(time.Now().Unix())

	charSet := "abcdedfghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var output strings.Builder
	length := 10

	var src CryptoSource
	// nolint
	rnd := rand.New(src)

	for i := 0; i < length; i++ {
		random := rnd.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}
	return output.String()
}

func (s CryptoSource) Seed(seed int64) {}

func (s CryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s CryptoSource) Uint64() (v uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &v)
	if err != nil {
		log.Error().Err(err).Caller().Msg("Failed to read from random number generator.")
	}
	return v
}
