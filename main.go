package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	out, err := exec.Command("java", "-version").Output()

	if err != nil {
		installDir, err := os.UserHomeDir()

		os.Mkdir(installDir+"java", os.ModePerm)

		downloadPath := installDir + "\\java\\openjdk-18.0.2.1_windows-x64_bin.zip"

		fmt.Println("Installing Java zip in " + downloadPath)
		out, err := os.Create(downloadPath)
		defer out.Close()

		// Java not installed
		resp, err := http.Get("https://download.java.net/java/GA/jdk18.0.2.1/db379da656dc47308e138f21b33976fa/1/GPL/openjdk-18.0.2.1_windows-x64_bin.zip")

		if err != nil {
			fmt.Println("Error downloading Java: " + err.Error())
		}

		defer resp.Body.Close()

		downloadedBytes, copyErr := io.Copy(out, resp.Body)

		if copyErr != nil {
			fmt.Println("Error copying: " + copyErr.Error())
		}

		fmt.Println("Downloaded bytes: " + string(downloadedBytes))
	}

	fmt.Println(out)
}
