package filtergo

type GHFilter struct {
	// GH filter
	x  float64
	dx float64
	dt float64
	g  float64
	h  float64
}

func (p *GHFilter) Update(z float64) (x, dx float64) {
	// prediction
	x_pred := p.x + (p.dx * p.dt)
	//update
	residual := z - x_pred
	p.dx += p.h * (residual) / p.dt
	p.x = x_pred + p.g*residual
	return p.x, p.dx
}
