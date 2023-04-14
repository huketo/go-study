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
9. 컨텍스트

---

## 1. 슬라이스(Slice)

`슬라이스(Slice)`는 배열의 뷰로서, 원본 배열의 일부를 가리키는 포인터, 길이, 용량 정보를 가지고 있습니다.

```go
var mySlice []T
```

여기서 T는 슬라이스의 원소 타입입니다.

---

## 1. 슬라이스(Slice)

**슬라이스(Slice)** 는 배열의 일부를 참조하기 때문에, 배열의 길이보다 작거나 같아야 합니다. 슬라이스를 생성할 때는 배열에서의 시작 인덱스와 끝 인덱스를 지정하여 배열의 특정 부분을 슬라이스로 추출할 수 있습니다. 

```go
myArray := [5]int{1, 2, 3, 4, 5} // [1, 2, 3, 4, 5]
mySlice := myArray[1:4] // [2, 3, 4]
```
---

## 1. 슬라이스(Slice)

**특징**

1. 슬라이스는 동적으로 크기가 조절될 수 있기 때문에 유연하게 배열 데이터를 조작할 수 있습니다.
2. 슬라이스는 원본 배열의 일부를 참조하고 있기 때문에, 슬라이스를 수정하면 원본 배열의 데이터도 변경되며, 원본 배열의 변경도 슬라이스에 반영됩니다.
3. 슬라이스는 매우 효율적인 메모리 관리를 제공하며, 가비지 컬렉션에도 좋은 성능을 보입니다.

---
## 1. 슬라이스(Slice) - 사용 예: 가비지 컬렉션

슬라이스는 가변적인 길이를 가지므로, 슬라이스의 길이가 변경되더라도 가비지 컬렉션은 적절하게 처리됩니다. 만약 슬라이스의 길이가 줄어들면, 가비지 컬렉션은 슬라이스의 길이가 줄어든 부분에 할당된 메모리를 회수하여 낭비를 방지합니다. 

```sh
# slice/gc_slice.go 실행결과
초기 HeapAlloc: 235088 bytes
슬라이스 초기 상태: [0 0 0 0 0]
슬라이스에 값 할당 후: [1 2 3 4 5]
슬라이스 할당 후 HeapAlloc: 238256 bytes
슬라이스 길이 변경 후: [1 2 3]
가비지 컬렉션을 통한 메모리 회수 후 HeapAlloc: 237152 bytes
HeapSys: 3801088 bytes
```

---

## 1. 슬라이스(Slice) - 사용 예: 복사

슬라이스는 참조 타입이기 때문에, 슬라이스를 복사하면 기존 슬라이스와 동일한 데이터를 참조하게 됩니다. 이를 통해 메모리 복사를 피하고 효율적으로 데이터를 공유할 수 있습니다.

```sh
# slice/copy_slice.go 실행결과
slice1: [1 2 3 4 5]
slice1_copy: [1 2 3 4 5]
=> slice1_copy 값 변경
slice1: [99 2 3 4 5]
slice1_copy: [99 2 3 4 5]
=> slice1_new_copy 값 변경
slice1: [99 2 3 4 5]
slice1_copy: [99 2 3 4 5]
slice1_new_copy: [1 2 3 4 5]
```

---

## 2. 메서드(Method)

**메서드(Method)** 는 특정한 타입(Type)에 연결된 함수(Function)입니다. 메서드는 구조체(Struct)를 포함한 모든 사용자 정의 타입(Type)에 연결할 수 있습니다. 메서드는 특정한 타입의 인스턴스에 대해 호출되며, 해당 인스턴스에 대한 작업을 수행하는 함수입니다.

---
## 2. 메서드(Method)

```go
type Rectangle struct {
	width  float64
	height float64
}

// Rectangle 타입에 대한 메서드 정의
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	rectangle := Rectangle{width: 10, height: 5}
	fmt.Println("넓이:", rectangle.Area()) // 메서드 호출
}
```

---
## 2. 메서드(Method)

메서드를 왜 사용하는가? - *OOP(객체 지향 프로그래밍)*

1. 코드의 응집성
2. 가독성
3. 재사용성
4. 확장성

---
## 2. 메서드(Method) - 종류

- 포인터 메서드(Pointer Methods)
- 값 타입 메서드(Value Receiver Methods)

→ 이 두 유형의 메서드는 동작과 동작을 수행하는 방식, 그리고 메서드가 호출될 때의 동작 등에서 차이가 있습니다.

---
## 2. 메서드(Method) - Pointer Methods

```go
type MyStruct struct {
	value int
}

func (m *MyStruct) PointerMethod() {
	// m은 MyStruct 객체의 포인터를 참조
	// MyStruct 객체의 상태를 변경할 수 있음
}

func main() {
	obj := &MyStruct{}
	obj.PointerMethod() // 포인터 메서드 호출
}
```

---
## 2. 메서드(Method) - Value Receiver Methods

```go
type MyStruct struct {
	value int
}

func (m MyStruct) ValueMethod() {
	// m은 MyStruct 객체의 복사본을 참조
	// MyStruct 객체의 상태를 변경할 수 없음
}

func main() {
	obj := MyStruct{}
	obj.ValueMethod() // 값 타입 메서드 호출
}
```
---
## 2. 메서드(Method) - 예제코드

```sh
# method/pointer_method.go 실행결과
[0] personA: &{Jinsu 23}
=> ChangeName_A() : Value Receiver Methods
[1] personA: &{Jinsu 23}
=> ChangeName_B() : Pointer Methods
[2] personA: &{Minsu 23}
```

→ 포인터 메서드는 객체의 포인터를 참조하여 객체의 상태를 변경할 수 있고, 값 타입 메서드는 객체의 복사본을 참조하여 객체의 상태를 변경할 수 없습니다.

---
## 3. 인터페이스(Interface)

**인터페이스(Interface)** 는 메서드의 집합을 나타내는 타입입니다. 인터페이스는 메서드의 시그니처만을 정의하고, 메서드의 구현은 인터페이스를 사용하는 타입에서 수행됩니다. 다른 언어의 추상 클래스와 유사한 개념으로 볼 수 있습니다.