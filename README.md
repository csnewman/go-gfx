# go-gfx
<img align="right" width="25%" src="mascot.png">
Cross-platform graphics framework for Go.

> [!WARNING]
> Work in progress.

## Structure

The repository contains the following packages:
- `gfx` - User facing API. Will eventually be stable.
- `hal` - Internal hardware abstraction layer. Expect breaking changes.
- `internal/`
  - `appkit` - macOS windowing backend.
  - `metal` - Metal rendering backend.

## Platforms

|           | macOS                    | Linux                 | Windows                | iOS | Android | Web |
|-----------|--------------------------|-----------------------|------------------------|-----|---------|-----|
| Windowing | 🏗️ AppKit               | ⌛ Wayland </br> ⌛ X11 | ⌛                      | 💤  | 💤      | 💤  | 
| Rendering | 🏗️ Metal </br> ⌛ Vulkan | ⌛ Vulkan              | ⌛ Vulkan <br/> 💤 DX12 | 💤  | 💤      | 💤  | 

✅ = Supported.  
🏗️ = Work in progress.  
⌛ = Future.  
💤 = No near term plans.

Platforms not listed here, such as those with licensing restrictions, can be supported by implementing the `hal` layer.
This allows you to benefit from the `gfx` abstraction and a unified codebase.