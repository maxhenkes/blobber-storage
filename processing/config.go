package processing

type Image_config struct {
	height int
	width  int
	name   string
}

type Image_crop struct {
	height int
	width  int
}

func GetConfigurations() *[4]Image_config {
	thumb := Image_config{800, 800, "thumb"}
	small := Image_config{1200, 1200, "small"}
	medium := Image_config{1800, 1800, "medium"}
	large := Image_config{2400, 2400, "large"}

	configurations := [4]Image_config{thumb, small, medium, large}

	return &configurations
}
