package routes

import "github.com/doduykhang/muse/pkg/controllers"

var (
	appControllers *controllers.Controllers
)

func init() {
	appControllers = controllers.GetControllers()
}
