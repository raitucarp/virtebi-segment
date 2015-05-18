package virtebi

import (
    "testing"
    "os"
)

func TestVirtebiSegment(t *testing.T) {
    corpusPath := os.Getenv("GOPATH")
    c := NewCorpus()
    c.LoadCorpus(corpusPath + "/src/github.com/raitucarp/corpus/corpus.txt")
    cases := []struct {
        origin, result string
    }{
        // facebook is awesome
        {"facebookisawesome", "facebook is awesome"},
        // the facebook
        {"thefacebook", "the facebook"},
        // toyota service manual
        {"toyotaservicemanual", "toyota service manual"},
        // packt publishing
        {"packtpublishing", "packt publishing"},
        // this is a text
        {"thisisatext", "this is a text"},
    }

    for _, text := range cases {
        segment, _ := c.Segment(text.origin)
        if text.result != segment {
            t.Errorf("expected result != result, %s != %s", text.result, segment)
        }
    }
}

