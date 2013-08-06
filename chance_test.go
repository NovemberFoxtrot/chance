package chance

import (
	"chance"
	"testing"
	"time"
)

var testScoreTable = []struct {
	ups      int
	downs    int
	expected int
}{
	{100, 0, 100},
	{1000, 0, 1000},
	{0, 0, 0},
	{0, 100, -100},
	{100, 99, 1},
}

func TestScore(t *testing.T) {
	for _, testdata := range testScoreTable {
		score := chance.Score(testdata.ups, testdata.downs)

		if score != testdata.expected {
			t.Errorf("Expexted %d from score(%d,%d)", testdata.expected, testdata.ups, testdata.downs)
		}
	}
}

var testHotTable = []struct {
	ups      int
	downs    int
	date     time.Time
	expected float64
}{
	{100, 0, time.Date(2013, time.August, 8, 6, 0, 0, 0, time.UTC), 5377.857711111111},
	{1000, 900, time.Date(2013, time.August, 8, 6, 0, 0, 0, time.UTC), 5377.857711111111},
	{100, 0, time.Date(2013, time.August, 8, 5, 0, 0, 0, time.UTC), 5377.777711111111},
	{100, 0, time.Date(2012, time.August, 8, 5, 0, 0, 0, time.UTC), 4676.977711111111},
}

func TestHot(t *testing.T) {
	for _, testdata := range testHotTable {
		hottness := chance.Hot(testdata.ups, testdata.downs, testdata.date)

		if hottness != testdata.expected {
			t.Error("Hotness", hottness, "Expected", testdata.expected)
		}
	}
}

var testConfidenceTable = []struct {
	ups      int
	downs    int
	expected float64
}{
	{100, 0, 0.9925588255880455},
	{1000, 0, 0.9992505931958449},
	{0, 0, 0},
	{0, 100, 0.06983532663200934},
	{100, 99, 0.7062135168679761},
}

func TestConfidence(t *testing.T) {
	for _, testdata := range testConfidenceTable {
		score := chance.Confidence(testdata.ups, testdata.downs)

		if score != testdata.expected {
			t.Errorf("Expexted %d from confidence(%d,%d)", testdata.expected, testdata.ups, testdata.downs)
		}
	}
}
