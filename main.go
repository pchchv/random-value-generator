package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	letters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers       = "0123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // Number of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())
var collection *mongo.Collection

func init() {
	// Load values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Panic("No .env file found")
	}
}

func getEnvValue(v string) string {
	value, exist := os.LookupEnv(v)
	if !exist {
		log.Panicf("Value %v does not exist", v)
	}
	return value
}

func valueGeneration(valType string, valLength int) string {
	var symbols string
	if valLength == 0 {
		valLength = int(src.Int63())
		for valLength > 66 {
			valLength /= 3
		}
		for valLength < 4 {
			valLength *= 3
		}
	}
	switch valType {
	case "num":
		symbols = numbers
	case "str":
		symbols = letters
	case "alp":
		symbols = letters + numbers
	}
	value := generator(valLength, symbols)
	return value
}

func generator(n int, symbols string) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(symbols) {
			b[i] = symbols[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func toDB(value string) string {
	v := bson.D{{Key: "value", Value: value}}
	result, err := collection.InsertOne(context.TODO(), v)
	if err != nil {
		log.Panic(err)
	}
	return fmt.Sprint(result.InsertedID)
}

func fromDB(id string) string {
	// TODO: Get value from database by id
	return id
}

func getJSON(pre string, str string) []byte {
	var res []byte
	s, err := json.MarshalIndent(str, "\t", "\t")
	if err != nil {
		log.Panic(err)
	}
	if pre != "" {
		pr, err := json.MarshalIndent(pre, "\t", "\t")
		if err != nil {
			log.Panic(err)
		}
		s = append(pr, s...)
	}
	res = append(res, s...)
	return res
}

func main() {
	server()
	db()
}
