package dragonfly

import (
	"os"
)

const Stub string = "W1siZmYiLCIvVXNlcnMvZGVubmlzL3dvcmtzcGFjZS96aXZpdHkvcHVibGljL2NvbnRlbnQvcGhvdG9zZXRzL29yaWdpbmFsc19hcmNoaXZlLzAwMC8wMDAwMDAvMDAwMDAwMDA3LzAwMDAwMDAwMjQtaC1vcmlnLmpwZyJdLFsicCIsInRodW1iIiwiMjB4MjAiXV0"

func ImageFor(jobstr string) (*os.File, error) {
	job, err := Decode(jobstr)

	if err != nil {
		panic(err)
	}

	file, err := job.Apply()

	return file, err
}
