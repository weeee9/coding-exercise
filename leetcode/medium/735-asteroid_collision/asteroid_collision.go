package asteroidcollision

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0, len(asteroids))

	for _, asteroid := range asteroids {
		if len(stack) > 0 {
			// 如果堆疊中有小行星，則處理碰撞事件
			stack = calculateCollide(stack, asteroid)
		} else {
			// 如果堆疊中沒有小行星，將當前小行星推入堆疊
			stack = append(stack, asteroid)
		}
	}

	return stack
}

// 處理小行星碰撞事件
func calculateCollide(asteroids []int, current int) []int {
	for len(asteroids) > 0 {
		// 獲取堆疊頂部的小行星
		peak := asteroids[len(asteroids)-1]

		// 當前是向右運動的小行星絕對不會發生和之前的小星行碰撞
		if current > 0 {
			return append(asteroids, current)
		}

		// 如果頂部的小行星是向左運動，且當前小行星也是向左運動
		// 它們會一起向左運動，不會發生碰撞
		if peak < 0 && current < 0 {
			return append(asteroids, current)
		}

		// 如果頂部的小行星是向右運動，當前小行星是向左運動，則判斷碰撞的結果
		if peak > 0 && current < 0 {
			// 當前小行星的絕對值等於頂部小行星，則兩個一起消失
			if current*-1 == peak {
				return asteroids[:len(asteroids)-1]
			}

			// 當前小行星的絕對值小於頂部小行星的絕對值，則當前小行星直接消失
			// 且不用做後續的判斷
			if current*-1 < peak {
				return asteroids
			}

			// 當前小行星的絕對值大於頂部小行星的絕對值，則頂部小行星消失
			if current*-1 > peak {
				asteroids = asteroids[:len(asteroids)-1]
			}
		}
	}

	// 所有小行星都消失了，只剩下當前小行星
	return []int{current}
}
