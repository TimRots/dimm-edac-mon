// dimm-edac-mon - Monitor Dimms for Error Detection And Correction events on Linux
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
)

const (
	EdacPath         = "/sys/devices/system/edac/mc/"
	MemoryController = "mc[0-9]*"
	Dimm             = "dimm[0-9]*"
)

// edacAttributes maps an attribute key to it's description value
type edacAttributes map[string]string

// edac attributes and descriptions from https://www.kernel.org/doc/Documentation/ABI/testing/sysfs-devices-edac
var (
	memoryControllerAttributes = edacAttributes{
		"ce_count":            "total count of correctable errors that have occurred on this memory controller",
		"ce_noinfo_count":     "number of CEs that have occurred on this memory controller wherewith no information as to which DIMM slot is having errors",
		"mc_name":             "type of memory controller that is being utilized",
		"seconds_since_reset": "elapsed seconds since the last counter reset",
		"size_mb":             "count of megabytes, of memory that this memory controller manages",
		"ue_count":            "total count of uncorrectable errors that have occurred on this memory controller",
		"ue_noinfo_count":     "the number of UEs that have occurred on this memory controller with no information as to which DIMM slot is having errors.",
	}
	dimmAttributes = edacAttributes{
		"dimm_ce_count":  "total count of correctable errors that have occurred on this DIMM",
		"dimm_dev_type":  "display what type of DRAM device is being utilized on this DIMM (x1, x2, x4, x8, ...).",
		"dimm_edac_mode": "type of Error detection and correction is being utilized",
		"dimm_location":  "the location (csrow/channel, branch/channel/slot or channel/slot) of the dimm",
		"dimm_mem_type":  "display what type of memory is currently on this csrow",
		"dimm_ue_count":  "total count of uncorrectable errors that have occurred on this DIMM.",
		"size":           "DIMM memory size in MB",
	}
)

// sorted takes a string/string-map as input and returns a sorted string slice of it's keys
func (a edacAttributes) sorted() []string {
	var attributes []string
	for k := range a {
		attributes = append(attributes, k)
	}
	sort.Strings(attributes)
	return attributes
}

// get fetches edac attributes for each memory controller
func (a edacAttributes) get() {
	// find all memory controllers
	memoryControllers, err := filepath.Glob(filepath.Join(EdacPath, MemoryController))
	if err != nil {
		log.Fatal(err)
	}
	switch len(memoryControllers) {
	case 0:
		log.Fatal("no memory controllers detected")
	default:
		// iterate over each memory controller
		for _, mc := range memoryControllers {
			memoryControllerBasePath := mc
			fmt.Println(memoryControllerBasePath)

			// fetch mc edac attributes
			for _, attribute := range memoryControllerAttributes.sorted() {
				attributeValue, err := ioutil.ReadFile(filepath.Join(memoryControllerBasePath, attribute))
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%v: %v", attribute, string(attributeValue))
			}
			fmt.Println()

			// find all dimms on this memory controller
			dimms, err := filepath.Glob(filepath.Join(memoryControllerBasePath, Dimm))
			if err != nil {
				log.Fatal(err)
			}
			// iterate over each dimm on mcX
			for _, dimm := range dimms {
				fmt.Println(dimm)
				// get the value for each key in dimmAttributes
				for _, attribute := range dimmAttributes.sorted() {
					attributeValue, err := ioutil.ReadFile(filepath.Join(dimm, attribute))
					if err != nil {
						fmt.Printf("%v: Not available", attribute)
					}
					fmt.Printf("%v: %v", attribute, string(attributeValue))
				}
				fmt.Println()
			}
		}
	}
}

func main() {
	memoryControllerAttributes.get()
}
