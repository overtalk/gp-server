package file

import (
	"os"
	"time"

	"github.com/qinhan-shu/gp-server/logger"
)

var (
	period = 20 * time.Minute
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
				logger.Sugar.Infof(" expired files : %v", files)
				for _, file := range files {
					if err := os.RemoveAll(f.path + file); err != nil {
						logger.Sugar.Errorf("failed to del dir[%s] : %v", f.path+file, err)
					}
					if err := f.cache.DelFileItem(file); err != nil {
						logger.Sugar.Errorf("failed to del file item[%s] : %v", file, err)
					}
				}
			}
		default:
			{
				time.Sleep(60 * time.Second)
			}
		}
	}
}
