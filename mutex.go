package main

import (
	"fmt"
	"sync"
)

type Vault struct {
	secrets map[string]string
	mutex   sync.Mutex
}

// Note without the locks it can result in the race condition "concurrent map writes" where underlying structure is being
// reorganised (due to write) while another is trying to read!
func (vault *Vault) Store(key string, secret string) {
	vault.mutex.Lock()
	defer vault.mutex.Unlock()
	vault.secrets[key] = secret
}

func (vault *Vault) Read(key string) (string, bool) {
	vault.mutex.Lock()
	defer vault.mutex.Unlock()
	secret, ok := vault.secrets[key]
	return secret, ok
}

func startMutex() {
	fmt.Println("MUTEX STARTING...")
	myVault := Vault{
		secrets: make(map[string]string),
	}

	var wg sync.WaitGroup

	for x := 0; x < 10; x++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			key := fmt.Sprintf("key_%d", x)
			val := fmt.Sprintf("val_%d", x)
			myVault.Store(key, val)
		}()
	}

	for x := 0; x < 10; x++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			key := fmt.Sprintf("key_%d", x)
			val, ok := myVault.Read(key)

			if ok {
				fmt.Printf("Key: %s = %s\n", key, val)
			} else {
				fmt.Printf("Key does not exist: %s\n", key)
			}
		}()
	}

	wg.Wait()
	fmt.Printf("MUTEX FINISHED.\n\n")
}
