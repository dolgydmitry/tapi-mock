package token

import (
	"fmt"
	"tapi-controller/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateToken(t *testing.T) {
	config, _ := utils.LoadConfigForTest()
	keys, err := utils.LoadKeys(config)
	require.NoError(t, err)
	tokenMaker := JwtMaker{PrivateKey: keys.PrivateKey, PublicKey: keys.PublicKey}
	accessToken, err := tokenMaker.CreateToken("test", time.Second*10)
	require.NoError(t, err)
	require.NotEmpty(t, accessToken)
	fmt.Println(accessToken)
}

func TestVerifeToken(t *testing.T) {
	config, _ := utils.LoadConfigForTest()
	keys, err := utils.LoadKeys(config)
	require.NoError(t, err)
	tokenMaker := JwtMaker{PrivateKey: keys.PrivateKey, PublicKey: keys.PublicKey}
	accessToken, err := tokenMaker.CreateToken("test", time.Second*20)
	require.NoError(t, err)
	wrongToken := "xyIhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImVkYzY1NDhmLWJhNjQtNDM0NS05ODcwLTdjZDJhM2EwZmE4YyIsInVzZXJuYW1lIjoidGVzdCIsImlzc3VlZF9hdCI6IjIwMjQtMDEtMDZUMTY6NDI6NTUuMDI5MTQyKzAxOjAwIiwiZXhwaXJlZF9hdCI6IjIwMjQtMDEtMDZUMTY6NDM6MDUuMDI5MTQyKzAxOjAwIn0.W5Btk9UiNg-2SSEFPPeK42Ypcv7KKtukx3nVjpxS9KLYGVSZB_BtaORHAwDuIrFOp928ExbTk2Rr4VxiExhM86ChlV_lo6eDxv29AqEgG8Hz5qg4urRPlPrZ787QHT_B9NTp5yuDkTiS-t7NhXTw9ms4vakpd1uHz_-yXD9HrBbBzjGKKTKuj4WBuP5H7cKx36IfxzOMDL2Bt7RRnrxhuUsgMTq80E8LRpgaM3WBrO9Kodlxt9pv8WlJG4sXi602j_m1jahUvg4iH_7e3nBf9wO2RRWOz3sRlWCNeacF5JMSPfr4e0vJznGZi_asHLj64TapLLtu69_dbZTn62RIOKBXEQWH0l9pTyOvAFFzXgbofnZnxa2qaryX08VAtaw8qLZ8JmZd1oiV1Il34771zkbdsZ_zIN44tAyCmV_CL7jkAad84vCKkncS2MqfavtsqG-66kIUkXZrEd0vWEVd5KjxccF2mtpmpDtFUsjtmqcLXJNBv_9F84raxLKeXUZRDsRf3kK_Uz7gyTG97UHmZ-79REs7CfzmXXtXNADzl50qGmYbD1jDNm34aFbh88kNlBuLOL6Ax67Oe0553SSbbTh77wI60h8TziCAnzByCijYMJ8ybqT4BCt7Tz4tBJt00ZVaFbwzRfnw9ASybj7c76lZbeXjkcGbApQou-WXNo4"
	expiredToken := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImVkYzY1NDhmLWJhNjQtNDM0NS05ODcwLTdjZDJhM2EwZmE4YyIsInVzZXJuYW1lIjoidGVzdCIsImlzc3VlZF9hdCI6IjIwMjQtMDEtMDZUMTY6NDI6NTUuMDI5MTQyKzAxOjAwIiwiZXhwaXJlZF9hdCI6IjIwMjQtMDEtMDZUMTY6NDM6MDUuMDI5MTQyKzAxOjAwIn0.W5Btk9UiNg-2SSEFPPeK42Ypcv7KKtukx3nVjpxS9KLYGVSZB_BtaORHAwDuIrFOp928ExbTk2Rr4VxiExhM86ChlV_lo6eDxv29AqEgG8Hz5qg4urRPlPrZ787QHT_B9NTp5yuDkTiS-t7NhXTw9ms4vakpd1uHz_-yXD9HrBbBzjGKKTKuj4WBuP5H7cKx36IfxzOMDL2Bt7RRnrxhuUsgMTq80E8LRpgaM3WBrO9Kodlxt9pv8WlJG4sXi602j_m1jahUvg4iH_7e3nBf9wO2RRWOz3sRlWCNeacF5JMSPfr4e0vJznGZi_asHLj64TapLLtu69_dbZTn62RIOKBXEQWH0l9pTyOvAFFzXgbofnZnxa2qaryX08VAtaw8qLZ8JmZd1oiV1Il34771zkbdsZ_zIN44tAyCmV_CL7jkAad84vCKkncS2MqfavtsqG-66kIUkXZrEd0vWEVd5KjxccF2mtpmpDtFUsjtmqcLXJNBv_9F84raxLKeXUZRDsRf3kK_Uz7gyTG97UHmZ-79REs7CfzmXXtXNADzl50qGmYbD1jDNm34aFbh88kNlBuLOL6Ax67Oe0553SSbbTh77wI60h8TziCAnzByCijYMJ8ybqT4BCt7Tz4tBJt00ZVaFbwzRfnw9ASybj7c76lZbeXjkcGbApQou-WXNo4"

	testCases := []struct {
		name        string
		accessToken string
		checker     func(t *testing.T, payload *Payload, err error)
	}{
		{
			name:        "good day",
			accessToken: accessToken,
			checker: func(t *testing.T, payload *Payload, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, payload)
			},
		},
		{
			name:        "invalid token",
			accessToken: wrongToken,
			checker: func(t *testing.T, payload *Payload, err error) {
				require.ErrorIs(t, ErrInvalidToken, err)
			},
		},
		{
			name:        "expired token",
			accessToken: expiredToken,
			checker: func(t *testing.T, payload *Payload, err error) {
				require.ErrorIs(t, ErrExpiredToken, err)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		payload, err := tokenMaker.VerifeToken(tc.accessToken)
		tc.checker(t, payload, err)
	}
}
