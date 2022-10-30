package mp3duration

import (
	"math"
	"testing"
)

func round(number float64, decimals int) float64 {
	n := math.Pow(10, float64(decimals))
	return math.Round(number*n) / n
}

func TestReadFile(t *testing.T) {
	duration, err := ReadFile("testfile.mp3")
	if err != nil {
		t.Fatal(err)
	}

	k := Info{
		Name:         "testfile.mp3",
		Duration:     "00:00:03",
		TimeDuration: 3056326416.0,
		Seconds:      3.056326416,
		SecondsInt:   3,
		Length:       48900,
		Frames:       117,
	}

	switch {
	case duration.Name != k.Name:
		t.Errorf("Name %s is not %s", duration.Name, k.Name)
	case duration.Duration != k.Duration:
		t.Errorf("Duration %s is not %s", duration.Duration, k.Duration)
	case duration.TimeDuration != k.TimeDuration:
		t.Errorf("TimeDuration %d is not %d", int64(duration.TimeDuration), int64(k.TimeDuration))
	case round(duration.Seconds, 8) != round(k.Seconds, 8):
		t.Errorf("Seconds %.8f is not %.8f", duration.Seconds, k.Seconds)
	case duration.SecondsInt != k.SecondsInt:
		t.Errorf("SecondsInt %d is not %d", duration.SecondsInt, k.SecondsInt)
	case duration.Length != k.Length:
		t.Errorf("Length %d is not %d", duration.Length, k.Length)
	case duration.Frames != k.Frames:
		t.Errorf("Frames %d is not %d", duration.Frames, k.Frames)
	}
}
