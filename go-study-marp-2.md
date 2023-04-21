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
### 목차

1. 자료구조
2. 에러핸들링
3. 고루틴
4. 채널과 컨텍스트

---
### 1. 자료구조(data structure)

**목차**

1.1. 자료구조란?
1.2. 리스트(List)
1.3. 링(Ring)
1.4. 맵(Map)

---
### 1.1. 자료구조란?

**데이터 값의 모임**, 각 원소들이 논리적으로 정의된 규칙에 의해 나열되며 자료에 대한 처리를 효율적으로 수행할 수 있도록 자료를 구분하여 표현한 것.

ex) 데이터 -> 책, 메모리 -> 책장

책을 책장에 남는 공간이 없게 정리할 수 도 있을 것이고, 책을 찾기 쉽게 찾기 위해서 장르별로 분리하거나 제목 순서별로 정렬하는 등의 일련의 행위 -> 데이터를 목적에 부합하는 효율적인 방법을 위한 구조 [자료구조]



---
### 1.2. 리스트(List)

**리스트(List)** 는 일련의 항목들을 저장하고 조작하는 자료 구조입니다. 각 항목은 순서대로 나열되며, 각각의 항목은 데이터와 다음 항목에 대한 참조(포인터 또는 링크)를 가지고 있습니다.

![list](diagram/list.png)

리스트는 일반적으로 배열과는 다르게 동적으로 크기를 변경할 수 있습니다. 이러한 동적 크기 변경 능력 때문에, 리스트는 데이터를 자주 추가하거나 제거하는 상황에서 유용하게 사용됩니다. 또한, 리스트는 항목이 추가되거나 제거될 때마다 다시 할당할 필요가 없어 메모리를 효율적으로 사용할 수 있습니다.

---
### 1.2 리스트(List)

자료의 추가, 삭제, 검색

**활용**
- 데이터베이스의 인덱스
- 운영체제의 프로세스,스레드의 스케줄링

---
### 1.2 리스트(List)

```go
package main

import (
    "container/list"
    "fmt"
)

func main() {
    stack := list.New()

    // Push elements onto stack
    stack.PushBack(1)
    stack.PushBack(2)
    stack.PushBack(3)
	// 3 -> 2 -> 1

    // Pop elements from stack
    for stack.Len() > 0 {
        e := stack.Back()
        stack.Remove(e)
        fmt.Println(e.Value)
    }
}
```

---
### 1.3 링(Ring)

---
### 1.4 맵(Map)

---
### 2. 에러핸들링(error handling)

**목차**

1. 에러반환
2. 에러타입
3. 패닉

---
### 3. 고루틴(goroutine)

**목차**

1. 스레드란?
2. 고루틴 사용
3. 고루틴의 동작 방법
4. 동시성 프로그래밍 주의점
5. 뮤텍스와 데드락
6. 다른 자원 관리 기법

---
### 4. 채널과 컨텍스트

**목차**

1. 채널
2. 컨텍스트