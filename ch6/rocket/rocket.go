package rocket

import (
	"fmt"
	"time"
)

// Rocket is rocket
type Rocket struct {
	Name string
}

//Launch is print
func (r *Rocket) Launch() {
	fmt.Println("Launched!!!!")
}

func main() {
	//ロケット打ち上げ
	r := new(Rocket)
	time.AfterFunc(10*time.Second, func() {
		fmt.Println("start AfterFunc")
		r.Launch()
		fmt.Println("end AfterFunc")
	})

	//メソッド値の構文だともっと短くなる
	time.AfterFunc(10*time.Second, r.Launch)
}
