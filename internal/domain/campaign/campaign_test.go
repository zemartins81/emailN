package campaign

import "testing"

func TestNewCampaign(t *testing.T) {
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email@email.com", "email2@email.com"}

	campaign := NewCampaign(name, content, contacts)

	if campaign.ID != "1" {
		t.Errorf("expected 1")
	}

	if campaign.Name != name {
		t.Errorf("expected correct name")
	}

	if campaign.Content != content {
		t.Errorf("expected correct content")
	}

	if len(campaign.Contacts) != len(contacts) {
		t.Errorf("expected correct contacts")
	}
}
