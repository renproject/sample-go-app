package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/google/uuid"
)

var version = "undefined"

func main() {
	v1 := "v1.0.0-beta.10"
	v2 := "v1.0.0-stable-5f34fd8"
	fmt.Println("Hello world!")
	id := uuid.New()
	fmt.Printf("ver: %s, id: %s\n", version, id.String())
	newer, err := compareVersions(v1, v2)
	if err != nil {
		panic(err)
	}
	if newer {
		fmt.Printf("%s is newer than %s\n", v2, v1)
	} else {
		fmt.Printf("%s is older than %s\n", v2, v1)
	}
}

func compareVersions(curr, latest string) (bool, error) {
	rgx, err := regexp.Compile("v([0-9]+).([0-9]+).([0-9]+)(-(nightly|beta|stable)(.([0-9]+))?)?")
	if err != nil {
		return false, err
	}

	currShares := rgx.FindStringSubmatch(curr)
	latestShares := rgx.FindStringSubmatch(latest)

	currMajor, err := strconv.ParseInt(currShares[1], 10, 64)
	if err != nil {
		return false, err
	}
	currMinor, err := strconv.ParseInt(currShares[2], 10, 64)
	if err != nil {
		return false, err
	}
	currPatch, err := strconv.ParseInt(currShares[3], 10, 64)
	if err != nil {
		return false, err
	}

	latestMajor, err := strconv.ParseInt(latestShares[1], 10, 64)
	if err != nil {
		return false, err
	}

	latestMinor, err := strconv.ParseInt(latestShares[2], 10, 64)
	if err != nil {
		return false, err
	}

	latestPatch, err := strconv.ParseInt(latestShares[3], 10, 64)
	if err != nil {
		return false, err
	}

	if currShares[0] == latestShares[0] ||
		currMajor > latestMajor ||
		(currMajor == latestMajor && currMinor > latestMinor) ||
		(currMajor == latestMajor && currMinor == latestMinor && currPatch > latestPatch) {
		return false, nil
	}

	if currMajor == latestMajor && currMinor == latestMinor && currPatch == latestPatch {
		return compareTags(currShares[5:], latestShares[5:])
	}
	return true, nil
}

func compareTags(curr, latest []string) (bool, error) {
	currType, err := convertTypeToNumber(curr[0])
	if err != nil {
		return false, err
	}

	latestType, err := convertTypeToNumber(latest[0])
	if err != nil {
		return false, err
	}

	if currType > latestType {
		return false, nil
	} else if currType < latestType {
		return true, nil
	}

	if curr[1] != "" && latest[1] == "" {
		return false, nil
	} else if curr[1] == "" && latest[1] != "" {
		return true, nil
	}

	currTypePatch, err := strconv.ParseInt(curr[2], 10, 64)
	if err != nil {
		return false, err
	}

	latestTypePatch, err := strconv.ParseInt(latest[2], 10, 64)
	if err != nil {
		return false, err
	}

	if currTypePatch > latestTypePatch {
		return false, nil
	}
	return true, nil
}

func convertTypeToNumber(releaseType string) (int, error) {
	switch releaseType {
	case "nightly":
		return 1, nil
	case "beta":
		return 2, nil
	case "stable", "":
		return 3, nil
	default:
		return -1, fmt.Errorf("unknown release type: %v", releaseType)
	}
}
