package uptimerobot

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/terraform"
	"github.com/louy/terraform-provider-uptimerobot/uptimerobot/api"
)

func testAccCheckAlertContactDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(uptimerobotapi.UptimeRobotApiClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "uptimerobot_alert_contact" {
			continue
		}

		id, err := strconv.Atoi(rs.Primary.ID)
		if err != nil {
			return err
		}

		_, err = client.GetAlertContact(id)

		if err == nil {
			return fmt.Errorf("Alert contact still exists")
		}

		// Verify the error is what we want
		if strings.Contains(err.Error(), "test") {
			return err
		}
	}

	return nil
}
