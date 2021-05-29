package config

// output front html
type OutputData struct {
	ID         int    `json:"id"`
	URL        string `json:"url"`
	Title      string `json:"title"`
	StatusCode string `json:"statuscode"`
	BodyLength string `json:"bodylength"`
	CreateTime string `json:"createtime"`
	Image      string `json:"image"`
	HTTPDump   string `json:"httpdump"`
}

// terminal options
type Terminal struct {
	Probes     map[string]bool // all probe data
	Target     string          // single target
	Targets    string          // multiple targets
	Threads    int             // scan threads
	Proxy      string          // proxy
	Timeout    int             // http request timeout
	ChromePath string          // screenshot chrome path
	Output     string          // output file，default .html
}
