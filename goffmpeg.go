package maincon

import "github.com/giorgisio/goav/avformat"

func main() {

	filename := "/Users/yangaowei/waqu/crawler/93gi60z0863yeqg3_normal.mp4"

	// Register all formats and codecs
	avformat.AvRegisterAll()

	// Open video file
	if avformat.AvformatOpenInput(&ctxtFormat, filename, nil, nil) != 0 {
		log.Println("Error: Couldn't open file.")
		return
	}

	// Retrieve stream information
	if ctxtFormat.AvformatFindStreamInfo(nil) < 0 {
		log.Println("Error: Couldn't find stream information.")
		return
	}

	//...

}
