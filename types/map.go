package types

// MapKeyType map key的约束
type MapKeyType interface {
	~string |
		~uint8 | ~int8 |
		~uint16 | ~int16 |
		~int | ~uint |
		~uint32 | ~int32 |
		~int64 | ~uint64 |
		~float64
}

// GenericMap 泛型 map
type GenericMap[K MapKeyType, V any] map[K]V

// MapKeySelector  [K MapKeyType, T any]
// map key 选择器， MapKeyType 约束 key的类型，T 约束 value 的类型
type MapKeySelector[K MapKeyType, V any] func(v V) K

// Array2Map [K MapKeyType, V any]
// 泛型函数，将一个切片转换为map, 可以自定义 key 值，
// MapKeySelector 是自定义key的函数
// 示例：
//
//	type ob struct {
//		 Id int64 `json:"id"`
//	}
//
//	selector := func(i ob) int64 {
//		 return i.Id
//	}
//
// m := Array2Map(selector, []ob{{Id: 1}, {Id: 2}})
// fmt.Println(m)
// Output:
// map[1:{i} 2:{2}]
func Array2Map[K MapKeyType, V any](
	kSelector MapKeySelector[K, V], // key选择器，自定义函数
	array []V, // 数组
) GenericMap[K, V] {
	m := make(GenericMap[K, V], len(array))
	for _, a := range array {
		m[kSelector(a)] = a
	}
	return m
}
