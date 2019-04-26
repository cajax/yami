package yami

import "testing"

func TestGetFirstVideoTrack(t *testing.T) {

	t0 := &Track{Type: TrackGeneral, StreamSize: 0}
	t1 := &Track{Type: TrackAudio, StreamSize: 1}
	t2 := &Track{Type: TrackVideo, StreamSize: 2}
	t3 := &Track{Type: TrackVideo, StreamSize: 3}

	m := &MediaInfo{&Media{Tracks: []*Track{t0, t1, t2, t3}}}

	if m.GetFirstVideoTrack().StreamSize != t2.StreamSize {

		t.Error("Incorrectly determined first VideoTrack")
	}
}

func TestGetFirstAudioTrack(t *testing.T) {

	t0 := &Track{Type: TrackGeneral, StreamSize: 0}
	t1 := &Track{Type: TrackVideo, StreamSize: 1}
	t2 := &Track{Type: TrackAudio, StreamSize: 2}
	t3 := &Track{Type: TrackVideo, StreamSize: 3}

	m := &MediaInfo{&Media{Tracks: []*Track{t0, t1, t2, t3}}}

	if m.GetFirstAudioTrack().StreamSize != t2.StreamSize {

		t.Error("Incorrectly determined first AudioTrack")
	}

}

func TestGetTracksByType(t *testing.T) {

	m := &MediaInfo{&Media{Tracks: []*Track{
		{Type: TrackGeneral},
		{Type: TrackVideo},
		{Type: TrackAudio},
		{Type: TrackVideo},
		{Type: TrackVideo},
	}}}
	tracks := m.GetTracksByType(TrackVideo)
	if len(tracks) != 3 {
		t.Error("Unexpected number of returned tracks")
	}

	for _, track := range tracks {
		if track.Type != TrackVideo {
			t.Error("Incorrectly determined track type")
		}
	}
}
