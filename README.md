# LGlivephoto

[![Go Reference](https://pkg.go.dev/badge/github.com/ryanking13/lglivephoto.svg)](https://pkg.go.dev/github.com/ryanking13/lglivephoto)

**_LG smart phone is [dead](https://www.lgnewsroom.com/2021/04/lg-to-close-mobile-phone-business-worldwide/). Long live the Live Photo._**

A Simple golang utility to unpack and pack [LG Live Photo](https://www.lg.com/uk/support/product-help/CT00008356-20150844039308).

_NOTE: this module is tested with LG G7._
## Usage (binary)

For windows (x86_64) download executables below:

- [livephoto-unpack](https://github.com/ryanking13/lglivephoto/releases/latest/download/livephoto-pack.exe)
- [livephoto-pack](https://github.com/ryanking13/lglivephoto/releases/latest/download/livephoto-pack.exe)

For other platforms, download binaries at [releases](https://github.com/ryanking13/lglivephoto/releases) section.

```bat
livephoto-unpack <livephoto image or directory which contains livephoto images>
```

```bat
livephoto-pack <non-livephoto image> <video to be embedded>
```

## Usage (module)

```sh
go get -u github.com/ryanking13/lglivephoto
```

```go
package main

import (
	"io/ioutil"

	"github.com/ryanking13/lglivephoto"
)

func main() {
	image, video, _ := lglivephoto.Unpack("livephoto.jpg")
	ioutil.WriteFile("livephoto_image.jpg", image, 0644)
	ioutil.WriteFile("livephoto_video.mp4", video, 0644)

	livephoto, _ := lglivephoto.Pack("livephoto_image.jpg", "livephoto_video.mp4")
	ioutil.WriteFile("livephoto_repack.jpg", livephoto, 0644)
}
```

See [document](https://pkg.go.dev/github.com/ryanking13/lglivephoto) for more information.

## Related links

- [Bulk Export, Transfer, Backup, or Convert Live/Motion Photos](https://www.reddit.com/r/lgg7/comments/avmv78/comment/ehns29d)
- [lg-livephoto](https://github.com/coldmund/lg-livephoto)