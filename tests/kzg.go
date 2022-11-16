package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Inphi/eip4844-interop/shared"
)

// The test generates commitment, versionedHash, aggregatedProof for a small portion of generated data.
// The output is then compared to c-kzg's
func main() {
	data := make([]byte, 5875)
	for i := 0; i < len(data); i++ {
		data[i] = byte(i % 250)
	}
	blobs := shared.EncodeBlobs(data)
	commitments, versionedHashes, aggregatedProof, err := blobs.ComputeCommitmentsAndAggregatedProof()
	if err != nil {
		log.Fatalf("Error getting proofs: %v", err)
	}

	file, err := os.OpenFile("output.txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	for i := 0; i < len(commitments); i++ {
		fmt.Fprintln(file, commitments[i].String()[2:])
	}

	for i := 0; i < len(versionedHashes); i++ {
		fmt.Fprintln(file, versionedHashes[i].String()[2:])
	}

	fmt.Fprintln(file, aggregatedProof.String()[2:])
}
