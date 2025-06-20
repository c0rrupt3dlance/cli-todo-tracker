package errors

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

func CheckAddArgs(args []string) error {
	if len(args) != 3 {
		log.Printf("user putted more than 1 argument")
		return errors.New("invalid number of arguments for add function")
	}
	return nil
}

func CheckDeleteArgs(args []string) error {
	if len(args) != 3 {
		log.Printf("user putted more than 1 argument")
		return errors.New("invalid number of arguments for delete function")
	}
	_, err := strconv.Atoi(args[2])
	if err != nil {
		log.Printf("invalid id argument")
		return errors.New("invalid id argument")
	}
	return nil
}

func CheckUpdateArgs(args []string) error {
	if len(args) != 4 {
		log.Printf("user putted %v number of arguments", len(args))
		return errors.New("invalid number of arguments for update function")
	}
	_, err := strconv.Atoi(args[2])
	if err != nil {
		log.Printf("invalid id argument")
		return errors.New("invalid id argument")
	}
	return nil
}

func CheckStatusArgs(args []string, status string) error {
	if len(args) != 3 {
		log.Printf("user putted %v number of arguments", len(args))
		return errors.New(fmt.Sprintf("invalid number of arguments for %s status", status))
	}
	_, err := strconv.Atoi(args[2])
	if err != nil {
		log.Printf("invalid type of argument")
		return errors.New("invalid type of argument")
	}
	return nil
}
