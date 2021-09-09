package metrics

// TODO this will become AWS costs instead
import "github.com/tyrannosaurus-becks/team-dashboard/internal/models"

func newInfraCostPerBooking(config *models.Config) *infraCostPerBooking {
	// TODO
	return &infraCostPerBooking{}
}

type infraCostPerBooking struct {
	// TODO
}

func (s *infraCostPerBooking) Name() string {
	return "infrastructure-cost-per-booking"
}

func (s *infraCostPerBooking) Value() (float64, error) {
	// TODO
	return 0, nil
}

func (s *infraCostPerBooking) Tags() *[]string {
	return nil
}
