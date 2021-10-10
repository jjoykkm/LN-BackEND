package config

/*-------------------------------------------------------------------------------------------*/
//                                 LANGUAGE
/*-------------------------------------------------------------------------------------------*/

const (
	EN string = "EN"
	TH string = "TH"
)
type Language struct {
	En	string
	Th	string
}
func GetLanguage() Language {
	return Language{
		En: EN,
		Th: TH,
	}
}
