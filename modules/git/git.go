package git

import (
	"fmt"
	"os/exec"
)

func CloneRepo(repoUrl string, destinationPath string) error {
	cloneCmd := "sudo git clone ssh://" + repoUrl + " " + destinationPath
	fmt.Println(cloneCmd)
	_, err := exec.Command("/bin/sh", "-c", cloneCmd).Output()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
