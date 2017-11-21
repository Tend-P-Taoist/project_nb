package result

type Result struct {
	Code int
	Message string
	Data map[string]interface{}

}

func (this Result)JSON() map[string]interface{}{
	return map[string]interface{}{"code":this.Code,"message":this.Message,"data":this.Data}
}

