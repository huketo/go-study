---
marp: true
theme: gaia
_class: lead
paginate: true
backgroundColor: #fff
backgroundImage: url('https://marp.app/assets/hero-background.svg')
---
![bg left:40% 80%](https://go.dev/blog/go-brand/Go-Logo/SVG/Go-Logo_Aqua.svg)

# **Golang Study**

## 장진수

---

# 목차

1. 슬라이스
2. 메서드
3. 인터페이스
4. 함수활용
5. 자료구조
6. 에러핸들링
7. 고루틴과 동시성 프로그래밍
8. 채널
9.  컨텍스트

---

# 1. 슬라이스(Slice)

`슬라이스(Slice)`는 배열의 뷰로서, 원본 배열의 일부를 가리키는 포인터, 길이, 용량 정보를 가지고 있습니다.

```go
var mySlice []T
```

여기서 T는 슬라이스의 원소 타입입니다.

---

