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

```bash
# 환경변수에 텔레그램 봇 토큰 및 GPT-3 API 키를 등록
export TELEGRAM_BOT_TOKEN=123456789:ABCDEFGHIJKLMNOPQRSTUVWXYZ
export OPENAI_API_KEY=ABCDEFGHIJKLMNOPQRSTUVWXYZ

# 실행

```