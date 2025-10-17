# go-gfx

<img align="right" width="25%" src="mascot.png">
Cross-platform graphics framework for Go, covering both mid-level and low-level components.

> [!WARNING]
> Work in progress.

## Core components

The repository contains the following packages, intended for low-level access and control:

- `vk` - Low-level Vulkan bindings.

These package are suitable for creating a custom engine from scratch.

## GFX framework

The repository contains the following packages, intended to provide a minimal Graphics framework:

- `gfx` - User facing API, acting as a hardware abstraction layer.
- `gfx/backends/`
    - `appkit` - macOS windowing backend.
    - `vulkan` - Vulkan rendering backend.
    - `windows` - Windows windowing backend.

These packages intend to cover the most common usages.

### Platforms

|           | macOS      | Linux                 | Windows    | iOS | Android | Web |
|-----------|------------|-----------------------|------------|-----|---------|-----|
| Windowing | 🏗️ AppKit | ⌛ Wayland </br> ⌛ X11 | 🏗️        | 💤  | 💤      | 💤  | 
| Rendering | 🏗️ Vulkan | ⌛ Vulkan              | 🏗️ Vulkan | 💤  | 💤      | 💤  | 

✅ = Supported.  
🏗️ = Work in progress.  
⌛ = Future.  
💤 = No near term plans.

