package main

import (
	"flag"
	"fmt"
	"image-deduplicator/imgdedup"
	"log"
)

func main() {
	// Define and parse flags
	recFlagPtr := flag.Bool("r", false, "Indicates whether the program should look for images in sub directories.")
	algoFlagPtr := flag.Int("algo", 1, "Indicates the algorithm that is to be used for hashing.\n1-aHash, 2-pHash, 3-dHash")
	verboseFlagPtr := flag.Bool("v", false, "Indicates whether the program should print out logs.")
	directoryFlag := flag.String("dir", "./", "Indicates the directory the program will run on.")
	thresholdFlag := flag.Uint64("t", 5, "Indicates the threshold value for the grouping, smaller threshold value yield to higher confidence in results")
	flag.Parse()

	// Assign flag values to corresponding variables
	imgdedup.HASHALGORITHM = *algoFlagPtr
	verbose := *verboseFlagPtr
	recursive := *recFlagPtr
	dir := *directoryFlag
	threshold := *thresholdFlag

	// Create ImageHash structs for files in the directory
	images, err := imgdedup.ListImagesInDir(dir, recursive)
	if err != nil {
		log.Fatal(err)
	}

	// (if verbose print out number of files and file names)
	if verbose {
		fmt.Printf("Found %d images in total.\n", len(images))
		for i, image := range images {
			fmt.Printf("\tFile %d: %s\n", i, image.FileName)
		}
	}

	//Hash images
	// (if verbose, print out progress)
	imgdedup.HashImages(images, verbose)

	// Group images with threshold
	// (if verbose, print out groups)
	imgdedup.GroupImages(images, threshold, verbose)
}
