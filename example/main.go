package main

import (
	"fmt"
	"time"

	"f-license/client"
)

// make the license configurable
var licenseKey = "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhZGRyZXNzIjoiSXN0YW5idWwsIFR1cmtleSIsInVzZXJuYW1lIjoiRnVya2FuIn0.bpdpgglPhYPQd-kKCP8uCtH_NqLLg183DxMiHKVbBE1hrZbIUhBQ4R7_otPF0DmrGKbG-XVtQG0IeG-i4uLFYA"

// pin this to your app
const secret = `-----BEGIN CERTIFICATE-----
MIIEpDCCAowCCQC14kMH/BQPvDANBgkqhkiG9w0BAQsFADAUMRIwEAYDVQQDDAls
b2NhbGhvc3QwHhcNMjAwNDE4MjAyMTQ5WhcNMjEwNDE4MjAyMTQ5WjAUMRIwEAYD
VQQDDAlsb2NhbGhvc3QwggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIKAoICAQDQ
0itWTslOpF43f+gc/7xW9VPaEn0+cAKW4TFyGZv5wLDiMq7ecPGrGJpclaLJKObF
7MrayvBFnFIRZ1FVg314h47rhdZy0O30h5PEFndPwtxEiC9hVJ3gqp8zV4LewtKq
uP/0C+1A1DC8WQgq5Uu817VLxzuCwmS/Ad8eOpiRHbpant9yHRKDiiT8tMZ3TVeu
kLrqmVhyIyQY31kevdHI9aONBumiqROEL2ehpD9LU3Ds0lpuQw74hJUbrK+AIpVi
HpJt57FdYe6mjant57cl2C6pfSYlqacceX4rMcDjqwkbzp0SlNgSv4fCGwp1lsm0
tPabiCO3T/O2zbZOGBiH9WZb7k0Y+2O87rLWnInrTagZEXywfUS7VgMv9Q4sxnws
oGbPMHn1rhOd3fPs68WS7UBZK4yJ9xaWWNgpZrzQjDiRl93sc4/efxBNhw1IQRK1
4DOr+OBPAyxpnYgfZJ2Bl3MTkF83bCN9oISLgxE3QJjxNQWn2zACktVQqphMG9jE
zQPrEyLiohJKMI6tRS9ZXQTKBONd+I1jTMSb+iITndXCCBwpgllAs7H/gxzaqWlV
M7/j+lzBcyGbodSO9+AJz6z3tae72lpGYcmYG3tjhKZun0Rt59D+PuBfwPQY+CLf
nw8yXoZw2ZQGj/Q15BTUVNPgN4z3Ahn3yIForMrJtQIDAQABMA0GCSqGSIb3DQEB
CwUAA4ICAQC2qqFmES2jjoxfhzSnsW0XbwQ1GsgB00mE54i207R4R20c16KXP8F9
UxeQ8rXgFZx4r9y2Yl+Yjr22dz7LNnliSimRwY+SesiZRShG390hM7J+xCcPjgpp
77jQqs2sgEzwD3nQWtAY6qQDOeapUtVj98Azk9UDAS3euwiBID06qa181e/5yZIr
x94Xbbm95C2/BKfKbrujJ9hsjwdnZnSIivUnzti01IsQcNZsb/L33pUcLD/U6O5J
v9Ut1nmEt66gJoUtLzahEaGT2xlqkhBd3ocU2am1y09FTDPZEB93rlbPBL9x3/wv
D30ocQys3DlSMSo8MgetwkDs9uzPYW7PGH7eIpq50Z8ArAEJvQiLtSOnpLTLugBC
tQxYqni+c6ohJAeWuULvj2adGdaydvPZyfesqM9+AMNOhqe5HSneaFhgewIUeTQ9
GluL9AhJQivwiQdNZzpEUJlg7x7cHwCsoeyByVcYxy8Kadb4N2wpqeJGknUU3+P0
pd0akjEcO0+lKNG1miVjjNKfc9dRPWIAeTT0bxevZ0/XXksS/AyUAdzmG1CVhS6U
XtJ6JsVTghGjZ4R8R4R9h5tb49T8vYLuS9Td9kqK2AFn3b0oaYHXYanpn4Jsd2/+
uvtCmm544fhjh2D/dGl5DFAOd4YWG/t7wUTn3fkEhiDYX6tv6A85Zw==
-----END CERTIFICATE-----`

var LicenseValid bool

func main() {
	//localVerification()
	remoteVerification()

	if LicenseValid {
		fmt.Println("An operation can be done if the license is valid")
	}
}

func localVerification() {
	for {
		LicenseValid, err := client.VerifyLocally(secret, licenseKey)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(LicenseValid)
		time.Sleep(2 * time.Second)
	}
}

func remoteVerification() {
	for {
		LicenseValid, err := client.VerifyRemotely("https://localhost:4242", secret, licenseKey)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(LicenseValid)
		time.Sleep(2 * time.Second)
	}
}
