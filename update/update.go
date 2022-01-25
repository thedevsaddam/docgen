package update

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"time"

	"github.com/Masterminds/semver"
	"github.com/briandowns/spinner"
)

const (
	layoutISO    = "2006-01-02"
	updateInDays = 7
)

// SelfUpdate update the application to its latest version
// if the current release is 3days old and has a new update
func SelfUpdate(ctx context.Context, buildDate, version string) error {
	if buildDate == "unknown" {
		return errors.New("update: unable to update based on unkown build date/version")
	}
	currBinaryReleaseDate, err := time.Parse(layoutISO, buildDate)
	if err != nil {
		return fmt.Errorf("update: %v", err)
	}
	if (time.Since(currBinaryReleaseDate).Hours() / 24) <= updateInDays {
		return nil
	}

	releaseInfo, err := fetchReleaseInfo(ctx)
	if err != nil {
		return fmt.Errorf("update: %v", err)
	}

	if releaseInfo.Draft || releaseInfo.Prerelease {
		return nil
	}

	c, err := semver.NewConstraint(">" + version)
	if err != nil {
		return fmt.Errorf("update: %v", err)
	}

	v, err := semver.NewVersion(releaseInfo.TagName)
	if err != nil {
		return fmt.Errorf("update: %v", err)
	}
	// Check if the version meets the constraints. The a variable will be true.
	if !c.Check(v) {
		return nil
	}

	os := runtime.GOOS
	arch := runtime.GOARCH
	fmt.Println("Found newer version:", releaseInfo.TagName)

	s := spinner.New(spinner.CharSets[70], 100*time.Millisecond, spinner.WithHiddenCursor(true)) // code:39 is earth for the lib
	s.Prefix = fmt.Sprintf("Updating from %s to %s ( ", version, releaseInfo.Name)
	s.Suffix = ")"
	s.Start()
	name := fmt.Sprintf("%s_%s", os, arch)
	if os == "windows" {
		name = name + ".exe"
	}
	err = updateBinary(ctx, releaseInfo.getDownloadURL(name))
	s.Stop()
	fmt.Println()
	fmt.Println("Update includes:")
	fmt.Print(releaseInfo.Body)
	fmt.Println()

	return err
}
