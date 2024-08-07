package config

type Config struct {

	// Database
	DBHost     string `mapstructure:"DB_HOST"`
	DBUsername string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_DBNAME"`
	DBPort     string `mapstructure:"DB_PORT"`
	SSLMode    string `mapstructure:"SSL_MODE"`

	// Mail
	SmtpHost      string `mapstructure:"SMTP_HOST"`
	SmtpPort      int    `mapstructure:"SMTP_PORT"`
	SenderEmail   string `mapstructure:"SENDER_EMAIL"`
	EmailPassword string `mapstructure:"EMAIL_PASSWORD"`

	// app
	AppPort string `mapstructure:"APP_PORT"`

	// middlewares
	SecretJWT string `mapstructure:"SECRET_JWT"`

	// oy
	BaseUrl  string `mapstructure:"BASEURL"`
	Username string `mapstructure:"USERNAME"`
	ApiKey   string `mapstructure:"API_KEY"`

	// Cloudinary
	CloudName   string `mapstructure:"CLOUD_NAME"`
	CloudKey    string `mapstructure:"CLOUD_KEY"`
	ApiSecret   string `mapstructure:"API_SECRET"`
	CloudFolder string `mapstructure:"CLOUD_FOLDER"`

	// IAK
	UsernameIak        string `mapstructure:"USERNAME_IAK"`
	ApiKeyIak          string `mapstructure:"API_KEY_IAK"`
	BaseUrlIakPostPaid string `mapstructure:"BASEURL_IAK_Post_Paid"`
}

var (
	AppConfig Config
)

func LoadConfig() *Config {
	// viper.SetConfigType("env")
	// viper.SetConfigName("public")
	// viper.AddConfigPath(".")

	// if err := viper.ReadInConfig(); err != nil {
	// 	panic(fmt.Errorf("fatal error config file: %w", err))
	// }

	// err := viper.Unmarshal(&AppConfig)
	// if err != nil {
	// 	panic(fmt.Errorf("fatal error config file: %w", err))
	// }
	AppConfig = Config{
		DBHost:     "103.174.115.239",
		DBUsername: "postgres",
		DBPassword: "ZfcnO9WYdcbNqvRtwH4WrA99nSax2t8yBdASIoNzKCNeIs8hXI8xy7lQWrD4DeyP",
		DBName:     "postgres",
		DBPort:     "5432",
		SSLMode:    "require",

		SmtpHost:      "smtp.gmail.com",
		SmtpPort:      587,
		SenderEmail:   "skuypay10@gmail.com",
		EmailPassword: "wrdvpewgnatmfnpa",
		AppPort:       "8080",
		SecretJWT:     "S3CR3t",

		BaseUrl:  "https://api-stg.oyindonesia.com/api",
		Username: "darulfh",
		ApiKey:   "6197039c-6a7d-47fa-9fa2-094639f9d0f1",

		CloudName:   "doizrq2wc",
		CloudKey:    "494258757634754",
		ApiSecret:   "UmpmAQ17amfjHx-HMb8a5_DsNrg",
		CloudFolder: "ppob",

		UsernameIak:        "081383566428",
		ApiKeyIak:          "45666815bf56abb1PGMd",
		BaseUrlIakPostPaid: "https://testpostpaid.mobilepulsa.net",
	}

	return &AppConfig
}
