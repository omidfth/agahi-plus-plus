package app

import (
	bService "agahi-plus-plus/internal/service"
	"go.uber.org/zap"
)

type service struct {
	oauthService       bService.OAuthService
	userService        bService.UserService
	zarinpalService    bService.ZarinpalService
	userPaymentService bService.UserPaymentService
	postService        bService.PostService
	planService        bService.PlanService
	divarService       bService.DivarService
	promptService      bService.PromptService
}

func (a *application) InitService(repo *repository, logger *zap.Logger) *service {
	var srv service
	srv.oauthService = bService.NewOAuthService(repo.oauthRepository, a.config, logger)
	srv.userPaymentService = bService.NewUserPaymentService(repo.userPaymentRepository, logger)
	srv.userService = bService.NewUserService(srv.oauthService, repo.userRepository, srv.userPaymentService, repo.divarRepository, a.config, logger)
	srv.planService = bService.NewPlanService(repo.planRepository, srv.userService, a.config, logger)
	srv.zarinpalService = bService.NewZarinpalService(repo.zarinpalApiRepository, srv.userPaymentService, srv.userService, srv.planService, a.config, logger)
	srv.postService = bService.NewPostService(repo.postApiRepository, repo.postDBRepository, srv.userService, a.config, logger)
	srv.promptService = bService.NewPromptService(repo.promptRepository, srv.postService, srv.userService, logger)
	srv.divarService = bService.NewDivarService(repo.divarRepository, srv.postService, srv.userService, a.config, logger)

	return &srv
}
