package file

import (
	"os"
	"time"

	"github.com/qinhan-shu/gp-server/logger"
)

var (
	period = 30 * time.Minute
)

// Daemon : del expired file
func (f *File) Daemon() {
	ticker := time.NewTicker(period)
	for {
		select {
		case <-ticker.C:
			{
				files, err := f.cache.GetExpiredFile()
				if err != nil {
					logger.Sugar.Errorf("failed to get expired file : %v", err)
					continue
				}
				for _, file := range files {
					os.RemoveAll(f.path + file)
					f.cache.DelFileItem(file)
				}
			}
		default:
			{
				time.Sleep(60 * time.Second)
			}
		}
	}
}
