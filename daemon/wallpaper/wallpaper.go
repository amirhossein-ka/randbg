package wallpaper

import (
	"fmt"
	"math/rand"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/amirhossein-ka/randbg/daemon/notify"
)

// ChangeWall changes the wallpaper and send a notification
func ChangeWall(pics []string) error {
	rand.Seed(time.Now().UnixNano())
	pic := pics[rand.Intn(len(pics))]
	cmd := exec.Command("feh", "--bg-fill", "--no-fehbg", pic)

    desc := fmt.Sprintf("Wallpaper has changed to %s\n", filepath.Base(pic))
    notify.Notify("Wallpaper changed", desc, pic)

	return cmd.Run()
}
