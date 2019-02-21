package middlewares

/*
func Logger(out io.Writer) func(next HandlerFn) HandlerFn {
	return func(next HandlerFn) HandlerFn {
		logger := log.New(out, "*goback*", 0)
		return func(w http.ResponseWriter, req *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, req)
			end := time.Since(start)
			recorder := httptest.NewRecorder()
			next.ServeHTTP(recorder, req)
			logger.Printf(" -- %s - %v \"%s %s %d\" \"%s\" \"%s\" - %v",
				strings.Split(req.RemoteAddr, ":")[0], start.Format("2006-01-02 15:04:05"),
				req.Method, req.URL.Path, recorder.Code, req.Referer(), req.UserAgent(), end)
		}
	}
}*/
