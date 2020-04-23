package cybuf

var (
	//debugLog *log.Logger
	// errorLog *log.Logger
	marshalSep byte = ' '
)

func init() {
	//debugLog = log.New(os.Stdout, "Debug ", log.LstdFlags|log.Lshortfile)
	// errorLog = log.New(ioutil.Discard, "Error ", log.LstdFlags|log.Lshortfile)
}

func SetMarshalSep(sep byte) {
	if sep != '\n' && sep != '\t' && sep != ' ' {
		marshalSep = sep
	}
}
