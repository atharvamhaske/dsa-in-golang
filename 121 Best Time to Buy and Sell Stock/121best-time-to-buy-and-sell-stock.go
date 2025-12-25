func maxProfit(prices []int) int {
   profit, min, max := 0, prices[0], prices[0]

   for _, v := range prices {
    if v < min {
        max = v
        min = v
    } else if v > max {
        max = v
    }
    if (max - min) > profit {
        profit = max-min
    }
   }
   return profit
}