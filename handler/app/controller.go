package app

import (
	bController "agahi-plus-plus/handler/controller"
	"go.uber.org/zap"
)

type controller struct {
	userController    bController.UserController
	paymentController bController.PaymentController
	postController    bController.PostController
	planController    bController.PlanController
	divarController   bController.DivarController
	promptController  bController.PromptController
}

func (a *application) InitController(srv *service, repo *repository, logger *zap.Logger) *controller {
	var ctrl controller
	ctrl.userController = bController.NewUserController(srv.userService, repo.appLogRepository, a.config, logger)
	ctrl.paymentController = bController.NewPaymentController(srv.zarinpalService, a.config, logger)
	ctrl.postController = bController.NewPostController(srv.postService, a.config, logger)
	ctrl.planController = bController.NewPlanController(srv.planService, a.config, logger)
	ctrl.divarController = bController.NewDivarController(srv.divarService, a.config, logger)
	ctrl.promptController = bController.NewPromptController(srv.promptService, a.config, logger)

	return &ctrl
}
