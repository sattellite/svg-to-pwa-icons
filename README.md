# Generate icons for PWA from SVG file

Project created as a helper utility for some of my projects with PWA (progressive web app). It creates all needed images to PWA for all platforms.

## Installation

Download the latest version from [releases page]
(/sattellite/svg-to-pwa-icons/releases).

Or install it manually. Project written in [Golang](https://go.dev/) and you 
need install it first. Then you can install app with command:

    $ go install github.com/sattellite/svg-to-pwa-icons@latest

## Usage

    $ svg-to-pwa-icons logo.png

App will create "out" directory with all predefined images.

## Images

App have predefined images set that it will be generate:

- `android-chrome-192x192.png`
- `android-chrome-512x512.png`
- `apple-touch-icon-120x120.png`
- `apple-touch-icon-152x152.png`
- `apple-touch-icon-180x180.png`
- `apple-touch-icon-60x60.png`
- `apple-touch-icon-76x76.png`
- `apple-touch-icon.png`
- `favicon-16x16.png`
- `favicon-32x32.png`
- `favicon.ico`
- `icon.png`
- `msapplication-icon-144x144.png`
- `mstile-150x150.png`
