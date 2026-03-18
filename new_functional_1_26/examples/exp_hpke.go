package examples

import (
	"crypto/ecdh"
	"crypto/hpke"
	"fmt"
	"log"
)

func Exphpke() {
	kem := hpke.DHKEM(ecdh.X25519())
	kdf := hpke.HKDFSHA256()
	aead := hpke.AES256GCM()

	privateKey, err := kem.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.PublicKey()

	message := []byte("Hello Go 1.26!")
	info := []byte("example context")

	ciphertext, err := hpke.Seal(publicKey, kdf, aead, info, message)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Enc: %x\n", ciphertext)

	plaintext, err := hpke.Open(privateKey, kdf, aead, info, ciphertext)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Dec: %s\n", plaintext)
}

func ExphpkeSenderRecipient() {
	kem := hpke.DHKEM(ecdh.X25519())
	kdf := hpke.HKDFSHA256()
	aead := hpke.ChaCha20Poly1305()

	privateKey, _ := kem.GenerateKey()
	publicKey := privateKey.PublicKey()

	// Sender
	enc, sender, err := hpke.NewSender(publicKey, kdf, aead, []byte("session info"))
	if err != nil {
		log.Fatal(err)
	}

	// Recipient
	recipient, err := hpke.NewRecipient(enc, privateKey, kdf, aead, []byte("session info"))
	if err != nil {
		log.Fatal(err)
	}

	messages := []string{"Message 1", "Message 2", "Message 3"}

	for _, msg := range messages {
		// Encryption
		ct, err := sender.Seal(nil, []byte(msg))
		if err != nil {
			log.Fatal(err)
		}

		// Deciphering
		pt, err := recipient.Open(nil, ct)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Sent: %s, received: %s\n", msg, pt)
	}
}

func ExphpkeKEM() {
	// Hybrid KEM: ML-KEM-768 + X25519 (X-Wing)
	kem := hpke.MLKEM768X25519()
	kdf := hpke.HKDFSHA256()
	aead := hpke.AES256GCM()

	privateKey, err := kem.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	message := []byte("Post-quantum message")
	ct, err := hpke.Seal(privateKey.PublicKey(), kdf, aead, nil, message)
	if err != nil {
		log.Fatal(err)
	}

	pt, err := hpke.Open(privateKey, kdf, aead, nil, ct)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Public key size: %d bytes\n", len(privateKey.PublicKey().Bytes()))
	fmt.Printf("Ciphertext size: %d bytes\n", len(ct))
	fmt.Printf("Message: %s\n", pt)
}
