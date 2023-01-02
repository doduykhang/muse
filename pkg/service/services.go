package service

type BusinessService struct {
	SongService SongService
	AuthService AuthService
}

type ComponentService struct {
	CacheService CacheService
	PasswordService PasswordService
	FileService FileService
	CookieService CookieService
}

var (
	businessService = &BusinessService{
		SongService: NewSongService(),
		AuthService: NewAuthService(),
	}

	componentService = &ComponentService{
		CacheService: NewCacheService(),
		PasswordService: NewPasswordService(),
		FileService: NewFileService(),
		CookieService: NewCookieService(),
	}
)

func GetBusinessServices() *BusinessService {
	return businessService
}

func GetComponentServices() *ComponentService {
	return componentService
}
