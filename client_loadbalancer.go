package qingcloud

func (c *Client) AddLoadBalancerBackends(params Params) ([]byte, error) {
	return c.Get("AddLoadBalancerBackends", params)
}
