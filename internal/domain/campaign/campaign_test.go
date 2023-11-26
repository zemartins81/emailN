package campaign

import "testing"

func TestNewCampaign(t *testing.T) {
	// Define variables for the campaign
	name := "Campaign x"
	content := "BOdy"
	contacts := []string{"email@e.com", "email2@e.com"}

	// Call the function to create a new campaign
	campaign := NewCampaign(name, content, contacts)

	// Test if the ID of the new campaign is as expected
	if campaign.ID != "1" {
		t.Errorf("expected 1 got %s", campaign.ID)
	}

	// Test if the campaign name matches the one assigned
	if campaign.Name != name {
		t.Errorf("expected %s got %s", name, campaign.Name)
	}

	// Test if the campaign content matches the one assigned
	if campaign.Content != content {
		t.Errorf("expected %s got %s", content, campaign.Content)
	}

	// Test if the campaign contact list size matches the one assigned
	if len(campaign.Contacts) != len(contacts) {
		t.Errorf("expected %d got %d", len(contacts), len(campaign.Contacts))
	}
}
