package gfx

import "github.com/csnewman/go-gfx/hal"

type Shader struct {
	shader hal.Shader
}

type ShaderConfig struct {
	SPIRV []byte
}

func (a *Application) LoadShader(cfg ShaderConfig) (*Shader, error) {
	shader, err := a.graphics.CreateShader(hal.ShaderConfig{
		SPIRV: cfg.SPIRV,
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
