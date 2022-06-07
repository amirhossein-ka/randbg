package lib

import (
	"context"
	"fmt"
	"math/rand"
	"os/exec"
	"time"
)

func ChangeWall(ctx context.Context, pics []string) error {
	rand.Seed(time.Now().UnixNano())
	pic := pics[rand.Intn(len(pics))]
	fmt.Println(pic)
	cmd := exec.CommandContext(ctx, "feh", "--bg-fill", pic)

	return cmd.Run()
}
