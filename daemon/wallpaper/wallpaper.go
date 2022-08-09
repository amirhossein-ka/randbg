package wallpaper

import (
	"fmt"
	"math/rand"
	"os/exec"
	"time"
)

func ChangeWall(pics []string) error {
	rand.Seed(time.Now().UnixNano())
	pic := pics[rand.Intn(len(pics))]
	fmt.Println(pic)
	cmd := exec.Command("feh", "--bg-fill", "--no-fehbg", pic)

	return cmd.Run()
}
