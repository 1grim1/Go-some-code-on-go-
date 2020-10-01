package main

import (
	"os"
	"bytes"
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "errors"
    "fmt"
    "io"
    "log"
)

func main() {
	var str string
	fmt.Scanf("%s", &str)
    text := []byte(str)
    key := []byte("abcdefghijklmnopqrstuvwxyz")

    ciphertext, err := encrypt(text, key)
    if err != nil {
        // TODO: Properly handle error
        log.Fatal(err)
    }
    fmt.Printf("%s => %x\n", text, ciphertext)

    f, err := os.Create("pass.txt")
    if err != nil {
        panic(err.Error())
    }
    _, err = io.Copy(f, bytes.NewReader(ciphertext))
    if err != nil {
        panic(err.Error())
    }


    /*plaintext, err := decrypt(ciphertext, key)
    if err != nil {
        // TODO: Properly handle error
        log.Fatal(err)
    }
    fmt.Printf("%x => %s\n", ciphertext, plaintext)*/
}

func encrypt(plaintext []byte, key []byte) ([]byte, error) {
    c, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    gcm, err := cipher.NewGCM(c)
    if err != nil {
        return nil, err
    }

    nonce := make([]byte, gcm.NonceSize())
    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, err
    }

    return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

/*func decrypt(ciphertext []byte, key []byte) ([]byte, error) {
    c, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    gcm, err := cipher.NewGCM(c)
    if err != nil {
        return nil, err
    }

    nonceSize := gcm.NonceSize()
    if len(ciphertext) < nonceSize {
        return nil, errors.New("ciphertext too short")
    }

    nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
    return gcm.Open(nil, nonce, ciphertext, nil)
}*/
