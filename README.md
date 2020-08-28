terraform-provider-rio
---

A terraform provider which supplies functionality of [rio](https://github.com/polydawn/rio).


## Update rio and timeless API versions to head
```
go mod edit -replace=go.polydawn.net/go-timeless-api=github.com/polydawn/go-timeless-api@$(git ls-remote git@github.com:polydawn/go-timeless-api.git HEAD | awk '{print $1}'); go mod tidy

go mod edit -replace=go.polydawn.net/rio=github.com/polydawn/rio@$(git ls-remote git@github.com:polydawn/rio.git HEAD | awk '{print $1}'); go mod tidy

go mod edit -replace=xi2.org/x/xz=github.com/xi2/xz@$(git ls-remote git@github.com:xi2/xz.git HEAD | awk '{print $1}'); go mod tidy
```
