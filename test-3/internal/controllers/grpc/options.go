package grpc

type Option func(*GrpcCtrl)

// SetUsecase -.
func SetUsecase(usecaseName string, usecase interface{}) Option {
	return func(c *GrpcCtrl) {
		// if _, ok := c.usecase[usecaseName]; !ok {
		// 	c.usecase[usecaseName] = make(map[string]interface{})
		// }
		c.usecase[usecaseName] = usecase
	}
}
