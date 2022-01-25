package update

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"time"

	"github.com/Masterminds/semver"
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
	fmt.Printf("Updating from %s to %s\n", version, releaseInfo.Name)

	name := fmt.Sprintf("%s_%s", os, arch)
	if os == "windows" {
		name = name + ".exe"
	}

	err = updateBinary(ctx, releaseInfo.getDownloadURL(name))
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("Update includes:")
	fmt.Print(releaseInfo.Body)
	fmt.Println()

	return err
}
