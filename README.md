# zteMF823Client
Golang client for ZTE MF823 4G modem

### Usage:
Get package:

`go get github.com/devcentr/zteMF823Client``

Create client

```
import modem "github.com/devcentr/zteMF823Client"

// IP of your modem
const modemIp string = "192.168.0.1"

func main() {
	client := modem.NewClient(modemIp)
    // your code
}
```

or with custom http client

```
import modem "github.com/devcentr/zteMF823Client"

// IP of your modem
const modemIp string = "192.168.0.1"

func main() {
    // your custom http client
    customHttpClient := &http.Client{}

	client := modem.NewClient(modemIp)
    client.SetCustomClient(customHttpClient)
	// your code
}
```

#### API
##### SMS
###### Send SMS
SendSMS(r SendSMSRequest) (*SendSMSResponse, error)
```
response, err := client.SendSMS(modem.SendSMSRequest{
    Text: "Hello",
    Phone: "+79876543210",
})
```
