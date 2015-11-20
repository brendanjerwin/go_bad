package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"gopkg.in/redis.v3"
)

var _ = Describe("Redis", func() {
	var err error
	var client *redis.Client

	BeforeEach(func() {
		client = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	})

	Describe("PING", func() {
		var pong string
		BeforeEach(func() {
			pong, err = client.Ping().Result()
		})

		It("should return PONG", func() {
			Expect(pong).To(Equal("PONG"))
		})
	})
})
