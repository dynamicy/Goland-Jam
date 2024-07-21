# 使用官方的 Golang 鏡像作為基礎鏡像
FROM golang:1.22-alpine

# 設置工作目錄
WORKDIR /app

# 將 go.mod 和 go.sum 複製到工作目錄
COPY go.mod ./
COPY go.sum ./

# 下載依賴
RUN go mod download

# 複製所有源代碼到工作目錄
COPY ./cmd ./cmd
COPY ./pkg ./pkg
COPY ./docs ./docs

# 編譯 Go 程式
RUN go build -o /Goland-Jam ./cmd/main.go

# 設置運行時的命令
CMD ["/Goland-Jam"]