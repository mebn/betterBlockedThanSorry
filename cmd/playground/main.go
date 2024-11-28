package main

import "github.com/mebn/betterBlockedThanSorry/internal/updater"

func main() {
	u, err := updater.NewUpdater()
	if err != nil {
		println(err)
		return
	}

	p, err := u.DownloadLatestBinary()
	if err != nil {
		println(err)
		return
	}

	println(p)
}
