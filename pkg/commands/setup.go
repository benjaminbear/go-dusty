package commands

import (
	"fmt"
	"os/user"

	"github.com/shirou/gopsutil/mem"

	"github.com/benjaminbear/go-dusty/constants"

	"github.com/benjaminbear/go-dusty/pkg/inputparse"

	"github.com/benjaminbear/go-dusty/pkg/rpcclient"
)

func SetupDustyConfig(username string, specsRepo string, vmMemory int, sshKeyPath string, update bool) (err error) {
	fmt.Println("We just need to verify a few settings before we get started.")

	if username != "" {
		fmt.Printf("Setting username to %s based on flag\n", username)
	} else {
		username, err = getUsername()
		if err != nil {
			return err
		}
	}

	err = verifyUsername(username)
	if err != nil {
		return err
	}

	if sshKeyPath != "" {
		fmt.Printf("Setting ssh_key_path to %s based on flag\n", sshKeyPath)
	} else {
		sshKeyPath, err = getSSHKeyPath()
		if err != nil {
			return err
		}
	}

	if specsRepo != "" {
		fmt.Printf("Setting specs_repo to %s based on flag\n", specsRepo)
	} else {
		specsRepo, err = getSpecsRepo()
		if err != nil {
			return err
		}
	}

	if vmMemory != 0 {
		fmt.Printf("Setting vm_memory to %v based on flag\n", vmMemory)
	} else {
		vmMemory, err = getVmSize()
		if err != nil {
			return err
		}
	}

	conn, err := rpcclient.CreateConnection()
	if err != nil {
		return err
	}

	err = conn.Setup(username, specsRepo, int32(vmMemory), sshKeyPath, update)
	if err != nil {
		return err
	}

	err = conn.Close()
	if err != nil {
		return err
	}
	return nil
}

func verifyUsername(username string) error {
	_, err := user.Lookup(username)
	if err != nil {
		return err
	}

	return nil
}

func getUsername() (userName string, err error) {
	userObj, err := user.Current()
	if err != nil {
		return userName, err
	}

	userName = userObj.Username

	fmt.Println()
	fmt.Printf("Is %s the username under which you'll primarily use Dusty?", userName)

	yes, err := inputparse.AskForYesNo(true)
	if err != nil {
		return userName, err
	}

	if !yes {
		fmt.Print("Enter your actual username: ")

		_, err = fmt.Scanf("%s\n", &userName)
		if err != nil {
			return userName, err
		}
	}

	return userName, nil
}

func getSpecsRepo() (specsRepo string, err error) {
	fmt.Println()
	fmt.Println("Repos may be specified with a URL (e.g. github.com/org/repo) or an absolute file path to a local repository")
	fmt.Print("Input the path to your specs repo, or leave blank to start with the example specs")

	specsRepo, err = inputparse.AskForWord(false)
	if err != nil {
		return specsRepo, err
	}

	if specsRepo == "" {
		fmt.Println("Using example specs repo", constants.ExampleRepo)
		specsRepo = constants.ExampleRepo
	}

	return specsRepo, nil
}

func getVmSize() (vmMegs int, err error) {
	virtualMem, err := mem.VirtualMemory()
	if err != nil {
		return vmMegs, err
	}

	memoryMegs := int(virtualMem.Total / (1 << 20))
	vmMegs = getRecommendedVmSize(memoryMegs)

	fmt.Println()
	fmt.Printf("Your system seems to have %v megabytes of memory. We would like to allocate %v to your vm. Is that ok?", memoryMegs, vmMegs)

	allocate, err := inputparse.AskForYesNo(true)
	if err != nil {
		return vmMegs, err
	}

	if !allocate {
		fmt.Print("Please input the number of megabytes to allocate to the vm")
		vmMegs, err = inputparse.AskForInteger(true)
		if err != nil {
			return vmMegs, err
		}
	}

	return vmMegs, nil
}

func getRecommendedVmSize(sysMemory int) int {
	// all math is done in megabytes
	if sysMemory >= 16*(1<<10) {
		return 6 * (1 << 10)
	} else if sysMemory >= 8*(1<<10) {
		return 4 * (1 << 10)
	} else {
		return 2 * (1 << 10)
	}
}

func getSSHKeyPath() (string, error) {
	fmt.Println()
	fmt.Println("The private ssh key will be used to checkout git repositories")
	fmt.Print("Input the path to your private ssh key, or leave blank to use the default one ($HOME/.ssh/id_rsa) or none")

	sshKeyPath, err := inputparse.AskForWord(false)
	if err != nil {
		return sshKeyPath, err
	}

	if sshKeyPath == "" {
		fmt.Println("Using default ssh key", constants.DefaultSSHKeyPath)
		sshKeyPath = constants.DefaultSSHKeyPath
	}

	return sshKeyPath, nil
}
