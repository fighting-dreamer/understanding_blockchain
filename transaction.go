package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"k8s.io/apimachinery/pkg/util/json"
)

type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

func NewTransaction(sender, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	println("Sender: " + t.senderBlockchainAddress)
	println("Recipient: " + t.recipientBlockchainAddress)
	println("Value: " + strconv.FormatFloat(float64(t.value), 'f', 6, 64))
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	// return []byte(fmt.Sprintf(`{"sender": "%s", "recipient": "%s", "value": %f}`,
	// 	t.senderBlockchainAddress, t.recipientBlockchainAddress, t.value)), nil
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderBlockchainAddress,
		Recipient: t.recipientBlockchainAddress,
		Value:     t.value,
	})
}
