# go-emacs-go-base64

This package is alternative to [emacs-module-go-base64](https://github.com/tsuu32/emacs-module-go-base64) powered by [go-emacs](https://github.com/sigma/go-emacs).

## Build
```sh
make
```

## Usage
```sh
emacs -L .
```

and

```emacs-lisp
(require 'go-emacs-go-base64)

(go-emacs-go-base64-encode "hello")
;; => "aGVsbG8="

(go-emacs-go-base64-decode "aGVsbG8=")
;; => "hello"
```

## note
The performance of go-emacs-go-base64 functions is the same as emacs-module-go-base64's.
