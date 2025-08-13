package discord

import(
	"net/http"
	"fmt"    
	"crypto/ed25519"
    "encoding/hex"

)

func RespPong() InteractionResponse{
	return InteractionResponse{
		Type: InteractionTypePing,
	}
}


func RespMessage(content string, ephemeral bool) InteractionResponse {
	flags := 0
	if ephemeral {
		flags = EphemeralMessageFlag
	}

	return InteractionResponse{
		Type: InteractionTypeChannelMessage,
		Data: &InteractionCallbackData{
			Content: content,
			Flags: flags,
		},
	}
}
 
func VerifySignature(r *http.Request, publicKey string, body []byte) bool {
	sigHex := r.Header.Get("X-Signature-Ed25519")
	ts := r.Header.Get("X-Signature-Timestamp")

	sigBytes, err := hex.DecodeString(sigHex)
	if (err != nil) {
		fmt.Println("Error decoding signature hex: ", err)
		return false
	}

	pubKeyBytes, err := hex.DecodeString(publicKey)
	if (err != nil) {
		fmt.Println("Error decoding public key hex: ", err)
		return false
	}
	
	msg := append([]byte(ts), body...)
	isValid := ed25519.Verify(pubKeyBytes, msg, sigBytes)
	fmt.Println("Signature verification: ", isValid)
	return isValid
}
