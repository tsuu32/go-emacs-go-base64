(defvar go-emacs-go-base64-test/test-path
  (f-dirname load-file-name))

(defvar go-emacs-go-base64-test/root-path
  (f-parent go-emacs-go-base64-test/test-path))

(module-load (f-expand "go-emacs-go-base64.so" go-emacs-go-base64-test/root-path))
