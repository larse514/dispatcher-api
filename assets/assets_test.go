package assets

import "testing"

func TestGetAsset(t *testing.T) {
	s, _ := getAsset("sam.yml")

	if *s == "" {
		t.Log("error retrieving assets")
		t.Fail()
	}

}
func TestGetAssetFail(t *testing.T) {
	_, err := getAsset("doesntexist.yml")

	if err == nil {
		t.Log("error not returned")
		t.Fail()
	}

}
