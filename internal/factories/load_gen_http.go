package factories

import (
	"github.com/clean-arc/internal/controllers"
)

func MakeLoadHttpController() *controllers.LoadHttpController {
	return controllers.NewLoadHttpController()
}
