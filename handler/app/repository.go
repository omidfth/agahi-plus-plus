package app

import (
	bRepository "agahi-plus-plus/internal/repository"
	"agahi-plus-plus/internal/repository/api"
	"agahi-plus-plus/internal/repository/db"
	"go.uber.org/zap"
)

type repository struct {
	oauthRepository       bRepository.OAuthRepository
	userRepository        bRepository.UserRepository
	zarinpalApiRepository bRepository.ZarinpalRepository
	userPaymentRepository bRepository.UserPaymentRepository
	postApiRepository     bRepository.PostApiRepo
	postDBRepository      bRepository.PostDbRepo
	pricingRepository     bRepository.PricingRepository
	divarRepository       bRepository.DivarRepository
	promptRepository      bRepository.PromptRepository
	profileRepository     bRepository.ProfileRepository
	appLogRepository      bRepository.AppLogRepository
	configDbRepository    bRepository.ConfigRepository
}

func (a *application) InitRepository(logger *zap.Logger) *repository {
	var repo repository
	repo.oauthRepository = api.NewOAuthRepository(logger)
	repo.userRepository = db.NewUserRepository(a.db, logger)
	repo.zarinpalApiRepository = api.NewZarinpal(a.config.Zarinpal.MerchantID, a.config.Zarinpal.Sandbox)
	repo.userPaymentRepository = db.NewUserPaymentRepository(a.db, logger)
	repo.postDBRepository = db.NewPostDb(a.db, logger)
	repo.pricingRepository = db.NewPricingDB(a.db, logger)
	repo.divarRepository = api.NewDivarApi(a.config, logger)
	repo.promptRepository = api.NewPromptApi(logger, a.config)
	repo.profileRepository = db.NewProfileDB(a.db, logger)
	repo.appLogRepository = db.NewAppLog(a.db, logger)
	repo.configDbRepository = db.NewConfigDb(a.db, logger)
	repo.postApiRepository = api.NewPostApi(a.config, logger)

	return &repo
}
