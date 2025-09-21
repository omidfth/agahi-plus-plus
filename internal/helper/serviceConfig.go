package helper

import (
	"agahi-plus-plus/internal/constant"
	"agahi-plus-plus/internal/database"
)

type ServiceConfig struct {
	System   systemConfig    `mapstructure:"system"`
	Database database.Config `mapstructure:"database"`
	Http     httpConfig      `mapstructure:"http"`
	Divar    divarConfig     `mapstructure:"divar"`
	App      app             `mapstructure:"app"`
	JWT      jwtConfig       `mapstructure:"jwt"`
	Zarinpal zarinpal        `mapstructure:"zarinpal"`
	Yektanet yekranet        `mapstructure:"yektanet"`
	Prompt   prompt          `mapstructure:"prompt"`
}

type systemConfig struct {
	DevelopMode     bool   `mapstructure:"develop_mode"`
	LogPath         string `mapstructure:"log_path"`
	ShutdownTimeout int    `mapstructure:"shutdown_timeout"`
}

type httpConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type divarConfig struct {
	Apartment divar    `mapstructure:"apartment"`
	Car       divar    `mapstructure:"car"`
	General   divar    `mapstructure:"general"`
	Api       divarApi `mapstructure:"api"`
}

type divarApi struct {
	GetPosts     string `mapstructure:"get_posts"`
	GetPost      string `mapstructure:"get_post"`
	EditPost     string `mapstructure:"edit_post"`
	DeleteWidget string `mapstructure:"delete_widget"`
	UploadImage  string `mapstructure:"upload_image"`
	Addons       string `mapstructure:"addons"`
}

type divar struct {
	ClientID         string                `mapstructure:"client_id"`
	ClientSecret     string                `mapstructure:"client_secret"`
	ApiKey           string                `mapstructure:"api_key"`
	RedirectUrl      string                `mapstructure:"redirect_uri"`
	OAuth            divarOauth            `mapstructure:"oauth"`
	OAuthToken       divarOauthToken       `mapstructure:"oauth_token"`
	OAuthPhoneNumber divarOauthPhoneNumber `mapstructure:"oauth_phone_number"`
	Scopes           string                `mapstructure:"scopes"`
}

type divarOauth struct {
	BaseUrl      string `mapstructure:"base_url"`
	ResponseType string `mapstructure:"response_type"`
}

type divarOauthToken struct {
	BaseUrl   string `mapstructure:"base_url"`
	GrantType string `mapstructure:"grant_type"`
}

type divarOauthPhoneNumber struct {
	BaseUrl string `mapstructure:"base_url"`
}

type app struct {
	Salt                         string      `mapstructure:"salt"`
	InquiryCost                  int         `mapstructure:"inquiry_cost"`
	InquiryRepo                  InquiryRepo `mapstructure:"inquiry_repo"`
	FrontEndLoginRedirect        string      `mapstructure:"front_end_login_redirect"`
	FrontEndPurchaseRedirect     string      `mapstructure:"front_end_purchase_redirect"`
	FrontEndAccessDeniedRedirect string      `mapstructure:"front_end_access_denied_redirect"`
	FrontEndEntryRedirect        string      `mapstructure:"front_end_entry_redirect"`
	Hash                         Hash        `mapstructure:"hash"`
	Test                         bool        `mapstructure:"test"`
	Kenar                        string      `mapstructure:"kenar_redirect_base_url"`
	KenarButton                  string      `mapstructure:"kenar_divar_button"`
	LLMApiKey                    string      `mapstructure:"llm_api_key"`
	LLMUrl                       string      `mapstructure:"llm_url"`
}

type InquiryRepo struct {
	BarcodeInquiryUrl string `mapstructure:"inquiry_by_barcode_url"`
	GreenInquiryUrl   string `mapstructure:"inquiry_by_green_number_url"`
	InquiryToken      string `mapstructure:"inquiry_api_token"`
	OwnerNationalID   string `mapstructure:"owner_national_id"`
	WalletIdentifier  string `mapstructure:"wallet_identifier"`
}

type jwtConfig struct {
	Secret     string `mapstructure:"secret"`
	ExpireHour int    `mapstructure:"expire_hour"`
}

type zarinpal struct {
	MerchantID  string `mapstructure:"merchant_id"`
	CallbackUrl string `mapstructure:"callback_url"`
	Sandbox     bool   `mapstructure:"sandbox"`
}

type Hash struct {
	Salt      string `mapstructure:"salt"`
	Alphabet  string `mapstructure:"alphabet"`
	MinLength int    `mapstructure:"min_length"`
}

type yekranet struct {
	FrontRedirectUrl string              `mapstructure:"front_redirect_url"`
	Apartment        yektanetDivarConfig `mapstructure:"apartment"`
	Car              yektanetDivarConfig `mapstructure:"car"`
}

type yektanetDivarConfig struct {
	RedirectUrl  string `mapstructure:"redirect_url"`
	ClientID     string `mapstructure:"client_id"`
	ResponseType string `mapstructure:"response_type"`
}

type prompt struct {
	ApiKey     string `mapstructure:"api_key"`
	Url        string `mapstructure:"url"`
	Role       string `mapstructure:"role"`
	Message    string `mapstructure:"message"`
	OutputPath string `mapstructure:"output_path"`
}

func (c *ServiceConfig) GetDivarConfig(service string) divar {
	var config divar
	switch service {
	case constant.ApartmentServiceName:
		config = c.Divar.Apartment
	}

	return config
}
