package gfx

import "github.com/csnewman/go-gfx/hal"

type Shader struct {
	shader hal.Shader
}

type ShaderConfig struct {
	Source string
}

func (a *Application) LoadShader(cfg ShaderConfig) (*Shader, error) {
	shader, err := a.graphics.CreateShader(hal.ShaderConfig{
		Source: cfg.Source,
	})
	if err != nil {
		return nil, err
	}

	return &Shader{
		shader: shader,
	}, nil
}

type ShaderFunction struct {
	function hal.ShaderFunction
}

func (s *Shader) Function(name string) (*ShaderFunction, error) {
	sf, err := s.shader.ResolveFunction(name)
	if err != nil {
		return nil, err
	}

	return &ShaderFunction{
		function: sf,
	}, nil
}
