package banner_test

import (
	"encoding/json"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/domain/banner"
	"github.com/guregu/null"
	"testing"
)

func TestBanner_NewBanner(t *testing.T) {
	b, err := banner.NewBanner()
	if err != nil {
		t.Errorf("Error while creating banner: %v", err)
	}
	if b == nil {
		t.Errorf("Error while creating banner")
	}
}

func TestBanner_SetContent(t *testing.T) {
	b, _ := banner.NewBanner()
	content := json.RawMessage(`{"name": "test name", "text": "Test text"}`)
	b.SetContent(content)

	if string(b.GetContent()) != string(content) {
		t.Errorf("Expected content to be %v, got %v", string(content), string(b.GetContent()))
	}
}

func TestBanner_AddTag_RemoveTag(t *testing.T) {
	b, _ := banner.NewBanner()
	tags := []int{1, 2, 3}

	for _, tag := range tags {
		b.AddTag(tag)
	}

	if len(b.GetTag()) != len(tags) {
		t.Errorf("Expected %d tags to be added, got %d", len(tags), len(b.GetTag()))
	}

	for _, tag := range tags {
		found := false
		for _, bTag := range b.GetTag() {
			if bTag == tag {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected tag %d to be added", tag)
		}
	}

	for _, tag := range tags {
		b.RemoveTag(tag)
	}

	if len(b.GetTag()) != 0 {
		t.Errorf("Expected all tags to be removed, but %d tags remain", len(b.GetTag()))
	}
}

func TestBanner_SetFeatures(t *testing.T) {
	b, _ := banner.NewBanner()

	featureID := int64(100)
	b.SetFeatures(null.IntFrom(featureID))

	if !b.GetFeatures().Valid || b.GetFeatures().Int64 != int64(featureID) {
		t.Errorf("Expected Features to be %d, got %d", featureID, b.GetFeatures().Int64)
	}

	// Проверка установки NULL
	b.SetFeatures(null.Int{})
	if b.GetFeatures().Valid {
		t.Errorf("Expected Features to be NULL")
	}
}

func TestBanner_SetIsActive(t *testing.T) {
	b, _ := banner.NewBanner()

	b.SetIsActive(true)
	if !b.GetIsActive() {
		t.Errorf("Expected IsActive to be true")
	}

	b.SetIsActive(false)
	if b.GetIsActive() {
		t.Errorf("Expected IsActive to be false")
	}
}
