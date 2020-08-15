package main

import (
	"fmt"
	"testing"
	"time"
)

func TestFuture_Get(t *testing.T) {

	t.Run("should wait for unfinished future", func(t *testing.T) {
		f1 := NewFuture(func() (string, error) {
			time.Sleep(1000)
			return "OK!", nil
		})

		fmt.Printf("ready with %v \n", f1.Get())
	})

	t.Run("should get result from finished future", func(t *testing.T) {

	})

	t.Run("should return an error", func(t *testing.T) {
		f2 := NewFuture(func() (string, error) {
			time.Sleep(1000)
			return "", fmt.Errorf("something went wrong")
		})

		fmt.Printf("ready with %v \n", f2.Get())
	})
}

func TestFuture_Cancel(t *testing.T) {
	t.Run("should be able to cancel a future", func(t *testing.T) {

	})
}
