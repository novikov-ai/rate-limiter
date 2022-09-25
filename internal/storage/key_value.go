package storage

import (
	"context"
)

type KeyValue interface {
	Connect(ctx context.Context) error
	Close() error

	FindAtBlackList(ip string) (bool, error)
	FindAtWhiteList(ip string) (bool, error)

	OverflowAttemptsLogin(login string) (bool, error)
	OverflowAttemptsPasswords(password string) (bool, error)

	Add(set, key, value string) error
	Remove(set, key string) error

	RemoveAllLoginsAttempts(logins []string) error
	RemoveAllAddressesAttempts(ips []string) error
}

const (
	WhiteList = "white" // global set
	BlackList = "black" // global set

	SetPasswords = "passwords" // global set

	SetLogins    = "logins"
	SetAddresses = "ips"
)
