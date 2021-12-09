package g42vbcloud

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/helper/pathorcontents"
)

var (
	G42VB_REGION_NAME                = os.Getenv("G42VB_REGION_NAME")
	G42VB_ACCOUNT_NAME               = os.Getenv("G42VB_ACCOUNT_NAME")
	G42VB_ADMIN                      = os.Getenv("G42VB_ADMIN")
	G42VB_ENTERPRISE_PROJECT_ID_TEST = os.Getenv("G42VB_ENTERPRISE_PROJECT_ID_TEST")
	G42VB_ACCESS_KEY                 = os.Getenv("G42VB_ACCESS_KEY")
	G42VB_SECRET_KEY                 = os.Getenv("G42VB_SECRET_KEY")
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"g42vbcloud": testAccProvider,
	}
}

func testAccPreCheckRequiredEnvVars(t *testing.T) {
	if G42VB_REGION_NAME == "" {
		t.Fatal("G42VB_REGION_NAME must be set for acceptance tests")
	}
}

func testAccPreCheck(t *testing.T) {
	testAccPreCheckRequiredEnvVars(t)
}

func testAccPreCheckAdminOnly(t *testing.T) {
	if G42VB_ADMIN == "" {
		t.Skip("G42VB_ADMIN must be set for acceptance tests")
	}
}

func testAccPreCheckEpsID(t *testing.T) {
	if G42VB_ENTERPRISE_PROJECT_ID_TEST == "" {
		t.Skip("This environment does not support EPS_ID tests")
	}
}

func testAccPreCheckOBS(t *testing.T) {
	if G42VB_ACCESS_KEY == "" || G42VB_SECRET_KEY == "" {
		t.Skip("G42VB_ACCESS_KEY and G42VB_SECRET_KEY must be set for OBS acceptance tests")
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func envVarContents(varName string) (string, error) {
	contents, _, err := pathorcontents.Read(os.Getenv(varName))
	if err != nil {
		return "", fmt.Errorf("Error reading %s: %s", varName, err)
	}
	return contents, nil
}

func envVarFile(varName string) (string, error) {
	contents, err := envVarContents(varName)
	if err != nil {
		return "", err
	}

	tmpFile, err := ioutil.TempFile("", varName)
	if err != nil {
		return "", fmt.Errorf("Error creating temp file: %s", err)
	}
	if _, err := tmpFile.Write([]byte(contents)); err != nil {
		_ = os.Remove(tmpFile.Name())
		return "", fmt.Errorf("Error writing temp file: %s", err)
	}
	if err := tmpFile.Close(); err != nil {
		_ = os.Remove(tmpFile.Name())
		return "", fmt.Errorf("Error closing temp file: %s", err)
	}
	return tmpFile.Name(), nil
}
