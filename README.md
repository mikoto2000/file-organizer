# file-organizer

## Usage:

```sh
Usage: file-organizer [options...] TARGET_DIRS...
  -dryrun
        dry run flag
  -letter_num int
        文字数。整頓タイプ「letter_num」で使用。ここで指定した長さの先頭文字列が整頓先ディレクトリになります (default 8)
  -pattern string
        正規表現パターン。整頓タイプ「pattern」で使用。正規表現マッチの $1 にマッチする文字列が整頓先ディレクトリになります (default "^(.*?)_")
  -type string
        整頓タイプ(letter_num or pattern). (default "pattern")
  -verbose
        verbose mode flag
```

## Install:

`go install` コマンドでインストールするか、 Rlease からバイナリをダウンロードしてください。

```sh
go install github.com/mikoto2000/file-organizer@latest
```

## License:

Copyright (C) 2022 mikoto2000

This software is released under the MIT License, see LICENSE

このソフトウェアは MIT ライセンスの下で公開されています。 LICENSE を参照してください。


## Author:

mikoto2000 <mikoto2000@gmail.com>

