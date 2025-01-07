package confog

type (
	Config struct {
		App      App
		Db       Db
		Jwt      Jwt
		Kafka    Kafka
		Grpc     Grpc
		Paginate Paginate
	}

	App struct {
		Name  string
		Url   string
		Stage string
	}

	Db struct {
		Url string
	}

	Jwt struct {
		AccessSecretKey  string
		RefreshSecretKey string
		ApiSceretKey     string
		AccessDuration   int64
		RefreshDuration  int64
		ApiDuration      int64
	}

	Kafka struct {
		Url    string
		ApiKey string
		Secret string
	}

	Grpc struct {
		AuthUrl      string
		PlayerUrl    string
		ItemUrl      string
		InventoryUrl string
		PaymentUrl   string
	}

	Paginate struct {
		ItemNextPageBasedUrl      string
		InventoryNextPageBasedUrl string
	}
)
