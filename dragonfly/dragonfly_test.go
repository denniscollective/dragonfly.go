package dragonfly_test

import (
	//"fmt"
	"github.com/denniscollective/dragonfly.go/dragonfly"
	"testing"
)

const stubB64Job string = "W1siZmYiLCIvVXNlcnMvZGVubmlzL3dvcmtzcGFjZS96aXZpdHkvcHVibGljL2NvbnRlbnQvcGhvdG9zZXRzL29yaWdpbmFsc19hcmNoaXZlLzAwMC8wMDAwMDAvMDAwMDAwMDA3LzAwMDAwMDAwMjQtaC1vcmlnLmpwZyJdLFsicCIsInRodW1iIiwiMjB4MjAiXV0"

func TestFetch(t *testing.T) {
	file, err := dragonfly.ImageFor(stubB64Job)

	if err != nil {
		t.Errorf("ImgFor(stub) failed %s", err)
	}

	if file == nil || len(file.Name()) < 10 {
		t.Error("expected a file Object")
	}

}

func TestFirstStepFailingErrorPropigation(t *testing.T) {
	jobstr := "W1siZmYiLCJwYXJ0eV90aW1lIl0sWyJwIiwidGh1bWIiLCIyMHgyMCJdXQ" //fetches a nonexistent file called partytime in step one
	_, err := dragonfly.ImageFor(jobstr)
	if err == nil {
		t.Error("nonexistent file party_time is supposed to fail fetching")
		return
	}
	if err.Error() != "open party_time: no such file or directory" {
		t.Errorf("Deconde job should have gotten fetch file failed, got %s", err)
	}
}
func TestDecodeThingThatNeedsTwoEquals(t *testing.T) {
	jobstr := "W1siZmYiLCIvVXNlcnMvZGVubmlzL3dvcmtzcGFjZS96aXZpdHkvcHVibGljL2ltYWdlcy9pY29ucy9kZWZhdWx0XzI1Ni5qcGciXSxbInAiLCJ0aHVtYiIsIjgweDgwIyJdXQ"
	job, err := dragonfly.Decode(jobstr)

	if err != nil {
		t.Errorf("Deconde job got error %s", err)
	}

	if len(job.Steps) != 2 {
		t.Error("job should have two steps")
	}

}

func TestDecodeDragonfly(t *testing.T) {
	job, err := dragonfly.Decode(stubB64Job)

	if err != nil {
		t.Errorf("Deconde job got error %s", err)
	}

	if len(job.Steps) != 2 {
		t.Error("job should have two steps")
	}

	//if job.Steps[0].Command != "ff" {
	//t.Error("the first test of the stub is supposed to be fetch File")
	//}

	//if args := job.Steps[1].Args; args[0] != "thumb" && args[1] != "20x20" {
	//t.Error("second step should be a resize to thumbnail 20x20 job")
	//}
}

func TestDecodeFailse(t *testing.T) {
	job, err := dragonfly.Decode("this is y i'm hawt")
	if err == nil {
		t.Error("Decode errors aren't propagating")
	}

	if len(job.Steps) != 0 {
		t.Error("Decode return a nil job")
	}
}
