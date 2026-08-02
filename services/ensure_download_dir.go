package services

import "os"

func ensureDownloadDir(dataDir string) error {
	_, err := os.Stat(dataDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dataDir, 0755)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}
