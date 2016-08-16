# Go BareMetal SDK ![](https://circleci.com/gh/MustWin/baremetal-sdk-go.svg?style=shield&circle-token=fa06ce2af6b594812e3a756f5451a9e101d7b9f5)

Package baremetal provides access to the Oracle BareMetal Cloud APIs.

## Usage

To use the Go BareMetal SDK instantiate a baremetal.Client, supplying  
your tenancy OCID, user OCID, RSA public key fingerprint, and RSA private key.
Then call functions as the example below illustrates.  Note that error
handling has been omitted to add clarity. See [API Docs](https://docs.us-az-phoenix-1.oracleiaas.com/) for more information.


```
import (
  "fmt"  
  "crypto/rsa"
  "github.com/mustwin/baremetal-sdk-go"
)

func main() {
  privateKey, _ := baremetal.PrivateKeyFromFile("/path/to/key.pem", "keyPassword")

  client := baremetal.New(
    "ocid1.tenancy.oc1..aaaaaaaaq3hulfjvrouw3e6qx2ncxtp256aq7etiabqqtzunnhxjslzkfyxq",
    "ocid1.user.oc1..aaaaaaaaflxvsdpjs5ztahmsf7vjxy5kdqnuzyqpvwnncbkfhavexwd4w5ra",
    "b4:8a:7d:54:e6:81:04:b2:99:8e:b3:ed:10:e2:12:2b",
    privateKey,
  )

  availabilityDomains, _ := client.ListAvailablityDomains()

  for _, ad := range availabilityDomains {
    fmt.Println(ad.Name)
  }
}
```
## Unit Testing
Some of the tests rely on GOPATH to build a path where a test private key is located. If
for some reason you have a composite GOPATH i.e /home/foo/go-projects:/usr/stuff
these tests will break.  In that case you export an environment variable with an
explicit path to the test private key.

```
export BAREMETAL_SDK_PEM_DATA_PATH="/home/foo/go-projects/src/github.com/../test/data/private.pem"
```

[Regression Tests](test/README.md)


# Vendoring
This project uses the [Go vendor folder](https://blog.gopheracademy.com/advent-2015/vendor-folder/) for dependencies.
If you need to add or update dependency, please review the [go
vendor docs](https://github.com/kardianos/govendor).

# References
[Oracle Bare Metal Iaas API Docs](https://docs.us-az-phoenix-1.oracleiaas.com/)
