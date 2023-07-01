package userscore

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"feather-ai/service-core/lib/users"
	"fmt"
	"io"
	"time"

	uuid "github.com/satori/go.uuid"
)

var gEncryptionKey = "tDX@a$o%JiVW9uSfQPJoHGOUwEShp5OD" // MUST be 32 bytes - configure per environment
var gFeatherTokenSecret = "^9jr4LncOIiP4*K"             // :TODO: Get this from config - different per environment

func Base64Encode(message []byte) string {
	b := make([]byte, base64.StdEncoding.EncodedLen(len(message)))
	base64.StdEncoding.Encode(b, message)
	return string(b)
}

func Base64Decode(message string) (b []byte, err error) {
	var l int
	messageBytes := []byte(message)
	b = make([]byte, base64.StdEncoding.DecodedLen(len(messageBytes)))
	l, err = base64.StdEncoding.Decode(b, messageBytes)
	if err != nil {
		return
	}
	return b[:l], nil
}

func GenerateEncryptedFeatherToken(userId uuid.UUID) *users.EncryptedFeatherToken {
	expiry := time.Now().UTC().Add(time.Minute * 1)
	token := users.FeatherTokenData{
		UserID: userId,
		Expiry: expiry,
		Secret: gFeatherTokenSecret,
	}

	jsonBytes, err := json.Marshal(&token)
	if err != nil {
		return nil
	}

	c, err := aes.NewCipher([]byte(gEncryptionKey))
	if err != nil {
		return nil
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil
	}

	nonce := make([]byte, gcm.NonceSize())
	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	encryptedJson := gcm.Seal(nonce, nonce, jsonBytes, nil)
	return &users.EncryptedFeatherToken{
		Payload: Base64Encode(encryptedJson),
		Expiry:  expiry,
	}
}

func DecryptFeatherToken(token users.EncryptedFeatherToken) *users.FeatherTokenData {
	c, err := aes.NewCipher([]byte(gEncryptionKey))
	if err != nil {
		return nil
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil
	}

	encryptedPayload, err := Base64Decode(token.Payload)
	if err != nil {
		return nil
	}

	nonceSize := gcm.NonceSize()
	if len(encryptedPayload) < nonceSize {
		return nil
	}

	nonce, ciphertext := encryptedPayload[:nonceSize], encryptedPayload[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)

	unencyptedToken := users.FeatherTokenData{}
	if json.Unmarshal(plaintext, &unencyptedToken) != nil {
		return nil
	}
	return &unencyptedToken
}
