# file-organizer

指定したディレクトリ内のファイルを、ファイル名に対応したサブディレクトリへ移動します。

![file-organizer_002](https://user-images.githubusercontent.com/345832/201456013-98ce1deb-a644-4d20-b22e-bb17f895cc94.gif)

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

## Release

### Linux

```sh
GOOS=linux GOARCH=amd64 go build
mkdir file-organizer_v1.0.0_linux_amd64
cp LICENSE README.md file-organizer_v1.0.0_linux_amd64/
mv file-organizer file-organizer_v1.0.0_linux_amd64/
tar cvfz file-organizer_v1.0.0_linux_amd64.tar.gz file-organizer_v1.0.0_linux_amd64/
```

### Windows

```sh
GOOS=windows GOARCH=amd64 go build
mkdir file-organizer_v1.0.0_windows_amd64
cp LICENSE README.md file-organizer_v1.0.0_windows_amd64/
mv file-organizer.exe file-organizer_v1.0.0_windows_amd64/
zip -r file-organizer_v1.0.0_windows_amd64.zip file-organizer_v1.0.0_windows_amd64/
```

### Darwin

```sh
GOOS=darwin GOARCH=amd64 go build
mkdir file-organizer_v1.0.0_darwin_amd64
cp LICENSE README.md file-organizer_v1.0.0_darwin_amd64/
mv file-organizer file-organizer_v1.0.0_darwin_amd64/
zip -r file-organizer_v1.0.0_darwin_amd64.zip file-organizer_v1.0.0_darwin_amd64/
```

## License:

Copyright (C) 2022 mikoto2000

This software is released under the MIT License, see LICENSE

このソフトウェアは MIT ライセンスの下で公開されています。 LICENSE を参照してください。


## Author:

mikoto2000 <mikoto2000@gmail.com>

