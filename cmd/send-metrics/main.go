package send_metrics

import (
	"github.com/tyrannosaurus-becks/team-dashboard/internal/models"
	"log"

	"github.com/tyrannosaurus-becks/team-dashboard/internal"
)

func main() {
	if err := internal.Run(&models.Config{}); err != nil {
		log.Fatal(err)
	}
}
