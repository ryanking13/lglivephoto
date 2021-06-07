# LGlivephoto

[![Go Reference](https://pkg.go.dev/badge/github.com/ryanking13/lglivephoto.svg)](https://pkg.go.dev/github.com/ryanking13/lglivephoto)

**_LG smart phone is [dead](https://www.lgnewsroom.com/2021/04/lg-to-close-mobile-phone-business-worldwide/). Long live the Live Photo._**

A Simple golang utility to unpack and pack [LG Live Photo](https://www.lg.com/uk/support/product-help/CT00008356-20150844039308).

## Usage (binary)

Download the binary at [releases](https://github.com/ryanking13/lglivephoto/releases) section.

```bat
lglivephoto-unpack <livephoto image or directory which contains livephoto images>
```

```bat
lglivephoto-pack -template <non-livephoto image> -video <video to be embedded>
```

## Usage (module)

```sh
go get -u github.com/ryanking13/lglivephoto
```

```go
package main

import (
    "ioutil"

    "github.com/ryanking13/lglivephoto"
)

func main() {
    image, video, _ := lglivephoto.unpack("livephoto.jpg")
    ioutil.WriteFile(image, "livephoto_image.jpg", 0644)
    ioutil.WriteFile(video, "livephoto_video.mp4", 0644)

    livephoto, _ := lglivephoto.pack("livephoto_image.jpg", "livephoto_video.mp4")
    ioutil.WriteFile(livephoto, "livephoto_repack.jpg")
}
```

See [document](https://pkg.go.dev/github.com/ryanking13/lglivephoto) for more information.