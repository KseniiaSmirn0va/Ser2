DetectorType_JiraToken                     DetectorType = 120
Token 23417846781247eef
AKIASP2TPHJSQH3FJRUX

data:   []byte(fmt.Sprintf("25:KEY='rzp_live_SnDTaP1ncfliDt'\n\"rzp_secret\" : \" %s\", ", secretInactive)),

Rfc2898DeriveBytes rdb = new Rfc2898DeriveBytes(password, "salt", 1000000);
import (
    "golang.org/x/crypto/scrypt"
    "fmt"
)

func main() {

    dk, err := scrypt.Key(pwd, []byte{0xc8, 0x28, 0xf2, 0x58, 0xa7, 0x6a, 0xad, 0x7b}, 1<<15, 8, 1, 32)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(base64.StdEncoding.EncodeToString(dk))
}
