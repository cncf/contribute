// +build mage

// This is a magefile, and is a "makefile for go".
// See https://magefile.org/
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/carolynvs/magex/pkg"
	"github.com/carolynvs/magex/shx"
)

var (
	prodBranch = "website"
)

const (
	webhook = "https://api.netlify.com/build_hooks/6005c2758a2bba09b6563667"
)

// Ensure Mage is installed and on the PATH.
func EnsureMage() error {
	return pkg.EnsureMage("")
}

// Deploy triggers a Netlify build for the website.
func Deploy() error {
	// Put up a page on the preview that redirects to the live site
	os.MkdirAll("public", 0755)
	err := copy("build/deploy-redirect.html", "public/index.html")
	if err != nil {
		return err
	}

	return triggerNetlifyDeployment()
}

func copy(src string, dest string) error {
	srcF, err := os.Open(src)
	if err != nil {
		return err
	}

	destF, err := os.Create(dest)
	if err != nil {
		return err
	}

	_, err = io.Copy(destF, srcF)
	return err
}

func triggerNetlifyDeployment() error {
	emptyMsg := "{}"
	data := strings.NewReader(emptyMsg)
	fmt.Println("POST", webhook)
	fmt.Println(emptyMsg)

	r, err := http.Post(webhook, "application/json", data)
	if err != nil {
		return err
	}

	if r.StatusCode >= 300 {
		defer r.Body.Close()
		msg, _ := ioutil.ReadAll(r.Body)
		return fmt.Errorf("request failed (%d) %s: %s", r.StatusCode, r.Status, msg)
	}

	return nil
}

// DeployPreview builds the entire website for a preview.
func DeployPreview() error {
	pwd, _ := os.Getwd()
	websiteDir := filepath.Join(filepath.Dir(pwd), "sig-contributor-strategy")

	if _, err := os.Stat(websiteDir); os.IsNotExist(err) {
		err := shx.RunV("git", "clone", "-b", prodBranch, "--recurse-submodules", "https://github.com/cncf/sig-contributor-strategy.git", websiteDir)
		if err != nil {
			return err
		}
	}

	// Build the website with our using our local changes
	err := shx.Command("go", "mod", "edit", "-replace", "github.com/cncf/contribute="+pwd).
		In(websiteDir).RunV()
	if err != nil {
		return err
	}

	err = shx.Command("go", "run", "mage.go", "DeployPreview").
		In(websiteDir).RunV()
	if err != nil {
		return err
	}

	return shx.Command("bash", "-c", "cp -R website/public "+pwd).
		In(websiteDir).RunV()
}
