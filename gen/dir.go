package gen

import "os"

func GenRouterDir() error {
	return os.Mkdir("router", 0644)
}
func GenHandlerDir() error {
	return os.Mkdir("handler", 0644)
}
func GenConfigDir() error {
	return os.Mkdir("config", 0644)
}
func GenMiddlewareDir() error {
	return os.Mkdir("middleware", 0644)
}
func GenModelDir() error {
	return os.Mkdir("model", 0644)
}

