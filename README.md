# telegram-gpt-go

---

## go.mod 사용법
```bash
# go.mod 파일 생성
go mod init github.com/motolies/telegram-gpt-go

# go 의존성 추가
# 처음 한 번만 실행하면 됨
go get github.com/go-telegram-bot-api/telegram-bot-api
```

---

## 사용법

### 로컬에서 실행
```bash
# 환경변수에 텔레그램 봇 토큰 및 GPT-3 API 키를 등록
export TELEGRAM_BOT_TOKEN=123456789:ABCDEFGHIJKLMNOPQRSTUVWXYZ
export OPENAI_API_KEY=ABCDEFGHIJKLMNOPQRSTUVWXYZ

# 빌드
go build -a -o telegram-gpt-bot ./cmd/server

# 실행
./telegram-gpt-bot
```

### 도커로 실행
```bash
# 도커 이미지 빌드
docker build -t telegram-gpt-bot .

# 실행
docker run -d --restart=unless-stopped \
    --name telegram-gpt-bot \
    -e TELEGRAM_BOT_TOKEN=123456789:ABCDEFGHIJKLMNOPQRSTUVWXYZ \
    -e OPENAI_API_KEY=ABCDEFGHIJKLMNOPQRSTUVWXYZ \
    telegram-gpt-bot
```

---

## buildx

### multi-arch build
```bash
docker buildx build --platform linux/amd64,linux/arm64 --no-cache --push -t telegram-gpt-bot .
```