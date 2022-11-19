package controller

func (c *controller) Register() {
	c.hanler.Add("/me", c.CmdMe)
}
