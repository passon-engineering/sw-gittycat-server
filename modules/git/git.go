package git

import (
	"fmt"
	"os/exec"
)

func CloneRepo(repoUrl string, clonePath string) error {
	cloneCmd := "git clone " + repoUrl + " " + clonePath
	fmt.Println(cloneCmd)
	_, err := exec.Command("/bin/sh", "-c", cloneCmd).Output()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
