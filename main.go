package main

import "fmt"

func main() {
	a := []string{"a", "b", "c"}
	b := []string{"b", "c", "d"}

	resultA, resultB := Result(a, b)
	fmt.Printf("1. answer is : %v \n2. answer is : %v", resultA, resultB)
}

func Result(a []string, b []string) ([]string, []string) {
	resultA := ArrayUnionNoDuplicates(a, b)
	resultB := ArrayMergeCutDuplicates(a, b)
	return resultA, resultB
}

// array ที่รวมข้อมูลโดยไม่มีข้อมูลซ้ำ
func ArrayUnionNoDuplicates(a []string, b []string) []string {
	anb := mergeArrays(a, b)
	var result []string
	for i := 0; i < len(anb); i++ {
		// เช็คถ้าจำนวนสมาชิดใน Result มีค่า = 0 ให้เอาสมาชิกตัวแรกสุดใส่เข้าไปใน result ได้เลย
		if len(result) == 0 {
			result = append(result, anb[i])
		} else {
			found := false
			for j := 0; j < len(result); j++ {
				// นำสมาชิกตำแหน่งที่ i มาเช็คว่าเป็นสมาชิกอยู่ใน result แล้วหรือยัง
				if anb[i] == result[j] {
					found = true
					break
				}
			}
			// เมื่อทำการตรวจสอบแล้วว่าสมาชิดตำแหน่งที่ i ไม่ได้อยู่ใน result เลย ก็จะทำการใส่ค่าเข้าไปใน result
			if !found {
				result = append(result, anb[i])
			}
		}
	}
	return result
}

// array ที่รวมข้อมูลที่รับเข้าด้วยกันโดยตัดข้อมูลซ้ำ
func ArrayMergeCutDuplicates(a []string, b []string) []string {
	anb := mergeArrays(a, b)
	var result []string
	for i := 0; i < len(anb); i++ {
		// เซ็ตค่าเริ่มต้นเพื่อนับจำนวนที่ซ้ำของสมาชิกตำแหน่งที่ i
		count := 0
		for j := 0; j < len(anb); j++ {
			// เช็คค่าว่าตำแหน่งที่ i มีตัวซ้ำกี่ตัว ถ้าซ้ำกันให้ +1 จำนวนซ้ำของสมาชิกตำแหน่งที่ i
			if anb[i] == anb[j] {
				count += 1
			}
		}
		// เช็คว่าสมาชิกตำแหน่งที่ i มีตัวซ้ำน้อยกว่าหรือเท่ากับ 1 ไหม ซึ่งที่เป็น 1 เนื่องจากนับค่าตัวเองเท่านั้น เพราะทำการเทียบ ตัวมันเอง กับตัวมันเองที่ทำการ merge array มาก่อน ถ้ามากกว่า 1 ก็คือซ้ำ
		if count <= 1 {
			result = append(result, anb[i])
		}
	}
	return result
}

func mergeArrays(a, b []string) []string {
	// สร้าง array เพื่อเก็บสมาชิกจากอาเรย์ a และ b
	var result []string

	// รวบรวมข้อมูลจากอาร์เรย์ a
	for i := 0; i < len(a); i++ {
		result = append(result, a[i])
	}

	// รวบรวมข้อมูลจากอาร์เรย์ b
	for i := 0; i < len(b); i++ {
		result = append(result, b[i])
	}

	return result
}
