package routes

import (
	"net/http"

	"github.com/doduykhang/muse/pkg/constants"
	"github.com/doduykhang/muse/pkg/controllers"
	"github.com/doduykhang/muse/pkg/dtos"
	"github.com/doduykhang/muse/pkg/lib"
	"github.com/doduykhang/muse/pkg/utils"
	"github.com/go-chi/chi/v5"
)



func authRoute(r chi.Router) {
	lib.JsonRequest[dtos.RegisterRequest](
		r, 
		http.MethodPost, 
		constants.REGISTER_ENDPOINT, 
		controllers.RegisterController,
	)

	lib.JsonRequest[dtos.LoginRequest](
		r, 
		http.MethodPost, 
		constants.LOGIN_ENDPOINT, 
		controllers.LoginController,
		utils.Authorized(constants.USER_ROLE),
	)

	r.With(utils.Authorized(constants.USER_ROLE)).Get("/user", func (w http.ResponseWriter, r *http.Request){
		w.Write([]byte("user secrect"))
	})

	r.With(utils.Authorized(constants.ADMIN_ROLE)).Get("/admin", func (w http.ResponseWriter, r *http.Request){
		w.Write([]byte("admin secrect"))
	})

	r.With(utils.Authenticated()).Get("/free", func (w http.ResponseWriter, r *http.Request){
		w.Write([]byte("free for lofin"))
	})
}
